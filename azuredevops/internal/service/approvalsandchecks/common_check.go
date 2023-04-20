package approvalsandchecks

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/pipelineschecks"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/internal/client"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/internal/utils"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/internal/utils/converter"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/internal/utils/tfhelper"
	pipelineschecksv7 "github.com/microsoft/terraform-provider-azuredevops/azuredevops/utils/pipelineschecksextrasv7"
)

// NOTE: In theory the API should accept "agentpool" as well, but the API client requires a project ID
// so it doesn't seem to work and the website UI doesn't have it available
var targetResourceTypes = []string{"endpoint", "environment", "queue", "repository", "securefile", "variablegroup"}

type flatFunc func(d *schema.ResourceData, check *pipelineschecks.GenericCheckConfiguration, projectID string) error
type expandFunc func(d *schema.ResourceData) (*pipelineschecks.GenericCheckConfiguration, string, error)

type approvalAndCheckTypes struct {
	Approval         *pipelineschecks.CheckType
	BranchProtection *pipelineschecks.CheckType
	BusinessHours    *pipelineschecks.CheckType
	TaskCheck        *pipelineschecks.CheckType
}

var approvalAndCheckType = approvalAndCheckTypes{
	Approval: &pipelineschecks.CheckType{
		Id:   converter.UUID("8c6f20a7-a545-4486-9777-f762fafe0d4d"),
		Name: converter.ToPtr("Approval"),
	},
	TaskCheck: &pipelineschecks.CheckType{
		Id: converter.UUID("fe1de3ee-a436-41b4-bb20-f6eb4cb879a7"),
	},
	BranchProtection: &pipelineschecks.CheckType{
		Id: converter.UUID("fe1de3ee-a436-41b4-bb20-f6eb4cb879a7"),
	},
	BusinessHours: &pipelineschecks.CheckType{
		Id: converter.UUID("fe1de3ee-a436-41b4-bb20-f6eb4cb879a7"),
	},
}

// genBaseCheckResource creates a Resource with the common parts
// that all checks require.
func genBaseCheckResource(f flatFunc, e expandFunc) *schema.Resource {
	return &schema.Resource{
		Create: genCheckCreateFunc(f, e),
		Read:   genCheckReadFunc(f),
		Update: genCheckUpdateFunc(f, e),
		Delete: genCheckDeleteFunc(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(2 * time.Minute),
			Read:   schema.DefaultTimeout(1 * time.Minute),
			Update: schema.DefaultTimeout(2 * time.Minute),
			Delete: schema.DefaultTimeout(2 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},
			"target_resource_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"target_resource_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice(targetResourceTypes, false),
			},
		},
	}
}

// doBaseExpansion performs the expansion for the 'base' attributes that are defined in the schema, above
func doBaseExpansion(d *schema.ResourceData, checkType *pipelineschecks.CheckType, settings map[string]interface{}, timeout *int) (*pipelineschecks.GenericCheckConfiguration, string, error) {
	projectID := d.Get("project_id").(string)

	taskCheck := pipelineschecks.GenericCheckConfiguration{
		Type:     checkType,
		Settings: settings,
		Resource: &pipelineschecks.Resource{
			Id:   converter.String(d.Get("target_resource_id").(string)),
			Type: converter.String(d.Get("target_resource_type").(string)),
		},
	}

	if timeout != nil {
		taskCheck.Timeout = timeout
	}

	if d.Id() != "" {
		taskCheckId, err := strconv.Atoi(d.Id())
		if err != nil {
			return nil, "", fmt.Errorf("Error parsing task check ID: (%+v)", err)
		}
		taskCheck.Id = &taskCheckId
	}

	return &taskCheck, projectID, nil
}

// doBaseFlattening performs the flattening for the 'base' attributes that are defined in the schema, above
func doBaseFlattening(d *schema.ResourceData, check *pipelineschecks.GenericCheckConfiguration, projectID string) error {
	d.SetId(fmt.Sprintf("%d", *check.Id))

	d.Set("project_id", projectID)

	if check.Resource == nil {
		return fmt.Errorf("Resource nil")
	}

	d.Set("target_resource_id", check.Resource.Id)
	d.Set("target_resource_type", check.Resource.Type)

	if check.Settings == nil {
		return fmt.Errorf("Settings nil")
	}

	return nil
}

func genCheckCreateFunc(flatFunc flatFunc, expandFunc expandFunc) func(d *schema.ResourceData, m interface{}) error {
	return func(d *schema.ResourceData, m interface{}) error {
		clients := m.(*client.AggregatedClient)
		configuration, projectID, err := expandFunc(d)
		if err != nil {
			return fmt.Errorf(" failed in expandFunc. Error: %+v", err)
		}

		createdCheck, err := clients.V7PipelinesChecksClientExtras.AddGenericCheckConfiguration(clients.Ctx, pipelineschecksv7.AddGenericCheckConfigurationArgs{
			Project:       &projectID,
			Configuration: configuration,
		})
		if err != nil {
			return fmt.Errorf(" failed creating check, project ID: %s. Error: %+v", projectID, err)
		}

		err = flatFunc(d, createdCheck, projectID)
		if err != nil {
			return err
		}
		return genCheckReadFunc(flatFunc)(d, m)
	}
}

func genCheckReadFunc(flatFunc flatFunc) func(d *schema.ResourceData, m interface{}) error {
	return func(d *schema.ResourceData, m interface{}) error {
		clients := m.(*client.AggregatedClient)
		projectID, taskCheckId, err := tfhelper.ParseProjectIDAndResourceID(d)
		if err != nil {
			return err
		}

		taskCheck, err := clients.V7PipelinesChecksClientExtras.GetGenericCheckConfiguration(clients.Ctx, pipelineschecksv7.GetGenericCheckConfigurationArgs{
			Project: &projectID,
			Id:      &taskCheckId,
			Expand:  &pipelineschecks.CheckConfigurationExpandParameterValues.Settings,
		})

		if err != nil {
			if utils.ResponseWasNotFound(err) || strings.Contains(err.Error(), "does not exist.") {
				d.SetId("")
				return nil
			}
			return err
		}

		return flatFunc(d, taskCheck, projectID)
	}
}

func genCheckUpdateFunc(flatFunc flatFunc, expandFunc expandFunc) schema.UpdateFunc { //nolint:staticcheck
	return func(d *schema.ResourceData, m interface{}) error {
		clients := m.(*client.AggregatedClient)
		taskCheck, projectID, err := expandFunc(d)
		if err != nil {
			return err
		}

		updatedBusinessHours, err := clients.V7PipelinesChecksClientExtras.UpdateGenericCheckConfiguration(clients.Ctx,
			pipelineschecksv7.UpdateGenericCheckConfigurationArgs{
				Project:       &projectID,
				Configuration: taskCheck,
				Id:            taskCheck.Id,
			})

		if err != nil {
			return err
		}

		err = flatFunc(d, updatedBusinessHours, projectID)
		if err != nil {
			return err
		}
		return genCheckReadFunc(flatFunc)(d, m)
	}
}

func genCheckDeleteFunc() schema.DeleteFunc { //nolint:staticcheck
	return func(d *schema.ResourceData, m interface{}) error {
		if strings.EqualFold(d.Id(), "") {
			return nil
		}

		clients := m.(*client.AggregatedClient)
		projectID, BusinessHoursID, err := tfhelper.ParseProjectIDAndResourceID(d)
		if err != nil {
			return err
		}

		return clients.V7PipelinesChecksClientExtras.DeleteGenericCheckConfiguration(m.(*client.AggregatedClient).Ctx,
			pipelineschecksv7.DeleteGenericCheckConfigurationArgs{
				Project: &projectID,
				Id:      &BusinessHoursID,
			})
	}
}

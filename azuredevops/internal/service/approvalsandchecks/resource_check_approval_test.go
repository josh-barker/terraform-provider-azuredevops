//go:build (all || resource_check_branch_control) && !exclude_approvalsandchecks
// +build all resource_check_branch_control
// +build !exclude_approvalsandchecks

package approvalsandchecks

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/pipelineschecks"
	"github.com/microsoft/terraform-provider-azuredevops/azdosdkmocks"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/internal/client"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/internal/utils/converter"
	pipelineschecksv7 "github.com/microsoft/terraform-provider-azuredevops/azuredevops/utils/pipelineschecksextrasv7"
	"github.com/stretchr/testify/require"
)

var ApprovalCheckID = 123456789
var ApprovalEndpointID = uuid.New().String()
var ApprovalCheckProjectID = uuid.New().String()
var ApprovalCheckTestProjectID = &ApprovalCheckProjectID

var endpointTypeApproval = "endpoint"
var endpointResourceApproval = pipelineschecks.Resource{
	Id:   &ApprovalEndpointID,
	Type: &endpointTypeApproval,
}

var approver = map[string]interface{}{
	"id": "xxxx",
}
var approvers = []interface{}{approver}

var ApprovalCheckSettings = map[string]interface{}{
	"instructions":              "hello world",
	"minRequiredApprovers":      1,
	"requesterCannotBeApprover": true,
	"approvers":                 approvers,
}

var ApprovalCheckTest = pipelineschecks.GenericCheckConfiguration{
	Id:       &ApprovalCheckID,
	Type:     approvalAndCheckType.Approval,
	Settings: ApprovalCheckSettings,
	Timeout:  converter.ToPtr(20000),
	Resource: &endpointResourceApproval,
}

// verifies that the flatten/expand round trip yields the same branch control
func TestCheckApproval_ExpandFlatten_Roundtrip(t *testing.T) {
	resourceData := schema.TestResourceDataRaw(t, ResourceCheckApproval().Schema, nil)
	flattenCheckApproval(resourceData, &ApprovalCheckTest, ApprovalCheckProjectID)

	ApprovalCheckAfterRoundTrip, projectID, err := expandCheckApproval(resourceData)

	require.Equal(t, ApprovalCheckTest, *ApprovalCheckAfterRoundTrip)
	require.Equal(t, ApprovalCheckProjectID, projectID)
	require.Nil(t, err)
}

// verifies that if an error is produced on create, the error is not swallowed
func TestCheckApproval_Create_DoesNotSwallowError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := ResourceCheckApproval()
	resourceData := schema.TestResourceDataRaw(t, r.Schema, nil)
	flattenCheckApproval(resourceData, &ApprovalCheckTest, ApprovalCheckProjectID)

	pipelinesChecksClient := azdosdkmocks.NewPipelinesChecksClientV7(ctrl)
	clients := &client.AggregatedClient{V7PipelinesChecksClientExtras: pipelinesChecksClient, Ctx: context.Background()}

	expectedArgs := pipelineschecksv7.AddGenericCheckConfigurationArgs{Configuration: &ApprovalCheckTest, Project: &ApprovalCheckProjectID}
	pipelinesChecksClient.
		EXPECT().
		AddGenericCheckConfiguration(clients.Ctx, expectedArgs).
		Return(nil, errors.New("AddCheckConfiguration() Failed")).
		Times(1)

	err := r.Create(resourceData, clients)
	require.Contains(t, err.Error(), "AddCheckConfiguration() Failed")
}

// verifies that if an error is produced on a read, it is not swallowed
func TestCheckApproval_Read_DoesNotSwallowError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := ResourceCheckApproval()
	resourceData := schema.TestResourceDataRaw(t, r.Schema, nil)
	flattenCheckApproval(resourceData, &ApprovalCheckTest, ApprovalCheckProjectID)

	pipelinesChecksClient := azdosdkmocks.NewPipelinesChecksClientV7(ctrl)
	clients := &client.AggregatedClient{V7PipelinesChecksClientExtras: pipelinesChecksClient, Ctx: context.Background()}

	expectedArgs := pipelineschecksv7.GetGenericCheckConfigurationArgs{
		Id:      ApprovalCheckTest.Id,
		Project: &ApprovalCheckProjectID,
		Expand:  &pipelineschecks.CheckConfigurationExpandParameterValues.Settings,
	}

	pipelinesChecksClient.
		EXPECT().
		GetGenericCheckConfiguration(clients.Ctx, expectedArgs).
		Return(nil, errors.New("GetServiceEndpoint() Failed")).
		Times(1)

	err := r.Read(resourceData, clients)
	require.Contains(t, err.Error(), "GetServiceEndpoint() Failed")
}

// verifies that if an error is produced on a delete, it is not swallowed
func TestCheckApproval_Delete_DoesNotSwallowError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := ResourceCheckApproval()
	resourceData := schema.TestResourceDataRaw(t, r.Schema, nil)
	flattenCheckApproval(resourceData, &ApprovalCheckTest, ApprovalCheckProjectID)

	pipelinesChecksClient := azdosdkmocks.NewPipelinesChecksClientV7(ctrl)
	clients := &client.AggregatedClient{V7PipelinesChecksClientExtras: pipelinesChecksClient, Ctx: context.Background()}

	expectedArgs := pipelineschecksv7.DeleteGenericCheckConfigurationArgs{
		Id:      ApprovalCheckTest.Id,
		Project: &ApprovalCheckProjectID,
	}

	pipelinesChecksClient.
		EXPECT().
		DeleteGenericCheckConfiguration(clients.Ctx, expectedArgs).
		Return(errors.New("DeleteServiceEndpoint() Failed")).
		Times(1)

	err := r.Delete(resourceData, clients)
	require.Contains(t, err.Error(), "DeleteServiceEndpoint() Failed")
}

// verifies that if an error is produced on an update, it is not swallowed
func TestCheckApproval_Update_DoesNotSwallowError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := ResourceCheckApproval()
	resourceData := schema.TestResourceDataRaw(t, r.Schema, nil)
	flattenCheckApproval(resourceData, &ApprovalCheckTest, ApprovalCheckProjectID)

	pipelinesChecksClient := azdosdkmocks.NewPipelinesChecksClientV7(ctrl)
	clients := &client.AggregatedClient{V7PipelinesChecksClientExtras: pipelinesChecksClient, Ctx: context.Background()}

	expectedArgs := pipelineschecksv7.UpdateGenericCheckConfigurationArgs{
		Project:       &ApprovalCheckProjectID,
		Configuration: &ApprovalCheckTest,
		Id:            &ApprovalCheckID,
	}

	pipelinesChecksClient.
		EXPECT().
		UpdateGenericCheckConfiguration(clients.Ctx, expectedArgs).
		Return(nil, errors.New("UpdateServiceEndpoint() Failed")).
		Times(1)

	err := r.Update(resourceData, clients)
	require.Contains(t, err.Error(), "UpdateServiceEndpoint() Failed")
}

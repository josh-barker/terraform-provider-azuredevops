// This file creates functions that accept the GenericCheckConfiguration, which contains Settings
// It's based on github.com/microsoft/azure-devops-go-api/azuredevops/v7/pipelineschecks/client.go
// which does not have the Settings element.

package pipelineschecksv7

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/pipelineschecks"
)

var ResourceAreaId, _ = uuid.Parse("4a933897-0488-45af-bd82-6fd3ad33f46a")

type Client interface {
	// [Preview API] Add a check configuration
	AddGenericCheckConfiguration(context.Context, AddGenericCheckConfigurationArgs) (*pipelineschecks.GenericCheckConfiguration, error)
	// [Preview API] Delete check configuration by id
	DeleteGenericCheckConfiguration(context.Context, DeleteGenericCheckConfigurationArgs) error
	// [Preview API] Get Check configuration by Id
	GetGenericCheckConfiguration(context.Context, GetGenericCheckConfigurationArgs) (*pipelineschecks.GenericCheckConfiguration, error)
	// [Preview API] Get Check configuration by resource type and id
	GetGenericCheckConfigurationsOnResource(context.Context, GetGenericCheckConfigurationsOnResourceArgs) (*[]pipelineschecks.GenericCheckConfiguration, error)
	// [Preview API] Get details for a specific check evaluation
	UpdateGenericCheckConfiguration(context.Context, UpdateGenericCheckConfigurationArgs) (*pipelineschecks.GenericCheckConfiguration, error)
}

type ClientImpl struct {
	Client azuredevops.Client
}

func NewClient(ctx context.Context, connection *azuredevops.Connection) (Client, error) {
	client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
	if err != nil {
		return nil, err
	}
	return &ClientImpl{
		Client: *client,
	}, nil
}

// [Preview API] Add a check configuration
func (client *ClientImpl) AddGenericCheckConfiguration(ctx context.Context, args AddGenericCheckConfigurationArgs) (*pipelineschecks.GenericCheckConfiguration, error) {
	if args.Configuration == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Configuration"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	body, marshalErr := json.Marshal(*args.Configuration)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("86c8381e-5aee-4cde-8ae4-25c0c7f5eaea")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue pipelineschecks.GenericCheckConfiguration
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the AddGenericCheckConfiguration function
type AddGenericCheckConfigurationArgs struct {
	// (required)
	Configuration *pipelineschecks.GenericCheckConfiguration
	// (required) Project ID or project name
	Project *string
}

// [Preview API] Delete check configuration by id
func (client *ClientImpl) DeleteGenericCheckConfiguration(ctx context.Context, args DeleteGenericCheckConfigurationArgs) error {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.Id == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.Id"}
	}
	routeValues["id"] = strconv.Itoa(*args.Id)

	locationId, _ := uuid.Parse("86c8381e-5aee-4cde-8ae4-25c0c7f5eaea")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "7.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteGenericCheckConfiguration function
type DeleteGenericCheckConfigurationArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) check configuration id
	Id *int
}

// [Preview API] Get Check configuration by Id
func (client *ClientImpl) GetGenericCheckConfiguration(ctx context.Context, args GetGenericCheckConfigurationArgs) (*pipelineschecks.GenericCheckConfiguration, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.Id == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Id"}
	}
	routeValues["id"] = strconv.Itoa(*args.Id)

	queryParams := url.Values{}
	if args.Expand != nil {
		queryParams.Add("$expand", string(*args.Expand))
	}
	locationId, _ := uuid.Parse("86c8381e-5aee-4cde-8ae4-25c0c7f5eaea")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue pipelineschecks.GenericCheckConfiguration
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetGenericCheckConfiguration function
type GetGenericCheckConfigurationArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required)
	Id *int
	// (optional)
	Expand *pipelineschecks.CheckConfigurationExpandParameter
}

// [Preview API] Get Check configuration by resource type and id
func (client *ClientImpl) GetGenericCheckConfigurationsOnResource(ctx context.Context, args GetGenericCheckConfigurationsOnResourceArgs) (*[]pipelineschecks.GenericCheckConfiguration, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.ResourceType != nil {
		queryParams.Add("resourceType", *args.ResourceType)
	}
	if args.ResourceId != nil {
		queryParams.Add("resourceId", *args.ResourceId)
	}
	if args.Expand != nil {
		queryParams.Add("$expand", string(*args.Expand))
	}
	locationId, _ := uuid.Parse("86c8381e-5aee-4cde-8ae4-25c0c7f5eaea")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []pipelineschecks.GenericCheckConfiguration
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetGenericCheckConfigurationsOnResource function
type GetGenericCheckConfigurationsOnResourceArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) resource type
	ResourceType *string
	// (optional) resource id
	ResourceId *string
	// (optional)
	Expand *pipelineschecks.CheckConfigurationExpandParameter
}

// [Preview API] Update check configuration
func (client *ClientImpl) UpdateGenericCheckConfiguration(ctx context.Context, args UpdateGenericCheckConfigurationArgs) (*pipelineschecks.GenericCheckConfiguration, error) {
	if args.Configuration == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Configuration"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.Id == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Id"}
	}
	routeValues["id"] = strconv.Itoa(*args.Id)

	body, marshalErr := json.Marshal(*args.Configuration)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("86c8381e-5aee-4cde-8ae4-25c0c7f5eaea")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue pipelineschecks.GenericCheckConfiguration
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateGenericCheckConfiguration function
type UpdateGenericCheckConfigurationArgs struct {
	// (required) check configuration
	Configuration *pipelineschecks.GenericCheckConfiguration
	// (required) Project ID or project name
	Project *string
	// (required) check configuration id
	Id *int
}

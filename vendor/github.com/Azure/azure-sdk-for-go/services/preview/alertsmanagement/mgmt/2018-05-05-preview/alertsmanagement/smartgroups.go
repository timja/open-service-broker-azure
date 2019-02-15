package alertsmanagement

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// SmartGroupsClient is the alertsManagement Client
type SmartGroupsClient struct {
	BaseClient
}

// NewSmartGroupsClient creates an instance of the SmartGroupsClient client.
func NewSmartGroupsClient(subscriptionID string, monitorService1 MonitorService) SmartGroupsClient {
	return NewSmartGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID, monitorService1)
}

// NewSmartGroupsClientWithBaseURI creates an instance of the SmartGroupsClient client.
func NewSmartGroupsClientWithBaseURI(baseURI string, subscriptionID string, monitorService1 MonitorService) SmartGroupsClient {
	return SmartGroupsClient{NewWithBaseURI(baseURI, subscriptionID, monitorService1)}
}

// ChangeState change the state from unresolved to resolved and all the alerts within the smart group will also be
// resolved.
// Parameters:
// smartGroupID - smart Group Id
// newState - filter by state
func (client SmartGroupsClient) ChangeState(ctx context.Context, smartGroupID string, newState AlertState) (result SmartGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SmartGroupsClient.ChangeState")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ChangeStatePreparer(ctx, smartGroupID, newState)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.SmartGroupsClient", "ChangeState", nil, "Failure preparing request")
		return
	}

	resp, err := client.ChangeStateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "alertsmanagement.SmartGroupsClient", "ChangeState", resp, "Failure sending request")
		return
	}

	result, err = client.ChangeStateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.SmartGroupsClient", "ChangeState", resp, "Failure responding to request")
	}

	return
}

// ChangeStatePreparer prepares the ChangeState request.
func (client SmartGroupsClient) ChangeStatePreparer(ctx context.Context, smartGroupID string, newState AlertState) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"smartGroupId":   autorest.Encode("path", smartGroupID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-05-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"newState":    autorest.Encode("query", newState),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/smartGroups/{smartGroupId}/changeState", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ChangeStateSender sends the ChangeState request. The method will close the
// http.Response Body if it receives an error.
func (client SmartGroupsClient) ChangeStateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ChangeStateResponder handles the response to the ChangeState request. The method always
// closes the http.Response Body.
func (client SmartGroupsClient) ChangeStateResponder(resp *http.Response) (result SmartGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAll list all the smartGroups within the specified subscription.
// Parameters:
// targetResource - filter by target resource
// targetResourceGroup - filter by target resource group name
// targetResourceType - filter by target resource type
// monitorCondition - filter by monitor condition which is the state of the alert at monitor service
// severity - filter by severity
// smartGroupState - filter by state
// timeRange - filter by time range, default value is 1 day
// pageCount - number of items per page, default value is '25'.
// sortBy - sort the query results by input field, default value is 'lastModifiedDateTime'.
// sortOrder - sort the query results order in either ascending or descending, default value is 'desc' for time
// fields and 'asc' for others.
func (client SmartGroupsClient) GetAll(ctx context.Context, targetResource string, targetResourceGroup string, targetResourceType string, monitorCondition MonitorCondition, severity Severity, smartGroupState AlertState, timeRange TimeRange, pageCount *int32, sortBy SmartGroupsSortByFields, sortOrder string) (result SmartGroupsList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SmartGroupsClient.GetAll")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetAllPreparer(ctx, targetResource, targetResourceGroup, targetResourceType, monitorCondition, severity, smartGroupState, timeRange, pageCount, sortBy, sortOrder)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.SmartGroupsClient", "GetAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "alertsmanagement.SmartGroupsClient", "GetAll", resp, "Failure sending request")
		return
	}

	result, err = client.GetAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.SmartGroupsClient", "GetAll", resp, "Failure responding to request")
	}

	return
}

// GetAllPreparer prepares the GetAll request.
func (client SmartGroupsClient) GetAllPreparer(ctx context.Context, targetResource string, targetResourceGroup string, targetResourceType string, monitorCondition MonitorCondition, severity Severity, smartGroupState AlertState, timeRange TimeRange, pageCount *int32, sortBy SmartGroupsSortByFields, sortOrder string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-05-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(targetResource) > 0 {
		queryParameters["targetResource"] = autorest.Encode("query", targetResource)
	}
	if len(targetResourceGroup) > 0 {
		queryParameters["targetResourceGroup"] = autorest.Encode("query", targetResourceGroup)
	}
	if len(targetResourceType) > 0 {
		queryParameters["targetResourceType"] = autorest.Encode("query", targetResourceType)
	}
	if len(string(client.MonitorService1)) > 0 {
		queryParameters["monitorService"] = autorest.Encode("query", client.MonitorService1)
	}
	if len(string(monitorCondition)) > 0 {
		queryParameters["monitorCondition"] = autorest.Encode("query", monitorCondition)
	}
	if len(string(severity)) > 0 {
		queryParameters["severity"] = autorest.Encode("query", severity)
	}
	if len(string(smartGroupState)) > 0 {
		queryParameters["smartGroupState"] = autorest.Encode("query", smartGroupState)
	}
	if len(string(timeRange)) > 0 {
		queryParameters["timeRange"] = autorest.Encode("query", timeRange)
	}
	if pageCount != nil {
		queryParameters["pageCount"] = autorest.Encode("query", *pageCount)
	}
	if len(string(sortBy)) > 0 {
		queryParameters["sortBy"] = autorest.Encode("query", sortBy)
	}
	if len(string(sortOrder)) > 0 {
		queryParameters["sortOrder"] = autorest.Encode("query", sortOrder)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/smartGroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAllSender sends the GetAll request. The method will close the
// http.Response Body if it receives an error.
func (client SmartGroupsClient) GetAllSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetAllResponder handles the response to the GetAll request. The method always
// closes the http.Response Body.
func (client SmartGroupsClient) GetAllResponder(resp *http.Response) (result SmartGroupsList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByID get details of smart group.
// Parameters:
// smartGroupID - smart Group Id
func (client SmartGroupsClient) GetByID(ctx context.Context, smartGroupID string) (result SmartGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SmartGroupsClient.GetByID")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByIDPreparer(ctx, smartGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.SmartGroupsClient", "GetByID", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "alertsmanagement.SmartGroupsClient", "GetByID", resp, "Failure sending request")
		return
	}

	result, err = client.GetByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.SmartGroupsClient", "GetByID", resp, "Failure responding to request")
	}

	return
}

// GetByIDPreparer prepares the GetByID request.
func (client SmartGroupsClient) GetByIDPreparer(ctx context.Context, smartGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"smartGroupId":   autorest.Encode("path", smartGroupID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-05-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/smartGroups/{smartGroupId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByIDSender sends the GetByID request. The method will close the
// http.Response Body if it receives an error.
func (client SmartGroupsClient) GetByIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetByIDResponder handles the response to the GetByID request. The method always
// closes the http.Response Body.
func (client SmartGroupsClient) GetByIDResponder(resp *http.Response) (result SmartGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetHistory get the history of the changes of smart group.
// Parameters:
// smartGroupID - smart Group Id
func (client SmartGroupsClient) GetHistory(ctx context.Context, smartGroupID string) (result SmartGroupModification, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SmartGroupsClient.GetHistory")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetHistoryPreparer(ctx, smartGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.SmartGroupsClient", "GetHistory", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetHistorySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "alertsmanagement.SmartGroupsClient", "GetHistory", resp, "Failure sending request")
		return
	}

	result, err = client.GetHistoryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.SmartGroupsClient", "GetHistory", resp, "Failure responding to request")
	}

	return
}

// GetHistoryPreparer prepares the GetHistory request.
func (client SmartGroupsClient) GetHistoryPreparer(ctx context.Context, smartGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"smartGroupId":   autorest.Encode("path", smartGroupID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-05-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/smartGroups/{smartGroupId}/history", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetHistorySender sends the GetHistory request. The method will close the
// http.Response Body if it receives an error.
func (client SmartGroupsClient) GetHistorySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetHistoryResponder handles the response to the GetHistory request. The method always
// closes the http.Response Body.
func (client SmartGroupsClient) GetHistoryResponder(resp *http.Response) (result SmartGroupModification, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

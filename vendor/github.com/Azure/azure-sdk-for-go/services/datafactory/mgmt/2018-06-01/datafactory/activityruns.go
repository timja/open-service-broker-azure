package datafactory

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ActivityRunsClient is the the Azure Data Factory V2 management API provides a RESTful set of web services that
// interact with Azure Data Factory V2 services.
type ActivityRunsClient struct {
	BaseClient
}

// NewActivityRunsClient creates an instance of the ActivityRunsClient client.
func NewActivityRunsClient(subscriptionID string) ActivityRunsClient {
	return NewActivityRunsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewActivityRunsClientWithBaseURI creates an instance of the ActivityRunsClient client.
func NewActivityRunsClientWithBaseURI(baseURI string, subscriptionID string) ActivityRunsClient {
	return ActivityRunsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// QueryByPipelineRun query activity runs based on input filter conditions.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// runID - the pipeline run identifier.
// filterParameters - parameters to filter the activity runs.
func (client ActivityRunsClient) QueryByPipelineRun(ctx context.Context, resourceGroupName string, factoryName string, runID string, filterParameters RunFilterParameters) (result ActivityRunsQueryResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActivityRunsClient.QueryByPipelineRun")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: filterParameters,
			Constraints: []validation.Constraint{{Target: "filterParameters.LastUpdatedAfter", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "filterParameters.LastUpdatedBefore", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.ActivityRunsClient", "QueryByPipelineRun", err.Error())
	}

	req, err := client.QueryByPipelineRunPreparer(ctx, resourceGroupName, factoryName, runID, filterParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.ActivityRunsClient", "QueryByPipelineRun", nil, "Failure preparing request")
		return
	}

	resp, err := client.QueryByPipelineRunSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.ActivityRunsClient", "QueryByPipelineRun", resp, "Failure sending request")
		return
	}

	result, err = client.QueryByPipelineRunResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.ActivityRunsClient", "QueryByPipelineRun", resp, "Failure responding to request")
	}

	return
}

// QueryByPipelineRunPreparer prepares the QueryByPipelineRun request.
func (client ActivityRunsClient) QueryByPipelineRunPreparer(ctx context.Context, resourceGroupName string, factoryName string, runID string, filterParameters RunFilterParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"runId":             autorest.Encode("path", runID),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelineruns/{runId}/queryActivityruns", pathParameters),
		autorest.WithJSON(filterParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// QueryByPipelineRunSender sends the QueryByPipelineRun request. The method will close the
// http.Response Body if it receives an error.
func (client ActivityRunsClient) QueryByPipelineRunSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// QueryByPipelineRunResponder handles the response to the QueryByPipelineRun request. The method always
// closes the http.Response Body.
func (client ActivityRunsClient) QueryByPipelineRunResponder(resp *http.Response) (result ActivityRunsQueryResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

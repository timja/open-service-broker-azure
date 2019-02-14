package schedulerapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler"
	"github.com/Azure/go-autorest/autorest"
)

// JobCollectionsClientAPI contains the set of methods on the JobCollectionsClient type.
type JobCollectionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, jobCollectionName string, jobCollection scheduler.JobCollectionDefinition) (result scheduler.JobCollectionDefinition, err error)
	Delete(ctx context.Context, resourceGroupName string, jobCollectionName string) (result autorest.Response, err error)
	Disable(ctx context.Context, resourceGroupName string, jobCollectionName string) (result autorest.Response, err error)
	Enable(ctx context.Context, resourceGroupName string, jobCollectionName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, jobCollectionName string) (result scheduler.JobCollectionDefinition, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result scheduler.JobCollectionListResultPage, err error)
	ListBySubscription(ctx context.Context) (result scheduler.JobCollectionListResultPage, err error)
	Patch(ctx context.Context, resourceGroupName string, jobCollectionName string, jobCollection scheduler.JobCollectionDefinition) (result scheduler.JobCollectionDefinition, err error)
}

var _ JobCollectionsClientAPI = (*scheduler.JobCollectionsClient)(nil)

// JobsClientAPI contains the set of methods on the JobsClient type.
type JobsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, job scheduler.JobDefinition) (result scheduler.JobDefinition, err error)
	Delete(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string) (result scheduler.JobDefinition, err error)
	List(ctx context.Context, resourceGroupName string, jobCollectionName string, top *int32, skip *int32, filter string) (result scheduler.JobListResultPage, err error)
	ListJobHistory(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, top *int32, skip *int32, filter string) (result scheduler.JobHistoryListResultPage, err error)
	Patch(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, job scheduler.JobDefinition) (result scheduler.JobDefinition, err error)
	Run(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string) (result autorest.Response, err error)
}

var _ JobsClientAPI = (*scheduler.JobsClient)(nil)

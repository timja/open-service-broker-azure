package skus

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
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2017-09-01/skus"

// ResourceSkuCapacityScaleType enumerates the values for resource sku capacity scale type.
type ResourceSkuCapacityScaleType string

const (
	// Automatic ...
	Automatic ResourceSkuCapacityScaleType = "Automatic"
	// Manual ...
	Manual ResourceSkuCapacityScaleType = "Manual"
	// None ...
	None ResourceSkuCapacityScaleType = "None"
)

// PossibleResourceSkuCapacityScaleTypeValues returns an array of possible values for the ResourceSkuCapacityScaleType const type.
func PossibleResourceSkuCapacityScaleTypeValues() []ResourceSkuCapacityScaleType {
	return []ResourceSkuCapacityScaleType{Automatic, Manual, None}
}

// ResourceSkuRestrictionsReasonCode enumerates the values for resource sku restrictions reason code.
type ResourceSkuRestrictionsReasonCode string

const (
	// NotAvailableForSubscription ...
	NotAvailableForSubscription ResourceSkuRestrictionsReasonCode = "NotAvailableForSubscription"
	// QuotaID ...
	QuotaID ResourceSkuRestrictionsReasonCode = "QuotaId"
)

// PossibleResourceSkuRestrictionsReasonCodeValues returns an array of possible values for the ResourceSkuRestrictionsReasonCode const type.
func PossibleResourceSkuRestrictionsReasonCodeValues() []ResourceSkuRestrictionsReasonCode {
	return []ResourceSkuRestrictionsReasonCode{NotAvailableForSubscription, QuotaID}
}

// ResourceSkuRestrictionsType enumerates the values for resource sku restrictions type.
type ResourceSkuRestrictionsType string

const (
	// Location ...
	Location ResourceSkuRestrictionsType = "Location"
	// Zone ...
	Zone ResourceSkuRestrictionsType = "Zone"
)

// PossibleResourceSkuRestrictionsTypeValues returns an array of possible values for the ResourceSkuRestrictionsType const type.
func PossibleResourceSkuRestrictionsTypeValues() []ResourceSkuRestrictionsType {
	return []ResourceSkuRestrictionsType{Location, Zone}
}

// ResourceSku describes an available Compute SKU.
type ResourceSku struct {
	// ResourceType - The type of resource the SKU applies to.
	ResourceType *string `json:"resourceType,omitempty"`
	// Name - The name of SKU.
	Name *string `json:"name,omitempty"`
	// Tier - Specifies the tier of virtual machines in a scale set.<br /><br /> Possible Values:<br /><br /> **Standard**<br /><br /> **Basic**
	Tier *string `json:"tier,omitempty"`
	// Size - The Size of the SKU.
	Size *string `json:"size,omitempty"`
	// Family - The Family of this particular SKU.
	Family *string `json:"family,omitempty"`
	// Kind - The Kind of resources that are supported in this SKU.
	Kind *string `json:"kind,omitempty"`
	// Capacity - Specifies the number of virtual machines in the scale set.
	Capacity *ResourceSkuCapacity `json:"capacity,omitempty"`
	// Locations - The set of locations that the SKU is available.
	Locations *[]string `json:"locations,omitempty"`
	// LocationInfo - A list of locations and availability zones in those locations where the SKU is available.
	LocationInfo *[]ResourceSkuLocationInfo `json:"locationInfo,omitempty"`
	// APIVersions - The api versions that support this SKU.
	APIVersions *[]string `json:"apiVersions,omitempty"`
	// Costs - Metadata for retrieving price info.
	Costs *[]ResourceSkuCosts `json:"costs,omitempty"`
	// Capabilities - A name value pair to describe the capability.
	Capabilities *[]ResourceSkuCapabilities `json:"capabilities,omitempty"`
	// Restrictions - The restrictions because of which SKU cannot be used. This is empty if there are no restrictions.
	Restrictions *[]ResourceSkuRestrictions `json:"restrictions,omitempty"`
}

// ResourceSkuCapabilities describes The SKU capabilities object.
type ResourceSkuCapabilities struct {
	// Name - An invariant to describe the feature.
	Name *string `json:"name,omitempty"`
	// Value - An invariant if the feature is measured by quantity.
	Value *string `json:"value,omitempty"`
}

// ResourceSkuCapacity describes scaling information of a SKU.
type ResourceSkuCapacity struct {
	// Minimum - The minimum capacity.
	Minimum *int64 `json:"minimum,omitempty"`
	// Maximum - The maximum capacity that can be set.
	Maximum *int64 `json:"maximum,omitempty"`
	// Default - The default capacity.
	Default *int64 `json:"default,omitempty"`
	// ScaleType - The scale type applicable to the sku. Possible values include: 'Automatic', 'Manual', 'None'
	ScaleType ResourceSkuCapacityScaleType `json:"scaleType,omitempty"`
}

// ResourceSkuCosts describes metadata for retrieving price info.
type ResourceSkuCosts struct {
	// MeterID - Used for querying price from commerce.
	MeterID *string `json:"meterID,omitempty"`
	// Quantity - The multiplier is needed to extend the base metered cost.
	Quantity *int64 `json:"quantity,omitempty"`
	// ExtendedUnit - An invariant to show the extended unit.
	ExtendedUnit *string `json:"extendedUnit,omitempty"`
}

// ResourceSkuLocationInfo ...
type ResourceSkuLocationInfo struct {
	// Location - Location of the SKU
	Location *string `json:"location,omitempty"`
	// Zones - List of availability zones where the SKU is supported.
	Zones *[]string `json:"zones,omitempty"`
}

// ResourceSkuRestrictionInfo ...
type ResourceSkuRestrictionInfo struct {
	// Locations - Locations where the SKU is restricted
	Locations *[]string `json:"locations,omitempty"`
	// Zones - List of availability zones where the SKU is restricted.
	Zones *[]string `json:"zones,omitempty"`
}

// ResourceSkuRestrictions describes scaling information of a SKU.
type ResourceSkuRestrictions struct {
	// Type - The type of restrictions. Possible values include: 'Location', 'Zone'
	Type ResourceSkuRestrictionsType `json:"type,omitempty"`
	// Values - The value of restrictions. If the restriction type is set to location. This would be different locations where the SKU is restricted.
	Values *[]string `json:"values,omitempty"`
	// RestrictionInfo - The information about the restriction where the SKU cannot be used.
	RestrictionInfo *ResourceSkuRestrictionInfo `json:"restrictionInfo,omitempty"`
	// ReasonCode - The reason for restriction. Possible values include: 'QuotaID', 'NotAvailableForSubscription'
	ReasonCode ResourceSkuRestrictionsReasonCode `json:"reasonCode,omitempty"`
}

// ResourceSkusResult the Compute List Skus operation response.
type ResourceSkusResult struct {
	autorest.Response `json:"-"`
	// Value - The list of skus available for the subscription.
	Value *[]ResourceSku `json:"value,omitempty"`
	// NextLink - The uri to fetch the next page of Compute Skus. Call ListNext() with this to fetch the next page of VMSS Skus.
	NextLink *string `json:"nextLink,omitempty"`
}

// ResourceSkusResultIterator provides access to a complete listing of ResourceSku values.
type ResourceSkusResultIterator struct {
	i    int
	page ResourceSkusResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ResourceSkusResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceSkusResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *ResourceSkusResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ResourceSkusResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ResourceSkusResultIterator) Response() ResourceSkusResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ResourceSkusResultIterator) Value() ResourceSku {
	if !iter.page.NotDone() {
		return ResourceSku{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ResourceSkusResultIterator type.
func NewResourceSkusResultIterator(page ResourceSkusResultPage) ResourceSkusResultIterator {
	return ResourceSkusResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (rsr ResourceSkusResult) IsEmpty() bool {
	return rsr.Value == nil || len(*rsr.Value) == 0
}

// resourceSkusResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (rsr ResourceSkusResult) resourceSkusResultPreparer(ctx context.Context) (*http.Request, error) {
	if rsr.NextLink == nil || len(to.String(rsr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(rsr.NextLink)))
}

// ResourceSkusResultPage contains a page of ResourceSku values.
type ResourceSkusResultPage struct {
	fn  func(context.Context, ResourceSkusResult) (ResourceSkusResult, error)
	rsr ResourceSkusResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ResourceSkusResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceSkusResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.rsr)
	if err != nil {
		return err
	}
	page.rsr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ResourceSkusResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ResourceSkusResultPage) NotDone() bool {
	return !page.rsr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ResourceSkusResultPage) Response() ResourceSkusResult {
	return page.rsr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ResourceSkusResultPage) Values() []ResourceSku {
	if page.rsr.IsEmpty() {
		return nil
	}
	return *page.rsr.Value
}

// Creates a new instance of the ResourceSkusResultPage type.
func NewResourceSkusResultPage(getNextPage func(context.Context, ResourceSkusResult) (ResourceSkusResult, error)) ResourceSkusResultPage {
	return ResourceSkusResultPage{fn: getNextPage}
}

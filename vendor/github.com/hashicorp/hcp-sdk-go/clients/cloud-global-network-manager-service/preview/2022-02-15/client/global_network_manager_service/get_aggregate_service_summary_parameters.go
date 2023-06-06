// Code generated by go-swagger; DO NOT EDIT.

package global_network_manager_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetAggregateServiceSummaryParams creates a new GetAggregateServiceSummaryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAggregateServiceSummaryParams() *GetAggregateServiceSummaryParams {
	return &GetAggregateServiceSummaryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAggregateServiceSummaryParamsWithTimeout creates a new GetAggregateServiceSummaryParams object
// with the ability to set a timeout on a request.
func NewGetAggregateServiceSummaryParamsWithTimeout(timeout time.Duration) *GetAggregateServiceSummaryParams {
	return &GetAggregateServiceSummaryParams{
		timeout: timeout,
	}
}

// NewGetAggregateServiceSummaryParamsWithContext creates a new GetAggregateServiceSummaryParams object
// with the ability to set a context for a request.
func NewGetAggregateServiceSummaryParamsWithContext(ctx context.Context) *GetAggregateServiceSummaryParams {
	return &GetAggregateServiceSummaryParams{
		Context: ctx,
	}
}

// NewGetAggregateServiceSummaryParamsWithHTTPClient creates a new GetAggregateServiceSummaryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAggregateServiceSummaryParamsWithHTTPClient(client *http.Client) *GetAggregateServiceSummaryParams {
	return &GetAggregateServiceSummaryParams{
		HTTPClient: client,
	}
}

/*
GetAggregateServiceSummaryParams contains all the parameters to send to the API endpoint

	for the get aggregate service summary operation.

	Typically these are written to a http.Request.
*/
type GetAggregateServiceSummaryParams struct {

	/* FilterClusters.

	   filters summaries to the provided cluster names. If empty, all are included
	*/
	FilterClusters []string

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID string

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	/* LocationRegionProvider.

	   provider is the named cloud provider ("aws", "gcp", "azure")
	*/
	LocationRegionProvider *string

	/* LocationRegionRegion.

	   region is the cloud region ("us-west1", "us-east1")
	*/
	LocationRegionRegion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get aggregate service summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAggregateServiceSummaryParams) WithDefaults() *GetAggregateServiceSummaryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get aggregate service summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAggregateServiceSummaryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get aggregate service summary params
func (o *GetAggregateServiceSummaryParams) WithTimeout(timeout time.Duration) *GetAggregateServiceSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get aggregate service summary params
func (o *GetAggregateServiceSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get aggregate service summary params
func (o *GetAggregateServiceSummaryParams) WithContext(ctx context.Context) *GetAggregateServiceSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get aggregate service summary params
func (o *GetAggregateServiceSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get aggregate service summary params
func (o *GetAggregateServiceSummaryParams) WithHTTPClient(client *http.Client) *GetAggregateServiceSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get aggregate service summary params
func (o *GetAggregateServiceSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterClusters adds the filterClusters to the get aggregate service summary params
func (o *GetAggregateServiceSummaryParams) WithFilterClusters(filterClusters []string) *GetAggregateServiceSummaryParams {
	o.SetFilterClusters(filterClusters)
	return o
}

// SetFilterClusters adds the filterClusters to the get aggregate service summary params
func (o *GetAggregateServiceSummaryParams) SetFilterClusters(filterClusters []string) {
	o.FilterClusters = filterClusters
}

// WithLocationOrganizationID adds the locationOrganizationID to the get aggregate service summary params
func (o *GetAggregateServiceSummaryParams) WithLocationOrganizationID(locationOrganizationID string) *GetAggregateServiceSummaryParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the get aggregate service summary params
func (o *GetAggregateServiceSummaryParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the get aggregate service summary params
func (o *GetAggregateServiceSummaryParams) WithLocationProjectID(locationProjectID string) *GetAggregateServiceSummaryParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the get aggregate service summary params
func (o *GetAggregateServiceSummaryParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the get aggregate service summary params
func (o *GetAggregateServiceSummaryParams) WithLocationRegionProvider(locationRegionProvider *string) *GetAggregateServiceSummaryParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the get aggregate service summary params
func (o *GetAggregateServiceSummaryParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the get aggregate service summary params
func (o *GetAggregateServiceSummaryParams) WithLocationRegionRegion(locationRegionRegion *string) *GetAggregateServiceSummaryParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the get aggregate service summary params
func (o *GetAggregateServiceSummaryParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *GetAggregateServiceSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterClusters != nil {

		// binding items for filter.clusters
		joinedFilterClusters := o.bindParamFilterClusters(reg)

		// query array param filter.clusters
		if err := r.SetQueryParam("filter.clusters", joinedFilterClusters...); err != nil {
			return err
		}
	}

	// path param location.organization_id
	if err := r.SetPathParam("location.organization_id", o.LocationOrganizationID); err != nil {
		return err
	}

	// path param location.project_id
	if err := r.SetPathParam("location.project_id", o.LocationProjectID); err != nil {
		return err
	}

	if o.LocationRegionProvider != nil {

		// query param location.region.provider
		var qrLocationRegionProvider string

		if o.LocationRegionProvider != nil {
			qrLocationRegionProvider = *o.LocationRegionProvider
		}
		qLocationRegionProvider := qrLocationRegionProvider
		if qLocationRegionProvider != "" {

			if err := r.SetQueryParam("location.region.provider", qLocationRegionProvider); err != nil {
				return err
			}
		}
	}

	if o.LocationRegionRegion != nil {

		// query param location.region.region
		var qrLocationRegionRegion string

		if o.LocationRegionRegion != nil {
			qrLocationRegionRegion = *o.LocationRegionRegion
		}
		qLocationRegionRegion := qrLocationRegionRegion
		if qLocationRegionRegion != "" {

			if err := r.SetQueryParam("location.region.region", qLocationRegionRegion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetAggregateServiceSummary binds the parameter filter.clusters
func (o *GetAggregateServiceSummaryParams) bindParamFilterClusters(formats strfmt.Registry) []string {
	filterClustersIR := o.FilterClusters

	var filterClustersIC []string
	for _, filterClustersIIR := range filterClustersIR { // explode []string

		filterClustersIIV := filterClustersIIR // string as string
		filterClustersIC = append(filterClustersIC, filterClustersIIV)
	}

	// items.CollectionFormat: "multi"
	filterClustersIS := swag.JoinByFormat(filterClustersIC, "multi")

	return filterClustersIS
}

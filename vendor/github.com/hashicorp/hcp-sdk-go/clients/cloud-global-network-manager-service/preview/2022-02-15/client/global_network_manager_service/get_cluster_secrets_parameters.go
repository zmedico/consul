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
)

// NewGetClusterSecretsParams creates a new GetClusterSecretsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterSecretsParams() *GetClusterSecretsParams {
	return &GetClusterSecretsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterSecretsParamsWithTimeout creates a new GetClusterSecretsParams object
// with the ability to set a timeout on a request.
func NewGetClusterSecretsParamsWithTimeout(timeout time.Duration) *GetClusterSecretsParams {
	return &GetClusterSecretsParams{
		timeout: timeout,
	}
}

// NewGetClusterSecretsParamsWithContext creates a new GetClusterSecretsParams object
// with the ability to set a context for a request.
func NewGetClusterSecretsParamsWithContext(ctx context.Context) *GetClusterSecretsParams {
	return &GetClusterSecretsParams{
		Context: ctx,
	}
}

// NewGetClusterSecretsParamsWithHTTPClient creates a new GetClusterSecretsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterSecretsParamsWithHTTPClient(client *http.Client) *GetClusterSecretsParams {
	return &GetClusterSecretsParams{
		HTTPClient: client,
	}
}

/*
GetClusterSecretsParams contains all the parameters to send to the API endpoint

	for the get cluster secrets operation.

	Typically these are written to a http.Request.
*/
type GetClusterSecretsParams struct {

	/* ID.

	   id is the user-provided GNM cluster name
	*/
	ID string

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

// WithDefaults hydrates default values in the get cluster secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterSecretsParams) WithDefaults() *GetClusterSecretsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterSecretsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster secrets params
func (o *GetClusterSecretsParams) WithTimeout(timeout time.Duration) *GetClusterSecretsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster secrets params
func (o *GetClusterSecretsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster secrets params
func (o *GetClusterSecretsParams) WithContext(ctx context.Context) *GetClusterSecretsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster secrets params
func (o *GetClusterSecretsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster secrets params
func (o *GetClusterSecretsParams) WithHTTPClient(client *http.Client) *GetClusterSecretsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster secrets params
func (o *GetClusterSecretsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get cluster secrets params
func (o *GetClusterSecretsParams) WithID(id string) *GetClusterSecretsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get cluster secrets params
func (o *GetClusterSecretsParams) SetID(id string) {
	o.ID = id
}

// WithLocationOrganizationID adds the locationOrganizationID to the get cluster secrets params
func (o *GetClusterSecretsParams) WithLocationOrganizationID(locationOrganizationID string) *GetClusterSecretsParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the get cluster secrets params
func (o *GetClusterSecretsParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the get cluster secrets params
func (o *GetClusterSecretsParams) WithLocationProjectID(locationProjectID string) *GetClusterSecretsParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the get cluster secrets params
func (o *GetClusterSecretsParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the get cluster secrets params
func (o *GetClusterSecretsParams) WithLocationRegionProvider(locationRegionProvider *string) *GetClusterSecretsParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the get cluster secrets params
func (o *GetClusterSecretsParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the get cluster secrets params
func (o *GetClusterSecretsParams) WithLocationRegionRegion(locationRegionRegion *string) *GetClusterSecretsParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the get cluster secrets params
func (o *GetClusterSecretsParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterSecretsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
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

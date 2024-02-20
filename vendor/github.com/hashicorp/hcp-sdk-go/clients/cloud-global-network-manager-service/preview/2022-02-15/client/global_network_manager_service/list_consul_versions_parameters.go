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

// NewListConsulVersionsParams creates a new ListConsulVersionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListConsulVersionsParams() *ListConsulVersionsParams {
	return &ListConsulVersionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListConsulVersionsParamsWithTimeout creates a new ListConsulVersionsParams object
// with the ability to set a timeout on a request.
func NewListConsulVersionsParamsWithTimeout(timeout time.Duration) *ListConsulVersionsParams {
	return &ListConsulVersionsParams{
		timeout: timeout,
	}
}

// NewListConsulVersionsParamsWithContext creates a new ListConsulVersionsParams object
// with the ability to set a context for a request.
func NewListConsulVersionsParamsWithContext(ctx context.Context) *ListConsulVersionsParams {
	return &ListConsulVersionsParams{
		Context: ctx,
	}
}

// NewListConsulVersionsParamsWithHTTPClient creates a new ListConsulVersionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListConsulVersionsParamsWithHTTPClient(client *http.Client) *ListConsulVersionsParams {
	return &ListConsulVersionsParams{
		HTTPClient: client,
	}
}

/*
ListConsulVersionsParams contains all the parameters to send to the API endpoint

	for the list consul versions operation.

	Typically these are written to a http.Request.
*/
type ListConsulVersionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list consul versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListConsulVersionsParams) WithDefaults() *ListConsulVersionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list consul versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListConsulVersionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list consul versions params
func (o *ListConsulVersionsParams) WithTimeout(timeout time.Duration) *ListConsulVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list consul versions params
func (o *ListConsulVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list consul versions params
func (o *ListConsulVersionsParams) WithContext(ctx context.Context) *ListConsulVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list consul versions params
func (o *ListConsulVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list consul versions params
func (o *ListConsulVersionsParams) WithHTTPClient(client *http.Client) *ListConsulVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list consul versions params
func (o *ListConsulVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListConsulVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

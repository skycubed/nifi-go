// Code generated by go-swagger; DO NOT EDIT.

package controller

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

// NewGetControllerConfigParams creates a new GetControllerConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetControllerConfigParams() *GetControllerConfigParams {
	return &GetControllerConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetControllerConfigParamsWithTimeout creates a new GetControllerConfigParams object
// with the ability to set a timeout on a request.
func NewGetControllerConfigParamsWithTimeout(timeout time.Duration) *GetControllerConfigParams {
	return &GetControllerConfigParams{
		timeout: timeout,
	}
}

// NewGetControllerConfigParamsWithContext creates a new GetControllerConfigParams object
// with the ability to set a context for a request.
func NewGetControllerConfigParamsWithContext(ctx context.Context) *GetControllerConfigParams {
	return &GetControllerConfigParams{
		Context: ctx,
	}
}

// NewGetControllerConfigParamsWithHTTPClient creates a new GetControllerConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetControllerConfigParamsWithHTTPClient(client *http.Client) *GetControllerConfigParams {
	return &GetControllerConfigParams{
		HTTPClient: client,
	}
}

/* GetControllerConfigParams contains all the parameters to send to the API endpoint
   for the get controller config operation.

   Typically these are written to a http.Request.
*/
type GetControllerConfigParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get controller config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetControllerConfigParams) WithDefaults() *GetControllerConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get controller config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetControllerConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get controller config params
func (o *GetControllerConfigParams) WithTimeout(timeout time.Duration) *GetControllerConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get controller config params
func (o *GetControllerConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get controller config params
func (o *GetControllerConfigParams) WithContext(ctx context.Context) *GetControllerConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get controller config params
func (o *GetControllerConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get controller config params
func (o *GetControllerConfigParams) WithHTTPClient(client *http.Client) *GetControllerConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get controller config params
func (o *GetControllerConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetControllerConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
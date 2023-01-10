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

// NewGetFlowRegistryClientParams creates a new GetFlowRegistryClientParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFlowRegistryClientParams() *GetFlowRegistryClientParams {
	return &GetFlowRegistryClientParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFlowRegistryClientParamsWithTimeout creates a new GetFlowRegistryClientParams object
// with the ability to set a timeout on a request.
func NewGetFlowRegistryClientParamsWithTimeout(timeout time.Duration) *GetFlowRegistryClientParams {
	return &GetFlowRegistryClientParams{
		timeout: timeout,
	}
}

// NewGetFlowRegistryClientParamsWithContext creates a new GetFlowRegistryClientParams object
// with the ability to set a context for a request.
func NewGetFlowRegistryClientParamsWithContext(ctx context.Context) *GetFlowRegistryClientParams {
	return &GetFlowRegistryClientParams{
		Context: ctx,
	}
}

// NewGetFlowRegistryClientParamsWithHTTPClient creates a new GetFlowRegistryClientParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFlowRegistryClientParamsWithHTTPClient(client *http.Client) *GetFlowRegistryClientParams {
	return &GetFlowRegistryClientParams{
		HTTPClient: client,
	}
}

/*
GetFlowRegistryClientParams contains all the parameters to send to the API endpoint

	for the get flow registry client operation.

	Typically these are written to a http.Request.
*/
type GetFlowRegistryClientParams struct {

	/* ID.

	   The flow registry client id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get flow registry client params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFlowRegistryClientParams) WithDefaults() *GetFlowRegistryClientParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get flow registry client params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFlowRegistryClientParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get flow registry client params
func (o *GetFlowRegistryClientParams) WithTimeout(timeout time.Duration) *GetFlowRegistryClientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get flow registry client params
func (o *GetFlowRegistryClientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get flow registry client params
func (o *GetFlowRegistryClientParams) WithContext(ctx context.Context) *GetFlowRegistryClientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get flow registry client params
func (o *GetFlowRegistryClientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get flow registry client params
func (o *GetFlowRegistryClientParams) WithHTTPClient(client *http.Client) *GetFlowRegistryClientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get flow registry client params
func (o *GetFlowRegistryClientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get flow registry client params
func (o *GetFlowRegistryClientParams) WithID(id string) *GetFlowRegistryClientParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get flow registry client params
func (o *GetFlowRegistryClientParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetFlowRegistryClientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
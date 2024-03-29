// Code generated by go-swagger; DO NOT EDIT.

package process_groups

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

// NewGetInputPortsParams creates a new GetInputPortsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInputPortsParams() *GetInputPortsParams {
	return &GetInputPortsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInputPortsParamsWithTimeout creates a new GetInputPortsParams object
// with the ability to set a timeout on a request.
func NewGetInputPortsParamsWithTimeout(timeout time.Duration) *GetInputPortsParams {
	return &GetInputPortsParams{
		timeout: timeout,
	}
}

// NewGetInputPortsParamsWithContext creates a new GetInputPortsParams object
// with the ability to set a context for a request.
func NewGetInputPortsParamsWithContext(ctx context.Context) *GetInputPortsParams {
	return &GetInputPortsParams{
		Context: ctx,
	}
}

// NewGetInputPortsParamsWithHTTPClient creates a new GetInputPortsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInputPortsParamsWithHTTPClient(client *http.Client) *GetInputPortsParams {
	return &GetInputPortsParams{
		HTTPClient: client,
	}
}

/*
GetInputPortsParams contains all the parameters to send to the API endpoint

	for the get input ports operation.

	Typically these are written to a http.Request.
*/
type GetInputPortsParams struct {

	/* ID.

	   The process group id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get input ports params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInputPortsParams) WithDefaults() *GetInputPortsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get input ports params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInputPortsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get input ports params
func (o *GetInputPortsParams) WithTimeout(timeout time.Duration) *GetInputPortsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get input ports params
func (o *GetInputPortsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get input ports params
func (o *GetInputPortsParams) WithContext(ctx context.Context) *GetInputPortsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get input ports params
func (o *GetInputPortsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get input ports params
func (o *GetInputPortsParams) WithHTTPClient(client *http.Client) *GetInputPortsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get input ports params
func (o *GetInputPortsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get input ports params
func (o *GetInputPortsParams) WithID(id string) *GetInputPortsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get input ports params
func (o *GetInputPortsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetInputPortsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

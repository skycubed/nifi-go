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

// NewGetOutputPortsParams creates a new GetOutputPortsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOutputPortsParams() *GetOutputPortsParams {
	return &GetOutputPortsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutputPortsParamsWithTimeout creates a new GetOutputPortsParams object
// with the ability to set a timeout on a request.
func NewGetOutputPortsParamsWithTimeout(timeout time.Duration) *GetOutputPortsParams {
	return &GetOutputPortsParams{
		timeout: timeout,
	}
}

// NewGetOutputPortsParamsWithContext creates a new GetOutputPortsParams object
// with the ability to set a context for a request.
func NewGetOutputPortsParamsWithContext(ctx context.Context) *GetOutputPortsParams {
	return &GetOutputPortsParams{
		Context: ctx,
	}
}

// NewGetOutputPortsParamsWithHTTPClient creates a new GetOutputPortsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOutputPortsParamsWithHTTPClient(client *http.Client) *GetOutputPortsParams {
	return &GetOutputPortsParams{
		HTTPClient: client,
	}
}

/* GetOutputPortsParams contains all the parameters to send to the API endpoint
   for the get output ports operation.

   Typically these are written to a http.Request.
*/
type GetOutputPortsParams struct {

	/* ID.

	   The process group id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get output ports params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOutputPortsParams) WithDefaults() *GetOutputPortsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get output ports params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOutputPortsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get output ports params
func (o *GetOutputPortsParams) WithTimeout(timeout time.Duration) *GetOutputPortsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get output ports params
func (o *GetOutputPortsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get output ports params
func (o *GetOutputPortsParams) WithContext(ctx context.Context) *GetOutputPortsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get output ports params
func (o *GetOutputPortsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get output ports params
func (o *GetOutputPortsParams) WithHTTPClient(client *http.Client) *GetOutputPortsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get output ports params
func (o *GetOutputPortsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get output ports params
func (o *GetOutputPortsParams) WithID(id string) *GetOutputPortsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get output ports params
func (o *GetOutputPortsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutputPortsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
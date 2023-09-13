// Code generated by go-swagger; DO NOT EDIT.

package processors

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

// NewGetProcessorDiagnosticsParams creates a new GetProcessorDiagnosticsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProcessorDiagnosticsParams() *GetProcessorDiagnosticsParams {
	return &GetProcessorDiagnosticsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProcessorDiagnosticsParamsWithTimeout creates a new GetProcessorDiagnosticsParams object
// with the ability to set a timeout on a request.
func NewGetProcessorDiagnosticsParamsWithTimeout(timeout time.Duration) *GetProcessorDiagnosticsParams {
	return &GetProcessorDiagnosticsParams{
		timeout: timeout,
	}
}

// NewGetProcessorDiagnosticsParamsWithContext creates a new GetProcessorDiagnosticsParams object
// with the ability to set a context for a request.
func NewGetProcessorDiagnosticsParamsWithContext(ctx context.Context) *GetProcessorDiagnosticsParams {
	return &GetProcessorDiagnosticsParams{
		Context: ctx,
	}
}

// NewGetProcessorDiagnosticsParamsWithHTTPClient creates a new GetProcessorDiagnosticsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProcessorDiagnosticsParamsWithHTTPClient(client *http.Client) *GetProcessorDiagnosticsParams {
	return &GetProcessorDiagnosticsParams{
		HTTPClient: client,
	}
}

/*
GetProcessorDiagnosticsParams contains all the parameters to send to the API endpoint

	for the get processor diagnostics operation.

	Typically these are written to a http.Request.
*/
type GetProcessorDiagnosticsParams struct {

	/* ID.

	   The processor id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get processor diagnostics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProcessorDiagnosticsParams) WithDefaults() *GetProcessorDiagnosticsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get processor diagnostics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProcessorDiagnosticsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get processor diagnostics params
func (o *GetProcessorDiagnosticsParams) WithTimeout(timeout time.Duration) *GetProcessorDiagnosticsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get processor diagnostics params
func (o *GetProcessorDiagnosticsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get processor diagnostics params
func (o *GetProcessorDiagnosticsParams) WithContext(ctx context.Context) *GetProcessorDiagnosticsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get processor diagnostics params
func (o *GetProcessorDiagnosticsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get processor diagnostics params
func (o *GetProcessorDiagnosticsParams) WithHTTPClient(client *http.Client) *GetProcessorDiagnosticsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get processor diagnostics params
func (o *GetProcessorDiagnosticsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get processor diagnostics params
func (o *GetProcessorDiagnosticsParams) WithID(id string) *GetProcessorDiagnosticsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get processor diagnostics params
func (o *GetProcessorDiagnosticsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetProcessorDiagnosticsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

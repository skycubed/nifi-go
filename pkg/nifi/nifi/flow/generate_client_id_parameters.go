// Code generated by go-swagger; DO NOT EDIT.

package flow

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

// NewGenerateClientIDParams creates a new GenerateClientIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGenerateClientIDParams() *GenerateClientIDParams {
	return &GenerateClientIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGenerateClientIDParamsWithTimeout creates a new GenerateClientIDParams object
// with the ability to set a timeout on a request.
func NewGenerateClientIDParamsWithTimeout(timeout time.Duration) *GenerateClientIDParams {
	return &GenerateClientIDParams{
		timeout: timeout,
	}
}

// NewGenerateClientIDParamsWithContext creates a new GenerateClientIDParams object
// with the ability to set a context for a request.
func NewGenerateClientIDParamsWithContext(ctx context.Context) *GenerateClientIDParams {
	return &GenerateClientIDParams{
		Context: ctx,
	}
}

// NewGenerateClientIDParamsWithHTTPClient creates a new GenerateClientIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGenerateClientIDParamsWithHTTPClient(client *http.Client) *GenerateClientIDParams {
	return &GenerateClientIDParams{
		HTTPClient: client,
	}
}

/*
GenerateClientIDParams contains all the parameters to send to the API endpoint

	for the generate client Id operation.

	Typically these are written to a http.Request.
*/
type GenerateClientIDParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the generate client Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateClientIDParams) WithDefaults() *GenerateClientIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the generate client Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateClientIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the generate client Id params
func (o *GenerateClientIDParams) WithTimeout(timeout time.Duration) *GenerateClientIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generate client Id params
func (o *GenerateClientIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generate client Id params
func (o *GenerateClientIDParams) WithContext(ctx context.Context) *GenerateClientIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generate client Id params
func (o *GenerateClientIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the generate client Id params
func (o *GenerateClientIDParams) WithHTTPClient(client *http.Client) *GenerateClientIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generate client Id params
func (o *GenerateClientIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GenerateClientIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

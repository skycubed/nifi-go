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

// NewGetParameterContextsParams creates a new GetParameterContextsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetParameterContextsParams() *GetParameterContextsParams {
	return &GetParameterContextsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetParameterContextsParamsWithTimeout creates a new GetParameterContextsParams object
// with the ability to set a timeout on a request.
func NewGetParameterContextsParamsWithTimeout(timeout time.Duration) *GetParameterContextsParams {
	return &GetParameterContextsParams{
		timeout: timeout,
	}
}

// NewGetParameterContextsParamsWithContext creates a new GetParameterContextsParams object
// with the ability to set a context for a request.
func NewGetParameterContextsParamsWithContext(ctx context.Context) *GetParameterContextsParams {
	return &GetParameterContextsParams{
		Context: ctx,
	}
}

// NewGetParameterContextsParamsWithHTTPClient creates a new GetParameterContextsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetParameterContextsParamsWithHTTPClient(client *http.Client) *GetParameterContextsParams {
	return &GetParameterContextsParams{
		HTTPClient: client,
	}
}

/*
GetParameterContextsParams contains all the parameters to send to the API endpoint

	for the get parameter contexts operation.

	Typically these are written to a http.Request.
*/
type GetParameterContextsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get parameter contexts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetParameterContextsParams) WithDefaults() *GetParameterContextsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get parameter contexts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetParameterContextsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get parameter contexts params
func (o *GetParameterContextsParams) WithTimeout(timeout time.Duration) *GetParameterContextsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get parameter contexts params
func (o *GetParameterContextsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get parameter contexts params
func (o *GetParameterContextsParams) WithContext(ctx context.Context) *GetParameterContextsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get parameter contexts params
func (o *GetParameterContextsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get parameter contexts params
func (o *GetParameterContextsParams) WithHTTPClient(client *http.Client) *GetParameterContextsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get parameter contexts params
func (o *GetParameterContextsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetParameterContextsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

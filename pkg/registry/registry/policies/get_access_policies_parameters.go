// Code generated by go-swagger; DO NOT EDIT.

package policies

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

// NewGetAccessPoliciesParams creates a new GetAccessPoliciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccessPoliciesParams() *GetAccessPoliciesParams {
	return &GetAccessPoliciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccessPoliciesParamsWithTimeout creates a new GetAccessPoliciesParams object
// with the ability to set a timeout on a request.
func NewGetAccessPoliciesParamsWithTimeout(timeout time.Duration) *GetAccessPoliciesParams {
	return &GetAccessPoliciesParams{
		timeout: timeout,
	}
}

// NewGetAccessPoliciesParamsWithContext creates a new GetAccessPoliciesParams object
// with the ability to set a context for a request.
func NewGetAccessPoliciesParamsWithContext(ctx context.Context) *GetAccessPoliciesParams {
	return &GetAccessPoliciesParams{
		Context: ctx,
	}
}

// NewGetAccessPoliciesParamsWithHTTPClient creates a new GetAccessPoliciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccessPoliciesParamsWithHTTPClient(client *http.Client) *GetAccessPoliciesParams {
	return &GetAccessPoliciesParams{
		HTTPClient: client,
	}
}

/* GetAccessPoliciesParams contains all the parameters to send to the API endpoint
   for the get access policies operation.

   Typically these are written to a http.Request.
*/
type GetAccessPoliciesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get access policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccessPoliciesParams) WithDefaults() *GetAccessPoliciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get access policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccessPoliciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get access policies params
func (o *GetAccessPoliciesParams) WithTimeout(timeout time.Duration) *GetAccessPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get access policies params
func (o *GetAccessPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get access policies params
func (o *GetAccessPoliciesParams) WithContext(ctx context.Context) *GetAccessPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get access policies params
func (o *GetAccessPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get access policies params
func (o *GetAccessPoliciesParams) WithHTTPClient(client *http.Client) *GetAccessPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get access policies params
func (o *GetAccessPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccessPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
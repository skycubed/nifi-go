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

// NewGetAccessPolicyParams creates a new GetAccessPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccessPolicyParams() *GetAccessPolicyParams {
	return &GetAccessPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccessPolicyParamsWithTimeout creates a new GetAccessPolicyParams object
// with the ability to set a timeout on a request.
func NewGetAccessPolicyParamsWithTimeout(timeout time.Duration) *GetAccessPolicyParams {
	return &GetAccessPolicyParams{
		timeout: timeout,
	}
}

// NewGetAccessPolicyParamsWithContext creates a new GetAccessPolicyParams object
// with the ability to set a context for a request.
func NewGetAccessPolicyParamsWithContext(ctx context.Context) *GetAccessPolicyParams {
	return &GetAccessPolicyParams{
		Context: ctx,
	}
}

// NewGetAccessPolicyParamsWithHTTPClient creates a new GetAccessPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccessPolicyParamsWithHTTPClient(client *http.Client) *GetAccessPolicyParams {
	return &GetAccessPolicyParams{
		HTTPClient: client,
	}
}

/*
GetAccessPolicyParams contains all the parameters to send to the API endpoint

	for the get access policy operation.

	Typically these are written to a http.Request.
*/
type GetAccessPolicyParams struct {

	/* ID.

	   The access policy id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get access policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccessPolicyParams) WithDefaults() *GetAccessPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get access policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccessPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get access policy params
func (o *GetAccessPolicyParams) WithTimeout(timeout time.Duration) *GetAccessPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get access policy params
func (o *GetAccessPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get access policy params
func (o *GetAccessPolicyParams) WithContext(ctx context.Context) *GetAccessPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get access policy params
func (o *GetAccessPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get access policy params
func (o *GetAccessPolicyParams) WithHTTPClient(client *http.Client) *GetAccessPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get access policy params
func (o *GetAccessPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get access policy params
func (o *GetAccessPolicyParams) WithID(id string) *GetAccessPolicyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get access policy params
func (o *GetAccessPolicyParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccessPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

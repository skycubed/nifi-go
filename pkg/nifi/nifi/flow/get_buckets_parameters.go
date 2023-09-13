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

// NewGetBucketsParams creates a new GetBucketsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBucketsParams() *GetBucketsParams {
	return &GetBucketsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBucketsParamsWithTimeout creates a new GetBucketsParams object
// with the ability to set a timeout on a request.
func NewGetBucketsParamsWithTimeout(timeout time.Duration) *GetBucketsParams {
	return &GetBucketsParams{
		timeout: timeout,
	}
}

// NewGetBucketsParamsWithContext creates a new GetBucketsParams object
// with the ability to set a context for a request.
func NewGetBucketsParamsWithContext(ctx context.Context) *GetBucketsParams {
	return &GetBucketsParams{
		Context: ctx,
	}
}

// NewGetBucketsParamsWithHTTPClient creates a new GetBucketsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBucketsParamsWithHTTPClient(client *http.Client) *GetBucketsParams {
	return &GetBucketsParams{
		HTTPClient: client,
	}
}

/*
GetBucketsParams contains all the parameters to send to the API endpoint

	for the get buckets operation.

	Typically these are written to a http.Request.
*/
type GetBucketsParams struct {

	/* ID.

	   The registry id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get buckets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBucketsParams) WithDefaults() *GetBucketsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get buckets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBucketsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get buckets params
func (o *GetBucketsParams) WithTimeout(timeout time.Duration) *GetBucketsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get buckets params
func (o *GetBucketsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get buckets params
func (o *GetBucketsParams) WithContext(ctx context.Context) *GetBucketsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get buckets params
func (o *GetBucketsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get buckets params
func (o *GetBucketsParams) WithHTTPClient(client *http.Client) *GetBucketsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get buckets params
func (o *GetBucketsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get buckets params
func (o *GetBucketsParams) WithID(id string) *GetBucketsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get buckets params
func (o *GetBucketsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetBucketsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

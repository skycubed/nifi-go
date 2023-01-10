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

// NewGetFlowsParams creates a new GetFlowsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFlowsParams() *GetFlowsParams {
	return &GetFlowsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFlowsParamsWithTimeout creates a new GetFlowsParams object
// with the ability to set a timeout on a request.
func NewGetFlowsParamsWithTimeout(timeout time.Duration) *GetFlowsParams {
	return &GetFlowsParams{
		timeout: timeout,
	}
}

// NewGetFlowsParamsWithContext creates a new GetFlowsParams object
// with the ability to set a context for a request.
func NewGetFlowsParamsWithContext(ctx context.Context) *GetFlowsParams {
	return &GetFlowsParams{
		Context: ctx,
	}
}

// NewGetFlowsParamsWithHTTPClient creates a new GetFlowsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFlowsParamsWithHTTPClient(client *http.Client) *GetFlowsParams {
	return &GetFlowsParams{
		HTTPClient: client,
	}
}

/*
GetFlowsParams contains all the parameters to send to the API endpoint

	for the get flows operation.

	Typically these are written to a http.Request.
*/
type GetFlowsParams struct {

	/* BucketID.

	   The bucket id.
	*/
	BucketID string

	/* RegistryID.

	   The registry client id.
	*/
	RegistryID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get flows params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFlowsParams) WithDefaults() *GetFlowsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get flows params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFlowsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get flows params
func (o *GetFlowsParams) WithTimeout(timeout time.Duration) *GetFlowsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get flows params
func (o *GetFlowsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get flows params
func (o *GetFlowsParams) WithContext(ctx context.Context) *GetFlowsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get flows params
func (o *GetFlowsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get flows params
func (o *GetFlowsParams) WithHTTPClient(client *http.Client) *GetFlowsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get flows params
func (o *GetFlowsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketID adds the bucketID to the get flows params
func (o *GetFlowsParams) WithBucketID(bucketID string) *GetFlowsParams {
	o.SetBucketID(bucketID)
	return o
}

// SetBucketID adds the bucketId to the get flows params
func (o *GetFlowsParams) SetBucketID(bucketID string) {
	o.BucketID = bucketID
}

// WithRegistryID adds the registryID to the get flows params
func (o *GetFlowsParams) WithRegistryID(registryID string) *GetFlowsParams {
	o.SetRegistryID(registryID)
	return o
}

// SetRegistryID adds the registryId to the get flows params
func (o *GetFlowsParams) SetRegistryID(registryID string) {
	o.RegistryID = registryID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFlowsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bucket-id
	if err := r.SetPathParam("bucket-id", o.BucketID); err != nil {
		return err
	}

	// path param registry-id
	if err := r.SetPathParam("registry-id", o.RegistryID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

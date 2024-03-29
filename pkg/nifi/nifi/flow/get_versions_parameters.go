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

// NewGetVersionsParams creates a new GetVersionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVersionsParams() *GetVersionsParams {
	return &GetVersionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVersionsParamsWithTimeout creates a new GetVersionsParams object
// with the ability to set a timeout on a request.
func NewGetVersionsParamsWithTimeout(timeout time.Duration) *GetVersionsParams {
	return &GetVersionsParams{
		timeout: timeout,
	}
}

// NewGetVersionsParamsWithContext creates a new GetVersionsParams object
// with the ability to set a context for a request.
func NewGetVersionsParamsWithContext(ctx context.Context) *GetVersionsParams {
	return &GetVersionsParams{
		Context: ctx,
	}
}

// NewGetVersionsParamsWithHTTPClient creates a new GetVersionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVersionsParamsWithHTTPClient(client *http.Client) *GetVersionsParams {
	return &GetVersionsParams{
		HTTPClient: client,
	}
}

/*
GetVersionsParams contains all the parameters to send to the API endpoint

	for the get versions operation.

	Typically these are written to a http.Request.
*/
type GetVersionsParams struct {

	/* BucketID.

	   The bucket id.
	*/
	BucketID string

	/* FlowID.

	   The flow id.
	*/
	FlowID string

	/* RegistryID.

	   The registry id.
	*/
	RegistryID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVersionsParams) WithDefaults() *GetVersionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVersionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get versions params
func (o *GetVersionsParams) WithTimeout(timeout time.Duration) *GetVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get versions params
func (o *GetVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get versions params
func (o *GetVersionsParams) WithContext(ctx context.Context) *GetVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get versions params
func (o *GetVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get versions params
func (o *GetVersionsParams) WithHTTPClient(client *http.Client) *GetVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get versions params
func (o *GetVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketID adds the bucketID to the get versions params
func (o *GetVersionsParams) WithBucketID(bucketID string) *GetVersionsParams {
	o.SetBucketID(bucketID)
	return o
}

// SetBucketID adds the bucketId to the get versions params
func (o *GetVersionsParams) SetBucketID(bucketID string) {
	o.BucketID = bucketID
}

// WithFlowID adds the flowID to the get versions params
func (o *GetVersionsParams) WithFlowID(flowID string) *GetVersionsParams {
	o.SetFlowID(flowID)
	return o
}

// SetFlowID adds the flowId to the get versions params
func (o *GetVersionsParams) SetFlowID(flowID string) {
	o.FlowID = flowID
}

// WithRegistryID adds the registryID to the get versions params
func (o *GetVersionsParams) WithRegistryID(registryID string) *GetVersionsParams {
	o.SetRegistryID(registryID)
	return o
}

// SetRegistryID adds the registryId to the get versions params
func (o *GetVersionsParams) SetRegistryID(registryID string) {
	o.RegistryID = registryID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bucket-id
	if err := r.SetPathParam("bucket-id", o.BucketID); err != nil {
		return err
	}

	// path param flow-id
	if err := r.SetPathParam("flow-id", o.FlowID); err != nil {
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

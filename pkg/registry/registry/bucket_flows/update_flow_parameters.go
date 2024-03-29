// Code generated by go-swagger; DO NOT EDIT.

package bucket_flows

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

	"github.com/skycubed/nifi-go/pkg/registry/models"
)

// NewUpdateFlowParams creates a new UpdateFlowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateFlowParams() *UpdateFlowParams {
	return &UpdateFlowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateFlowParamsWithTimeout creates a new UpdateFlowParams object
// with the ability to set a timeout on a request.
func NewUpdateFlowParamsWithTimeout(timeout time.Duration) *UpdateFlowParams {
	return &UpdateFlowParams{
		timeout: timeout,
	}
}

// NewUpdateFlowParamsWithContext creates a new UpdateFlowParams object
// with the ability to set a context for a request.
func NewUpdateFlowParamsWithContext(ctx context.Context) *UpdateFlowParams {
	return &UpdateFlowParams{
		Context: ctx,
	}
}

// NewUpdateFlowParamsWithHTTPClient creates a new UpdateFlowParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateFlowParamsWithHTTPClient(client *http.Client) *UpdateFlowParams {
	return &UpdateFlowParams{
		HTTPClient: client,
	}
}

/*
UpdateFlowParams contains all the parameters to send to the API endpoint

	for the update flow operation.

	Typically these are written to a http.Request.
*/
type UpdateFlowParams struct {

	/* Body.

	   The updated flow
	*/
	Body *models.VersionedFlow

	/* BucketID.

	   The bucket identifier
	*/
	BucketID string

	/* FlowID.

	   The flow identifier
	*/
	FlowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateFlowParams) WithDefaults() *UpdateFlowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateFlowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update flow params
func (o *UpdateFlowParams) WithTimeout(timeout time.Duration) *UpdateFlowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update flow params
func (o *UpdateFlowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update flow params
func (o *UpdateFlowParams) WithContext(ctx context.Context) *UpdateFlowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update flow params
func (o *UpdateFlowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update flow params
func (o *UpdateFlowParams) WithHTTPClient(client *http.Client) *UpdateFlowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update flow params
func (o *UpdateFlowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update flow params
func (o *UpdateFlowParams) WithBody(body *models.VersionedFlow) *UpdateFlowParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update flow params
func (o *UpdateFlowParams) SetBody(body *models.VersionedFlow) {
	o.Body = body
}

// WithBucketID adds the bucketID to the update flow params
func (o *UpdateFlowParams) WithBucketID(bucketID string) *UpdateFlowParams {
	o.SetBucketID(bucketID)
	return o
}

// SetBucketID adds the bucketId to the update flow params
func (o *UpdateFlowParams) SetBucketID(bucketID string) {
	o.BucketID = bucketID
}

// WithFlowID adds the flowID to the update flow params
func (o *UpdateFlowParams) WithFlowID(flowID string) *UpdateFlowParams {
	o.SetFlowID(flowID)
	return o
}

// SetFlowID adds the flowId to the update flow params
func (o *UpdateFlowParams) SetFlowID(flowID string) {
	o.FlowID = flowID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateFlowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param bucketId
	if err := r.SetPathParam("bucketId", o.BucketID); err != nil {
		return err
	}

	// path param flowId
	if err := r.SetPathParam("flowId", o.FlowID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

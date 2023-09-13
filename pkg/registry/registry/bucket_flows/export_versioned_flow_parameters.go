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
	"github.com/go-openapi/swag"
)

// NewExportVersionedFlowParams creates a new ExportVersionedFlowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExportVersionedFlowParams() *ExportVersionedFlowParams {
	return &ExportVersionedFlowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExportVersionedFlowParamsWithTimeout creates a new ExportVersionedFlowParams object
// with the ability to set a timeout on a request.
func NewExportVersionedFlowParamsWithTimeout(timeout time.Duration) *ExportVersionedFlowParams {
	return &ExportVersionedFlowParams{
		timeout: timeout,
	}
}

// NewExportVersionedFlowParamsWithContext creates a new ExportVersionedFlowParams object
// with the ability to set a context for a request.
func NewExportVersionedFlowParamsWithContext(ctx context.Context) *ExportVersionedFlowParams {
	return &ExportVersionedFlowParams{
		Context: ctx,
	}
}

// NewExportVersionedFlowParamsWithHTTPClient creates a new ExportVersionedFlowParams object
// with the ability to set a custom HTTPClient for a request.
func NewExportVersionedFlowParamsWithHTTPClient(client *http.Client) *ExportVersionedFlowParams {
	return &ExportVersionedFlowParams{
		HTTPClient: client,
	}
}

/*
ExportVersionedFlowParams contains all the parameters to send to the API endpoint

	for the export versioned flow operation.

	Typically these are written to a http.Request.
*/
type ExportVersionedFlowParams struct {

	/* BucketID.

	   The bucket identifier
	*/
	BucketID string

	/* FlowID.

	   The flow identifier
	*/
	FlowID string

	/* VersionNumber.

	   The version number

	   Format: int32
	*/
	VersionNumber int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the export versioned flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportVersionedFlowParams) WithDefaults() *ExportVersionedFlowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the export versioned flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportVersionedFlowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the export versioned flow params
func (o *ExportVersionedFlowParams) WithTimeout(timeout time.Duration) *ExportVersionedFlowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export versioned flow params
func (o *ExportVersionedFlowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export versioned flow params
func (o *ExportVersionedFlowParams) WithContext(ctx context.Context) *ExportVersionedFlowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export versioned flow params
func (o *ExportVersionedFlowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export versioned flow params
func (o *ExportVersionedFlowParams) WithHTTPClient(client *http.Client) *ExportVersionedFlowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export versioned flow params
func (o *ExportVersionedFlowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketID adds the bucketID to the export versioned flow params
func (o *ExportVersionedFlowParams) WithBucketID(bucketID string) *ExportVersionedFlowParams {
	o.SetBucketID(bucketID)
	return o
}

// SetBucketID adds the bucketId to the export versioned flow params
func (o *ExportVersionedFlowParams) SetBucketID(bucketID string) {
	o.BucketID = bucketID
}

// WithFlowID adds the flowID to the export versioned flow params
func (o *ExportVersionedFlowParams) WithFlowID(flowID string) *ExportVersionedFlowParams {
	o.SetFlowID(flowID)
	return o
}

// SetFlowID adds the flowId to the export versioned flow params
func (o *ExportVersionedFlowParams) SetFlowID(flowID string) {
	o.FlowID = flowID
}

// WithVersionNumber adds the versionNumber to the export versioned flow params
func (o *ExportVersionedFlowParams) WithVersionNumber(versionNumber int32) *ExportVersionedFlowParams {
	o.SetVersionNumber(versionNumber)
	return o
}

// SetVersionNumber adds the versionNumber to the export versioned flow params
func (o *ExportVersionedFlowParams) SetVersionNumber(versionNumber int32) {
	o.VersionNumber = versionNumber
}

// WriteToRequest writes these params to a swagger request
func (o *ExportVersionedFlowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bucketId
	if err := r.SetPathParam("bucketId", o.BucketID); err != nil {
		return err
	}

	// path param flowId
	if err := r.SetPathParam("flowId", o.FlowID); err != nil {
		return err
	}

	// path param versionNumber
	if err := r.SetPathParam("versionNumber", swag.FormatInt32(o.VersionNumber)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

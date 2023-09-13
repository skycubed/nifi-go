// Code generated by go-swagger; DO NOT EDIT.

package bucket_bundles

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

// NewCreateExtensionBundleVersionParams creates a new CreateExtensionBundleVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateExtensionBundleVersionParams() *CreateExtensionBundleVersionParams {
	return &CreateExtensionBundleVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateExtensionBundleVersionParamsWithTimeout creates a new CreateExtensionBundleVersionParams object
// with the ability to set a timeout on a request.
func NewCreateExtensionBundleVersionParamsWithTimeout(timeout time.Duration) *CreateExtensionBundleVersionParams {
	return &CreateExtensionBundleVersionParams{
		timeout: timeout,
	}
}

// NewCreateExtensionBundleVersionParamsWithContext creates a new CreateExtensionBundleVersionParams object
// with the ability to set a context for a request.
func NewCreateExtensionBundleVersionParamsWithContext(ctx context.Context) *CreateExtensionBundleVersionParams {
	return &CreateExtensionBundleVersionParams{
		Context: ctx,
	}
}

// NewCreateExtensionBundleVersionParamsWithHTTPClient creates a new CreateExtensionBundleVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateExtensionBundleVersionParamsWithHTTPClient(client *http.Client) *CreateExtensionBundleVersionParams {
	return &CreateExtensionBundleVersionParams{
		HTTPClient: client,
	}
}

/*
CreateExtensionBundleVersionParams contains all the parameters to send to the API endpoint

	for the create extension bundle version operation.

	Typically these are written to a http.Request.
*/
type CreateExtensionBundleVersionParams struct {

	/* BucketID.

	   The bucket identifier
	*/
	BucketID string

	/* BundleType.

	   The type of the bundle
	*/
	BundleType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create extension bundle version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateExtensionBundleVersionParams) WithDefaults() *CreateExtensionBundleVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create extension bundle version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateExtensionBundleVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create extension bundle version params
func (o *CreateExtensionBundleVersionParams) WithTimeout(timeout time.Duration) *CreateExtensionBundleVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create extension bundle version params
func (o *CreateExtensionBundleVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create extension bundle version params
func (o *CreateExtensionBundleVersionParams) WithContext(ctx context.Context) *CreateExtensionBundleVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create extension bundle version params
func (o *CreateExtensionBundleVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create extension bundle version params
func (o *CreateExtensionBundleVersionParams) WithHTTPClient(client *http.Client) *CreateExtensionBundleVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create extension bundle version params
func (o *CreateExtensionBundleVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketID adds the bucketID to the create extension bundle version params
func (o *CreateExtensionBundleVersionParams) WithBucketID(bucketID string) *CreateExtensionBundleVersionParams {
	o.SetBucketID(bucketID)
	return o
}

// SetBucketID adds the bucketId to the create extension bundle version params
func (o *CreateExtensionBundleVersionParams) SetBucketID(bucketID string) {
	o.BucketID = bucketID
}

// WithBundleType adds the bundleType to the create extension bundle version params
func (o *CreateExtensionBundleVersionParams) WithBundleType(bundleType string) *CreateExtensionBundleVersionParams {
	o.SetBundleType(bundleType)
	return o
}

// SetBundleType adds the bundleType to the create extension bundle version params
func (o *CreateExtensionBundleVersionParams) SetBundleType(bundleType string) {
	o.BundleType = bundleType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateExtensionBundleVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bucketId
	if err := r.SetPathParam("bucketId", o.BucketID); err != nil {
		return err
	}

	// path param bundleType
	if err := r.SetPathParam("bundleType", o.BundleType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

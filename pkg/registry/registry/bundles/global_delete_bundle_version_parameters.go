// Code generated by go-swagger; DO NOT EDIT.

package bundles

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

// NewGlobalDeleteBundleVersionParams creates a new GlobalDeleteBundleVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGlobalDeleteBundleVersionParams() *GlobalDeleteBundleVersionParams {
	return &GlobalDeleteBundleVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGlobalDeleteBundleVersionParamsWithTimeout creates a new GlobalDeleteBundleVersionParams object
// with the ability to set a timeout on a request.
func NewGlobalDeleteBundleVersionParamsWithTimeout(timeout time.Duration) *GlobalDeleteBundleVersionParams {
	return &GlobalDeleteBundleVersionParams{
		timeout: timeout,
	}
}

// NewGlobalDeleteBundleVersionParamsWithContext creates a new GlobalDeleteBundleVersionParams object
// with the ability to set a context for a request.
func NewGlobalDeleteBundleVersionParamsWithContext(ctx context.Context) *GlobalDeleteBundleVersionParams {
	return &GlobalDeleteBundleVersionParams{
		Context: ctx,
	}
}

// NewGlobalDeleteBundleVersionParamsWithHTTPClient creates a new GlobalDeleteBundleVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGlobalDeleteBundleVersionParamsWithHTTPClient(client *http.Client) *GlobalDeleteBundleVersionParams {
	return &GlobalDeleteBundleVersionParams{
		HTTPClient: client,
	}
}

/*
GlobalDeleteBundleVersionParams contains all the parameters to send to the API endpoint

	for the global delete bundle version operation.

	Typically these are written to a http.Request.
*/
type GlobalDeleteBundleVersionParams struct {

	/* BundleID.

	   The extension bundle identifier
	*/
	BundleID string

	/* Version.

	   The version of the bundle
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the global delete bundle version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GlobalDeleteBundleVersionParams) WithDefaults() *GlobalDeleteBundleVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the global delete bundle version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GlobalDeleteBundleVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the global delete bundle version params
func (o *GlobalDeleteBundleVersionParams) WithTimeout(timeout time.Duration) *GlobalDeleteBundleVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the global delete bundle version params
func (o *GlobalDeleteBundleVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the global delete bundle version params
func (o *GlobalDeleteBundleVersionParams) WithContext(ctx context.Context) *GlobalDeleteBundleVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the global delete bundle version params
func (o *GlobalDeleteBundleVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the global delete bundle version params
func (o *GlobalDeleteBundleVersionParams) WithHTTPClient(client *http.Client) *GlobalDeleteBundleVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the global delete bundle version params
func (o *GlobalDeleteBundleVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBundleID adds the bundleID to the global delete bundle version params
func (o *GlobalDeleteBundleVersionParams) WithBundleID(bundleID string) *GlobalDeleteBundleVersionParams {
	o.SetBundleID(bundleID)
	return o
}

// SetBundleID adds the bundleId to the global delete bundle version params
func (o *GlobalDeleteBundleVersionParams) SetBundleID(bundleID string) {
	o.BundleID = bundleID
}

// WithVersion adds the version to the global delete bundle version params
func (o *GlobalDeleteBundleVersionParams) WithVersion(version string) *GlobalDeleteBundleVersionParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the global delete bundle version params
func (o *GlobalDeleteBundleVersionParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GlobalDeleteBundleVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bundleId
	if err := r.SetPathParam("bundleId", o.BundleID); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

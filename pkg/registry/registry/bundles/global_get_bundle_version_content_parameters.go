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

// NewGlobalGetBundleVersionContentParams creates a new GlobalGetBundleVersionContentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGlobalGetBundleVersionContentParams() *GlobalGetBundleVersionContentParams {
	return &GlobalGetBundleVersionContentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGlobalGetBundleVersionContentParamsWithTimeout creates a new GlobalGetBundleVersionContentParams object
// with the ability to set a timeout on a request.
func NewGlobalGetBundleVersionContentParamsWithTimeout(timeout time.Duration) *GlobalGetBundleVersionContentParams {
	return &GlobalGetBundleVersionContentParams{
		timeout: timeout,
	}
}

// NewGlobalGetBundleVersionContentParamsWithContext creates a new GlobalGetBundleVersionContentParams object
// with the ability to set a context for a request.
func NewGlobalGetBundleVersionContentParamsWithContext(ctx context.Context) *GlobalGetBundleVersionContentParams {
	return &GlobalGetBundleVersionContentParams{
		Context: ctx,
	}
}

// NewGlobalGetBundleVersionContentParamsWithHTTPClient creates a new GlobalGetBundleVersionContentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGlobalGetBundleVersionContentParamsWithHTTPClient(client *http.Client) *GlobalGetBundleVersionContentParams {
	return &GlobalGetBundleVersionContentParams{
		HTTPClient: client,
	}
}

/*
GlobalGetBundleVersionContentParams contains all the parameters to send to the API endpoint

	for the global get bundle version content operation.

	Typically these are written to a http.Request.
*/
type GlobalGetBundleVersionContentParams struct {

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

// WithDefaults hydrates default values in the global get bundle version content params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GlobalGetBundleVersionContentParams) WithDefaults() *GlobalGetBundleVersionContentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the global get bundle version content params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GlobalGetBundleVersionContentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the global get bundle version content params
func (o *GlobalGetBundleVersionContentParams) WithTimeout(timeout time.Duration) *GlobalGetBundleVersionContentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the global get bundle version content params
func (o *GlobalGetBundleVersionContentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the global get bundle version content params
func (o *GlobalGetBundleVersionContentParams) WithContext(ctx context.Context) *GlobalGetBundleVersionContentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the global get bundle version content params
func (o *GlobalGetBundleVersionContentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the global get bundle version content params
func (o *GlobalGetBundleVersionContentParams) WithHTTPClient(client *http.Client) *GlobalGetBundleVersionContentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the global get bundle version content params
func (o *GlobalGetBundleVersionContentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBundleID adds the bundleID to the global get bundle version content params
func (o *GlobalGetBundleVersionContentParams) WithBundleID(bundleID string) *GlobalGetBundleVersionContentParams {
	o.SetBundleID(bundleID)
	return o
}

// SetBundleID adds the bundleId to the global get bundle version content params
func (o *GlobalGetBundleVersionContentParams) SetBundleID(bundleID string) {
	o.BundleID = bundleID
}

// WithVersion adds the version to the global get bundle version content params
func (o *GlobalGetBundleVersionContentParams) WithVersion(version string) *GlobalGetBundleVersionContentParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the global get bundle version content params
func (o *GlobalGetBundleVersionContentParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GlobalGetBundleVersionContentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

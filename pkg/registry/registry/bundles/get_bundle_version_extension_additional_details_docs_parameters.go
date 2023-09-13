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

// NewGetBundleVersionExtensionAdditionalDetailsDocsParams creates a new GetBundleVersionExtensionAdditionalDetailsDocsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBundleVersionExtensionAdditionalDetailsDocsParams() *GetBundleVersionExtensionAdditionalDetailsDocsParams {
	return &GetBundleVersionExtensionAdditionalDetailsDocsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBundleVersionExtensionAdditionalDetailsDocsParamsWithTimeout creates a new GetBundleVersionExtensionAdditionalDetailsDocsParams object
// with the ability to set a timeout on a request.
func NewGetBundleVersionExtensionAdditionalDetailsDocsParamsWithTimeout(timeout time.Duration) *GetBundleVersionExtensionAdditionalDetailsDocsParams {
	return &GetBundleVersionExtensionAdditionalDetailsDocsParams{
		timeout: timeout,
	}
}

// NewGetBundleVersionExtensionAdditionalDetailsDocsParamsWithContext creates a new GetBundleVersionExtensionAdditionalDetailsDocsParams object
// with the ability to set a context for a request.
func NewGetBundleVersionExtensionAdditionalDetailsDocsParamsWithContext(ctx context.Context) *GetBundleVersionExtensionAdditionalDetailsDocsParams {
	return &GetBundleVersionExtensionAdditionalDetailsDocsParams{
		Context: ctx,
	}
}

// NewGetBundleVersionExtensionAdditionalDetailsDocsParamsWithHTTPClient creates a new GetBundleVersionExtensionAdditionalDetailsDocsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBundleVersionExtensionAdditionalDetailsDocsParamsWithHTTPClient(client *http.Client) *GetBundleVersionExtensionAdditionalDetailsDocsParams {
	return &GetBundleVersionExtensionAdditionalDetailsDocsParams{
		HTTPClient: client,
	}
}

/*
GetBundleVersionExtensionAdditionalDetailsDocsParams contains all the parameters to send to the API endpoint

	for the get bundle version extension additional details docs operation.

	Typically these are written to a http.Request.
*/
type GetBundleVersionExtensionAdditionalDetailsDocsParams struct {

	/* BundleID.

	   The extension bundle identifier
	*/
	BundleID string

	/* Name.

	   The fully qualified name of the extension
	*/
	Name string

	/* Version.

	   The version of the bundle
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get bundle version extension additional details docs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBundleVersionExtensionAdditionalDetailsDocsParams) WithDefaults() *GetBundleVersionExtensionAdditionalDetailsDocsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get bundle version extension additional details docs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBundleVersionExtensionAdditionalDetailsDocsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get bundle version extension additional details docs params
func (o *GetBundleVersionExtensionAdditionalDetailsDocsParams) WithTimeout(timeout time.Duration) *GetBundleVersionExtensionAdditionalDetailsDocsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bundle version extension additional details docs params
func (o *GetBundleVersionExtensionAdditionalDetailsDocsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bundle version extension additional details docs params
func (o *GetBundleVersionExtensionAdditionalDetailsDocsParams) WithContext(ctx context.Context) *GetBundleVersionExtensionAdditionalDetailsDocsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bundle version extension additional details docs params
func (o *GetBundleVersionExtensionAdditionalDetailsDocsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bundle version extension additional details docs params
func (o *GetBundleVersionExtensionAdditionalDetailsDocsParams) WithHTTPClient(client *http.Client) *GetBundleVersionExtensionAdditionalDetailsDocsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bundle version extension additional details docs params
func (o *GetBundleVersionExtensionAdditionalDetailsDocsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBundleID adds the bundleID to the get bundle version extension additional details docs params
func (o *GetBundleVersionExtensionAdditionalDetailsDocsParams) WithBundleID(bundleID string) *GetBundleVersionExtensionAdditionalDetailsDocsParams {
	o.SetBundleID(bundleID)
	return o
}

// SetBundleID adds the bundleId to the get bundle version extension additional details docs params
func (o *GetBundleVersionExtensionAdditionalDetailsDocsParams) SetBundleID(bundleID string) {
	o.BundleID = bundleID
}

// WithName adds the name to the get bundle version extension additional details docs params
func (o *GetBundleVersionExtensionAdditionalDetailsDocsParams) WithName(name string) *GetBundleVersionExtensionAdditionalDetailsDocsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get bundle version extension additional details docs params
func (o *GetBundleVersionExtensionAdditionalDetailsDocsParams) SetName(name string) {
	o.Name = name
}

// WithVersion adds the version to the get bundle version extension additional details docs params
func (o *GetBundleVersionExtensionAdditionalDetailsDocsParams) WithVersion(version string) *GetBundleVersionExtensionAdditionalDetailsDocsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get bundle version extension additional details docs params
func (o *GetBundleVersionExtensionAdditionalDetailsDocsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetBundleVersionExtensionAdditionalDetailsDocsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bundleId
	if err := r.SetPathParam("bundleId", o.BundleID); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
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

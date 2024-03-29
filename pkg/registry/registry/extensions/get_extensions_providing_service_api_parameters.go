// Code generated by go-swagger; DO NOT EDIT.

package extensions

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

// NewGetExtensionsProvidingServiceAPIParams creates a new GetExtensionsProvidingServiceAPIParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetExtensionsProvidingServiceAPIParams() *GetExtensionsProvidingServiceAPIParams {
	return &GetExtensionsProvidingServiceAPIParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetExtensionsProvidingServiceAPIParamsWithTimeout creates a new GetExtensionsProvidingServiceAPIParams object
// with the ability to set a timeout on a request.
func NewGetExtensionsProvidingServiceAPIParamsWithTimeout(timeout time.Duration) *GetExtensionsProvidingServiceAPIParams {
	return &GetExtensionsProvidingServiceAPIParams{
		timeout: timeout,
	}
}

// NewGetExtensionsProvidingServiceAPIParamsWithContext creates a new GetExtensionsProvidingServiceAPIParams object
// with the ability to set a context for a request.
func NewGetExtensionsProvidingServiceAPIParamsWithContext(ctx context.Context) *GetExtensionsProvidingServiceAPIParams {
	return &GetExtensionsProvidingServiceAPIParams{
		Context: ctx,
	}
}

// NewGetExtensionsProvidingServiceAPIParamsWithHTTPClient creates a new GetExtensionsProvidingServiceAPIParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetExtensionsProvidingServiceAPIParamsWithHTTPClient(client *http.Client) *GetExtensionsProvidingServiceAPIParams {
	return &GetExtensionsProvidingServiceAPIParams{
		HTTPClient: client,
	}
}

/*
GetExtensionsProvidingServiceAPIParams contains all the parameters to send to the API endpoint

	for the get extensions providing service API operation.

	Typically these are written to a http.Request.
*/
type GetExtensionsProvidingServiceAPIParams struct {

	/* ArtifactID.

	   The artifactId of the bundle containing the service API class
	*/
	ArtifactID string

	/* ClassName.

	   The name of the service API class
	*/
	ClassName string

	/* GroupID.

	   The groupId of the bundle containing the service API class
	*/
	GroupID string

	/* Version.

	   The version of the bundle containing the service API class
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get extensions providing service API params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExtensionsProvidingServiceAPIParams) WithDefaults() *GetExtensionsProvidingServiceAPIParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get extensions providing service API params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExtensionsProvidingServiceAPIParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get extensions providing service API params
func (o *GetExtensionsProvidingServiceAPIParams) WithTimeout(timeout time.Duration) *GetExtensionsProvidingServiceAPIParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get extensions providing service API params
func (o *GetExtensionsProvidingServiceAPIParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get extensions providing service API params
func (o *GetExtensionsProvidingServiceAPIParams) WithContext(ctx context.Context) *GetExtensionsProvidingServiceAPIParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get extensions providing service API params
func (o *GetExtensionsProvidingServiceAPIParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get extensions providing service API params
func (o *GetExtensionsProvidingServiceAPIParams) WithHTTPClient(client *http.Client) *GetExtensionsProvidingServiceAPIParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get extensions providing service API params
func (o *GetExtensionsProvidingServiceAPIParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArtifactID adds the artifactID to the get extensions providing service API params
func (o *GetExtensionsProvidingServiceAPIParams) WithArtifactID(artifactID string) *GetExtensionsProvidingServiceAPIParams {
	o.SetArtifactID(artifactID)
	return o
}

// SetArtifactID adds the artifactId to the get extensions providing service API params
func (o *GetExtensionsProvidingServiceAPIParams) SetArtifactID(artifactID string) {
	o.ArtifactID = artifactID
}

// WithClassName adds the className to the get extensions providing service API params
func (o *GetExtensionsProvidingServiceAPIParams) WithClassName(className string) *GetExtensionsProvidingServiceAPIParams {
	o.SetClassName(className)
	return o
}

// SetClassName adds the className to the get extensions providing service API params
func (o *GetExtensionsProvidingServiceAPIParams) SetClassName(className string) {
	o.ClassName = className
}

// WithGroupID adds the groupID to the get extensions providing service API params
func (o *GetExtensionsProvidingServiceAPIParams) WithGroupID(groupID string) *GetExtensionsProvidingServiceAPIParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get extensions providing service API params
func (o *GetExtensionsProvidingServiceAPIParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithVersion adds the version to the get extensions providing service API params
func (o *GetExtensionsProvidingServiceAPIParams) WithVersion(version string) *GetExtensionsProvidingServiceAPIParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get extensions providing service API params
func (o *GetExtensionsProvidingServiceAPIParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetExtensionsProvidingServiceAPIParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param artifactId
	qrArtifactID := o.ArtifactID
	qArtifactID := qrArtifactID
	if qArtifactID != "" {

		if err := r.SetQueryParam("artifactId", qArtifactID); err != nil {
			return err
		}
	}

	// query param className
	qrClassName := o.ClassName
	qClassName := qrClassName
	if qClassName != "" {

		if err := r.SetQueryParam("className", qClassName); err != nil {
			return err
		}
	}

	// query param groupId
	qrGroupID := o.GroupID
	qGroupID := qrGroupID
	if qGroupID != "" {

		if err := r.SetQueryParam("groupId", qGroupID); err != nil {
			return err
		}
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {

		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

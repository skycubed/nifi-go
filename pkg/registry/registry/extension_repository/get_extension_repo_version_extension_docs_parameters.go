// Code generated by go-swagger; DO NOT EDIT.

package extension_repository

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

// NewGetExtensionRepoVersionExtensionDocsParams creates a new GetExtensionRepoVersionExtensionDocsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetExtensionRepoVersionExtensionDocsParams() *GetExtensionRepoVersionExtensionDocsParams {
	return &GetExtensionRepoVersionExtensionDocsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetExtensionRepoVersionExtensionDocsParamsWithTimeout creates a new GetExtensionRepoVersionExtensionDocsParams object
// with the ability to set a timeout on a request.
func NewGetExtensionRepoVersionExtensionDocsParamsWithTimeout(timeout time.Duration) *GetExtensionRepoVersionExtensionDocsParams {
	return &GetExtensionRepoVersionExtensionDocsParams{
		timeout: timeout,
	}
}

// NewGetExtensionRepoVersionExtensionDocsParamsWithContext creates a new GetExtensionRepoVersionExtensionDocsParams object
// with the ability to set a context for a request.
func NewGetExtensionRepoVersionExtensionDocsParamsWithContext(ctx context.Context) *GetExtensionRepoVersionExtensionDocsParams {
	return &GetExtensionRepoVersionExtensionDocsParams{
		Context: ctx,
	}
}

// NewGetExtensionRepoVersionExtensionDocsParamsWithHTTPClient creates a new GetExtensionRepoVersionExtensionDocsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetExtensionRepoVersionExtensionDocsParamsWithHTTPClient(client *http.Client) *GetExtensionRepoVersionExtensionDocsParams {
	return &GetExtensionRepoVersionExtensionDocsParams{
		HTTPClient: client,
	}
}

/*
GetExtensionRepoVersionExtensionDocsParams contains all the parameters to send to the API endpoint

	for the get extension repo version extension docs operation.

	Typically these are written to a http.Request.
*/
type GetExtensionRepoVersionExtensionDocsParams struct {

	/* ArtifactID.

	   The artifact identifier
	*/
	ArtifactID string

	/* BucketName.

	   The bucket name
	*/
	BucketName string

	/* GroupID.

	   The group identifier
	*/
	GroupID string

	/* Name.

	   The fully qualified name of the extension
	*/
	Name string

	/* Version.

	   The version
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get extension repo version extension docs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExtensionRepoVersionExtensionDocsParams) WithDefaults() *GetExtensionRepoVersionExtensionDocsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get extension repo version extension docs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExtensionRepoVersionExtensionDocsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get extension repo version extension docs params
func (o *GetExtensionRepoVersionExtensionDocsParams) WithTimeout(timeout time.Duration) *GetExtensionRepoVersionExtensionDocsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get extension repo version extension docs params
func (o *GetExtensionRepoVersionExtensionDocsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get extension repo version extension docs params
func (o *GetExtensionRepoVersionExtensionDocsParams) WithContext(ctx context.Context) *GetExtensionRepoVersionExtensionDocsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get extension repo version extension docs params
func (o *GetExtensionRepoVersionExtensionDocsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get extension repo version extension docs params
func (o *GetExtensionRepoVersionExtensionDocsParams) WithHTTPClient(client *http.Client) *GetExtensionRepoVersionExtensionDocsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get extension repo version extension docs params
func (o *GetExtensionRepoVersionExtensionDocsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArtifactID adds the artifactID to the get extension repo version extension docs params
func (o *GetExtensionRepoVersionExtensionDocsParams) WithArtifactID(artifactID string) *GetExtensionRepoVersionExtensionDocsParams {
	o.SetArtifactID(artifactID)
	return o
}

// SetArtifactID adds the artifactId to the get extension repo version extension docs params
func (o *GetExtensionRepoVersionExtensionDocsParams) SetArtifactID(artifactID string) {
	o.ArtifactID = artifactID
}

// WithBucketName adds the bucketName to the get extension repo version extension docs params
func (o *GetExtensionRepoVersionExtensionDocsParams) WithBucketName(bucketName string) *GetExtensionRepoVersionExtensionDocsParams {
	o.SetBucketName(bucketName)
	return o
}

// SetBucketName adds the bucketName to the get extension repo version extension docs params
func (o *GetExtensionRepoVersionExtensionDocsParams) SetBucketName(bucketName string) {
	o.BucketName = bucketName
}

// WithGroupID adds the groupID to the get extension repo version extension docs params
func (o *GetExtensionRepoVersionExtensionDocsParams) WithGroupID(groupID string) *GetExtensionRepoVersionExtensionDocsParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get extension repo version extension docs params
func (o *GetExtensionRepoVersionExtensionDocsParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithName adds the name to the get extension repo version extension docs params
func (o *GetExtensionRepoVersionExtensionDocsParams) WithName(name string) *GetExtensionRepoVersionExtensionDocsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get extension repo version extension docs params
func (o *GetExtensionRepoVersionExtensionDocsParams) SetName(name string) {
	o.Name = name
}

// WithVersion adds the version to the get extension repo version extension docs params
func (o *GetExtensionRepoVersionExtensionDocsParams) WithVersion(version string) *GetExtensionRepoVersionExtensionDocsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get extension repo version extension docs params
func (o *GetExtensionRepoVersionExtensionDocsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetExtensionRepoVersionExtensionDocsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param artifactId
	if err := r.SetPathParam("artifactId", o.ArtifactID); err != nil {
		return err
	}

	// path param bucketName
	if err := r.SetPathParam("bucketName", o.BucketName); err != nil {
		return err
	}

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
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

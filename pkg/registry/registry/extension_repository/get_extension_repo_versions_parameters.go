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

// NewGetExtensionRepoVersionsParams creates a new GetExtensionRepoVersionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetExtensionRepoVersionsParams() *GetExtensionRepoVersionsParams {
	return &GetExtensionRepoVersionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetExtensionRepoVersionsParamsWithTimeout creates a new GetExtensionRepoVersionsParams object
// with the ability to set a timeout on a request.
func NewGetExtensionRepoVersionsParamsWithTimeout(timeout time.Duration) *GetExtensionRepoVersionsParams {
	return &GetExtensionRepoVersionsParams{
		timeout: timeout,
	}
}

// NewGetExtensionRepoVersionsParamsWithContext creates a new GetExtensionRepoVersionsParams object
// with the ability to set a context for a request.
func NewGetExtensionRepoVersionsParamsWithContext(ctx context.Context) *GetExtensionRepoVersionsParams {
	return &GetExtensionRepoVersionsParams{
		Context: ctx,
	}
}

// NewGetExtensionRepoVersionsParamsWithHTTPClient creates a new GetExtensionRepoVersionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetExtensionRepoVersionsParamsWithHTTPClient(client *http.Client) *GetExtensionRepoVersionsParams {
	return &GetExtensionRepoVersionsParams{
		HTTPClient: client,
	}
}

/* GetExtensionRepoVersionsParams contains all the parameters to send to the API endpoint
   for the get extension repo versions operation.

   Typically these are written to a http.Request.
*/
type GetExtensionRepoVersionsParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get extension repo versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExtensionRepoVersionsParams) WithDefaults() *GetExtensionRepoVersionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get extension repo versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExtensionRepoVersionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get extension repo versions params
func (o *GetExtensionRepoVersionsParams) WithTimeout(timeout time.Duration) *GetExtensionRepoVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get extension repo versions params
func (o *GetExtensionRepoVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get extension repo versions params
func (o *GetExtensionRepoVersionsParams) WithContext(ctx context.Context) *GetExtensionRepoVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get extension repo versions params
func (o *GetExtensionRepoVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get extension repo versions params
func (o *GetExtensionRepoVersionsParams) WithHTTPClient(client *http.Client) *GetExtensionRepoVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get extension repo versions params
func (o *GetExtensionRepoVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArtifactID adds the artifactID to the get extension repo versions params
func (o *GetExtensionRepoVersionsParams) WithArtifactID(artifactID string) *GetExtensionRepoVersionsParams {
	o.SetArtifactID(artifactID)
	return o
}

// SetArtifactID adds the artifactId to the get extension repo versions params
func (o *GetExtensionRepoVersionsParams) SetArtifactID(artifactID string) {
	o.ArtifactID = artifactID
}

// WithBucketName adds the bucketName to the get extension repo versions params
func (o *GetExtensionRepoVersionsParams) WithBucketName(bucketName string) *GetExtensionRepoVersionsParams {
	o.SetBucketName(bucketName)
	return o
}

// SetBucketName adds the bucketName to the get extension repo versions params
func (o *GetExtensionRepoVersionsParams) SetBucketName(bucketName string) {
	o.BucketName = bucketName
}

// WithGroupID adds the groupID to the get extension repo versions params
func (o *GetExtensionRepoVersionsParams) WithGroupID(groupID string) *GetExtensionRepoVersionsParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get extension repo versions params
func (o *GetExtensionRepoVersionsParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetExtensionRepoVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

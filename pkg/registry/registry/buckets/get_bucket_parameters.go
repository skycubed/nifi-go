// Code generated by go-swagger; DO NOT EDIT.

package buckets

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

// NewGetBucketParams creates a new GetBucketParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBucketParams() *GetBucketParams {
	return &GetBucketParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBucketParamsWithTimeout creates a new GetBucketParams object
// with the ability to set a timeout on a request.
func NewGetBucketParamsWithTimeout(timeout time.Duration) *GetBucketParams {
	return &GetBucketParams{
		timeout: timeout,
	}
}

// NewGetBucketParamsWithContext creates a new GetBucketParams object
// with the ability to set a context for a request.
func NewGetBucketParamsWithContext(ctx context.Context) *GetBucketParams {
	return &GetBucketParams{
		Context: ctx,
	}
}

// NewGetBucketParamsWithHTTPClient creates a new GetBucketParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBucketParamsWithHTTPClient(client *http.Client) *GetBucketParams {
	return &GetBucketParams{
		HTTPClient: client,
	}
}

/* GetBucketParams contains all the parameters to send to the API endpoint
   for the get bucket operation.

   Typically these are written to a http.Request.
*/
type GetBucketParams struct {

	/* BucketID.

	   The bucket identifier
	*/
	BucketID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get bucket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBucketParams) WithDefaults() *GetBucketParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get bucket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBucketParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get bucket params
func (o *GetBucketParams) WithTimeout(timeout time.Duration) *GetBucketParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bucket params
func (o *GetBucketParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bucket params
func (o *GetBucketParams) WithContext(ctx context.Context) *GetBucketParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bucket params
func (o *GetBucketParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bucket params
func (o *GetBucketParams) WithHTTPClient(client *http.Client) *GetBucketParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bucket params
func (o *GetBucketParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketID adds the bucketID to the get bucket params
func (o *GetBucketParams) WithBucketID(bucketID string) *GetBucketParams {
	o.SetBucketID(bucketID)
	return o
}

// SetBucketID adds the bucketId to the get bucket params
func (o *GetBucketParams) SetBucketID(bucketID string) {
	o.BucketID = bucketID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBucketParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bucketId
	if err := r.SetPathParam("bucketId", o.BucketID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
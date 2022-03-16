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

// NewGetAvailableBucketFieldsParams creates a new GetAvailableBucketFieldsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAvailableBucketFieldsParams() *GetAvailableBucketFieldsParams {
	return &GetAvailableBucketFieldsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAvailableBucketFieldsParamsWithTimeout creates a new GetAvailableBucketFieldsParams object
// with the ability to set a timeout on a request.
func NewGetAvailableBucketFieldsParamsWithTimeout(timeout time.Duration) *GetAvailableBucketFieldsParams {
	return &GetAvailableBucketFieldsParams{
		timeout: timeout,
	}
}

// NewGetAvailableBucketFieldsParamsWithContext creates a new GetAvailableBucketFieldsParams object
// with the ability to set a context for a request.
func NewGetAvailableBucketFieldsParamsWithContext(ctx context.Context) *GetAvailableBucketFieldsParams {
	return &GetAvailableBucketFieldsParams{
		Context: ctx,
	}
}

// NewGetAvailableBucketFieldsParamsWithHTTPClient creates a new GetAvailableBucketFieldsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAvailableBucketFieldsParamsWithHTTPClient(client *http.Client) *GetAvailableBucketFieldsParams {
	return &GetAvailableBucketFieldsParams{
		HTTPClient: client,
	}
}

/* GetAvailableBucketFieldsParams contains all the parameters to send to the API endpoint
   for the get available bucket fields operation.

   Typically these are written to a http.Request.
*/
type GetAvailableBucketFieldsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get available bucket fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAvailableBucketFieldsParams) WithDefaults() *GetAvailableBucketFieldsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get available bucket fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAvailableBucketFieldsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get available bucket fields params
func (o *GetAvailableBucketFieldsParams) WithTimeout(timeout time.Duration) *GetAvailableBucketFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get available bucket fields params
func (o *GetAvailableBucketFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get available bucket fields params
func (o *GetAvailableBucketFieldsParams) WithContext(ctx context.Context) *GetAvailableBucketFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get available bucket fields params
func (o *GetAvailableBucketFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get available bucket fields params
func (o *GetAvailableBucketFieldsParams) WithHTTPClient(client *http.Client) *GetAvailableBucketFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get available bucket fields params
func (o *GetAvailableBucketFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAvailableBucketFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package access

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

// NewGetAccessTokenExpirationParams creates a new GetAccessTokenExpirationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccessTokenExpirationParams() *GetAccessTokenExpirationParams {
	return &GetAccessTokenExpirationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccessTokenExpirationParamsWithTimeout creates a new GetAccessTokenExpirationParams object
// with the ability to set a timeout on a request.
func NewGetAccessTokenExpirationParamsWithTimeout(timeout time.Duration) *GetAccessTokenExpirationParams {
	return &GetAccessTokenExpirationParams{
		timeout: timeout,
	}
}

// NewGetAccessTokenExpirationParamsWithContext creates a new GetAccessTokenExpirationParams object
// with the ability to set a context for a request.
func NewGetAccessTokenExpirationParamsWithContext(ctx context.Context) *GetAccessTokenExpirationParams {
	return &GetAccessTokenExpirationParams{
		Context: ctx,
	}
}

// NewGetAccessTokenExpirationParamsWithHTTPClient creates a new GetAccessTokenExpirationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccessTokenExpirationParamsWithHTTPClient(client *http.Client) *GetAccessTokenExpirationParams {
	return &GetAccessTokenExpirationParams{
		HTTPClient: client,
	}
}

/*
GetAccessTokenExpirationParams contains all the parameters to send to the API endpoint

	for the get access token expiration operation.

	Typically these are written to a http.Request.
*/
type GetAccessTokenExpirationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get access token expiration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccessTokenExpirationParams) WithDefaults() *GetAccessTokenExpirationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get access token expiration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccessTokenExpirationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get access token expiration params
func (o *GetAccessTokenExpirationParams) WithTimeout(timeout time.Duration) *GetAccessTokenExpirationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get access token expiration params
func (o *GetAccessTokenExpirationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get access token expiration params
func (o *GetAccessTokenExpirationParams) WithContext(ctx context.Context) *GetAccessTokenExpirationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get access token expiration params
func (o *GetAccessTokenExpirationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get access token expiration params
func (o *GetAccessTokenExpirationParams) WithHTTPClient(client *http.Client) *GetAccessTokenExpirationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get access token expiration params
func (o *GetAccessTokenExpirationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccessTokenExpirationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
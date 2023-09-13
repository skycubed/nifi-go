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

// NewTestIdentityProviderRecognizesCredentialsFormatParams creates a new TestIdentityProviderRecognizesCredentialsFormatParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTestIdentityProviderRecognizesCredentialsFormatParams() *TestIdentityProviderRecognizesCredentialsFormatParams {
	return &TestIdentityProviderRecognizesCredentialsFormatParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTestIdentityProviderRecognizesCredentialsFormatParamsWithTimeout creates a new TestIdentityProviderRecognizesCredentialsFormatParams object
// with the ability to set a timeout on a request.
func NewTestIdentityProviderRecognizesCredentialsFormatParamsWithTimeout(timeout time.Duration) *TestIdentityProviderRecognizesCredentialsFormatParams {
	return &TestIdentityProviderRecognizesCredentialsFormatParams{
		timeout: timeout,
	}
}

// NewTestIdentityProviderRecognizesCredentialsFormatParamsWithContext creates a new TestIdentityProviderRecognizesCredentialsFormatParams object
// with the ability to set a context for a request.
func NewTestIdentityProviderRecognizesCredentialsFormatParamsWithContext(ctx context.Context) *TestIdentityProviderRecognizesCredentialsFormatParams {
	return &TestIdentityProviderRecognizesCredentialsFormatParams{
		Context: ctx,
	}
}

// NewTestIdentityProviderRecognizesCredentialsFormatParamsWithHTTPClient creates a new TestIdentityProviderRecognizesCredentialsFormatParams object
// with the ability to set a custom HTTPClient for a request.
func NewTestIdentityProviderRecognizesCredentialsFormatParamsWithHTTPClient(client *http.Client) *TestIdentityProviderRecognizesCredentialsFormatParams {
	return &TestIdentityProviderRecognizesCredentialsFormatParams{
		HTTPClient: client,
	}
}

/*
TestIdentityProviderRecognizesCredentialsFormatParams contains all the parameters to send to the API endpoint

	for the test identity provider recognizes credentials format operation.

	Typically these are written to a http.Request.
*/
type TestIdentityProviderRecognizesCredentialsFormatParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the test identity provider recognizes credentials format params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestIdentityProviderRecognizesCredentialsFormatParams) WithDefaults() *TestIdentityProviderRecognizesCredentialsFormatParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the test identity provider recognizes credentials format params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestIdentityProviderRecognizesCredentialsFormatParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the test identity provider recognizes credentials format params
func (o *TestIdentityProviderRecognizesCredentialsFormatParams) WithTimeout(timeout time.Duration) *TestIdentityProviderRecognizesCredentialsFormatParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test identity provider recognizes credentials format params
func (o *TestIdentityProviderRecognizesCredentialsFormatParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test identity provider recognizes credentials format params
func (o *TestIdentityProviderRecognizesCredentialsFormatParams) WithContext(ctx context.Context) *TestIdentityProviderRecognizesCredentialsFormatParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test identity provider recognizes credentials format params
func (o *TestIdentityProviderRecognizesCredentialsFormatParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test identity provider recognizes credentials format params
func (o *TestIdentityProviderRecognizesCredentialsFormatParams) WithHTTPClient(client *http.Client) *TestIdentityProviderRecognizesCredentialsFormatParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test identity provider recognizes credentials format params
func (o *TestIdentityProviderRecognizesCredentialsFormatParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *TestIdentityProviderRecognizesCredentialsFormatParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package accessoidc

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

// NewOidcLogoutParams creates a new OidcLogoutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOidcLogoutParams() *OidcLogoutParams {
	return &OidcLogoutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOidcLogoutParamsWithTimeout creates a new OidcLogoutParams object
// with the ability to set a timeout on a request.
func NewOidcLogoutParamsWithTimeout(timeout time.Duration) *OidcLogoutParams {
	return &OidcLogoutParams{
		timeout: timeout,
	}
}

// NewOidcLogoutParamsWithContext creates a new OidcLogoutParams object
// with the ability to set a context for a request.
func NewOidcLogoutParamsWithContext(ctx context.Context) *OidcLogoutParams {
	return &OidcLogoutParams{
		Context: ctx,
	}
}

// NewOidcLogoutParamsWithHTTPClient creates a new OidcLogoutParams object
// with the ability to set a custom HTTPClient for a request.
func NewOidcLogoutParamsWithHTTPClient(client *http.Client) *OidcLogoutParams {
	return &OidcLogoutParams{
		HTTPClient: client,
	}
}

/* OidcLogoutParams contains all the parameters to send to the API endpoint
   for the oidc logout operation.

   Typically these are written to a http.Request.
*/
type OidcLogoutParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the oidc logout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OidcLogoutParams) WithDefaults() *OidcLogoutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the oidc logout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OidcLogoutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the oidc logout params
func (o *OidcLogoutParams) WithTimeout(timeout time.Duration) *OidcLogoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the oidc logout params
func (o *OidcLogoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the oidc logout params
func (o *OidcLogoutParams) WithContext(ctx context.Context) *OidcLogoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the oidc logout params
func (o *OidcLogoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the oidc logout params
func (o *OidcLogoutParams) WithHTTPClient(client *http.Client) *OidcLogoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the oidc logout params
func (o *OidcLogoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *OidcLogoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
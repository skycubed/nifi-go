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

// NewOidcExchangeParams creates a new OidcExchangeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOidcExchangeParams() *OidcExchangeParams {
	return &OidcExchangeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOidcExchangeParamsWithTimeout creates a new OidcExchangeParams object
// with the ability to set a timeout on a request.
func NewOidcExchangeParamsWithTimeout(timeout time.Duration) *OidcExchangeParams {
	return &OidcExchangeParams{
		timeout: timeout,
	}
}

// NewOidcExchangeParamsWithContext creates a new OidcExchangeParams object
// with the ability to set a context for a request.
func NewOidcExchangeParamsWithContext(ctx context.Context) *OidcExchangeParams {
	return &OidcExchangeParams{
		Context: ctx,
	}
}

// NewOidcExchangeParamsWithHTTPClient creates a new OidcExchangeParams object
// with the ability to set a custom HTTPClient for a request.
func NewOidcExchangeParamsWithHTTPClient(client *http.Client) *OidcExchangeParams {
	return &OidcExchangeParams{
		HTTPClient: client,
	}
}

/*
OidcExchangeParams contains all the parameters to send to the API endpoint

	for the oidc exchange operation.

	Typically these are written to a http.Request.
*/
type OidcExchangeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the oidc exchange params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OidcExchangeParams) WithDefaults() *OidcExchangeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the oidc exchange params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OidcExchangeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the oidc exchange params
func (o *OidcExchangeParams) WithTimeout(timeout time.Duration) *OidcExchangeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the oidc exchange params
func (o *OidcExchangeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the oidc exchange params
func (o *OidcExchangeParams) WithContext(ctx context.Context) *OidcExchangeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the oidc exchange params
func (o *OidcExchangeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the oidc exchange params
func (o *OidcExchangeParams) WithHTTPClient(client *http.Client) *OidcExchangeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the oidc exchange params
func (o *OidcExchangeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *OidcExchangeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

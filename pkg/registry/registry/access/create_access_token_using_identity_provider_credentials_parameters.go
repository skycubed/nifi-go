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

// NewCreateAccessTokenUsingIdentityProviderCredentialsParams creates a new CreateAccessTokenUsingIdentityProviderCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAccessTokenUsingIdentityProviderCredentialsParams() *CreateAccessTokenUsingIdentityProviderCredentialsParams {
	return &CreateAccessTokenUsingIdentityProviderCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAccessTokenUsingIdentityProviderCredentialsParamsWithTimeout creates a new CreateAccessTokenUsingIdentityProviderCredentialsParams object
// with the ability to set a timeout on a request.
func NewCreateAccessTokenUsingIdentityProviderCredentialsParamsWithTimeout(timeout time.Duration) *CreateAccessTokenUsingIdentityProviderCredentialsParams {
	return &CreateAccessTokenUsingIdentityProviderCredentialsParams{
		timeout: timeout,
	}
}

// NewCreateAccessTokenUsingIdentityProviderCredentialsParamsWithContext creates a new CreateAccessTokenUsingIdentityProviderCredentialsParams object
// with the ability to set a context for a request.
func NewCreateAccessTokenUsingIdentityProviderCredentialsParamsWithContext(ctx context.Context) *CreateAccessTokenUsingIdentityProviderCredentialsParams {
	return &CreateAccessTokenUsingIdentityProviderCredentialsParams{
		Context: ctx,
	}
}

// NewCreateAccessTokenUsingIdentityProviderCredentialsParamsWithHTTPClient creates a new CreateAccessTokenUsingIdentityProviderCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAccessTokenUsingIdentityProviderCredentialsParamsWithHTTPClient(client *http.Client) *CreateAccessTokenUsingIdentityProviderCredentialsParams {
	return &CreateAccessTokenUsingIdentityProviderCredentialsParams{
		HTTPClient: client,
	}
}

/* CreateAccessTokenUsingIdentityProviderCredentialsParams contains all the parameters to send to the API endpoint
   for the create access token using identity provider credentials operation.

   Typically these are written to a http.Request.
*/
type CreateAccessTokenUsingIdentityProviderCredentialsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create access token using identity provider credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAccessTokenUsingIdentityProviderCredentialsParams) WithDefaults() *CreateAccessTokenUsingIdentityProviderCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create access token using identity provider credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAccessTokenUsingIdentityProviderCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create access token using identity provider credentials params
func (o *CreateAccessTokenUsingIdentityProviderCredentialsParams) WithTimeout(timeout time.Duration) *CreateAccessTokenUsingIdentityProviderCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create access token using identity provider credentials params
func (o *CreateAccessTokenUsingIdentityProviderCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create access token using identity provider credentials params
func (o *CreateAccessTokenUsingIdentityProviderCredentialsParams) WithContext(ctx context.Context) *CreateAccessTokenUsingIdentityProviderCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create access token using identity provider credentials params
func (o *CreateAccessTokenUsingIdentityProviderCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create access token using identity provider credentials params
func (o *CreateAccessTokenUsingIdentityProviderCredentialsParams) WithHTTPClient(client *http.Client) *CreateAccessTokenUsingIdentityProviderCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create access token using identity provider credentials params
func (o *CreateAccessTokenUsingIdentityProviderCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAccessTokenUsingIdentityProviderCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

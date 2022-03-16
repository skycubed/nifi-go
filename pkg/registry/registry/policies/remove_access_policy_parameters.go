// Code generated by go-swagger; DO NOT EDIT.

package policies

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

// NewRemoveAccessPolicyParams creates a new RemoveAccessPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveAccessPolicyParams() *RemoveAccessPolicyParams {
	return &RemoveAccessPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveAccessPolicyParamsWithTimeout creates a new RemoveAccessPolicyParams object
// with the ability to set a timeout on a request.
func NewRemoveAccessPolicyParamsWithTimeout(timeout time.Duration) *RemoveAccessPolicyParams {
	return &RemoveAccessPolicyParams{
		timeout: timeout,
	}
}

// NewRemoveAccessPolicyParamsWithContext creates a new RemoveAccessPolicyParams object
// with the ability to set a context for a request.
func NewRemoveAccessPolicyParamsWithContext(ctx context.Context) *RemoveAccessPolicyParams {
	return &RemoveAccessPolicyParams{
		Context: ctx,
	}
}

// NewRemoveAccessPolicyParamsWithHTTPClient creates a new RemoveAccessPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveAccessPolicyParamsWithHTTPClient(client *http.Client) *RemoveAccessPolicyParams {
	return &RemoveAccessPolicyParams{
		HTTPClient: client,
	}
}

/* RemoveAccessPolicyParams contains all the parameters to send to the API endpoint
   for the remove access policy operation.

   Typically these are written to a http.Request.
*/
type RemoveAccessPolicyParams struct {

	/* ClientID.

	   If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response.
	*/
	ClientID *string

	/* ID.

	   The access policy id.
	*/
	ID string

	/* Version.

	   The version is used to verify the client is working with the latest version of the entity.
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove access policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveAccessPolicyParams) WithDefaults() *RemoveAccessPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove access policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveAccessPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove access policy params
func (o *RemoveAccessPolicyParams) WithTimeout(timeout time.Duration) *RemoveAccessPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove access policy params
func (o *RemoveAccessPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove access policy params
func (o *RemoveAccessPolicyParams) WithContext(ctx context.Context) *RemoveAccessPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove access policy params
func (o *RemoveAccessPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove access policy params
func (o *RemoveAccessPolicyParams) WithHTTPClient(client *http.Client) *RemoveAccessPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove access policy params
func (o *RemoveAccessPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the remove access policy params
func (o *RemoveAccessPolicyParams) WithClientID(clientID *string) *RemoveAccessPolicyParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the remove access policy params
func (o *RemoveAccessPolicyParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithID adds the id to the remove access policy params
func (o *RemoveAccessPolicyParams) WithID(id string) *RemoveAccessPolicyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the remove access policy params
func (o *RemoveAccessPolicyParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the remove access policy params
func (o *RemoveAccessPolicyParams) WithVersion(version string) *RemoveAccessPolicyParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the remove access policy params
func (o *RemoveAccessPolicyParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveAccessPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientID != nil {

		// query param clientId
		var qrClientID string

		if o.ClientID != nil {
			qrClientID = *o.ClientID
		}
		qClientID := qrClientID
		if qClientID != "" {

			if err := r.SetQueryParam("clientId", qClientID); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
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
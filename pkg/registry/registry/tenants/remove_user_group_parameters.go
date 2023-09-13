// Code generated by go-swagger; DO NOT EDIT.

package tenants

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

// NewRemoveUserGroupParams creates a new RemoveUserGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveUserGroupParams() *RemoveUserGroupParams {
	return &RemoveUserGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveUserGroupParamsWithTimeout creates a new RemoveUserGroupParams object
// with the ability to set a timeout on a request.
func NewRemoveUserGroupParamsWithTimeout(timeout time.Duration) *RemoveUserGroupParams {
	return &RemoveUserGroupParams{
		timeout: timeout,
	}
}

// NewRemoveUserGroupParamsWithContext creates a new RemoveUserGroupParams object
// with the ability to set a context for a request.
func NewRemoveUserGroupParamsWithContext(ctx context.Context) *RemoveUserGroupParams {
	return &RemoveUserGroupParams{
		Context: ctx,
	}
}

// NewRemoveUserGroupParamsWithHTTPClient creates a new RemoveUserGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveUserGroupParamsWithHTTPClient(client *http.Client) *RemoveUserGroupParams {
	return &RemoveUserGroupParams{
		HTTPClient: client,
	}
}

/*
RemoveUserGroupParams contains all the parameters to send to the API endpoint

	for the remove user group operation.

	Typically these are written to a http.Request.
*/
type RemoveUserGroupParams struct {

	/* ClientID.

	   If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response.
	*/
	ClientID *string

	/* ID.

	   The user group id.
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

// WithDefaults hydrates default values in the remove user group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveUserGroupParams) WithDefaults() *RemoveUserGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove user group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveUserGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove user group params
func (o *RemoveUserGroupParams) WithTimeout(timeout time.Duration) *RemoveUserGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove user group params
func (o *RemoveUserGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove user group params
func (o *RemoveUserGroupParams) WithContext(ctx context.Context) *RemoveUserGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove user group params
func (o *RemoveUserGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove user group params
func (o *RemoveUserGroupParams) WithHTTPClient(client *http.Client) *RemoveUserGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove user group params
func (o *RemoveUserGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the remove user group params
func (o *RemoveUserGroupParams) WithClientID(clientID *string) *RemoveUserGroupParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the remove user group params
func (o *RemoveUserGroupParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithID adds the id to the remove user group params
func (o *RemoveUserGroupParams) WithID(id string) *RemoveUserGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the remove user group params
func (o *RemoveUserGroupParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the remove user group params
func (o *RemoveUserGroupParams) WithVersion(version string) *RemoveUserGroupParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the remove user group params
func (o *RemoveUserGroupParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveUserGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Code generated by go-swagger; DO NOT EDIT.

package controller

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
	"github.com/go-openapi/swag"
)

// NewDeleteFlowRegistryClientParams creates a new DeleteFlowRegistryClientParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteFlowRegistryClientParams() *DeleteFlowRegistryClientParams {
	return &DeleteFlowRegistryClientParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFlowRegistryClientParamsWithTimeout creates a new DeleteFlowRegistryClientParams object
// with the ability to set a timeout on a request.
func NewDeleteFlowRegistryClientParamsWithTimeout(timeout time.Duration) *DeleteFlowRegistryClientParams {
	return &DeleteFlowRegistryClientParams{
		timeout: timeout,
	}
}

// NewDeleteFlowRegistryClientParamsWithContext creates a new DeleteFlowRegistryClientParams object
// with the ability to set a context for a request.
func NewDeleteFlowRegistryClientParamsWithContext(ctx context.Context) *DeleteFlowRegistryClientParams {
	return &DeleteFlowRegistryClientParams{
		Context: ctx,
	}
}

// NewDeleteFlowRegistryClientParamsWithHTTPClient creates a new DeleteFlowRegistryClientParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteFlowRegistryClientParamsWithHTTPClient(client *http.Client) *DeleteFlowRegistryClientParams {
	return &DeleteFlowRegistryClientParams{
		HTTPClient: client,
	}
}

/*
DeleteFlowRegistryClientParams contains all the parameters to send to the API endpoint

	for the delete flow registry client operation.

	Typically these are written to a http.Request.
*/
type DeleteFlowRegistryClientParams struct {

	/* ClientID.

	   If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response.
	*/
	ClientID *string

	/* DisconnectedNodeAcknowledged.

	   Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	*/
	DisconnectedNodeAcknowledged *bool

	/* ID.

	   The flow registry client id.
	*/
	ID string

	/* Version.

	   The revision is used to verify the client is working with the latest version of the flow.
	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete flow registry client params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteFlowRegistryClientParams) WithDefaults() *DeleteFlowRegistryClientParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete flow registry client params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteFlowRegistryClientParams) SetDefaults() {
	var (
		disconnectedNodeAcknowledgedDefault = bool(false)
	)

	val := DeleteFlowRegistryClientParams{
		DisconnectedNodeAcknowledged: &disconnectedNodeAcknowledgedDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete flow registry client params
func (o *DeleteFlowRegistryClientParams) WithTimeout(timeout time.Duration) *DeleteFlowRegistryClientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete flow registry client params
func (o *DeleteFlowRegistryClientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete flow registry client params
func (o *DeleteFlowRegistryClientParams) WithContext(ctx context.Context) *DeleteFlowRegistryClientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete flow registry client params
func (o *DeleteFlowRegistryClientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete flow registry client params
func (o *DeleteFlowRegistryClientParams) WithHTTPClient(client *http.Client) *DeleteFlowRegistryClientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete flow registry client params
func (o *DeleteFlowRegistryClientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the delete flow registry client params
func (o *DeleteFlowRegistryClientParams) WithClientID(clientID *string) *DeleteFlowRegistryClientParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the delete flow registry client params
func (o *DeleteFlowRegistryClientParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithDisconnectedNodeAcknowledged adds the disconnectedNodeAcknowledged to the delete flow registry client params
func (o *DeleteFlowRegistryClientParams) WithDisconnectedNodeAcknowledged(disconnectedNodeAcknowledged *bool) *DeleteFlowRegistryClientParams {
	o.SetDisconnectedNodeAcknowledged(disconnectedNodeAcknowledged)
	return o
}

// SetDisconnectedNodeAcknowledged adds the disconnectedNodeAcknowledged to the delete flow registry client params
func (o *DeleteFlowRegistryClientParams) SetDisconnectedNodeAcknowledged(disconnectedNodeAcknowledged *bool) {
	o.DisconnectedNodeAcknowledged = disconnectedNodeAcknowledged
}

// WithID adds the id to the delete flow registry client params
func (o *DeleteFlowRegistryClientParams) WithID(id string) *DeleteFlowRegistryClientParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete flow registry client params
func (o *DeleteFlowRegistryClientParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the delete flow registry client params
func (o *DeleteFlowRegistryClientParams) WithVersion(version *string) *DeleteFlowRegistryClientParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete flow registry client params
func (o *DeleteFlowRegistryClientParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFlowRegistryClientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DisconnectedNodeAcknowledged != nil {

		// query param disconnectedNodeAcknowledged
		var qrDisconnectedNodeAcknowledged bool

		if o.DisconnectedNodeAcknowledged != nil {
			qrDisconnectedNodeAcknowledged = *o.DisconnectedNodeAcknowledged
		}
		qDisconnectedNodeAcknowledged := swag.FormatBool(qrDisconnectedNodeAcknowledged)
		if qDisconnectedNodeAcknowledged != "" {

			if err := r.SetQueryParam("disconnectedNodeAcknowledged", qDisconnectedNodeAcknowledged); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion string

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package input_ports

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

// NewRemoveInputPortParams creates a new RemoveInputPortParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveInputPortParams() *RemoveInputPortParams {
	return &RemoveInputPortParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveInputPortParamsWithTimeout creates a new RemoveInputPortParams object
// with the ability to set a timeout on a request.
func NewRemoveInputPortParamsWithTimeout(timeout time.Duration) *RemoveInputPortParams {
	return &RemoveInputPortParams{
		timeout: timeout,
	}
}

// NewRemoveInputPortParamsWithContext creates a new RemoveInputPortParams object
// with the ability to set a context for a request.
func NewRemoveInputPortParamsWithContext(ctx context.Context) *RemoveInputPortParams {
	return &RemoveInputPortParams{
		Context: ctx,
	}
}

// NewRemoveInputPortParamsWithHTTPClient creates a new RemoveInputPortParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveInputPortParamsWithHTTPClient(client *http.Client) *RemoveInputPortParams {
	return &RemoveInputPortParams{
		HTTPClient: client,
	}
}

/*
RemoveInputPortParams contains all the parameters to send to the API endpoint

	for the remove input port operation.

	Typically these are written to a http.Request.
*/
type RemoveInputPortParams struct {

	/* ClientID.

	   If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response.
	*/
	ClientID *string

	/* DisconnectedNodeAcknowledged.

	   Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	*/
	DisconnectedNodeAcknowledged *bool

	/* ID.

	   The input port id.
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

// WithDefaults hydrates default values in the remove input port params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveInputPortParams) WithDefaults() *RemoveInputPortParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove input port params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveInputPortParams) SetDefaults() {
	var (
		disconnectedNodeAcknowledgedDefault = bool(false)
	)

	val := RemoveInputPortParams{
		DisconnectedNodeAcknowledged: &disconnectedNodeAcknowledgedDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the remove input port params
func (o *RemoveInputPortParams) WithTimeout(timeout time.Duration) *RemoveInputPortParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove input port params
func (o *RemoveInputPortParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove input port params
func (o *RemoveInputPortParams) WithContext(ctx context.Context) *RemoveInputPortParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove input port params
func (o *RemoveInputPortParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove input port params
func (o *RemoveInputPortParams) WithHTTPClient(client *http.Client) *RemoveInputPortParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove input port params
func (o *RemoveInputPortParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the remove input port params
func (o *RemoveInputPortParams) WithClientID(clientID *string) *RemoveInputPortParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the remove input port params
func (o *RemoveInputPortParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithDisconnectedNodeAcknowledged adds the disconnectedNodeAcknowledged to the remove input port params
func (o *RemoveInputPortParams) WithDisconnectedNodeAcknowledged(disconnectedNodeAcknowledged *bool) *RemoveInputPortParams {
	o.SetDisconnectedNodeAcknowledged(disconnectedNodeAcknowledged)
	return o
}

// SetDisconnectedNodeAcknowledged adds the disconnectedNodeAcknowledged to the remove input port params
func (o *RemoveInputPortParams) SetDisconnectedNodeAcknowledged(disconnectedNodeAcknowledged *bool) {
	o.DisconnectedNodeAcknowledged = disconnectedNodeAcknowledged
}

// WithID adds the id to the remove input port params
func (o *RemoveInputPortParams) WithID(id string) *RemoveInputPortParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the remove input port params
func (o *RemoveInputPortParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the remove input port params
func (o *RemoveInputPortParams) WithVersion(version *string) *RemoveInputPortParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the remove input port params
func (o *RemoveInputPortParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveInputPortParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

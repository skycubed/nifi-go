// Code generated by go-swagger; DO NOT EDIT.

package versions

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

// NewDeleteRevertRequestParams creates a new DeleteRevertRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRevertRequestParams() *DeleteRevertRequestParams {
	return &DeleteRevertRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRevertRequestParamsWithTimeout creates a new DeleteRevertRequestParams object
// with the ability to set a timeout on a request.
func NewDeleteRevertRequestParamsWithTimeout(timeout time.Duration) *DeleteRevertRequestParams {
	return &DeleteRevertRequestParams{
		timeout: timeout,
	}
}

// NewDeleteRevertRequestParamsWithContext creates a new DeleteRevertRequestParams object
// with the ability to set a context for a request.
func NewDeleteRevertRequestParamsWithContext(ctx context.Context) *DeleteRevertRequestParams {
	return &DeleteRevertRequestParams{
		Context: ctx,
	}
}

// NewDeleteRevertRequestParamsWithHTTPClient creates a new DeleteRevertRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRevertRequestParamsWithHTTPClient(client *http.Client) *DeleteRevertRequestParams {
	return &DeleteRevertRequestParams{
		HTTPClient: client,
	}
}

/*
DeleteRevertRequestParams contains all the parameters to send to the API endpoint

	for the delete revert request operation.

	Typically these are written to a http.Request.
*/
type DeleteRevertRequestParams struct {

	/* DisconnectedNodeAcknowledged.

	   Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	*/
	DisconnectedNodeAcknowledged *bool

	/* ID.

	   The ID of the Revert Request
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete revert request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRevertRequestParams) WithDefaults() *DeleteRevertRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete revert request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRevertRequestParams) SetDefaults() {
	var (
		disconnectedNodeAcknowledgedDefault = bool(false)
	)

	val := DeleteRevertRequestParams{
		DisconnectedNodeAcknowledged: &disconnectedNodeAcknowledgedDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete revert request params
func (o *DeleteRevertRequestParams) WithTimeout(timeout time.Duration) *DeleteRevertRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete revert request params
func (o *DeleteRevertRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete revert request params
func (o *DeleteRevertRequestParams) WithContext(ctx context.Context) *DeleteRevertRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete revert request params
func (o *DeleteRevertRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete revert request params
func (o *DeleteRevertRequestParams) WithHTTPClient(client *http.Client) *DeleteRevertRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete revert request params
func (o *DeleteRevertRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisconnectedNodeAcknowledged adds the disconnectedNodeAcknowledged to the delete revert request params
func (o *DeleteRevertRequestParams) WithDisconnectedNodeAcknowledged(disconnectedNodeAcknowledged *bool) *DeleteRevertRequestParams {
	o.SetDisconnectedNodeAcknowledged(disconnectedNodeAcknowledged)
	return o
}

// SetDisconnectedNodeAcknowledged adds the disconnectedNodeAcknowledged to the delete revert request params
func (o *DeleteRevertRequestParams) SetDisconnectedNodeAcknowledged(disconnectedNodeAcknowledged *bool) {
	o.DisconnectedNodeAcknowledged = disconnectedNodeAcknowledged
}

// WithID adds the id to the delete revert request params
func (o *DeleteRevertRequestParams) WithID(id string) *DeleteRevertRequestParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete revert request params
func (o *DeleteRevertRequestParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRevertRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package process_groups

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

// NewDeleteVariableRegistryUpdateRequestParams creates a new DeleteVariableRegistryUpdateRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVariableRegistryUpdateRequestParams() *DeleteVariableRegistryUpdateRequestParams {
	return &DeleteVariableRegistryUpdateRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVariableRegistryUpdateRequestParamsWithTimeout creates a new DeleteVariableRegistryUpdateRequestParams object
// with the ability to set a timeout on a request.
func NewDeleteVariableRegistryUpdateRequestParamsWithTimeout(timeout time.Duration) *DeleteVariableRegistryUpdateRequestParams {
	return &DeleteVariableRegistryUpdateRequestParams{
		timeout: timeout,
	}
}

// NewDeleteVariableRegistryUpdateRequestParamsWithContext creates a new DeleteVariableRegistryUpdateRequestParams object
// with the ability to set a context for a request.
func NewDeleteVariableRegistryUpdateRequestParamsWithContext(ctx context.Context) *DeleteVariableRegistryUpdateRequestParams {
	return &DeleteVariableRegistryUpdateRequestParams{
		Context: ctx,
	}
}

// NewDeleteVariableRegistryUpdateRequestParamsWithHTTPClient creates a new DeleteVariableRegistryUpdateRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVariableRegistryUpdateRequestParamsWithHTTPClient(client *http.Client) *DeleteVariableRegistryUpdateRequestParams {
	return &DeleteVariableRegistryUpdateRequestParams{
		HTTPClient: client,
	}
}

/* DeleteVariableRegistryUpdateRequestParams contains all the parameters to send to the API endpoint
   for the delete variable registry update request operation.

   Typically these are written to a http.Request.
*/
type DeleteVariableRegistryUpdateRequestParams struct {

	/* DisconnectedNodeAcknowledged.

	   Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	*/
	DisconnectedNodeAcknowledged *bool

	/* GroupID.

	   The process group id.
	*/
	GroupID string

	/* UpdateID.

	   The ID of the Variable Registry Update Request
	*/
	UpdateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete variable registry update request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVariableRegistryUpdateRequestParams) WithDefaults() *DeleteVariableRegistryUpdateRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete variable registry update request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVariableRegistryUpdateRequestParams) SetDefaults() {
	var (
		disconnectedNodeAcknowledgedDefault = bool(false)
	)

	val := DeleteVariableRegistryUpdateRequestParams{
		DisconnectedNodeAcknowledged: &disconnectedNodeAcknowledgedDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete variable registry update request params
func (o *DeleteVariableRegistryUpdateRequestParams) WithTimeout(timeout time.Duration) *DeleteVariableRegistryUpdateRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete variable registry update request params
func (o *DeleteVariableRegistryUpdateRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete variable registry update request params
func (o *DeleteVariableRegistryUpdateRequestParams) WithContext(ctx context.Context) *DeleteVariableRegistryUpdateRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete variable registry update request params
func (o *DeleteVariableRegistryUpdateRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete variable registry update request params
func (o *DeleteVariableRegistryUpdateRequestParams) WithHTTPClient(client *http.Client) *DeleteVariableRegistryUpdateRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete variable registry update request params
func (o *DeleteVariableRegistryUpdateRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisconnectedNodeAcknowledged adds the disconnectedNodeAcknowledged to the delete variable registry update request params
func (o *DeleteVariableRegistryUpdateRequestParams) WithDisconnectedNodeAcknowledged(disconnectedNodeAcknowledged *bool) *DeleteVariableRegistryUpdateRequestParams {
	o.SetDisconnectedNodeAcknowledged(disconnectedNodeAcknowledged)
	return o
}

// SetDisconnectedNodeAcknowledged adds the disconnectedNodeAcknowledged to the delete variable registry update request params
func (o *DeleteVariableRegistryUpdateRequestParams) SetDisconnectedNodeAcknowledged(disconnectedNodeAcknowledged *bool) {
	o.DisconnectedNodeAcknowledged = disconnectedNodeAcknowledged
}

// WithGroupID adds the groupID to the delete variable registry update request params
func (o *DeleteVariableRegistryUpdateRequestParams) WithGroupID(groupID string) *DeleteVariableRegistryUpdateRequestParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the delete variable registry update request params
func (o *DeleteVariableRegistryUpdateRequestParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithUpdateID adds the updateID to the delete variable registry update request params
func (o *DeleteVariableRegistryUpdateRequestParams) WithUpdateID(updateID string) *DeleteVariableRegistryUpdateRequestParams {
	o.SetUpdateID(updateID)
	return o
}

// SetUpdateID adds the updateId to the delete variable registry update request params
func (o *DeleteVariableRegistryUpdateRequestParams) SetUpdateID(updateID string) {
	o.UpdateID = updateID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVariableRegistryUpdateRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
		return err
	}

	// path param updateId
	if err := r.SetPathParam("updateId", o.UpdateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
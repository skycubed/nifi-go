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
)

// NewRemoveDropRequestParams creates a new RemoveDropRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveDropRequestParams() *RemoveDropRequestParams {
	return &RemoveDropRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveDropRequestParamsWithTimeout creates a new RemoveDropRequestParams object
// with the ability to set a timeout on a request.
func NewRemoveDropRequestParamsWithTimeout(timeout time.Duration) *RemoveDropRequestParams {
	return &RemoveDropRequestParams{
		timeout: timeout,
	}
}

// NewRemoveDropRequestParamsWithContext creates a new RemoveDropRequestParams object
// with the ability to set a context for a request.
func NewRemoveDropRequestParamsWithContext(ctx context.Context) *RemoveDropRequestParams {
	return &RemoveDropRequestParams{
		Context: ctx,
	}
}

// NewRemoveDropRequestParamsWithHTTPClient creates a new RemoveDropRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveDropRequestParamsWithHTTPClient(client *http.Client) *RemoveDropRequestParams {
	return &RemoveDropRequestParams{
		HTTPClient: client,
	}
}

/*
RemoveDropRequestParams contains all the parameters to send to the API endpoint

	for the remove drop request operation.

	Typically these are written to a http.Request.
*/
type RemoveDropRequestParams struct {

	/* DropRequestID.

	   The drop request id.
	*/
	DropRequestID string

	/* ID.

	   The process group id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove drop request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveDropRequestParams) WithDefaults() *RemoveDropRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove drop request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveDropRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove drop request params
func (o *RemoveDropRequestParams) WithTimeout(timeout time.Duration) *RemoveDropRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove drop request params
func (o *RemoveDropRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove drop request params
func (o *RemoveDropRequestParams) WithContext(ctx context.Context) *RemoveDropRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove drop request params
func (o *RemoveDropRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove drop request params
func (o *RemoveDropRequestParams) WithHTTPClient(client *http.Client) *RemoveDropRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove drop request params
func (o *RemoveDropRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDropRequestID adds the dropRequestID to the remove drop request params
func (o *RemoveDropRequestParams) WithDropRequestID(dropRequestID string) *RemoveDropRequestParams {
	o.SetDropRequestID(dropRequestID)
	return o
}

// SetDropRequestID adds the dropRequestId to the remove drop request params
func (o *RemoveDropRequestParams) SetDropRequestID(dropRequestID string) {
	o.DropRequestID = dropRequestID
}

// WithID adds the id to the remove drop request params
func (o *RemoveDropRequestParams) WithID(id string) *RemoveDropRequestParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the remove drop request params
func (o *RemoveDropRequestParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveDropRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param drop-request-id
	if err := r.SetPathParam("drop-request-id", o.DropRequestID); err != nil {
		return err
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

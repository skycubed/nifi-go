// Code generated by go-swagger; DO NOT EDIT.

package flowfile_queues

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

// NewGetDropRequestParams creates a new GetDropRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDropRequestParams() *GetDropRequestParams {
	return &GetDropRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDropRequestParamsWithTimeout creates a new GetDropRequestParams object
// with the ability to set a timeout on a request.
func NewGetDropRequestParamsWithTimeout(timeout time.Duration) *GetDropRequestParams {
	return &GetDropRequestParams{
		timeout: timeout,
	}
}

// NewGetDropRequestParamsWithContext creates a new GetDropRequestParams object
// with the ability to set a context for a request.
func NewGetDropRequestParamsWithContext(ctx context.Context) *GetDropRequestParams {
	return &GetDropRequestParams{
		Context: ctx,
	}
}

// NewGetDropRequestParamsWithHTTPClient creates a new GetDropRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDropRequestParamsWithHTTPClient(client *http.Client) *GetDropRequestParams {
	return &GetDropRequestParams{
		HTTPClient: client,
	}
}

/* GetDropRequestParams contains all the parameters to send to the API endpoint
   for the get drop request operation.

   Typically these are written to a http.Request.
*/
type GetDropRequestParams struct {

	/* DropRequestID.

	   The drop request id.
	*/
	DropRequestID string

	/* ID.

	   The connection id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get drop request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDropRequestParams) WithDefaults() *GetDropRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get drop request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDropRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get drop request params
func (o *GetDropRequestParams) WithTimeout(timeout time.Duration) *GetDropRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get drop request params
func (o *GetDropRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get drop request params
func (o *GetDropRequestParams) WithContext(ctx context.Context) *GetDropRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get drop request params
func (o *GetDropRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get drop request params
func (o *GetDropRequestParams) WithHTTPClient(client *http.Client) *GetDropRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get drop request params
func (o *GetDropRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDropRequestID adds the dropRequestID to the get drop request params
func (o *GetDropRequestParams) WithDropRequestID(dropRequestID string) *GetDropRequestParams {
	o.SetDropRequestID(dropRequestID)
	return o
}

// SetDropRequestID adds the dropRequestId to the get drop request params
func (o *GetDropRequestParams) SetDropRequestID(dropRequestID string) {
	o.DropRequestID = dropRequestID
}

// WithID adds the id to the get drop request params
func (o *GetDropRequestParams) WithID(id string) *GetDropRequestParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get drop request params
func (o *GetDropRequestParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDropRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
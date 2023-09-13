// Code generated by go-swagger; DO NOT EDIT.

package counters

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

// NewUpdateCounterParams creates a new UpdateCounterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCounterParams() *UpdateCounterParams {
	return &UpdateCounterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCounterParamsWithTimeout creates a new UpdateCounterParams object
// with the ability to set a timeout on a request.
func NewUpdateCounterParamsWithTimeout(timeout time.Duration) *UpdateCounterParams {
	return &UpdateCounterParams{
		timeout: timeout,
	}
}

// NewUpdateCounterParamsWithContext creates a new UpdateCounterParams object
// with the ability to set a context for a request.
func NewUpdateCounterParamsWithContext(ctx context.Context) *UpdateCounterParams {
	return &UpdateCounterParams{
		Context: ctx,
	}
}

// NewUpdateCounterParamsWithHTTPClient creates a new UpdateCounterParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCounterParamsWithHTTPClient(client *http.Client) *UpdateCounterParams {
	return &UpdateCounterParams{
		HTTPClient: client,
	}
}

/*
UpdateCounterParams contains all the parameters to send to the API endpoint

	for the update counter operation.

	Typically these are written to a http.Request.
*/
type UpdateCounterParams struct {

	/* ID.

	   The id of the counter.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update counter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCounterParams) WithDefaults() *UpdateCounterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update counter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCounterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update counter params
func (o *UpdateCounterParams) WithTimeout(timeout time.Duration) *UpdateCounterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update counter params
func (o *UpdateCounterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update counter params
func (o *UpdateCounterParams) WithContext(ctx context.Context) *UpdateCounterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update counter params
func (o *UpdateCounterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update counter params
func (o *UpdateCounterParams) WithHTTPClient(client *http.Client) *UpdateCounterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update counter params
func (o *UpdateCounterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update counter params
func (o *UpdateCounterParams) WithID(id string) *UpdateCounterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update counter params
func (o *UpdateCounterParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCounterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

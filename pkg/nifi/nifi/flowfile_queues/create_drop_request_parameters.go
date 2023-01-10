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

// NewCreateDropRequestParams creates a new CreateDropRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateDropRequestParams() *CreateDropRequestParams {
	return &CreateDropRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDropRequestParamsWithTimeout creates a new CreateDropRequestParams object
// with the ability to set a timeout on a request.
func NewCreateDropRequestParamsWithTimeout(timeout time.Duration) *CreateDropRequestParams {
	return &CreateDropRequestParams{
		timeout: timeout,
	}
}

// NewCreateDropRequestParamsWithContext creates a new CreateDropRequestParams object
// with the ability to set a context for a request.
func NewCreateDropRequestParamsWithContext(ctx context.Context) *CreateDropRequestParams {
	return &CreateDropRequestParams{
		Context: ctx,
	}
}

// NewCreateDropRequestParamsWithHTTPClient creates a new CreateDropRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateDropRequestParamsWithHTTPClient(client *http.Client) *CreateDropRequestParams {
	return &CreateDropRequestParams{
		HTTPClient: client,
	}
}

/*
CreateDropRequestParams contains all the parameters to send to the API endpoint

	for the create drop request operation.

	Typically these are written to a http.Request.
*/
type CreateDropRequestParams struct {

	/* ID.

	   The connection id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create drop request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDropRequestParams) WithDefaults() *CreateDropRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create drop request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDropRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create drop request params
func (o *CreateDropRequestParams) WithTimeout(timeout time.Duration) *CreateDropRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create drop request params
func (o *CreateDropRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create drop request params
func (o *CreateDropRequestParams) WithContext(ctx context.Context) *CreateDropRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create drop request params
func (o *CreateDropRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create drop request params
func (o *CreateDropRequestParams) WithHTTPClient(client *http.Client) *CreateDropRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create drop request params
func (o *CreateDropRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the create drop request params
func (o *CreateDropRequestParams) WithID(id string) *CreateDropRequestParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create drop request params
func (o *CreateDropRequestParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDropRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

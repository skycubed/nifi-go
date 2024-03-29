// Code generated by go-swagger; DO NOT EDIT.

package access

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

// NewLogOutParams creates a new LogOutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLogOutParams() *LogOutParams {
	return &LogOutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLogOutParamsWithTimeout creates a new LogOutParams object
// with the ability to set a timeout on a request.
func NewLogOutParamsWithTimeout(timeout time.Duration) *LogOutParams {
	return &LogOutParams{
		timeout: timeout,
	}
}

// NewLogOutParamsWithContext creates a new LogOutParams object
// with the ability to set a context for a request.
func NewLogOutParamsWithContext(ctx context.Context) *LogOutParams {
	return &LogOutParams{
		Context: ctx,
	}
}

// NewLogOutParamsWithHTTPClient creates a new LogOutParams object
// with the ability to set a custom HTTPClient for a request.
func NewLogOutParamsWithHTTPClient(client *http.Client) *LogOutParams {
	return &LogOutParams{
		HTTPClient: client,
	}
}

/*
LogOutParams contains all the parameters to send to the API endpoint

	for the log out operation.

	Typically these are written to a http.Request.
*/
type LogOutParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the log out params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogOutParams) WithDefaults() *LogOutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the log out params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogOutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the log out params
func (o *LogOutParams) WithTimeout(timeout time.Duration) *LogOutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the log out params
func (o *LogOutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the log out params
func (o *LogOutParams) WithContext(ctx context.Context) *LogOutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the log out params
func (o *LogOutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the log out params
func (o *LogOutParams) WithHTTPClient(client *http.Client) *LogOutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the log out params
func (o *LogOutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *LogOutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package reporting_tasks

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

// NewGetReportingTaskParams creates a new GetReportingTaskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetReportingTaskParams() *GetReportingTaskParams {
	return &GetReportingTaskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetReportingTaskParamsWithTimeout creates a new GetReportingTaskParams object
// with the ability to set a timeout on a request.
func NewGetReportingTaskParamsWithTimeout(timeout time.Duration) *GetReportingTaskParams {
	return &GetReportingTaskParams{
		timeout: timeout,
	}
}

// NewGetReportingTaskParamsWithContext creates a new GetReportingTaskParams object
// with the ability to set a context for a request.
func NewGetReportingTaskParamsWithContext(ctx context.Context) *GetReportingTaskParams {
	return &GetReportingTaskParams{
		Context: ctx,
	}
}

// NewGetReportingTaskParamsWithHTTPClient creates a new GetReportingTaskParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetReportingTaskParamsWithHTTPClient(client *http.Client) *GetReportingTaskParams {
	return &GetReportingTaskParams{
		HTTPClient: client,
	}
}

/*
GetReportingTaskParams contains all the parameters to send to the API endpoint

	for the get reporting task operation.

	Typically these are written to a http.Request.
*/
type GetReportingTaskParams struct {

	/* ID.

	   The reporting task id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get reporting task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReportingTaskParams) WithDefaults() *GetReportingTaskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get reporting task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReportingTaskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get reporting task params
func (o *GetReportingTaskParams) WithTimeout(timeout time.Duration) *GetReportingTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get reporting task params
func (o *GetReportingTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get reporting task params
func (o *GetReportingTaskParams) WithContext(ctx context.Context) *GetReportingTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get reporting task params
func (o *GetReportingTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get reporting task params
func (o *GetReportingTaskParams) WithHTTPClient(client *http.Client) *GetReportingTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get reporting task params
func (o *GetReportingTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get reporting task params
func (o *GetReportingTaskParams) WithID(id string) *GetReportingTaskParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get reporting task params
func (o *GetReportingTaskParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetReportingTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

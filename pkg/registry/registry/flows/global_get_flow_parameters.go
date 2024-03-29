// Code generated by go-swagger; DO NOT EDIT.

package flows

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

// NewGlobalGetFlowParams creates a new GlobalGetFlowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGlobalGetFlowParams() *GlobalGetFlowParams {
	return &GlobalGetFlowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGlobalGetFlowParamsWithTimeout creates a new GlobalGetFlowParams object
// with the ability to set a timeout on a request.
func NewGlobalGetFlowParamsWithTimeout(timeout time.Duration) *GlobalGetFlowParams {
	return &GlobalGetFlowParams{
		timeout: timeout,
	}
}

// NewGlobalGetFlowParamsWithContext creates a new GlobalGetFlowParams object
// with the ability to set a context for a request.
func NewGlobalGetFlowParamsWithContext(ctx context.Context) *GlobalGetFlowParams {
	return &GlobalGetFlowParams{
		Context: ctx,
	}
}

// NewGlobalGetFlowParamsWithHTTPClient creates a new GlobalGetFlowParams object
// with the ability to set a custom HTTPClient for a request.
func NewGlobalGetFlowParamsWithHTTPClient(client *http.Client) *GlobalGetFlowParams {
	return &GlobalGetFlowParams{
		HTTPClient: client,
	}
}

/*
GlobalGetFlowParams contains all the parameters to send to the API endpoint

	for the global get flow operation.

	Typically these are written to a http.Request.
*/
type GlobalGetFlowParams struct {

	/* FlowID.

	   The flow identifier
	*/
	FlowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the global get flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GlobalGetFlowParams) WithDefaults() *GlobalGetFlowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the global get flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GlobalGetFlowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the global get flow params
func (o *GlobalGetFlowParams) WithTimeout(timeout time.Duration) *GlobalGetFlowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the global get flow params
func (o *GlobalGetFlowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the global get flow params
func (o *GlobalGetFlowParams) WithContext(ctx context.Context) *GlobalGetFlowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the global get flow params
func (o *GlobalGetFlowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the global get flow params
func (o *GlobalGetFlowParams) WithHTTPClient(client *http.Client) *GlobalGetFlowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the global get flow params
func (o *GlobalGetFlowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFlowID adds the flowID to the global get flow params
func (o *GlobalGetFlowParams) WithFlowID(flowID string) *GlobalGetFlowParams {
	o.SetFlowID(flowID)
	return o
}

// SetFlowID adds the flowId to the global get flow params
func (o *GlobalGetFlowParams) SetFlowID(flowID string) {
	o.FlowID = flowID
}

// WriteToRequest writes these params to a swagger request
func (o *GlobalGetFlowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param flowId
	if err := r.SetPathParam("flowId", o.FlowID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

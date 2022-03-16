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

// NewGlobalGetFlowVersionsParams creates a new GlobalGetFlowVersionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGlobalGetFlowVersionsParams() *GlobalGetFlowVersionsParams {
	return &GlobalGetFlowVersionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGlobalGetFlowVersionsParamsWithTimeout creates a new GlobalGetFlowVersionsParams object
// with the ability to set a timeout on a request.
func NewGlobalGetFlowVersionsParamsWithTimeout(timeout time.Duration) *GlobalGetFlowVersionsParams {
	return &GlobalGetFlowVersionsParams{
		timeout: timeout,
	}
}

// NewGlobalGetFlowVersionsParamsWithContext creates a new GlobalGetFlowVersionsParams object
// with the ability to set a context for a request.
func NewGlobalGetFlowVersionsParamsWithContext(ctx context.Context) *GlobalGetFlowVersionsParams {
	return &GlobalGetFlowVersionsParams{
		Context: ctx,
	}
}

// NewGlobalGetFlowVersionsParamsWithHTTPClient creates a new GlobalGetFlowVersionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGlobalGetFlowVersionsParamsWithHTTPClient(client *http.Client) *GlobalGetFlowVersionsParams {
	return &GlobalGetFlowVersionsParams{
		HTTPClient: client,
	}
}

/* GlobalGetFlowVersionsParams contains all the parameters to send to the API endpoint
   for the global get flow versions operation.

   Typically these are written to a http.Request.
*/
type GlobalGetFlowVersionsParams struct {

	/* FlowID.

	   The flow identifier
	*/
	FlowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the global get flow versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GlobalGetFlowVersionsParams) WithDefaults() *GlobalGetFlowVersionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the global get flow versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GlobalGetFlowVersionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the global get flow versions params
func (o *GlobalGetFlowVersionsParams) WithTimeout(timeout time.Duration) *GlobalGetFlowVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the global get flow versions params
func (o *GlobalGetFlowVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the global get flow versions params
func (o *GlobalGetFlowVersionsParams) WithContext(ctx context.Context) *GlobalGetFlowVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the global get flow versions params
func (o *GlobalGetFlowVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the global get flow versions params
func (o *GlobalGetFlowVersionsParams) WithHTTPClient(client *http.Client) *GlobalGetFlowVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the global get flow versions params
func (o *GlobalGetFlowVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFlowID adds the flowID to the global get flow versions params
func (o *GlobalGetFlowVersionsParams) WithFlowID(flowID string) *GlobalGetFlowVersionsParams {
	o.SetFlowID(flowID)
	return o
}

// SetFlowID adds the flowId to the global get flow versions params
func (o *GlobalGetFlowVersionsParams) SetFlowID(flowID string) {
	o.FlowID = flowID
}

// WriteToRequest writes these params to a swagger request
func (o *GlobalGetFlowVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

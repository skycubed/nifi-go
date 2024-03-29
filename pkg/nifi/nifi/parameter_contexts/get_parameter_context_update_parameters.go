// Code generated by go-swagger; DO NOT EDIT.

package parameter_contexts

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

// NewGetParameterContextUpdateParams creates a new GetParameterContextUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetParameterContextUpdateParams() *GetParameterContextUpdateParams {
	return &GetParameterContextUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetParameterContextUpdateParamsWithTimeout creates a new GetParameterContextUpdateParams object
// with the ability to set a timeout on a request.
func NewGetParameterContextUpdateParamsWithTimeout(timeout time.Duration) *GetParameterContextUpdateParams {
	return &GetParameterContextUpdateParams{
		timeout: timeout,
	}
}

// NewGetParameterContextUpdateParamsWithContext creates a new GetParameterContextUpdateParams object
// with the ability to set a context for a request.
func NewGetParameterContextUpdateParamsWithContext(ctx context.Context) *GetParameterContextUpdateParams {
	return &GetParameterContextUpdateParams{
		Context: ctx,
	}
}

// NewGetParameterContextUpdateParamsWithHTTPClient creates a new GetParameterContextUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetParameterContextUpdateParamsWithHTTPClient(client *http.Client) *GetParameterContextUpdateParams {
	return &GetParameterContextUpdateParams{
		HTTPClient: client,
	}
}

/*
GetParameterContextUpdateParams contains all the parameters to send to the API endpoint

	for the get parameter context update operation.

	Typically these are written to a http.Request.
*/
type GetParameterContextUpdateParams struct {

	/* ContextID.

	   The ID of the Parameter Context
	*/
	ContextID string

	/* RequestID.

	   The ID of the Update Request
	*/
	RequestID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get parameter context update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetParameterContextUpdateParams) WithDefaults() *GetParameterContextUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get parameter context update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetParameterContextUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get parameter context update params
func (o *GetParameterContextUpdateParams) WithTimeout(timeout time.Duration) *GetParameterContextUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get parameter context update params
func (o *GetParameterContextUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get parameter context update params
func (o *GetParameterContextUpdateParams) WithContext(ctx context.Context) *GetParameterContextUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get parameter context update params
func (o *GetParameterContextUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get parameter context update params
func (o *GetParameterContextUpdateParams) WithHTTPClient(client *http.Client) *GetParameterContextUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get parameter context update params
func (o *GetParameterContextUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContextID adds the contextID to the get parameter context update params
func (o *GetParameterContextUpdateParams) WithContextID(contextID string) *GetParameterContextUpdateParams {
	o.SetContextID(contextID)
	return o
}

// SetContextID adds the contextId to the get parameter context update params
func (o *GetParameterContextUpdateParams) SetContextID(contextID string) {
	o.ContextID = contextID
}

// WithRequestID adds the requestID to the get parameter context update params
func (o *GetParameterContextUpdateParams) WithRequestID(requestID string) *GetParameterContextUpdateParams {
	o.SetRequestID(requestID)
	return o
}

// SetRequestID adds the requestId to the get parameter context update params
func (o *GetParameterContextUpdateParams) SetRequestID(requestID string) {
	o.RequestID = requestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetParameterContextUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contextId
	if err := r.SetPathParam("contextId", o.ContextID); err != nil {
		return err
	}

	// path param requestId
	if err := r.SetPathParam("requestId", o.RequestID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

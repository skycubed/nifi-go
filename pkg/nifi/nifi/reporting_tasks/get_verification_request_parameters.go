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

// NewGetVerificationRequestParams creates a new GetVerificationRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVerificationRequestParams() *GetVerificationRequestParams {
	return &GetVerificationRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVerificationRequestParamsWithTimeout creates a new GetVerificationRequestParams object
// with the ability to set a timeout on a request.
func NewGetVerificationRequestParamsWithTimeout(timeout time.Duration) *GetVerificationRequestParams {
	return &GetVerificationRequestParams{
		timeout: timeout,
	}
}

// NewGetVerificationRequestParamsWithContext creates a new GetVerificationRequestParams object
// with the ability to set a context for a request.
func NewGetVerificationRequestParamsWithContext(ctx context.Context) *GetVerificationRequestParams {
	return &GetVerificationRequestParams{
		Context: ctx,
	}
}

// NewGetVerificationRequestParamsWithHTTPClient creates a new GetVerificationRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVerificationRequestParamsWithHTTPClient(client *http.Client) *GetVerificationRequestParams {
	return &GetVerificationRequestParams{
		HTTPClient: client,
	}
}

/* GetVerificationRequestParams contains all the parameters to send to the API endpoint
   for the get verification request operation.

   Typically these are written to a http.Request.
*/
type GetVerificationRequestParams struct {

	/* ID.

	   The ID of the Reporting Task
	*/
	ID string

	/* RequestID.

	   The ID of the Verification Request
	*/
	RequestID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get verification request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVerificationRequestParams) WithDefaults() *GetVerificationRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get verification request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVerificationRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get verification request params
func (o *GetVerificationRequestParams) WithTimeout(timeout time.Duration) *GetVerificationRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get verification request params
func (o *GetVerificationRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get verification request params
func (o *GetVerificationRequestParams) WithContext(ctx context.Context) *GetVerificationRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get verification request params
func (o *GetVerificationRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get verification request params
func (o *GetVerificationRequestParams) WithHTTPClient(client *http.Client) *GetVerificationRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get verification request params
func (o *GetVerificationRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get verification request params
func (o *GetVerificationRequestParams) WithID(id string) *GetVerificationRequestParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get verification request params
func (o *GetVerificationRequestParams) SetID(id string) {
	o.ID = id
}

// WithRequestID adds the requestID to the get verification request params
func (o *GetVerificationRequestParams) WithRequestID(requestID string) *GetVerificationRequestParams {
	o.SetRequestID(requestID)
	return o
}

// SetRequestID adds the requestId to the get verification request params
func (o *GetVerificationRequestParams) SetRequestID(requestID string) {
	o.RequestID = requestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVerificationRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
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
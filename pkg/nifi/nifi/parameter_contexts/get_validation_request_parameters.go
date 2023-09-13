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

// NewGetValidationRequestParams creates a new GetValidationRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetValidationRequestParams() *GetValidationRequestParams {
	return &GetValidationRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetValidationRequestParamsWithTimeout creates a new GetValidationRequestParams object
// with the ability to set a timeout on a request.
func NewGetValidationRequestParamsWithTimeout(timeout time.Duration) *GetValidationRequestParams {
	return &GetValidationRequestParams{
		timeout: timeout,
	}
}

// NewGetValidationRequestParamsWithContext creates a new GetValidationRequestParams object
// with the ability to set a context for a request.
func NewGetValidationRequestParamsWithContext(ctx context.Context) *GetValidationRequestParams {
	return &GetValidationRequestParams{
		Context: ctx,
	}
}

// NewGetValidationRequestParamsWithHTTPClient creates a new GetValidationRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetValidationRequestParamsWithHTTPClient(client *http.Client) *GetValidationRequestParams {
	return &GetValidationRequestParams{
		HTTPClient: client,
	}
}

/*
GetValidationRequestParams contains all the parameters to send to the API endpoint

	for the get validation request operation.

	Typically these are written to a http.Request.
*/
type GetValidationRequestParams struct {

	/* ContextID.

	   The ID of the Parameter Context
	*/
	ContextID string

	/* ID.

	   The ID of the Validation Request
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get validation request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetValidationRequestParams) WithDefaults() *GetValidationRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get validation request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetValidationRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get validation request params
func (o *GetValidationRequestParams) WithTimeout(timeout time.Duration) *GetValidationRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get validation request params
func (o *GetValidationRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get validation request params
func (o *GetValidationRequestParams) WithContext(ctx context.Context) *GetValidationRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get validation request params
func (o *GetValidationRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get validation request params
func (o *GetValidationRequestParams) WithHTTPClient(client *http.Client) *GetValidationRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get validation request params
func (o *GetValidationRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContextID adds the contextID to the get validation request params
func (o *GetValidationRequestParams) WithContextID(contextID string) *GetValidationRequestParams {
	o.SetContextID(contextID)
	return o
}

// SetContextID adds the contextId to the get validation request params
func (o *GetValidationRequestParams) SetContextID(contextID string) {
	o.ContextID = contextID
}

// WithID adds the id to the get validation request params
func (o *GetValidationRequestParams) WithID(id string) *GetValidationRequestParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get validation request params
func (o *GetValidationRequestParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetValidationRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contextId
	if err := r.SetPathParam("contextId", o.ContextID); err != nil {
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

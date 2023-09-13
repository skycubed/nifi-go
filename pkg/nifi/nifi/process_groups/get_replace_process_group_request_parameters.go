// Code generated by go-swagger; DO NOT EDIT.

package process_groups

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

// NewGetReplaceProcessGroupRequestParams creates a new GetReplaceProcessGroupRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetReplaceProcessGroupRequestParams() *GetReplaceProcessGroupRequestParams {
	return &GetReplaceProcessGroupRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetReplaceProcessGroupRequestParamsWithTimeout creates a new GetReplaceProcessGroupRequestParams object
// with the ability to set a timeout on a request.
func NewGetReplaceProcessGroupRequestParamsWithTimeout(timeout time.Duration) *GetReplaceProcessGroupRequestParams {
	return &GetReplaceProcessGroupRequestParams{
		timeout: timeout,
	}
}

// NewGetReplaceProcessGroupRequestParamsWithContext creates a new GetReplaceProcessGroupRequestParams object
// with the ability to set a context for a request.
func NewGetReplaceProcessGroupRequestParamsWithContext(ctx context.Context) *GetReplaceProcessGroupRequestParams {
	return &GetReplaceProcessGroupRequestParams{
		Context: ctx,
	}
}

// NewGetReplaceProcessGroupRequestParamsWithHTTPClient creates a new GetReplaceProcessGroupRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetReplaceProcessGroupRequestParamsWithHTTPClient(client *http.Client) *GetReplaceProcessGroupRequestParams {
	return &GetReplaceProcessGroupRequestParams{
		HTTPClient: client,
	}
}

/*
GetReplaceProcessGroupRequestParams contains all the parameters to send to the API endpoint

	for the get replace process group request operation.

	Typically these are written to a http.Request.
*/
type GetReplaceProcessGroupRequestParams struct {

	/* ID.

	   The ID of the Replace Request
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get replace process group request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReplaceProcessGroupRequestParams) WithDefaults() *GetReplaceProcessGroupRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get replace process group request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReplaceProcessGroupRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get replace process group request params
func (o *GetReplaceProcessGroupRequestParams) WithTimeout(timeout time.Duration) *GetReplaceProcessGroupRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get replace process group request params
func (o *GetReplaceProcessGroupRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get replace process group request params
func (o *GetReplaceProcessGroupRequestParams) WithContext(ctx context.Context) *GetReplaceProcessGroupRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get replace process group request params
func (o *GetReplaceProcessGroupRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get replace process group request params
func (o *GetReplaceProcessGroupRequestParams) WithHTTPClient(client *http.Client) *GetReplaceProcessGroupRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get replace process group request params
func (o *GetReplaceProcessGroupRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get replace process group request params
func (o *GetReplaceProcessGroupRequestParams) WithID(id string) *GetReplaceProcessGroupRequestParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get replace process group request params
func (o *GetReplaceProcessGroupRequestParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetReplaceProcessGroupRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

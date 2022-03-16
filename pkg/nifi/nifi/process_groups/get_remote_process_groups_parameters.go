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

// NewGetRemoteProcessGroupsParams creates a new GetRemoteProcessGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRemoteProcessGroupsParams() *GetRemoteProcessGroupsParams {
	return &GetRemoteProcessGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRemoteProcessGroupsParamsWithTimeout creates a new GetRemoteProcessGroupsParams object
// with the ability to set a timeout on a request.
func NewGetRemoteProcessGroupsParamsWithTimeout(timeout time.Duration) *GetRemoteProcessGroupsParams {
	return &GetRemoteProcessGroupsParams{
		timeout: timeout,
	}
}

// NewGetRemoteProcessGroupsParamsWithContext creates a new GetRemoteProcessGroupsParams object
// with the ability to set a context for a request.
func NewGetRemoteProcessGroupsParamsWithContext(ctx context.Context) *GetRemoteProcessGroupsParams {
	return &GetRemoteProcessGroupsParams{
		Context: ctx,
	}
}

// NewGetRemoteProcessGroupsParamsWithHTTPClient creates a new GetRemoteProcessGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRemoteProcessGroupsParamsWithHTTPClient(client *http.Client) *GetRemoteProcessGroupsParams {
	return &GetRemoteProcessGroupsParams{
		HTTPClient: client,
	}
}

/* GetRemoteProcessGroupsParams contains all the parameters to send to the API endpoint
   for the get remote process groups operation.

   Typically these are written to a http.Request.
*/
type GetRemoteProcessGroupsParams struct {

	/* ID.

	   The process group id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get remote process groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRemoteProcessGroupsParams) WithDefaults() *GetRemoteProcessGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get remote process groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRemoteProcessGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get remote process groups params
func (o *GetRemoteProcessGroupsParams) WithTimeout(timeout time.Duration) *GetRemoteProcessGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get remote process groups params
func (o *GetRemoteProcessGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get remote process groups params
func (o *GetRemoteProcessGroupsParams) WithContext(ctx context.Context) *GetRemoteProcessGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get remote process groups params
func (o *GetRemoteProcessGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get remote process groups params
func (o *GetRemoteProcessGroupsParams) WithHTTPClient(client *http.Client) *GetRemoteProcessGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get remote process groups params
func (o *GetRemoteProcessGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get remote process groups params
func (o *GetRemoteProcessGroupsParams) WithID(id string) *GetRemoteProcessGroupsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get remote process groups params
func (o *GetRemoteProcessGroupsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetRemoteProcessGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
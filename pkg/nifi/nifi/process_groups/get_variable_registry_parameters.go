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
	"github.com/go-openapi/swag"
)

// NewGetVariableRegistryParams creates a new GetVariableRegistryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVariableRegistryParams() *GetVariableRegistryParams {
	return &GetVariableRegistryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVariableRegistryParamsWithTimeout creates a new GetVariableRegistryParams object
// with the ability to set a timeout on a request.
func NewGetVariableRegistryParamsWithTimeout(timeout time.Duration) *GetVariableRegistryParams {
	return &GetVariableRegistryParams{
		timeout: timeout,
	}
}

// NewGetVariableRegistryParamsWithContext creates a new GetVariableRegistryParams object
// with the ability to set a context for a request.
func NewGetVariableRegistryParamsWithContext(ctx context.Context) *GetVariableRegistryParams {
	return &GetVariableRegistryParams{
		Context: ctx,
	}
}

// NewGetVariableRegistryParamsWithHTTPClient creates a new GetVariableRegistryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVariableRegistryParamsWithHTTPClient(client *http.Client) *GetVariableRegistryParams {
	return &GetVariableRegistryParams{
		HTTPClient: client,
	}
}

/* GetVariableRegistryParams contains all the parameters to send to the API endpoint
   for the get variable registry operation.

   Typically these are written to a http.Request.
*/
type GetVariableRegistryParams struct {

	/* ID.

	   The process group id.
	*/
	ID string

	/* IncludeAncestorGroups.

	   Whether or not to include ancestor groups

	   Default: true
	*/
	IncludeAncestorGroups *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get variable registry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVariableRegistryParams) WithDefaults() *GetVariableRegistryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get variable registry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVariableRegistryParams) SetDefaults() {
	var (
		includeAncestorGroupsDefault = bool(true)
	)

	val := GetVariableRegistryParams{
		IncludeAncestorGroups: &includeAncestorGroupsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get variable registry params
func (o *GetVariableRegistryParams) WithTimeout(timeout time.Duration) *GetVariableRegistryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get variable registry params
func (o *GetVariableRegistryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get variable registry params
func (o *GetVariableRegistryParams) WithContext(ctx context.Context) *GetVariableRegistryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get variable registry params
func (o *GetVariableRegistryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get variable registry params
func (o *GetVariableRegistryParams) WithHTTPClient(client *http.Client) *GetVariableRegistryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get variable registry params
func (o *GetVariableRegistryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get variable registry params
func (o *GetVariableRegistryParams) WithID(id string) *GetVariableRegistryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get variable registry params
func (o *GetVariableRegistryParams) SetID(id string) {
	o.ID = id
}

// WithIncludeAncestorGroups adds the includeAncestorGroups to the get variable registry params
func (o *GetVariableRegistryParams) WithIncludeAncestorGroups(includeAncestorGroups *bool) *GetVariableRegistryParams {
	o.SetIncludeAncestorGroups(includeAncestorGroups)
	return o
}

// SetIncludeAncestorGroups adds the includeAncestorGroups to the get variable registry params
func (o *GetVariableRegistryParams) SetIncludeAncestorGroups(includeAncestorGroups *bool) {
	o.IncludeAncestorGroups = includeAncestorGroups
}

// WriteToRequest writes these params to a swagger request
func (o *GetVariableRegistryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.IncludeAncestorGroups != nil {

		// query param includeAncestorGroups
		var qrIncludeAncestorGroups bool

		if o.IncludeAncestorGroups != nil {
			qrIncludeAncestorGroups = *o.IncludeAncestorGroups
		}
		qIncludeAncestorGroups := swag.FormatBool(qrIncludeAncestorGroups)
		if qIncludeAncestorGroups != "" {

			if err := r.SetQueryParam("includeAncestorGroups", qIncludeAncestorGroups); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
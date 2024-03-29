// Code generated by go-swagger; DO NOT EDIT.

package flow

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

// NewGetControllerServicesFromGroupParams creates a new GetControllerServicesFromGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetControllerServicesFromGroupParams() *GetControllerServicesFromGroupParams {
	return &GetControllerServicesFromGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetControllerServicesFromGroupParamsWithTimeout creates a new GetControllerServicesFromGroupParams object
// with the ability to set a timeout on a request.
func NewGetControllerServicesFromGroupParamsWithTimeout(timeout time.Duration) *GetControllerServicesFromGroupParams {
	return &GetControllerServicesFromGroupParams{
		timeout: timeout,
	}
}

// NewGetControllerServicesFromGroupParamsWithContext creates a new GetControllerServicesFromGroupParams object
// with the ability to set a context for a request.
func NewGetControllerServicesFromGroupParamsWithContext(ctx context.Context) *GetControllerServicesFromGroupParams {
	return &GetControllerServicesFromGroupParams{
		Context: ctx,
	}
}

// NewGetControllerServicesFromGroupParamsWithHTTPClient creates a new GetControllerServicesFromGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetControllerServicesFromGroupParamsWithHTTPClient(client *http.Client) *GetControllerServicesFromGroupParams {
	return &GetControllerServicesFromGroupParams{
		HTTPClient: client,
	}
}

/*
GetControllerServicesFromGroupParams contains all the parameters to send to the API endpoint

	for the get controller services from group operation.

	Typically these are written to a http.Request.
*/
type GetControllerServicesFromGroupParams struct {

	/* ID.

	   The process group id.
	*/
	ID string

	/* IncludeAncestorGroups.

	   Whether or not to include parent/ancestory process groups

	   Default: true
	*/
	IncludeAncestorGroups *bool

	/* IncludeDescendantGroups.

	   Whether or not to include descendant process groups
	*/
	IncludeDescendantGroups *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get controller services from group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetControllerServicesFromGroupParams) WithDefaults() *GetControllerServicesFromGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get controller services from group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetControllerServicesFromGroupParams) SetDefaults() {
	var (
		includeAncestorGroupsDefault = bool(true)

		includeDescendantGroupsDefault = bool(false)
	)

	val := GetControllerServicesFromGroupParams{
		IncludeAncestorGroups:   &includeAncestorGroupsDefault,
		IncludeDescendantGroups: &includeDescendantGroupsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get controller services from group params
func (o *GetControllerServicesFromGroupParams) WithTimeout(timeout time.Duration) *GetControllerServicesFromGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get controller services from group params
func (o *GetControllerServicesFromGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get controller services from group params
func (o *GetControllerServicesFromGroupParams) WithContext(ctx context.Context) *GetControllerServicesFromGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get controller services from group params
func (o *GetControllerServicesFromGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get controller services from group params
func (o *GetControllerServicesFromGroupParams) WithHTTPClient(client *http.Client) *GetControllerServicesFromGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get controller services from group params
func (o *GetControllerServicesFromGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get controller services from group params
func (o *GetControllerServicesFromGroupParams) WithID(id string) *GetControllerServicesFromGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get controller services from group params
func (o *GetControllerServicesFromGroupParams) SetID(id string) {
	o.ID = id
}

// WithIncludeAncestorGroups adds the includeAncestorGroups to the get controller services from group params
func (o *GetControllerServicesFromGroupParams) WithIncludeAncestorGroups(includeAncestorGroups *bool) *GetControllerServicesFromGroupParams {
	o.SetIncludeAncestorGroups(includeAncestorGroups)
	return o
}

// SetIncludeAncestorGroups adds the includeAncestorGroups to the get controller services from group params
func (o *GetControllerServicesFromGroupParams) SetIncludeAncestorGroups(includeAncestorGroups *bool) {
	o.IncludeAncestorGroups = includeAncestorGroups
}

// WithIncludeDescendantGroups adds the includeDescendantGroups to the get controller services from group params
func (o *GetControllerServicesFromGroupParams) WithIncludeDescendantGroups(includeDescendantGroups *bool) *GetControllerServicesFromGroupParams {
	o.SetIncludeDescendantGroups(includeDescendantGroups)
	return o
}

// SetIncludeDescendantGroups adds the includeDescendantGroups to the get controller services from group params
func (o *GetControllerServicesFromGroupParams) SetIncludeDescendantGroups(includeDescendantGroups *bool) {
	o.IncludeDescendantGroups = includeDescendantGroups
}

// WriteToRequest writes these params to a swagger request
func (o *GetControllerServicesFromGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IncludeDescendantGroups != nil {

		// query param includeDescendantGroups
		var qrIncludeDescendantGroups bool

		if o.IncludeDescendantGroups != nil {
			qrIncludeDescendantGroups = *o.IncludeDescendantGroups
		}
		qIncludeDescendantGroups := swag.FormatBool(qrIncludeDescendantGroups)
		if qIncludeDescendantGroups != "" {

			if err := r.SetQueryParam("includeDescendantGroups", qIncludeDescendantGroups); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

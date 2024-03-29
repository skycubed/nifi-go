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

// NewGetPropertyDescriptorParams creates a new GetPropertyDescriptorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPropertyDescriptorParams() *GetPropertyDescriptorParams {
	return &GetPropertyDescriptorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPropertyDescriptorParamsWithTimeout creates a new GetPropertyDescriptorParams object
// with the ability to set a timeout on a request.
func NewGetPropertyDescriptorParamsWithTimeout(timeout time.Duration) *GetPropertyDescriptorParams {
	return &GetPropertyDescriptorParams{
		timeout: timeout,
	}
}

// NewGetPropertyDescriptorParamsWithContext creates a new GetPropertyDescriptorParams object
// with the ability to set a context for a request.
func NewGetPropertyDescriptorParamsWithContext(ctx context.Context) *GetPropertyDescriptorParams {
	return &GetPropertyDescriptorParams{
		Context: ctx,
	}
}

// NewGetPropertyDescriptorParamsWithHTTPClient creates a new GetPropertyDescriptorParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPropertyDescriptorParamsWithHTTPClient(client *http.Client) *GetPropertyDescriptorParams {
	return &GetPropertyDescriptorParams{
		HTTPClient: client,
	}
}

/*
GetPropertyDescriptorParams contains all the parameters to send to the API endpoint

	for the get property descriptor operation.

	Typically these are written to a http.Request.
*/
type GetPropertyDescriptorParams struct {

	/* ID.

	   The reporting task id.
	*/
	ID string

	/* PropertyName.

	   The property name.
	*/
	PropertyName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get property descriptor params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPropertyDescriptorParams) WithDefaults() *GetPropertyDescriptorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get property descriptor params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPropertyDescriptorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get property descriptor params
func (o *GetPropertyDescriptorParams) WithTimeout(timeout time.Duration) *GetPropertyDescriptorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get property descriptor params
func (o *GetPropertyDescriptorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get property descriptor params
func (o *GetPropertyDescriptorParams) WithContext(ctx context.Context) *GetPropertyDescriptorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get property descriptor params
func (o *GetPropertyDescriptorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get property descriptor params
func (o *GetPropertyDescriptorParams) WithHTTPClient(client *http.Client) *GetPropertyDescriptorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get property descriptor params
func (o *GetPropertyDescriptorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get property descriptor params
func (o *GetPropertyDescriptorParams) WithID(id string) *GetPropertyDescriptorParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get property descriptor params
func (o *GetPropertyDescriptorParams) SetID(id string) {
	o.ID = id
}

// WithPropertyName adds the propertyName to the get property descriptor params
func (o *GetPropertyDescriptorParams) WithPropertyName(propertyName string) *GetPropertyDescriptorParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the get property descriptor params
func (o *GetPropertyDescriptorParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WriteToRequest writes these params to a swagger request
func (o *GetPropertyDescriptorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param propertyName
	qrPropertyName := o.PropertyName
	qPropertyName := qrPropertyName
	if qPropertyName != "" {

		if err := r.SetQueryParam("propertyName", qPropertyName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

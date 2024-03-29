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

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// NewUpdateVariableRegistryParams creates a new UpdateVariableRegistryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateVariableRegistryParams() *UpdateVariableRegistryParams {
	return &UpdateVariableRegistryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVariableRegistryParamsWithTimeout creates a new UpdateVariableRegistryParams object
// with the ability to set a timeout on a request.
func NewUpdateVariableRegistryParamsWithTimeout(timeout time.Duration) *UpdateVariableRegistryParams {
	return &UpdateVariableRegistryParams{
		timeout: timeout,
	}
}

// NewUpdateVariableRegistryParamsWithContext creates a new UpdateVariableRegistryParams object
// with the ability to set a context for a request.
func NewUpdateVariableRegistryParamsWithContext(ctx context.Context) *UpdateVariableRegistryParams {
	return &UpdateVariableRegistryParams{
		Context: ctx,
	}
}

// NewUpdateVariableRegistryParamsWithHTTPClient creates a new UpdateVariableRegistryParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateVariableRegistryParamsWithHTTPClient(client *http.Client) *UpdateVariableRegistryParams {
	return &UpdateVariableRegistryParams{
		HTTPClient: client,
	}
}

/*
UpdateVariableRegistryParams contains all the parameters to send to the API endpoint

	for the update variable registry operation.

	Typically these are written to a http.Request.
*/
type UpdateVariableRegistryParams struct {

	/* Body.

	   The variable registry configuration details.
	*/
	Body *models.VariableRegistryEntity

	/* ID.

	   The process group id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update variable registry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVariableRegistryParams) WithDefaults() *UpdateVariableRegistryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update variable registry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVariableRegistryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update variable registry params
func (o *UpdateVariableRegistryParams) WithTimeout(timeout time.Duration) *UpdateVariableRegistryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update variable registry params
func (o *UpdateVariableRegistryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update variable registry params
func (o *UpdateVariableRegistryParams) WithContext(ctx context.Context) *UpdateVariableRegistryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update variable registry params
func (o *UpdateVariableRegistryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update variable registry params
func (o *UpdateVariableRegistryParams) WithHTTPClient(client *http.Client) *UpdateVariableRegistryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update variable registry params
func (o *UpdateVariableRegistryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update variable registry params
func (o *UpdateVariableRegistryParams) WithBody(body *models.VariableRegistryEntity) *UpdateVariableRegistryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update variable registry params
func (o *UpdateVariableRegistryParams) SetBody(body *models.VariableRegistryEntity) {
	o.Body = body
}

// WithID adds the id to the update variable registry params
func (o *UpdateVariableRegistryParams) WithID(id string) *UpdateVariableRegistryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update variable registry params
func (o *UpdateVariableRegistryParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVariableRegistryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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

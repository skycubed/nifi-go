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

// NewUpdateProcessGroupParams creates a new UpdateProcessGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateProcessGroupParams() *UpdateProcessGroupParams {
	return &UpdateProcessGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProcessGroupParamsWithTimeout creates a new UpdateProcessGroupParams object
// with the ability to set a timeout on a request.
func NewUpdateProcessGroupParamsWithTimeout(timeout time.Duration) *UpdateProcessGroupParams {
	return &UpdateProcessGroupParams{
		timeout: timeout,
	}
}

// NewUpdateProcessGroupParamsWithContext creates a new UpdateProcessGroupParams object
// with the ability to set a context for a request.
func NewUpdateProcessGroupParamsWithContext(ctx context.Context) *UpdateProcessGroupParams {
	return &UpdateProcessGroupParams{
		Context: ctx,
	}
}

// NewUpdateProcessGroupParamsWithHTTPClient creates a new UpdateProcessGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateProcessGroupParamsWithHTTPClient(client *http.Client) *UpdateProcessGroupParams {
	return &UpdateProcessGroupParams{
		HTTPClient: client,
	}
}

/*
UpdateProcessGroupParams contains all the parameters to send to the API endpoint

	for the update process group operation.

	Typically these are written to a http.Request.
*/
type UpdateProcessGroupParams struct {

	/* Body.

	   The process group configuration details.
	*/
	Body *models.ProcessGroupEntity

	/* ID.

	   The process group id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update process group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProcessGroupParams) WithDefaults() *UpdateProcessGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update process group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProcessGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update process group params
func (o *UpdateProcessGroupParams) WithTimeout(timeout time.Duration) *UpdateProcessGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update process group params
func (o *UpdateProcessGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update process group params
func (o *UpdateProcessGroupParams) WithContext(ctx context.Context) *UpdateProcessGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update process group params
func (o *UpdateProcessGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update process group params
func (o *UpdateProcessGroupParams) WithHTTPClient(client *http.Client) *UpdateProcessGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update process group params
func (o *UpdateProcessGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update process group params
func (o *UpdateProcessGroupParams) WithBody(body *models.ProcessGroupEntity) *UpdateProcessGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update process group params
func (o *UpdateProcessGroupParams) SetBody(body *models.ProcessGroupEntity) {
	o.Body = body
}

// WithID adds the id to the update process group params
func (o *UpdateProcessGroupParams) WithID(id string) *UpdateProcessGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update process group params
func (o *UpdateProcessGroupParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProcessGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

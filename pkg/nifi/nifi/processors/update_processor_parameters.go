// Code generated by go-swagger; DO NOT EDIT.

package processors

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

// NewUpdateProcessorParams creates a new UpdateProcessorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateProcessorParams() *UpdateProcessorParams {
	return &UpdateProcessorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProcessorParamsWithTimeout creates a new UpdateProcessorParams object
// with the ability to set a timeout on a request.
func NewUpdateProcessorParamsWithTimeout(timeout time.Duration) *UpdateProcessorParams {
	return &UpdateProcessorParams{
		timeout: timeout,
	}
}

// NewUpdateProcessorParamsWithContext creates a new UpdateProcessorParams object
// with the ability to set a context for a request.
func NewUpdateProcessorParamsWithContext(ctx context.Context) *UpdateProcessorParams {
	return &UpdateProcessorParams{
		Context: ctx,
	}
}

// NewUpdateProcessorParamsWithHTTPClient creates a new UpdateProcessorParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateProcessorParamsWithHTTPClient(client *http.Client) *UpdateProcessorParams {
	return &UpdateProcessorParams{
		HTTPClient: client,
	}
}

/* UpdateProcessorParams contains all the parameters to send to the API endpoint
   for the update processor operation.

   Typically these are written to a http.Request.
*/
type UpdateProcessorParams struct {

	/* Body.

	   The processor configuration details.
	*/
	Body *models.ProcessorEntity

	/* ID.

	   The processor id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update processor params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProcessorParams) WithDefaults() *UpdateProcessorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update processor params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProcessorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update processor params
func (o *UpdateProcessorParams) WithTimeout(timeout time.Duration) *UpdateProcessorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update processor params
func (o *UpdateProcessorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update processor params
func (o *UpdateProcessorParams) WithContext(ctx context.Context) *UpdateProcessorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update processor params
func (o *UpdateProcessorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update processor params
func (o *UpdateProcessorParams) WithHTTPClient(client *http.Client) *UpdateProcessorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update processor params
func (o *UpdateProcessorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update processor params
func (o *UpdateProcessorParams) WithBody(body *models.ProcessorEntity) *UpdateProcessorParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update processor params
func (o *UpdateProcessorParams) SetBody(body *models.ProcessorEntity) {
	o.Body = body
}

// WithID adds the id to the update processor params
func (o *UpdateProcessorParams) WithID(id string) *UpdateProcessorParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update processor params
func (o *UpdateProcessorParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProcessorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

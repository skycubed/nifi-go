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

// NewCreateProcessorParams creates a new CreateProcessorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateProcessorParams() *CreateProcessorParams {
	return &CreateProcessorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProcessorParamsWithTimeout creates a new CreateProcessorParams object
// with the ability to set a timeout on a request.
func NewCreateProcessorParamsWithTimeout(timeout time.Duration) *CreateProcessorParams {
	return &CreateProcessorParams{
		timeout: timeout,
	}
}

// NewCreateProcessorParamsWithContext creates a new CreateProcessorParams object
// with the ability to set a context for a request.
func NewCreateProcessorParamsWithContext(ctx context.Context) *CreateProcessorParams {
	return &CreateProcessorParams{
		Context: ctx,
	}
}

// NewCreateProcessorParamsWithHTTPClient creates a new CreateProcessorParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateProcessorParamsWithHTTPClient(client *http.Client) *CreateProcessorParams {
	return &CreateProcessorParams{
		HTTPClient: client,
	}
}

/* CreateProcessorParams contains all the parameters to send to the API endpoint
   for the create processor operation.

   Typically these are written to a http.Request.
*/
type CreateProcessorParams struct {

	/* Body.

	   The processor configuration details.
	*/
	Body *models.ProcessorEntity

	/* ID.

	   The process group id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create processor params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProcessorParams) WithDefaults() *CreateProcessorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create processor params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProcessorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create processor params
func (o *CreateProcessorParams) WithTimeout(timeout time.Duration) *CreateProcessorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create processor params
func (o *CreateProcessorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create processor params
func (o *CreateProcessorParams) WithContext(ctx context.Context) *CreateProcessorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create processor params
func (o *CreateProcessorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create processor params
func (o *CreateProcessorParams) WithHTTPClient(client *http.Client) *CreateProcessorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create processor params
func (o *CreateProcessorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create processor params
func (o *CreateProcessorParams) WithBody(body *models.ProcessorEntity) *CreateProcessorParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create processor params
func (o *CreateProcessorParams) SetBody(body *models.ProcessorEntity) {
	o.Body = body
}

// WithID adds the id to the create processor params
func (o *CreateProcessorParams) WithID(id string) *CreateProcessorParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create processor params
func (o *CreateProcessorParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateProcessorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
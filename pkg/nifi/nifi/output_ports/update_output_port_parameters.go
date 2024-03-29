// Code generated by go-swagger; DO NOT EDIT.

package output_ports

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

// NewUpdateOutputPortParams creates a new UpdateOutputPortParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOutputPortParams() *UpdateOutputPortParams {
	return &UpdateOutputPortParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOutputPortParamsWithTimeout creates a new UpdateOutputPortParams object
// with the ability to set a timeout on a request.
func NewUpdateOutputPortParamsWithTimeout(timeout time.Duration) *UpdateOutputPortParams {
	return &UpdateOutputPortParams{
		timeout: timeout,
	}
}

// NewUpdateOutputPortParamsWithContext creates a new UpdateOutputPortParams object
// with the ability to set a context for a request.
func NewUpdateOutputPortParamsWithContext(ctx context.Context) *UpdateOutputPortParams {
	return &UpdateOutputPortParams{
		Context: ctx,
	}
}

// NewUpdateOutputPortParamsWithHTTPClient creates a new UpdateOutputPortParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOutputPortParamsWithHTTPClient(client *http.Client) *UpdateOutputPortParams {
	return &UpdateOutputPortParams{
		HTTPClient: client,
	}
}

/*
UpdateOutputPortParams contains all the parameters to send to the API endpoint

	for the update output port operation.

	Typically these are written to a http.Request.
*/
type UpdateOutputPortParams struct {

	/* Body.

	   The output port configuration details.
	*/
	Body *models.PortEntity

	/* ID.

	   The output port id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update output port params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOutputPortParams) WithDefaults() *UpdateOutputPortParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update output port params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOutputPortParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update output port params
func (o *UpdateOutputPortParams) WithTimeout(timeout time.Duration) *UpdateOutputPortParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update output port params
func (o *UpdateOutputPortParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update output port params
func (o *UpdateOutputPortParams) WithContext(ctx context.Context) *UpdateOutputPortParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update output port params
func (o *UpdateOutputPortParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update output port params
func (o *UpdateOutputPortParams) WithHTTPClient(client *http.Client) *UpdateOutputPortParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update output port params
func (o *UpdateOutputPortParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update output port params
func (o *UpdateOutputPortParams) WithBody(body *models.PortEntity) *UpdateOutputPortParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update output port params
func (o *UpdateOutputPortParams) SetBody(body *models.PortEntity) {
	o.Body = body
}

// WithID adds the id to the update output port params
func (o *UpdateOutputPortParams) WithID(id string) *UpdateOutputPortParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update output port params
func (o *UpdateOutputPortParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOutputPortParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

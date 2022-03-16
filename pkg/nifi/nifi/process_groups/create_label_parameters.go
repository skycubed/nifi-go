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

// NewCreateLabelParams creates a new CreateLabelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateLabelParams() *CreateLabelParams {
	return &CreateLabelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateLabelParamsWithTimeout creates a new CreateLabelParams object
// with the ability to set a timeout on a request.
func NewCreateLabelParamsWithTimeout(timeout time.Duration) *CreateLabelParams {
	return &CreateLabelParams{
		timeout: timeout,
	}
}

// NewCreateLabelParamsWithContext creates a new CreateLabelParams object
// with the ability to set a context for a request.
func NewCreateLabelParamsWithContext(ctx context.Context) *CreateLabelParams {
	return &CreateLabelParams{
		Context: ctx,
	}
}

// NewCreateLabelParamsWithHTTPClient creates a new CreateLabelParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateLabelParamsWithHTTPClient(client *http.Client) *CreateLabelParams {
	return &CreateLabelParams{
		HTTPClient: client,
	}
}

/* CreateLabelParams contains all the parameters to send to the API endpoint
   for the create label operation.

   Typically these are written to a http.Request.
*/
type CreateLabelParams struct {

	/* Body.

	   The label configuration details.
	*/
	Body *models.LabelEntity

	/* ID.

	   The process group id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create label params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateLabelParams) WithDefaults() *CreateLabelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create label params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateLabelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create label params
func (o *CreateLabelParams) WithTimeout(timeout time.Duration) *CreateLabelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create label params
func (o *CreateLabelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create label params
func (o *CreateLabelParams) WithContext(ctx context.Context) *CreateLabelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create label params
func (o *CreateLabelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create label params
func (o *CreateLabelParams) WithHTTPClient(client *http.Client) *CreateLabelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create label params
func (o *CreateLabelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create label params
func (o *CreateLabelParams) WithBody(body *models.LabelEntity) *CreateLabelParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create label params
func (o *CreateLabelParams) SetBody(body *models.LabelEntity) {
	o.Body = body
}

// WithID adds the id to the create label params
func (o *CreateLabelParams) WithID(id string) *CreateLabelParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create label params
func (o *CreateLabelParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateLabelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

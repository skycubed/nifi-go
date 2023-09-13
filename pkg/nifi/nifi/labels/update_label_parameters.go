// Code generated by go-swagger; DO NOT EDIT.

package labels

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

// NewUpdateLabelParams creates a new UpdateLabelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateLabelParams() *UpdateLabelParams {
	return &UpdateLabelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLabelParamsWithTimeout creates a new UpdateLabelParams object
// with the ability to set a timeout on a request.
func NewUpdateLabelParamsWithTimeout(timeout time.Duration) *UpdateLabelParams {
	return &UpdateLabelParams{
		timeout: timeout,
	}
}

// NewUpdateLabelParamsWithContext creates a new UpdateLabelParams object
// with the ability to set a context for a request.
func NewUpdateLabelParamsWithContext(ctx context.Context) *UpdateLabelParams {
	return &UpdateLabelParams{
		Context: ctx,
	}
}

// NewUpdateLabelParamsWithHTTPClient creates a new UpdateLabelParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateLabelParamsWithHTTPClient(client *http.Client) *UpdateLabelParams {
	return &UpdateLabelParams{
		HTTPClient: client,
	}
}

/*
UpdateLabelParams contains all the parameters to send to the API endpoint

	for the update label operation.

	Typically these are written to a http.Request.
*/
type UpdateLabelParams struct {

	/* Body.

	   The label configuration details.
	*/
	Body *models.LabelEntity

	/* ID.

	   The label id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update label params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLabelParams) WithDefaults() *UpdateLabelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update label params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLabelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update label params
func (o *UpdateLabelParams) WithTimeout(timeout time.Duration) *UpdateLabelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update label params
func (o *UpdateLabelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update label params
func (o *UpdateLabelParams) WithContext(ctx context.Context) *UpdateLabelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update label params
func (o *UpdateLabelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update label params
func (o *UpdateLabelParams) WithHTTPClient(client *http.Client) *UpdateLabelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update label params
func (o *UpdateLabelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update label params
func (o *UpdateLabelParams) WithBody(body *models.LabelEntity) *UpdateLabelParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update label params
func (o *UpdateLabelParams) SetBody(body *models.LabelEntity) {
	o.Body = body
}

// WithID adds the id to the update label params
func (o *UpdateLabelParams) WithID(id string) *UpdateLabelParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update label params
func (o *UpdateLabelParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLabelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

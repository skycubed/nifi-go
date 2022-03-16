// Code generated by go-swagger; DO NOT EDIT.

package controller_services

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

// NewUpdateControllerServiceReferencesParams creates a new UpdateControllerServiceReferencesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateControllerServiceReferencesParams() *UpdateControllerServiceReferencesParams {
	return &UpdateControllerServiceReferencesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateControllerServiceReferencesParamsWithTimeout creates a new UpdateControllerServiceReferencesParams object
// with the ability to set a timeout on a request.
func NewUpdateControllerServiceReferencesParamsWithTimeout(timeout time.Duration) *UpdateControllerServiceReferencesParams {
	return &UpdateControllerServiceReferencesParams{
		timeout: timeout,
	}
}

// NewUpdateControllerServiceReferencesParamsWithContext creates a new UpdateControllerServiceReferencesParams object
// with the ability to set a context for a request.
func NewUpdateControllerServiceReferencesParamsWithContext(ctx context.Context) *UpdateControllerServiceReferencesParams {
	return &UpdateControllerServiceReferencesParams{
		Context: ctx,
	}
}

// NewUpdateControllerServiceReferencesParamsWithHTTPClient creates a new UpdateControllerServiceReferencesParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateControllerServiceReferencesParamsWithHTTPClient(client *http.Client) *UpdateControllerServiceReferencesParams {
	return &UpdateControllerServiceReferencesParams{
		HTTPClient: client,
	}
}

/* UpdateControllerServiceReferencesParams contains all the parameters to send to the API endpoint
   for the update controller service references operation.

   Typically these are written to a http.Request.
*/
type UpdateControllerServiceReferencesParams struct {

	/* Body.

	   The controller service request update request.
	*/
	Body *models.UpdateControllerServiceReferenceRequestEntity

	/* ID.

	   The controller service id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update controller service references params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateControllerServiceReferencesParams) WithDefaults() *UpdateControllerServiceReferencesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update controller service references params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateControllerServiceReferencesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update controller service references params
func (o *UpdateControllerServiceReferencesParams) WithTimeout(timeout time.Duration) *UpdateControllerServiceReferencesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update controller service references params
func (o *UpdateControllerServiceReferencesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update controller service references params
func (o *UpdateControllerServiceReferencesParams) WithContext(ctx context.Context) *UpdateControllerServiceReferencesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update controller service references params
func (o *UpdateControllerServiceReferencesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update controller service references params
func (o *UpdateControllerServiceReferencesParams) WithHTTPClient(client *http.Client) *UpdateControllerServiceReferencesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update controller service references params
func (o *UpdateControllerServiceReferencesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update controller service references params
func (o *UpdateControllerServiceReferencesParams) WithBody(body *models.UpdateControllerServiceReferenceRequestEntity) *UpdateControllerServiceReferencesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update controller service references params
func (o *UpdateControllerServiceReferencesParams) SetBody(body *models.UpdateControllerServiceReferenceRequestEntity) {
	o.Body = body
}

// WithID adds the id to the update controller service references params
func (o *UpdateControllerServiceReferencesParams) WithID(id string) *UpdateControllerServiceReferencesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update controller service references params
func (o *UpdateControllerServiceReferencesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateControllerServiceReferencesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
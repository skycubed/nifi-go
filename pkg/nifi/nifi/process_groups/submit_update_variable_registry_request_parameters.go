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

// NewSubmitUpdateVariableRegistryRequestParams creates a new SubmitUpdateVariableRegistryRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubmitUpdateVariableRegistryRequestParams() *SubmitUpdateVariableRegistryRequestParams {
	return &SubmitUpdateVariableRegistryRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubmitUpdateVariableRegistryRequestParamsWithTimeout creates a new SubmitUpdateVariableRegistryRequestParams object
// with the ability to set a timeout on a request.
func NewSubmitUpdateVariableRegistryRequestParamsWithTimeout(timeout time.Duration) *SubmitUpdateVariableRegistryRequestParams {
	return &SubmitUpdateVariableRegistryRequestParams{
		timeout: timeout,
	}
}

// NewSubmitUpdateVariableRegistryRequestParamsWithContext creates a new SubmitUpdateVariableRegistryRequestParams object
// with the ability to set a context for a request.
func NewSubmitUpdateVariableRegistryRequestParamsWithContext(ctx context.Context) *SubmitUpdateVariableRegistryRequestParams {
	return &SubmitUpdateVariableRegistryRequestParams{
		Context: ctx,
	}
}

// NewSubmitUpdateVariableRegistryRequestParamsWithHTTPClient creates a new SubmitUpdateVariableRegistryRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubmitUpdateVariableRegistryRequestParamsWithHTTPClient(client *http.Client) *SubmitUpdateVariableRegistryRequestParams {
	return &SubmitUpdateVariableRegistryRequestParams{
		HTTPClient: client,
	}
}

/* SubmitUpdateVariableRegistryRequestParams contains all the parameters to send to the API endpoint
   for the submit update variable registry request operation.

   Typically these are written to a http.Request.
*/
type SubmitUpdateVariableRegistryRequestParams struct {

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

// WithDefaults hydrates default values in the submit update variable registry request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitUpdateVariableRegistryRequestParams) WithDefaults() *SubmitUpdateVariableRegistryRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the submit update variable registry request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitUpdateVariableRegistryRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the submit update variable registry request params
func (o *SubmitUpdateVariableRegistryRequestParams) WithTimeout(timeout time.Duration) *SubmitUpdateVariableRegistryRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the submit update variable registry request params
func (o *SubmitUpdateVariableRegistryRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the submit update variable registry request params
func (o *SubmitUpdateVariableRegistryRequestParams) WithContext(ctx context.Context) *SubmitUpdateVariableRegistryRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the submit update variable registry request params
func (o *SubmitUpdateVariableRegistryRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the submit update variable registry request params
func (o *SubmitUpdateVariableRegistryRequestParams) WithHTTPClient(client *http.Client) *SubmitUpdateVariableRegistryRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the submit update variable registry request params
func (o *SubmitUpdateVariableRegistryRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the submit update variable registry request params
func (o *SubmitUpdateVariableRegistryRequestParams) WithBody(body *models.VariableRegistryEntity) *SubmitUpdateVariableRegistryRequestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the submit update variable registry request params
func (o *SubmitUpdateVariableRegistryRequestParams) SetBody(body *models.VariableRegistryEntity) {
	o.Body = body
}

// WithID adds the id to the submit update variable registry request params
func (o *SubmitUpdateVariableRegistryRequestParams) WithID(id string) *SubmitUpdateVariableRegistryRequestParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the submit update variable registry request params
func (o *SubmitUpdateVariableRegistryRequestParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SubmitUpdateVariableRegistryRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
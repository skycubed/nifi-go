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

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// NewUpdateReportingTaskParams creates a new UpdateReportingTaskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateReportingTaskParams() *UpdateReportingTaskParams {
	return &UpdateReportingTaskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateReportingTaskParamsWithTimeout creates a new UpdateReportingTaskParams object
// with the ability to set a timeout on a request.
func NewUpdateReportingTaskParamsWithTimeout(timeout time.Duration) *UpdateReportingTaskParams {
	return &UpdateReportingTaskParams{
		timeout: timeout,
	}
}

// NewUpdateReportingTaskParamsWithContext creates a new UpdateReportingTaskParams object
// with the ability to set a context for a request.
func NewUpdateReportingTaskParamsWithContext(ctx context.Context) *UpdateReportingTaskParams {
	return &UpdateReportingTaskParams{
		Context: ctx,
	}
}

// NewUpdateReportingTaskParamsWithHTTPClient creates a new UpdateReportingTaskParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateReportingTaskParamsWithHTTPClient(client *http.Client) *UpdateReportingTaskParams {
	return &UpdateReportingTaskParams{
		HTTPClient: client,
	}
}

/* UpdateReportingTaskParams contains all the parameters to send to the API endpoint
   for the update reporting task operation.

   Typically these are written to a http.Request.
*/
type UpdateReportingTaskParams struct {

	/* Body.

	   The reporting task configuration details.
	*/
	Body *models.ReportingTaskEntity

	/* ID.

	   The reporting task id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update reporting task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateReportingTaskParams) WithDefaults() *UpdateReportingTaskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update reporting task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateReportingTaskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update reporting task params
func (o *UpdateReportingTaskParams) WithTimeout(timeout time.Duration) *UpdateReportingTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update reporting task params
func (o *UpdateReportingTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update reporting task params
func (o *UpdateReportingTaskParams) WithContext(ctx context.Context) *UpdateReportingTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update reporting task params
func (o *UpdateReportingTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update reporting task params
func (o *UpdateReportingTaskParams) WithHTTPClient(client *http.Client) *UpdateReportingTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update reporting task params
func (o *UpdateReportingTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update reporting task params
func (o *UpdateReportingTaskParams) WithBody(body *models.ReportingTaskEntity) *UpdateReportingTaskParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update reporting task params
func (o *UpdateReportingTaskParams) SetBody(body *models.ReportingTaskEntity) {
	o.Body = body
}

// WithID adds the id to the update reporting task params
func (o *UpdateReportingTaskParams) WithID(id string) *UpdateReportingTaskParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update reporting task params
func (o *UpdateReportingTaskParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateReportingTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
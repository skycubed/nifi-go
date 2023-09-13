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

// NewCreateProcessGroupParams creates a new CreateProcessGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateProcessGroupParams() *CreateProcessGroupParams {
	return &CreateProcessGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProcessGroupParamsWithTimeout creates a new CreateProcessGroupParams object
// with the ability to set a timeout on a request.
func NewCreateProcessGroupParamsWithTimeout(timeout time.Duration) *CreateProcessGroupParams {
	return &CreateProcessGroupParams{
		timeout: timeout,
	}
}

// NewCreateProcessGroupParamsWithContext creates a new CreateProcessGroupParams object
// with the ability to set a context for a request.
func NewCreateProcessGroupParamsWithContext(ctx context.Context) *CreateProcessGroupParams {
	return &CreateProcessGroupParams{
		Context: ctx,
	}
}

// NewCreateProcessGroupParamsWithHTTPClient creates a new CreateProcessGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateProcessGroupParamsWithHTTPClient(client *http.Client) *CreateProcessGroupParams {
	return &CreateProcessGroupParams{
		HTTPClient: client,
	}
}

/*
CreateProcessGroupParams contains all the parameters to send to the API endpoint

	for the create process group operation.

	Typically these are written to a http.Request.
*/
type CreateProcessGroupParams struct {

	/* Body.

	   The process group configuration details.
	*/
	Body *models.ProcessGroupEntity

	/* ID.

	   The process group id.
	*/
	ID string

	/* ParameterContextHandlingStrategy.

	   Handling Strategy controls whether to keep or replace Parameter Contexts

	   Default: "KEEP_EXISTING"
	*/
	ParameterContextHandlingStrategy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create process group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProcessGroupParams) WithDefaults() *CreateProcessGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create process group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProcessGroupParams) SetDefaults() {
	var (
		parameterContextHandlingStrategyDefault = string("KEEP_EXISTING")
	)

	val := CreateProcessGroupParams{
		ParameterContextHandlingStrategy: &parameterContextHandlingStrategyDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create process group params
func (o *CreateProcessGroupParams) WithTimeout(timeout time.Duration) *CreateProcessGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create process group params
func (o *CreateProcessGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create process group params
func (o *CreateProcessGroupParams) WithContext(ctx context.Context) *CreateProcessGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create process group params
func (o *CreateProcessGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create process group params
func (o *CreateProcessGroupParams) WithHTTPClient(client *http.Client) *CreateProcessGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create process group params
func (o *CreateProcessGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create process group params
func (o *CreateProcessGroupParams) WithBody(body *models.ProcessGroupEntity) *CreateProcessGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create process group params
func (o *CreateProcessGroupParams) SetBody(body *models.ProcessGroupEntity) {
	o.Body = body
}

// WithID adds the id to the create process group params
func (o *CreateProcessGroupParams) WithID(id string) *CreateProcessGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create process group params
func (o *CreateProcessGroupParams) SetID(id string) {
	o.ID = id
}

// WithParameterContextHandlingStrategy adds the parameterContextHandlingStrategy to the create process group params
func (o *CreateProcessGroupParams) WithParameterContextHandlingStrategy(parameterContextHandlingStrategy *string) *CreateProcessGroupParams {
	o.SetParameterContextHandlingStrategy(parameterContextHandlingStrategy)
	return o
}

// SetParameterContextHandlingStrategy adds the parameterContextHandlingStrategy to the create process group params
func (o *CreateProcessGroupParams) SetParameterContextHandlingStrategy(parameterContextHandlingStrategy *string) {
	o.ParameterContextHandlingStrategy = parameterContextHandlingStrategy
}

// WriteToRequest writes these params to a swagger request
func (o *CreateProcessGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ParameterContextHandlingStrategy != nil {

		// query param parameterContextHandlingStrategy
		var qrParameterContextHandlingStrategy string

		if o.ParameterContextHandlingStrategy != nil {
			qrParameterContextHandlingStrategy = *o.ParameterContextHandlingStrategy
		}
		qParameterContextHandlingStrategy := qrParameterContextHandlingStrategy
		if qParameterContextHandlingStrategy != "" {

			if err := r.SetQueryParam("parameterContextHandlingStrategy", qParameterContextHandlingStrategy); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

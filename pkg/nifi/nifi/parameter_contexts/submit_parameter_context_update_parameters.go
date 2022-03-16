// Code generated by go-swagger; DO NOT EDIT.

package parameter_contexts

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

// NewSubmitParameterContextUpdateParams creates a new SubmitParameterContextUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubmitParameterContextUpdateParams() *SubmitParameterContextUpdateParams {
	return &SubmitParameterContextUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubmitParameterContextUpdateParamsWithTimeout creates a new SubmitParameterContextUpdateParams object
// with the ability to set a timeout on a request.
func NewSubmitParameterContextUpdateParamsWithTimeout(timeout time.Duration) *SubmitParameterContextUpdateParams {
	return &SubmitParameterContextUpdateParams{
		timeout: timeout,
	}
}

// NewSubmitParameterContextUpdateParamsWithContext creates a new SubmitParameterContextUpdateParams object
// with the ability to set a context for a request.
func NewSubmitParameterContextUpdateParamsWithContext(ctx context.Context) *SubmitParameterContextUpdateParams {
	return &SubmitParameterContextUpdateParams{
		Context: ctx,
	}
}

// NewSubmitParameterContextUpdateParamsWithHTTPClient creates a new SubmitParameterContextUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubmitParameterContextUpdateParamsWithHTTPClient(client *http.Client) *SubmitParameterContextUpdateParams {
	return &SubmitParameterContextUpdateParams{
		HTTPClient: client,
	}
}

/* SubmitParameterContextUpdateParams contains all the parameters to send to the API endpoint
   for the submit parameter context update operation.

   Typically these are written to a http.Request.
*/
type SubmitParameterContextUpdateParams struct {

	/* Body.

	   The updated version of the parameter context.
	*/
	Body *models.ParameterContextEntity

	// ContextID.
	ContextID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the submit parameter context update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitParameterContextUpdateParams) WithDefaults() *SubmitParameterContextUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the submit parameter context update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitParameterContextUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the submit parameter context update params
func (o *SubmitParameterContextUpdateParams) WithTimeout(timeout time.Duration) *SubmitParameterContextUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the submit parameter context update params
func (o *SubmitParameterContextUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the submit parameter context update params
func (o *SubmitParameterContextUpdateParams) WithContext(ctx context.Context) *SubmitParameterContextUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the submit parameter context update params
func (o *SubmitParameterContextUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the submit parameter context update params
func (o *SubmitParameterContextUpdateParams) WithHTTPClient(client *http.Client) *SubmitParameterContextUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the submit parameter context update params
func (o *SubmitParameterContextUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the submit parameter context update params
func (o *SubmitParameterContextUpdateParams) WithBody(body *models.ParameterContextEntity) *SubmitParameterContextUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the submit parameter context update params
func (o *SubmitParameterContextUpdateParams) SetBody(body *models.ParameterContextEntity) {
	o.Body = body
}

// WithContextID adds the contextID to the submit parameter context update params
func (o *SubmitParameterContextUpdateParams) WithContextID(contextID string) *SubmitParameterContextUpdateParams {
	o.SetContextID(contextID)
	return o
}

// SetContextID adds the contextId to the submit parameter context update params
func (o *SubmitParameterContextUpdateParams) SetContextID(contextID string) {
	o.ContextID = contextID
}

// WriteToRequest writes these params to a swagger request
func (o *SubmitParameterContextUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param contextId
	if err := r.SetPathParam("contextId", o.ContextID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

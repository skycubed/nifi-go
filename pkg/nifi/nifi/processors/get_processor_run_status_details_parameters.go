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

// NewGetProcessorRunStatusDetailsParams creates a new GetProcessorRunStatusDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProcessorRunStatusDetailsParams() *GetProcessorRunStatusDetailsParams {
	return &GetProcessorRunStatusDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProcessorRunStatusDetailsParamsWithTimeout creates a new GetProcessorRunStatusDetailsParams object
// with the ability to set a timeout on a request.
func NewGetProcessorRunStatusDetailsParamsWithTimeout(timeout time.Duration) *GetProcessorRunStatusDetailsParams {
	return &GetProcessorRunStatusDetailsParams{
		timeout: timeout,
	}
}

// NewGetProcessorRunStatusDetailsParamsWithContext creates a new GetProcessorRunStatusDetailsParams object
// with the ability to set a context for a request.
func NewGetProcessorRunStatusDetailsParamsWithContext(ctx context.Context) *GetProcessorRunStatusDetailsParams {
	return &GetProcessorRunStatusDetailsParams{
		Context: ctx,
	}
}

// NewGetProcessorRunStatusDetailsParamsWithHTTPClient creates a new GetProcessorRunStatusDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProcessorRunStatusDetailsParamsWithHTTPClient(client *http.Client) *GetProcessorRunStatusDetailsParams {
	return &GetProcessorRunStatusDetailsParams{
		HTTPClient: client,
	}
}

/*
GetProcessorRunStatusDetailsParams contains all the parameters to send to the API endpoint

	for the get processor run status details operation.

	Typically these are written to a http.Request.
*/
type GetProcessorRunStatusDetailsParams struct {

	/* Body.

	   The request for the processors that should be included in the results
	*/
	Body *models.RunStatusDetailsRequestEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get processor run status details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProcessorRunStatusDetailsParams) WithDefaults() *GetProcessorRunStatusDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get processor run status details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProcessorRunStatusDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get processor run status details params
func (o *GetProcessorRunStatusDetailsParams) WithTimeout(timeout time.Duration) *GetProcessorRunStatusDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get processor run status details params
func (o *GetProcessorRunStatusDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get processor run status details params
func (o *GetProcessorRunStatusDetailsParams) WithContext(ctx context.Context) *GetProcessorRunStatusDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get processor run status details params
func (o *GetProcessorRunStatusDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get processor run status details params
func (o *GetProcessorRunStatusDetailsParams) WithHTTPClient(client *http.Client) *GetProcessorRunStatusDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get processor run status details params
func (o *GetProcessorRunStatusDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get processor run status details params
func (o *GetProcessorRunStatusDetailsParams) WithBody(body *models.RunStatusDetailsRequestEntity) *GetProcessorRunStatusDetailsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get processor run status details params
func (o *GetProcessorRunStatusDetailsParams) SetBody(body *models.RunStatusDetailsRequestEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetProcessorRunStatusDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package flow

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
)

// NewGetFlowMetricsParams creates a new GetFlowMetricsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFlowMetricsParams() *GetFlowMetricsParams {
	return &GetFlowMetricsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFlowMetricsParamsWithTimeout creates a new GetFlowMetricsParams object
// with the ability to set a timeout on a request.
func NewGetFlowMetricsParamsWithTimeout(timeout time.Duration) *GetFlowMetricsParams {
	return &GetFlowMetricsParams{
		timeout: timeout,
	}
}

// NewGetFlowMetricsParamsWithContext creates a new GetFlowMetricsParams object
// with the ability to set a context for a request.
func NewGetFlowMetricsParamsWithContext(ctx context.Context) *GetFlowMetricsParams {
	return &GetFlowMetricsParams{
		Context: ctx,
	}
}

// NewGetFlowMetricsParamsWithHTTPClient creates a new GetFlowMetricsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFlowMetricsParamsWithHTTPClient(client *http.Client) *GetFlowMetricsParams {
	return &GetFlowMetricsParams{
		HTTPClient: client,
	}
}

/* GetFlowMetricsParams contains all the parameters to send to the API endpoint
   for the get flow metrics operation.

   Typically these are written to a http.Request.
*/
type GetFlowMetricsParams struct {

	/* Producer.

	   The producer for flow file metrics. Each producer may have its own output format.
	*/
	Producer string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get flow metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFlowMetricsParams) WithDefaults() *GetFlowMetricsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get flow metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFlowMetricsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get flow metrics params
func (o *GetFlowMetricsParams) WithTimeout(timeout time.Duration) *GetFlowMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get flow metrics params
func (o *GetFlowMetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get flow metrics params
func (o *GetFlowMetricsParams) WithContext(ctx context.Context) *GetFlowMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get flow metrics params
func (o *GetFlowMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get flow metrics params
func (o *GetFlowMetricsParams) WithHTTPClient(client *http.Client) *GetFlowMetricsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get flow metrics params
func (o *GetFlowMetricsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProducer adds the producer to the get flow metrics params
func (o *GetFlowMetricsParams) WithProducer(producer string) *GetFlowMetricsParams {
	o.SetProducer(producer)
	return o
}

// SetProducer adds the producer to the get flow metrics params
func (o *GetFlowMetricsParams) SetProducer(producer string) {
	o.Producer = producer
}

// WriteToRequest writes these params to a swagger request
func (o *GetFlowMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param producer
	if err := r.SetPathParam("producer", o.Producer); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

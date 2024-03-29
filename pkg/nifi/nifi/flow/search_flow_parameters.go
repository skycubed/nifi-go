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

// NewSearchFlowParams creates a new SearchFlowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchFlowParams() *SearchFlowParams {
	return &SearchFlowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchFlowParamsWithTimeout creates a new SearchFlowParams object
// with the ability to set a timeout on a request.
func NewSearchFlowParamsWithTimeout(timeout time.Duration) *SearchFlowParams {
	return &SearchFlowParams{
		timeout: timeout,
	}
}

// NewSearchFlowParamsWithContext creates a new SearchFlowParams object
// with the ability to set a context for a request.
func NewSearchFlowParamsWithContext(ctx context.Context) *SearchFlowParams {
	return &SearchFlowParams{
		Context: ctx,
	}
}

// NewSearchFlowParamsWithHTTPClient creates a new SearchFlowParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchFlowParamsWithHTTPClient(client *http.Client) *SearchFlowParams {
	return &SearchFlowParams{
		HTTPClient: client,
	}
}

/*
SearchFlowParams contains all the parameters to send to the API endpoint

	for the search flow operation.

	Typically these are written to a http.Request.
*/
type SearchFlowParams struct {

	// A.
	A *string

	// Q.
	Q *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchFlowParams) WithDefaults() *SearchFlowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchFlowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search flow params
func (o *SearchFlowParams) WithTimeout(timeout time.Duration) *SearchFlowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search flow params
func (o *SearchFlowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search flow params
func (o *SearchFlowParams) WithContext(ctx context.Context) *SearchFlowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search flow params
func (o *SearchFlowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search flow params
func (o *SearchFlowParams) WithHTTPClient(client *http.Client) *SearchFlowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search flow params
func (o *SearchFlowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithA adds the a to the search flow params
func (o *SearchFlowParams) WithA(a *string) *SearchFlowParams {
	o.SetA(a)
	return o
}

// SetA adds the a to the search flow params
func (o *SearchFlowParams) SetA(a *string) {
	o.A = a
}

// WithQ adds the q to the search flow params
func (o *SearchFlowParams) WithQ(q *string) *SearchFlowParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the search flow params
func (o *SearchFlowParams) SetQ(q *string) {
	o.Q = q
}

// WriteToRequest writes these params to a swagger request
func (o *SearchFlowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.A != nil {

		// query param a
		var qrA string

		if o.A != nil {
			qrA = *o.A
		}
		qA := qrA
		if qA != "" {

			if err := r.SetQueryParam("a", qA); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

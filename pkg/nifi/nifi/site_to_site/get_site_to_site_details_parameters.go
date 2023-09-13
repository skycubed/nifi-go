// Code generated by go-swagger; DO NOT EDIT.

package site_to_site

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

// NewGetSiteToSiteDetailsParams creates a new GetSiteToSiteDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSiteToSiteDetailsParams() *GetSiteToSiteDetailsParams {
	return &GetSiteToSiteDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSiteToSiteDetailsParamsWithTimeout creates a new GetSiteToSiteDetailsParams object
// with the ability to set a timeout on a request.
func NewGetSiteToSiteDetailsParamsWithTimeout(timeout time.Duration) *GetSiteToSiteDetailsParams {
	return &GetSiteToSiteDetailsParams{
		timeout: timeout,
	}
}

// NewGetSiteToSiteDetailsParamsWithContext creates a new GetSiteToSiteDetailsParams object
// with the ability to set a context for a request.
func NewGetSiteToSiteDetailsParamsWithContext(ctx context.Context) *GetSiteToSiteDetailsParams {
	return &GetSiteToSiteDetailsParams{
		Context: ctx,
	}
}

// NewGetSiteToSiteDetailsParamsWithHTTPClient creates a new GetSiteToSiteDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSiteToSiteDetailsParamsWithHTTPClient(client *http.Client) *GetSiteToSiteDetailsParams {
	return &GetSiteToSiteDetailsParams{
		HTTPClient: client,
	}
}

/*
GetSiteToSiteDetailsParams contains all the parameters to send to the API endpoint

	for the get site to site details operation.

	Typically these are written to a http.Request.
*/
type GetSiteToSiteDetailsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get site to site details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSiteToSiteDetailsParams) WithDefaults() *GetSiteToSiteDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get site to site details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSiteToSiteDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get site to site details params
func (o *GetSiteToSiteDetailsParams) WithTimeout(timeout time.Duration) *GetSiteToSiteDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get site to site details params
func (o *GetSiteToSiteDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get site to site details params
func (o *GetSiteToSiteDetailsParams) WithContext(ctx context.Context) *GetSiteToSiteDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get site to site details params
func (o *GetSiteToSiteDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get site to site details params
func (o *GetSiteToSiteDetailsParams) WithHTTPClient(client *http.Client) *GetSiteToSiteDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get site to site details params
func (o *GetSiteToSiteDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSiteToSiteDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

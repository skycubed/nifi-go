// Code generated by go-swagger; DO NOT EDIT.

package controller

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

// NewDeleteNodeParams creates a new DeleteNodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNodeParams() *DeleteNodeParams {
	return &DeleteNodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNodeParamsWithTimeout creates a new DeleteNodeParams object
// with the ability to set a timeout on a request.
func NewDeleteNodeParamsWithTimeout(timeout time.Duration) *DeleteNodeParams {
	return &DeleteNodeParams{
		timeout: timeout,
	}
}

// NewDeleteNodeParamsWithContext creates a new DeleteNodeParams object
// with the ability to set a context for a request.
func NewDeleteNodeParamsWithContext(ctx context.Context) *DeleteNodeParams {
	return &DeleteNodeParams{
		Context: ctx,
	}
}

// NewDeleteNodeParamsWithHTTPClient creates a new DeleteNodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNodeParamsWithHTTPClient(client *http.Client) *DeleteNodeParams {
	return &DeleteNodeParams{
		HTTPClient: client,
	}
}

/*
DeleteNodeParams contains all the parameters to send to the API endpoint

	for the delete node operation.

	Typically these are written to a http.Request.
*/
type DeleteNodeParams struct {

	/* ID.

	   The node id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNodeParams) WithDefaults() *DeleteNodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete node params
func (o *DeleteNodeParams) WithTimeout(timeout time.Duration) *DeleteNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete node params
func (o *DeleteNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete node params
func (o *DeleteNodeParams) WithContext(ctx context.Context) *DeleteNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete node params
func (o *DeleteNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete node params
func (o *DeleteNodeParams) WithHTTPClient(client *http.Client) *DeleteNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete node params
func (o *DeleteNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete node params
func (o *DeleteNodeParams) WithID(id string) *DeleteNodeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete node params
func (o *DeleteNodeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

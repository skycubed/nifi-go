// Code generated by go-swagger; DO NOT EDIT.

package versions

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

// NewCreateVersionControlRequestParams creates a new CreateVersionControlRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateVersionControlRequestParams() *CreateVersionControlRequestParams {
	return &CreateVersionControlRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVersionControlRequestParamsWithTimeout creates a new CreateVersionControlRequestParams object
// with the ability to set a timeout on a request.
func NewCreateVersionControlRequestParamsWithTimeout(timeout time.Duration) *CreateVersionControlRequestParams {
	return &CreateVersionControlRequestParams{
		timeout: timeout,
	}
}

// NewCreateVersionControlRequestParamsWithContext creates a new CreateVersionControlRequestParams object
// with the ability to set a context for a request.
func NewCreateVersionControlRequestParamsWithContext(ctx context.Context) *CreateVersionControlRequestParams {
	return &CreateVersionControlRequestParams{
		Context: ctx,
	}
}

// NewCreateVersionControlRequestParamsWithHTTPClient creates a new CreateVersionControlRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateVersionControlRequestParamsWithHTTPClient(client *http.Client) *CreateVersionControlRequestParams {
	return &CreateVersionControlRequestParams{
		HTTPClient: client,
	}
}

/*
CreateVersionControlRequestParams contains all the parameters to send to the API endpoint

	for the create version control request operation.

	Typically these are written to a http.Request.
*/
type CreateVersionControlRequestParams struct {

	/* Body.

	   The versioned flow details.
	*/
	Body *models.CreateActiveRequestEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create version control request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVersionControlRequestParams) WithDefaults() *CreateVersionControlRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create version control request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVersionControlRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create version control request params
func (o *CreateVersionControlRequestParams) WithTimeout(timeout time.Duration) *CreateVersionControlRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create version control request params
func (o *CreateVersionControlRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create version control request params
func (o *CreateVersionControlRequestParams) WithContext(ctx context.Context) *CreateVersionControlRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create version control request params
func (o *CreateVersionControlRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create version control request params
func (o *CreateVersionControlRequestParams) WithHTTPClient(client *http.Client) *CreateVersionControlRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create version control request params
func (o *CreateVersionControlRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create version control request params
func (o *CreateVersionControlRequestParams) WithBody(body *models.CreateActiveRequestEntity) *CreateVersionControlRequestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create version control request params
func (o *CreateVersionControlRequestParams) SetBody(body *models.CreateActiveRequestEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVersionControlRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

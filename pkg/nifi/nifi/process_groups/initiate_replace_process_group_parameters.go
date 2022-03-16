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

// NewInitiateReplaceProcessGroupParams creates a new InitiateReplaceProcessGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInitiateReplaceProcessGroupParams() *InitiateReplaceProcessGroupParams {
	return &InitiateReplaceProcessGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInitiateReplaceProcessGroupParamsWithTimeout creates a new InitiateReplaceProcessGroupParams object
// with the ability to set a timeout on a request.
func NewInitiateReplaceProcessGroupParamsWithTimeout(timeout time.Duration) *InitiateReplaceProcessGroupParams {
	return &InitiateReplaceProcessGroupParams{
		timeout: timeout,
	}
}

// NewInitiateReplaceProcessGroupParamsWithContext creates a new InitiateReplaceProcessGroupParams object
// with the ability to set a context for a request.
func NewInitiateReplaceProcessGroupParamsWithContext(ctx context.Context) *InitiateReplaceProcessGroupParams {
	return &InitiateReplaceProcessGroupParams{
		Context: ctx,
	}
}

// NewInitiateReplaceProcessGroupParamsWithHTTPClient creates a new InitiateReplaceProcessGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewInitiateReplaceProcessGroupParamsWithHTTPClient(client *http.Client) *InitiateReplaceProcessGroupParams {
	return &InitiateReplaceProcessGroupParams{
		HTTPClient: client,
	}
}

/* InitiateReplaceProcessGroupParams contains all the parameters to send to the API endpoint
   for the initiate replace process group operation.

   Typically these are written to a http.Request.
*/
type InitiateReplaceProcessGroupParams struct {

	/* Body.

	   The process group replace request entity
	*/
	Body *models.ProcessGroupImportEntity

	/* ID.

	   The process group id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the initiate replace process group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitiateReplaceProcessGroupParams) WithDefaults() *InitiateReplaceProcessGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the initiate replace process group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitiateReplaceProcessGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the initiate replace process group params
func (o *InitiateReplaceProcessGroupParams) WithTimeout(timeout time.Duration) *InitiateReplaceProcessGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the initiate replace process group params
func (o *InitiateReplaceProcessGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the initiate replace process group params
func (o *InitiateReplaceProcessGroupParams) WithContext(ctx context.Context) *InitiateReplaceProcessGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the initiate replace process group params
func (o *InitiateReplaceProcessGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the initiate replace process group params
func (o *InitiateReplaceProcessGroupParams) WithHTTPClient(client *http.Client) *InitiateReplaceProcessGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the initiate replace process group params
func (o *InitiateReplaceProcessGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the initiate replace process group params
func (o *InitiateReplaceProcessGroupParams) WithBody(body *models.ProcessGroupImportEntity) *InitiateReplaceProcessGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the initiate replace process group params
func (o *InitiateReplaceProcessGroupParams) SetBody(body *models.ProcessGroupImportEntity) {
	o.Body = body
}

// WithID adds the id to the initiate replace process group params
func (o *InitiateReplaceProcessGroupParams) WithID(id string) *InitiateReplaceProcessGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the initiate replace process group params
func (o *InitiateReplaceProcessGroupParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *InitiateReplaceProcessGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
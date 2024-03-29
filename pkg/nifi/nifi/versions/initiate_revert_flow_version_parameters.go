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

// NewInitiateRevertFlowVersionParams creates a new InitiateRevertFlowVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInitiateRevertFlowVersionParams() *InitiateRevertFlowVersionParams {
	return &InitiateRevertFlowVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInitiateRevertFlowVersionParamsWithTimeout creates a new InitiateRevertFlowVersionParams object
// with the ability to set a timeout on a request.
func NewInitiateRevertFlowVersionParamsWithTimeout(timeout time.Duration) *InitiateRevertFlowVersionParams {
	return &InitiateRevertFlowVersionParams{
		timeout: timeout,
	}
}

// NewInitiateRevertFlowVersionParamsWithContext creates a new InitiateRevertFlowVersionParams object
// with the ability to set a context for a request.
func NewInitiateRevertFlowVersionParamsWithContext(ctx context.Context) *InitiateRevertFlowVersionParams {
	return &InitiateRevertFlowVersionParams{
		Context: ctx,
	}
}

// NewInitiateRevertFlowVersionParamsWithHTTPClient creates a new InitiateRevertFlowVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewInitiateRevertFlowVersionParamsWithHTTPClient(client *http.Client) *InitiateRevertFlowVersionParams {
	return &InitiateRevertFlowVersionParams{
		HTTPClient: client,
	}
}

/*
InitiateRevertFlowVersionParams contains all the parameters to send to the API endpoint

	for the initiate revert flow version operation.

	Typically these are written to a http.Request.
*/
type InitiateRevertFlowVersionParams struct {

	/* Body.

	   The controller service configuration details.
	*/
	Body *models.VersionControlInformationEntity

	/* ID.

	   The process group id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the initiate revert flow version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitiateRevertFlowVersionParams) WithDefaults() *InitiateRevertFlowVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the initiate revert flow version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitiateRevertFlowVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the initiate revert flow version params
func (o *InitiateRevertFlowVersionParams) WithTimeout(timeout time.Duration) *InitiateRevertFlowVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the initiate revert flow version params
func (o *InitiateRevertFlowVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the initiate revert flow version params
func (o *InitiateRevertFlowVersionParams) WithContext(ctx context.Context) *InitiateRevertFlowVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the initiate revert flow version params
func (o *InitiateRevertFlowVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the initiate revert flow version params
func (o *InitiateRevertFlowVersionParams) WithHTTPClient(client *http.Client) *InitiateRevertFlowVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the initiate revert flow version params
func (o *InitiateRevertFlowVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the initiate revert flow version params
func (o *InitiateRevertFlowVersionParams) WithBody(body *models.VersionControlInformationEntity) *InitiateRevertFlowVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the initiate revert flow version params
func (o *InitiateRevertFlowVersionParams) SetBody(body *models.VersionControlInformationEntity) {
	o.Body = body
}

// WithID adds the id to the initiate revert flow version params
func (o *InitiateRevertFlowVersionParams) WithID(id string) *InitiateRevertFlowVersionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the initiate revert flow version params
func (o *InitiateRevertFlowVersionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *InitiateRevertFlowVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

// NewInitiateVersionControlUpdateParams creates a new InitiateVersionControlUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInitiateVersionControlUpdateParams() *InitiateVersionControlUpdateParams {
	return &InitiateVersionControlUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInitiateVersionControlUpdateParamsWithTimeout creates a new InitiateVersionControlUpdateParams object
// with the ability to set a timeout on a request.
func NewInitiateVersionControlUpdateParamsWithTimeout(timeout time.Duration) *InitiateVersionControlUpdateParams {
	return &InitiateVersionControlUpdateParams{
		timeout: timeout,
	}
}

// NewInitiateVersionControlUpdateParamsWithContext creates a new InitiateVersionControlUpdateParams object
// with the ability to set a context for a request.
func NewInitiateVersionControlUpdateParamsWithContext(ctx context.Context) *InitiateVersionControlUpdateParams {
	return &InitiateVersionControlUpdateParams{
		Context: ctx,
	}
}

// NewInitiateVersionControlUpdateParamsWithHTTPClient creates a new InitiateVersionControlUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewInitiateVersionControlUpdateParamsWithHTTPClient(client *http.Client) *InitiateVersionControlUpdateParams {
	return &InitiateVersionControlUpdateParams{
		HTTPClient: client,
	}
}

/* InitiateVersionControlUpdateParams contains all the parameters to send to the API endpoint
   for the initiate version control update operation.

   Typically these are written to a http.Request.
*/
type InitiateVersionControlUpdateParams struct {

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

// WithDefaults hydrates default values in the initiate version control update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitiateVersionControlUpdateParams) WithDefaults() *InitiateVersionControlUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the initiate version control update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitiateVersionControlUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the initiate version control update params
func (o *InitiateVersionControlUpdateParams) WithTimeout(timeout time.Duration) *InitiateVersionControlUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the initiate version control update params
func (o *InitiateVersionControlUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the initiate version control update params
func (o *InitiateVersionControlUpdateParams) WithContext(ctx context.Context) *InitiateVersionControlUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the initiate version control update params
func (o *InitiateVersionControlUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the initiate version control update params
func (o *InitiateVersionControlUpdateParams) WithHTTPClient(client *http.Client) *InitiateVersionControlUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the initiate version control update params
func (o *InitiateVersionControlUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the initiate version control update params
func (o *InitiateVersionControlUpdateParams) WithBody(body *models.VersionControlInformationEntity) *InitiateVersionControlUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the initiate version control update params
func (o *InitiateVersionControlUpdateParams) SetBody(body *models.VersionControlInformationEntity) {
	o.Body = body
}

// WithID adds the id to the initiate version control update params
func (o *InitiateVersionControlUpdateParams) WithID(id string) *InitiateVersionControlUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the initiate version control update params
func (o *InitiateVersionControlUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *InitiateVersionControlUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
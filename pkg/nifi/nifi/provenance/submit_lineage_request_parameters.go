// Code generated by go-swagger; DO NOT EDIT.

package provenance

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

// NewSubmitLineageRequestParams creates a new SubmitLineageRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubmitLineageRequestParams() *SubmitLineageRequestParams {
	return &SubmitLineageRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubmitLineageRequestParamsWithTimeout creates a new SubmitLineageRequestParams object
// with the ability to set a timeout on a request.
func NewSubmitLineageRequestParamsWithTimeout(timeout time.Duration) *SubmitLineageRequestParams {
	return &SubmitLineageRequestParams{
		timeout: timeout,
	}
}

// NewSubmitLineageRequestParamsWithContext creates a new SubmitLineageRequestParams object
// with the ability to set a context for a request.
func NewSubmitLineageRequestParamsWithContext(ctx context.Context) *SubmitLineageRequestParams {
	return &SubmitLineageRequestParams{
		Context: ctx,
	}
}

// NewSubmitLineageRequestParamsWithHTTPClient creates a new SubmitLineageRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubmitLineageRequestParamsWithHTTPClient(client *http.Client) *SubmitLineageRequestParams {
	return &SubmitLineageRequestParams{
		HTTPClient: client,
	}
}

/* SubmitLineageRequestParams contains all the parameters to send to the API endpoint
   for the submit lineage request operation.

   Typically these are written to a http.Request.
*/
type SubmitLineageRequestParams struct {

	/* Body.

	   The lineage query details.
	*/
	Body *models.LineageEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the submit lineage request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitLineageRequestParams) WithDefaults() *SubmitLineageRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the submit lineage request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitLineageRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the submit lineage request params
func (o *SubmitLineageRequestParams) WithTimeout(timeout time.Duration) *SubmitLineageRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the submit lineage request params
func (o *SubmitLineageRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the submit lineage request params
func (o *SubmitLineageRequestParams) WithContext(ctx context.Context) *SubmitLineageRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the submit lineage request params
func (o *SubmitLineageRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the submit lineage request params
func (o *SubmitLineageRequestParams) WithHTTPClient(client *http.Client) *SubmitLineageRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the submit lineage request params
func (o *SubmitLineageRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the submit lineage request params
func (o *SubmitLineageRequestParams) WithBody(body *models.LineageEntity) *SubmitLineageRequestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the submit lineage request params
func (o *SubmitLineageRequestParams) SetBody(body *models.LineageEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SubmitLineageRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

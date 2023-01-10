// Code generated by go-swagger; DO NOT EDIT.

package remote_process_groups

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

// NewUpdateRemoteProcessGroupRunStatusesParams creates a new UpdateRemoteProcessGroupRunStatusesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRemoteProcessGroupRunStatusesParams() *UpdateRemoteProcessGroupRunStatusesParams {
	return &UpdateRemoteProcessGroupRunStatusesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRemoteProcessGroupRunStatusesParamsWithTimeout creates a new UpdateRemoteProcessGroupRunStatusesParams object
// with the ability to set a timeout on a request.
func NewUpdateRemoteProcessGroupRunStatusesParamsWithTimeout(timeout time.Duration) *UpdateRemoteProcessGroupRunStatusesParams {
	return &UpdateRemoteProcessGroupRunStatusesParams{
		timeout: timeout,
	}
}

// NewUpdateRemoteProcessGroupRunStatusesParamsWithContext creates a new UpdateRemoteProcessGroupRunStatusesParams object
// with the ability to set a context for a request.
func NewUpdateRemoteProcessGroupRunStatusesParamsWithContext(ctx context.Context) *UpdateRemoteProcessGroupRunStatusesParams {
	return &UpdateRemoteProcessGroupRunStatusesParams{
		Context: ctx,
	}
}

// NewUpdateRemoteProcessGroupRunStatusesParamsWithHTTPClient creates a new UpdateRemoteProcessGroupRunStatusesParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRemoteProcessGroupRunStatusesParamsWithHTTPClient(client *http.Client) *UpdateRemoteProcessGroupRunStatusesParams {
	return &UpdateRemoteProcessGroupRunStatusesParams{
		HTTPClient: client,
	}
}

/*
UpdateRemoteProcessGroupRunStatusesParams contains all the parameters to send to the API endpoint

	for the update remote process group run statuses operation.

	Typically these are written to a http.Request.
*/
type UpdateRemoteProcessGroupRunStatusesParams struct {

	/* Body.

	   The remote process groups run status.
	*/
	Body *models.RemotePortRunStatusEntity

	/* ID.

	   The process group id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update remote process group run statuses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRemoteProcessGroupRunStatusesParams) WithDefaults() *UpdateRemoteProcessGroupRunStatusesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update remote process group run statuses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRemoteProcessGroupRunStatusesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update remote process group run statuses params
func (o *UpdateRemoteProcessGroupRunStatusesParams) WithTimeout(timeout time.Duration) *UpdateRemoteProcessGroupRunStatusesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update remote process group run statuses params
func (o *UpdateRemoteProcessGroupRunStatusesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update remote process group run statuses params
func (o *UpdateRemoteProcessGroupRunStatusesParams) WithContext(ctx context.Context) *UpdateRemoteProcessGroupRunStatusesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update remote process group run statuses params
func (o *UpdateRemoteProcessGroupRunStatusesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update remote process group run statuses params
func (o *UpdateRemoteProcessGroupRunStatusesParams) WithHTTPClient(client *http.Client) *UpdateRemoteProcessGroupRunStatusesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update remote process group run statuses params
func (o *UpdateRemoteProcessGroupRunStatusesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update remote process group run statuses params
func (o *UpdateRemoteProcessGroupRunStatusesParams) WithBody(body *models.RemotePortRunStatusEntity) *UpdateRemoteProcessGroupRunStatusesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update remote process group run statuses params
func (o *UpdateRemoteProcessGroupRunStatusesParams) SetBody(body *models.RemotePortRunStatusEntity) {
	o.Body = body
}

// WithID adds the id to the update remote process group run statuses params
func (o *UpdateRemoteProcessGroupRunStatusesParams) WithID(id string) *UpdateRemoteProcessGroupRunStatusesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update remote process group run statuses params
func (o *UpdateRemoteProcessGroupRunStatusesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRemoteProcessGroupRunStatusesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

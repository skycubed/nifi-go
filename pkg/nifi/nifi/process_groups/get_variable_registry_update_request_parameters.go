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
)

// NewGetVariableRegistryUpdateRequestParams creates a new GetVariableRegistryUpdateRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVariableRegistryUpdateRequestParams() *GetVariableRegistryUpdateRequestParams {
	return &GetVariableRegistryUpdateRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVariableRegistryUpdateRequestParamsWithTimeout creates a new GetVariableRegistryUpdateRequestParams object
// with the ability to set a timeout on a request.
func NewGetVariableRegistryUpdateRequestParamsWithTimeout(timeout time.Duration) *GetVariableRegistryUpdateRequestParams {
	return &GetVariableRegistryUpdateRequestParams{
		timeout: timeout,
	}
}

// NewGetVariableRegistryUpdateRequestParamsWithContext creates a new GetVariableRegistryUpdateRequestParams object
// with the ability to set a context for a request.
func NewGetVariableRegistryUpdateRequestParamsWithContext(ctx context.Context) *GetVariableRegistryUpdateRequestParams {
	return &GetVariableRegistryUpdateRequestParams{
		Context: ctx,
	}
}

// NewGetVariableRegistryUpdateRequestParamsWithHTTPClient creates a new GetVariableRegistryUpdateRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVariableRegistryUpdateRequestParamsWithHTTPClient(client *http.Client) *GetVariableRegistryUpdateRequestParams {
	return &GetVariableRegistryUpdateRequestParams{
		HTTPClient: client,
	}
}

/*
GetVariableRegistryUpdateRequestParams contains all the parameters to send to the API endpoint

	for the get variable registry update request operation.

	Typically these are written to a http.Request.
*/
type GetVariableRegistryUpdateRequestParams struct {

	/* GroupID.

	   The process group id.
	*/
	GroupID string

	/* UpdateID.

	   The ID of the Variable Registry Update Request
	*/
	UpdateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get variable registry update request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVariableRegistryUpdateRequestParams) WithDefaults() *GetVariableRegistryUpdateRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get variable registry update request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVariableRegistryUpdateRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get variable registry update request params
func (o *GetVariableRegistryUpdateRequestParams) WithTimeout(timeout time.Duration) *GetVariableRegistryUpdateRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get variable registry update request params
func (o *GetVariableRegistryUpdateRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get variable registry update request params
func (o *GetVariableRegistryUpdateRequestParams) WithContext(ctx context.Context) *GetVariableRegistryUpdateRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get variable registry update request params
func (o *GetVariableRegistryUpdateRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get variable registry update request params
func (o *GetVariableRegistryUpdateRequestParams) WithHTTPClient(client *http.Client) *GetVariableRegistryUpdateRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get variable registry update request params
func (o *GetVariableRegistryUpdateRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the get variable registry update request params
func (o *GetVariableRegistryUpdateRequestParams) WithGroupID(groupID string) *GetVariableRegistryUpdateRequestParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get variable registry update request params
func (o *GetVariableRegistryUpdateRequestParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithUpdateID adds the updateID to the get variable registry update request params
func (o *GetVariableRegistryUpdateRequestParams) WithUpdateID(updateID string) *GetVariableRegistryUpdateRequestParams {
	o.SetUpdateID(updateID)
	return o
}

// SetUpdateID adds the updateId to the get variable registry update request params
func (o *GetVariableRegistryUpdateRequestParams) SetUpdateID(updateID string) {
	o.UpdateID = updateID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVariableRegistryUpdateRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
		return err
	}

	// path param updateId
	if err := r.SetPathParam("updateId", o.UpdateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

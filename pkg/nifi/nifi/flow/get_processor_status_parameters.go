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
	"github.com/go-openapi/swag"
)

// NewGetProcessorStatusParams creates a new GetProcessorStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProcessorStatusParams() *GetProcessorStatusParams {
	return &GetProcessorStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProcessorStatusParamsWithTimeout creates a new GetProcessorStatusParams object
// with the ability to set a timeout on a request.
func NewGetProcessorStatusParamsWithTimeout(timeout time.Duration) *GetProcessorStatusParams {
	return &GetProcessorStatusParams{
		timeout: timeout,
	}
}

// NewGetProcessorStatusParamsWithContext creates a new GetProcessorStatusParams object
// with the ability to set a context for a request.
func NewGetProcessorStatusParamsWithContext(ctx context.Context) *GetProcessorStatusParams {
	return &GetProcessorStatusParams{
		Context: ctx,
	}
}

// NewGetProcessorStatusParamsWithHTTPClient creates a new GetProcessorStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProcessorStatusParamsWithHTTPClient(client *http.Client) *GetProcessorStatusParams {
	return &GetProcessorStatusParams{
		HTTPClient: client,
	}
}

/* GetProcessorStatusParams contains all the parameters to send to the API endpoint
   for the get processor status operation.

   Typically these are written to a http.Request.
*/
type GetProcessorStatusParams struct {

	/* ClusterNodeID.

	   The id of the node where to get the status.
	*/
	ClusterNodeID *string

	/* ID.

	   The processor id.
	*/
	ID string

	/* Nodewise.

	   Whether or not to include the breakdown per node. Optional, defaults to false
	*/
	Nodewise *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get processor status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProcessorStatusParams) WithDefaults() *GetProcessorStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get processor status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProcessorStatusParams) SetDefaults() {
	var (
		nodewiseDefault = bool(false)
	)

	val := GetProcessorStatusParams{
		Nodewise: &nodewiseDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get processor status params
func (o *GetProcessorStatusParams) WithTimeout(timeout time.Duration) *GetProcessorStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get processor status params
func (o *GetProcessorStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get processor status params
func (o *GetProcessorStatusParams) WithContext(ctx context.Context) *GetProcessorStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get processor status params
func (o *GetProcessorStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get processor status params
func (o *GetProcessorStatusParams) WithHTTPClient(client *http.Client) *GetProcessorStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get processor status params
func (o *GetProcessorStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterNodeID adds the clusterNodeID to the get processor status params
func (o *GetProcessorStatusParams) WithClusterNodeID(clusterNodeID *string) *GetProcessorStatusParams {
	o.SetClusterNodeID(clusterNodeID)
	return o
}

// SetClusterNodeID adds the clusterNodeId to the get processor status params
func (o *GetProcessorStatusParams) SetClusterNodeID(clusterNodeID *string) {
	o.ClusterNodeID = clusterNodeID
}

// WithID adds the id to the get processor status params
func (o *GetProcessorStatusParams) WithID(id string) *GetProcessorStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get processor status params
func (o *GetProcessorStatusParams) SetID(id string) {
	o.ID = id
}

// WithNodewise adds the nodewise to the get processor status params
func (o *GetProcessorStatusParams) WithNodewise(nodewise *bool) *GetProcessorStatusParams {
	o.SetNodewise(nodewise)
	return o
}

// SetNodewise adds the nodewise to the get processor status params
func (o *GetProcessorStatusParams) SetNodewise(nodewise *bool) {
	o.Nodewise = nodewise
}

// WriteToRequest writes these params to a swagger request
func (o *GetProcessorStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClusterNodeID != nil {

		// query param clusterNodeId
		var qrClusterNodeID string

		if o.ClusterNodeID != nil {
			qrClusterNodeID = *o.ClusterNodeID
		}
		qClusterNodeID := qrClusterNodeID
		if qClusterNodeID != "" {

			if err := r.SetQueryParam("clusterNodeId", qClusterNodeID); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Nodewise != nil {

		// query param nodewise
		var qrNodewise bool

		if o.Nodewise != nil {
			qrNodewise = *o.Nodewise
		}
		qNodewise := swag.FormatBool(qrNodewise)
		if qNodewise != "" {

			if err := r.SetQueryParam("nodewise", qNodewise); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

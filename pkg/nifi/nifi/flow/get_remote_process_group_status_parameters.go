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

// NewGetRemoteProcessGroupStatusParams creates a new GetRemoteProcessGroupStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRemoteProcessGroupStatusParams() *GetRemoteProcessGroupStatusParams {
	return &GetRemoteProcessGroupStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRemoteProcessGroupStatusParamsWithTimeout creates a new GetRemoteProcessGroupStatusParams object
// with the ability to set a timeout on a request.
func NewGetRemoteProcessGroupStatusParamsWithTimeout(timeout time.Duration) *GetRemoteProcessGroupStatusParams {
	return &GetRemoteProcessGroupStatusParams{
		timeout: timeout,
	}
}

// NewGetRemoteProcessGroupStatusParamsWithContext creates a new GetRemoteProcessGroupStatusParams object
// with the ability to set a context for a request.
func NewGetRemoteProcessGroupStatusParamsWithContext(ctx context.Context) *GetRemoteProcessGroupStatusParams {
	return &GetRemoteProcessGroupStatusParams{
		Context: ctx,
	}
}

// NewGetRemoteProcessGroupStatusParamsWithHTTPClient creates a new GetRemoteProcessGroupStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRemoteProcessGroupStatusParamsWithHTTPClient(client *http.Client) *GetRemoteProcessGroupStatusParams {
	return &GetRemoteProcessGroupStatusParams{
		HTTPClient: client,
	}
}

/* GetRemoteProcessGroupStatusParams contains all the parameters to send to the API endpoint
   for the get remote process group status operation.

   Typically these are written to a http.Request.
*/
type GetRemoteProcessGroupStatusParams struct {

	/* ClusterNodeID.

	   The id of the node where to get the status.
	*/
	ClusterNodeID *string

	/* ID.

	   The remote process group id.
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

// WithDefaults hydrates default values in the get remote process group status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRemoteProcessGroupStatusParams) WithDefaults() *GetRemoteProcessGroupStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get remote process group status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRemoteProcessGroupStatusParams) SetDefaults() {
	var (
		nodewiseDefault = bool(false)
	)

	val := GetRemoteProcessGroupStatusParams{
		Nodewise: &nodewiseDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get remote process group status params
func (o *GetRemoteProcessGroupStatusParams) WithTimeout(timeout time.Duration) *GetRemoteProcessGroupStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get remote process group status params
func (o *GetRemoteProcessGroupStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get remote process group status params
func (o *GetRemoteProcessGroupStatusParams) WithContext(ctx context.Context) *GetRemoteProcessGroupStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get remote process group status params
func (o *GetRemoteProcessGroupStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get remote process group status params
func (o *GetRemoteProcessGroupStatusParams) WithHTTPClient(client *http.Client) *GetRemoteProcessGroupStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get remote process group status params
func (o *GetRemoteProcessGroupStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterNodeID adds the clusterNodeID to the get remote process group status params
func (o *GetRemoteProcessGroupStatusParams) WithClusterNodeID(clusterNodeID *string) *GetRemoteProcessGroupStatusParams {
	o.SetClusterNodeID(clusterNodeID)
	return o
}

// SetClusterNodeID adds the clusterNodeId to the get remote process group status params
func (o *GetRemoteProcessGroupStatusParams) SetClusterNodeID(clusterNodeID *string) {
	o.ClusterNodeID = clusterNodeID
}

// WithID adds the id to the get remote process group status params
func (o *GetRemoteProcessGroupStatusParams) WithID(id string) *GetRemoteProcessGroupStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get remote process group status params
func (o *GetRemoteProcessGroupStatusParams) SetID(id string) {
	o.ID = id
}

// WithNodewise adds the nodewise to the get remote process group status params
func (o *GetRemoteProcessGroupStatusParams) WithNodewise(nodewise *bool) *GetRemoteProcessGroupStatusParams {
	o.SetNodewise(nodewise)
	return o
}

// SetNodewise adds the nodewise to the get remote process group status params
func (o *GetRemoteProcessGroupStatusParams) SetNodewise(nodewise *bool) {
	o.Nodewise = nodewise
}

// WriteToRequest writes these params to a swagger request
func (o *GetRemoteProcessGroupStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

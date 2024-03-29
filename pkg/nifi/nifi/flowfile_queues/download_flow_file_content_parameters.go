// Code generated by go-swagger; DO NOT EDIT.

package flowfile_queues

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

// NewDownloadFlowFileContentParams creates a new DownloadFlowFileContentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDownloadFlowFileContentParams() *DownloadFlowFileContentParams {
	return &DownloadFlowFileContentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadFlowFileContentParamsWithTimeout creates a new DownloadFlowFileContentParams object
// with the ability to set a timeout on a request.
func NewDownloadFlowFileContentParamsWithTimeout(timeout time.Duration) *DownloadFlowFileContentParams {
	return &DownloadFlowFileContentParams{
		timeout: timeout,
	}
}

// NewDownloadFlowFileContentParamsWithContext creates a new DownloadFlowFileContentParams object
// with the ability to set a context for a request.
func NewDownloadFlowFileContentParamsWithContext(ctx context.Context) *DownloadFlowFileContentParams {
	return &DownloadFlowFileContentParams{
		Context: ctx,
	}
}

// NewDownloadFlowFileContentParamsWithHTTPClient creates a new DownloadFlowFileContentParams object
// with the ability to set a custom HTTPClient for a request.
func NewDownloadFlowFileContentParamsWithHTTPClient(client *http.Client) *DownloadFlowFileContentParams {
	return &DownloadFlowFileContentParams{
		HTTPClient: client,
	}
}

/*
DownloadFlowFileContentParams contains all the parameters to send to the API endpoint

	for the download flow file content operation.

	Typically these are written to a http.Request.
*/
type DownloadFlowFileContentParams struct {

	/* ClientID.

	   If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response.
	*/
	ClientID *string

	/* ClusterNodeID.

	   The id of the node where the content exists if clustered.
	*/
	ClusterNodeID *string

	/* FlowfileUUID.

	   The flowfile uuid.
	*/
	FlowfileUUID string

	/* ID.

	   The connection id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the download flow file content params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadFlowFileContentParams) WithDefaults() *DownloadFlowFileContentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the download flow file content params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadFlowFileContentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the download flow file content params
func (o *DownloadFlowFileContentParams) WithTimeout(timeout time.Duration) *DownloadFlowFileContentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download flow file content params
func (o *DownloadFlowFileContentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download flow file content params
func (o *DownloadFlowFileContentParams) WithContext(ctx context.Context) *DownloadFlowFileContentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download flow file content params
func (o *DownloadFlowFileContentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download flow file content params
func (o *DownloadFlowFileContentParams) WithHTTPClient(client *http.Client) *DownloadFlowFileContentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download flow file content params
func (o *DownloadFlowFileContentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the download flow file content params
func (o *DownloadFlowFileContentParams) WithClientID(clientID *string) *DownloadFlowFileContentParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the download flow file content params
func (o *DownloadFlowFileContentParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithClusterNodeID adds the clusterNodeID to the download flow file content params
func (o *DownloadFlowFileContentParams) WithClusterNodeID(clusterNodeID *string) *DownloadFlowFileContentParams {
	o.SetClusterNodeID(clusterNodeID)
	return o
}

// SetClusterNodeID adds the clusterNodeId to the download flow file content params
func (o *DownloadFlowFileContentParams) SetClusterNodeID(clusterNodeID *string) {
	o.ClusterNodeID = clusterNodeID
}

// WithFlowfileUUID adds the flowfileUUID to the download flow file content params
func (o *DownloadFlowFileContentParams) WithFlowfileUUID(flowfileUUID string) *DownloadFlowFileContentParams {
	o.SetFlowfileUUID(flowfileUUID)
	return o
}

// SetFlowfileUUID adds the flowfileUuid to the download flow file content params
func (o *DownloadFlowFileContentParams) SetFlowfileUUID(flowfileUUID string) {
	o.FlowfileUUID = flowfileUUID
}

// WithID adds the id to the download flow file content params
func (o *DownloadFlowFileContentParams) WithID(id string) *DownloadFlowFileContentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the download flow file content params
func (o *DownloadFlowFileContentParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadFlowFileContentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientID != nil {

		// query param clientId
		var qrClientID string

		if o.ClientID != nil {
			qrClientID = *o.ClientID
		}
		qClientID := qrClientID
		if qClientID != "" {

			if err := r.SetQueryParam("clientId", qClientID); err != nil {
				return err
			}
		}
	}

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

	// path param flowfile-uuid
	if err := r.SetPathParam("flowfile-uuid", o.FlowfileUUID); err != nil {
		return err
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

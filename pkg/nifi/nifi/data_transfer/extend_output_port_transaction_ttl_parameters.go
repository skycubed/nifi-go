// Code generated by go-swagger; DO NOT EDIT.

package data_transfer

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

// NewExtendOutputPortTransactionTTLParams creates a new ExtendOutputPortTransactionTTLParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtendOutputPortTransactionTTLParams() *ExtendOutputPortTransactionTTLParams {
	return &ExtendOutputPortTransactionTTLParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtendOutputPortTransactionTTLParamsWithTimeout creates a new ExtendOutputPortTransactionTTLParams object
// with the ability to set a timeout on a request.
func NewExtendOutputPortTransactionTTLParamsWithTimeout(timeout time.Duration) *ExtendOutputPortTransactionTTLParams {
	return &ExtendOutputPortTransactionTTLParams{
		timeout: timeout,
	}
}

// NewExtendOutputPortTransactionTTLParamsWithContext creates a new ExtendOutputPortTransactionTTLParams object
// with the ability to set a context for a request.
func NewExtendOutputPortTransactionTTLParamsWithContext(ctx context.Context) *ExtendOutputPortTransactionTTLParams {
	return &ExtendOutputPortTransactionTTLParams{
		Context: ctx,
	}
}

// NewExtendOutputPortTransactionTTLParamsWithHTTPClient creates a new ExtendOutputPortTransactionTTLParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtendOutputPortTransactionTTLParamsWithHTTPClient(client *http.Client) *ExtendOutputPortTransactionTTLParams {
	return &ExtendOutputPortTransactionTTLParams{
		HTTPClient: client,
	}
}

/* ExtendOutputPortTransactionTTLParams contains all the parameters to send to the API endpoint
   for the extend output port transaction TTL operation.

   Typically these are written to a http.Request.
*/
type ExtendOutputPortTransactionTTLParams struct {

	// PortID.
	PortID string

	// TransactionID.
	TransactionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extend output port transaction TTL params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtendOutputPortTransactionTTLParams) WithDefaults() *ExtendOutputPortTransactionTTLParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extend output port transaction TTL params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtendOutputPortTransactionTTLParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extend output port transaction TTL params
func (o *ExtendOutputPortTransactionTTLParams) WithTimeout(timeout time.Duration) *ExtendOutputPortTransactionTTLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extend output port transaction TTL params
func (o *ExtendOutputPortTransactionTTLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extend output port transaction TTL params
func (o *ExtendOutputPortTransactionTTLParams) WithContext(ctx context.Context) *ExtendOutputPortTransactionTTLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extend output port transaction TTL params
func (o *ExtendOutputPortTransactionTTLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extend output port transaction TTL params
func (o *ExtendOutputPortTransactionTTLParams) WithHTTPClient(client *http.Client) *ExtendOutputPortTransactionTTLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extend output port transaction TTL params
func (o *ExtendOutputPortTransactionTTLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPortID adds the portID to the extend output port transaction TTL params
func (o *ExtendOutputPortTransactionTTLParams) WithPortID(portID string) *ExtendOutputPortTransactionTTLParams {
	o.SetPortID(portID)
	return o
}

// SetPortID adds the portId to the extend output port transaction TTL params
func (o *ExtendOutputPortTransactionTTLParams) SetPortID(portID string) {
	o.PortID = portID
}

// WithTransactionID adds the transactionID to the extend output port transaction TTL params
func (o *ExtendOutputPortTransactionTTLParams) WithTransactionID(transactionID string) *ExtendOutputPortTransactionTTLParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the extend output port transaction TTL params
func (o *ExtendOutputPortTransactionTTLParams) SetTransactionID(transactionID string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *ExtendOutputPortTransactionTTLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param portId
	if err := r.SetPathParam("portId", o.PortID); err != nil {
		return err
	}

	// path param transactionId
	if err := r.SetPathParam("transactionId", o.TransactionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
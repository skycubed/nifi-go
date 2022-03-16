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

// NewGetListingRequestParams creates a new GetListingRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetListingRequestParams() *GetListingRequestParams {
	return &GetListingRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetListingRequestParamsWithTimeout creates a new GetListingRequestParams object
// with the ability to set a timeout on a request.
func NewGetListingRequestParamsWithTimeout(timeout time.Duration) *GetListingRequestParams {
	return &GetListingRequestParams{
		timeout: timeout,
	}
}

// NewGetListingRequestParamsWithContext creates a new GetListingRequestParams object
// with the ability to set a context for a request.
func NewGetListingRequestParamsWithContext(ctx context.Context) *GetListingRequestParams {
	return &GetListingRequestParams{
		Context: ctx,
	}
}

// NewGetListingRequestParamsWithHTTPClient creates a new GetListingRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetListingRequestParamsWithHTTPClient(client *http.Client) *GetListingRequestParams {
	return &GetListingRequestParams{
		HTTPClient: client,
	}
}

/* GetListingRequestParams contains all the parameters to send to the API endpoint
   for the get listing request operation.

   Typically these are written to a http.Request.
*/
type GetListingRequestParams struct {

	/* ID.

	   The connection id.
	*/
	ID string

	/* ListingRequestID.

	   The listing request id.
	*/
	ListingRequestID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get listing request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetListingRequestParams) WithDefaults() *GetListingRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get listing request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetListingRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get listing request params
func (o *GetListingRequestParams) WithTimeout(timeout time.Duration) *GetListingRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get listing request params
func (o *GetListingRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get listing request params
func (o *GetListingRequestParams) WithContext(ctx context.Context) *GetListingRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get listing request params
func (o *GetListingRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get listing request params
func (o *GetListingRequestParams) WithHTTPClient(client *http.Client) *GetListingRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get listing request params
func (o *GetListingRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get listing request params
func (o *GetListingRequestParams) WithID(id string) *GetListingRequestParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get listing request params
func (o *GetListingRequestParams) SetID(id string) {
	o.ID = id
}

// WithListingRequestID adds the listingRequestID to the get listing request params
func (o *GetListingRequestParams) WithListingRequestID(listingRequestID string) *GetListingRequestParams {
	o.SetListingRequestID(listingRequestID)
	return o
}

// SetListingRequestID adds the listingRequestId to the get listing request params
func (o *GetListingRequestParams) SetListingRequestID(listingRequestID string) {
	o.ListingRequestID = listingRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetListingRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param listing-request-id
	if err := r.SetPathParam("listing-request-id", o.ListingRequestID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
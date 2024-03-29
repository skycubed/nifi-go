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
)

// NewGetBulletinBoardParams creates a new GetBulletinBoardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBulletinBoardParams() *GetBulletinBoardParams {
	return &GetBulletinBoardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBulletinBoardParamsWithTimeout creates a new GetBulletinBoardParams object
// with the ability to set a timeout on a request.
func NewGetBulletinBoardParamsWithTimeout(timeout time.Duration) *GetBulletinBoardParams {
	return &GetBulletinBoardParams{
		timeout: timeout,
	}
}

// NewGetBulletinBoardParamsWithContext creates a new GetBulletinBoardParams object
// with the ability to set a context for a request.
func NewGetBulletinBoardParamsWithContext(ctx context.Context) *GetBulletinBoardParams {
	return &GetBulletinBoardParams{
		Context: ctx,
	}
}

// NewGetBulletinBoardParamsWithHTTPClient creates a new GetBulletinBoardParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBulletinBoardParamsWithHTTPClient(client *http.Client) *GetBulletinBoardParams {
	return &GetBulletinBoardParams{
		HTTPClient: client,
	}
}

/*
GetBulletinBoardParams contains all the parameters to send to the API endpoint

	for the get bulletin board operation.

	Typically these are written to a http.Request.
*/
type GetBulletinBoardParams struct {

	/* After.

	   Includes bulletins with an id after this value.
	*/
	After *string

	/* GroupID.

	   Includes bulletins originating from this sources whose group id match this regular expression.
	*/
	GroupID *string

	/* Limit.

	   The number of bulletins to limit the response to.
	*/
	Limit *string

	/* Message.

	   Includes bulletins whose message that match this regular expression.
	*/
	Message *string

	/* SourceID.

	   Includes bulletins originating from this sources whose id match this regular expression.
	*/
	SourceID *string

	/* SourceName.

	   Includes bulletins originating from this sources whose name match this regular expression.
	*/
	SourceName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get bulletin board params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBulletinBoardParams) WithDefaults() *GetBulletinBoardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get bulletin board params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBulletinBoardParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get bulletin board params
func (o *GetBulletinBoardParams) WithTimeout(timeout time.Duration) *GetBulletinBoardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bulletin board params
func (o *GetBulletinBoardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bulletin board params
func (o *GetBulletinBoardParams) WithContext(ctx context.Context) *GetBulletinBoardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bulletin board params
func (o *GetBulletinBoardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bulletin board params
func (o *GetBulletinBoardParams) WithHTTPClient(client *http.Client) *GetBulletinBoardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bulletin board params
func (o *GetBulletinBoardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAfter adds the after to the get bulletin board params
func (o *GetBulletinBoardParams) WithAfter(after *string) *GetBulletinBoardParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the get bulletin board params
func (o *GetBulletinBoardParams) SetAfter(after *string) {
	o.After = after
}

// WithGroupID adds the groupID to the get bulletin board params
func (o *GetBulletinBoardParams) WithGroupID(groupID *string) *GetBulletinBoardParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get bulletin board params
func (o *GetBulletinBoardParams) SetGroupID(groupID *string) {
	o.GroupID = groupID
}

// WithLimit adds the limit to the get bulletin board params
func (o *GetBulletinBoardParams) WithLimit(limit *string) *GetBulletinBoardParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get bulletin board params
func (o *GetBulletinBoardParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithMessage adds the message to the get bulletin board params
func (o *GetBulletinBoardParams) WithMessage(message *string) *GetBulletinBoardParams {
	o.SetMessage(message)
	return o
}

// SetMessage adds the message to the get bulletin board params
func (o *GetBulletinBoardParams) SetMessage(message *string) {
	o.Message = message
}

// WithSourceID adds the sourceID to the get bulletin board params
func (o *GetBulletinBoardParams) WithSourceID(sourceID *string) *GetBulletinBoardParams {
	o.SetSourceID(sourceID)
	return o
}

// SetSourceID adds the sourceId to the get bulletin board params
func (o *GetBulletinBoardParams) SetSourceID(sourceID *string) {
	o.SourceID = sourceID
}

// WithSourceName adds the sourceName to the get bulletin board params
func (o *GetBulletinBoardParams) WithSourceName(sourceName *string) *GetBulletinBoardParams {
	o.SetSourceName(sourceName)
	return o
}

// SetSourceName adds the sourceName to the get bulletin board params
func (o *GetBulletinBoardParams) SetSourceName(sourceName *string) {
	o.SourceName = sourceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetBulletinBoardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.After != nil {

		// query param after
		var qrAfter string

		if o.After != nil {
			qrAfter = *o.After
		}
		qAfter := qrAfter
		if qAfter != "" {

			if err := r.SetQueryParam("after", qAfter); err != nil {
				return err
			}
		}
	}

	if o.GroupID != nil {

		// query param groupId
		var qrGroupID string

		if o.GroupID != nil {
			qrGroupID = *o.GroupID
		}
		qGroupID := qrGroupID
		if qGroupID != "" {

			if err := r.SetQueryParam("groupId", qGroupID); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit string

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Message != nil {

		// query param message
		var qrMessage string

		if o.Message != nil {
			qrMessage = *o.Message
		}
		qMessage := qrMessage
		if qMessage != "" {

			if err := r.SetQueryParam("message", qMessage); err != nil {
				return err
			}
		}
	}

	if o.SourceID != nil {

		// query param sourceId
		var qrSourceID string

		if o.SourceID != nil {
			qrSourceID = *o.SourceID
		}
		qSourceID := qrSourceID
		if qSourceID != "" {

			if err := r.SetQueryParam("sourceId", qSourceID); err != nil {
				return err
			}
		}
	}

	if o.SourceName != nil {

		// query param sourceName
		var qrSourceName string

		if o.SourceName != nil {
			qrSourceName = *o.SourceName
		}
		qSourceName := qrSourceName
		if qSourceName != "" {

			if err := r.SetQueryParam("sourceName", qSourceName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

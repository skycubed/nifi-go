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

// NewQueryHistoryParams creates a new QueryHistoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryHistoryParams() *QueryHistoryParams {
	return &QueryHistoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryHistoryParamsWithTimeout creates a new QueryHistoryParams object
// with the ability to set a timeout on a request.
func NewQueryHistoryParamsWithTimeout(timeout time.Duration) *QueryHistoryParams {
	return &QueryHistoryParams{
		timeout: timeout,
	}
}

// NewQueryHistoryParamsWithContext creates a new QueryHistoryParams object
// with the ability to set a context for a request.
func NewQueryHistoryParamsWithContext(ctx context.Context) *QueryHistoryParams {
	return &QueryHistoryParams{
		Context: ctx,
	}
}

// NewQueryHistoryParamsWithHTTPClient creates a new QueryHistoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryHistoryParamsWithHTTPClient(client *http.Client) *QueryHistoryParams {
	return &QueryHistoryParams{
		HTTPClient: client,
	}
}

/*
QueryHistoryParams contains all the parameters to send to the API endpoint

	for the query history operation.

	Typically these are written to a http.Request.
*/
type QueryHistoryParams struct {

	/* Count.

	   The number of actions to return.
	*/
	Count string

	/* EndDate.

	   Include actions before this date.
	*/
	EndDate *string

	/* Offset.

	   The offset into the result set.
	*/
	Offset string

	/* SortColumn.

	   The field to sort on.
	*/
	SortColumn *string

	/* SortOrder.

	   The direction to sort.
	*/
	SortOrder *string

	/* SourceID.

	   Include actions on this component.
	*/
	SourceID *string

	/* StartDate.

	   Include actions after this date.
	*/
	StartDate *string

	/* UserIdentity.

	   Include actions performed by this user.
	*/
	UserIdentity *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryHistoryParams) WithDefaults() *QueryHistoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryHistoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query history params
func (o *QueryHistoryParams) WithTimeout(timeout time.Duration) *QueryHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query history params
func (o *QueryHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query history params
func (o *QueryHistoryParams) WithContext(ctx context.Context) *QueryHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query history params
func (o *QueryHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query history params
func (o *QueryHistoryParams) WithHTTPClient(client *http.Client) *QueryHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query history params
func (o *QueryHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the query history params
func (o *QueryHistoryParams) WithCount(count string) *QueryHistoryParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the query history params
func (o *QueryHistoryParams) SetCount(count string) {
	o.Count = count
}

// WithEndDate adds the endDate to the query history params
func (o *QueryHistoryParams) WithEndDate(endDate *string) *QueryHistoryParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the query history params
func (o *QueryHistoryParams) SetEndDate(endDate *string) {
	o.EndDate = endDate
}

// WithOffset adds the offset to the query history params
func (o *QueryHistoryParams) WithOffset(offset string) *QueryHistoryParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query history params
func (o *QueryHistoryParams) SetOffset(offset string) {
	o.Offset = offset
}

// WithSortColumn adds the sortColumn to the query history params
func (o *QueryHistoryParams) WithSortColumn(sortColumn *string) *QueryHistoryParams {
	o.SetSortColumn(sortColumn)
	return o
}

// SetSortColumn adds the sortColumn to the query history params
func (o *QueryHistoryParams) SetSortColumn(sortColumn *string) {
	o.SortColumn = sortColumn
}

// WithSortOrder adds the sortOrder to the query history params
func (o *QueryHistoryParams) WithSortOrder(sortOrder *string) *QueryHistoryParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the query history params
func (o *QueryHistoryParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WithSourceID adds the sourceID to the query history params
func (o *QueryHistoryParams) WithSourceID(sourceID *string) *QueryHistoryParams {
	o.SetSourceID(sourceID)
	return o
}

// SetSourceID adds the sourceId to the query history params
func (o *QueryHistoryParams) SetSourceID(sourceID *string) {
	o.SourceID = sourceID
}

// WithStartDate adds the startDate to the query history params
func (o *QueryHistoryParams) WithStartDate(startDate *string) *QueryHistoryParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the query history params
func (o *QueryHistoryParams) SetStartDate(startDate *string) {
	o.StartDate = startDate
}

// WithUserIdentity adds the userIdentity to the query history params
func (o *QueryHistoryParams) WithUserIdentity(userIdentity *string) *QueryHistoryParams {
	o.SetUserIdentity(userIdentity)
	return o
}

// SetUserIdentity adds the userIdentity to the query history params
func (o *QueryHistoryParams) SetUserIdentity(userIdentity *string) {
	o.UserIdentity = userIdentity
}

// WriteToRequest writes these params to a swagger request
func (o *QueryHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param count
	qrCount := o.Count
	qCount := qrCount
	if qCount != "" {

		if err := r.SetQueryParam("count", qCount); err != nil {
			return err
		}
	}

	if o.EndDate != nil {

		// query param endDate
		var qrEndDate string

		if o.EndDate != nil {
			qrEndDate = *o.EndDate
		}
		qEndDate := qrEndDate
		if qEndDate != "" {

			if err := r.SetQueryParam("endDate", qEndDate); err != nil {
				return err
			}
		}
	}

	// query param offset
	qrOffset := o.Offset
	qOffset := qrOffset
	if qOffset != "" {

		if err := r.SetQueryParam("offset", qOffset); err != nil {
			return err
		}
	}

	if o.SortColumn != nil {

		// query param sortColumn
		var qrSortColumn string

		if o.SortColumn != nil {
			qrSortColumn = *o.SortColumn
		}
		qSortColumn := qrSortColumn
		if qSortColumn != "" {

			if err := r.SetQueryParam("sortColumn", qSortColumn); err != nil {
				return err
			}
		}
	}

	if o.SortOrder != nil {

		// query param sortOrder
		var qrSortOrder string

		if o.SortOrder != nil {
			qrSortOrder = *o.SortOrder
		}
		qSortOrder := qrSortOrder
		if qSortOrder != "" {

			if err := r.SetQueryParam("sortOrder", qSortOrder); err != nil {
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

	if o.StartDate != nil {

		// query param startDate
		var qrStartDate string

		if o.StartDate != nil {
			qrStartDate = *o.StartDate
		}
		qStartDate := qrStartDate
		if qStartDate != "" {

			if err := r.SetQueryParam("startDate", qStartDate); err != nil {
				return err
			}
		}
	}

	if o.UserIdentity != nil {

		// query param userIdentity
		var qrUserIdentity string

		if o.UserIdentity != nil {
			qrUserIdentity = *o.UserIdentity
		}
		qUserIdentity := qrUserIdentity
		if qUserIdentity != "" {

			if err := r.SetQueryParam("userIdentity", qUserIdentity); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

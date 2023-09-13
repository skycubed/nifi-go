// Code generated by go-swagger; DO NOT EDIT.

package items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new items API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for items API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAvailableBucketItemFields(params *GetAvailableBucketItemFieldsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAvailableBucketItemFieldsOK, error)

	GetItems(params *GetItemsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetItemsOK, error)

	GetItemsInBucket(params *GetItemsInBucketParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetItemsInBucketOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetAvailableBucketItemFields gets item fields

Retrieves the item field names for searching or sorting on bucket items.
*/
func (a *Client) GetAvailableBucketItemFields(params *GetAvailableBucketItemFieldsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAvailableBucketItemFieldsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAvailableBucketItemFieldsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAvailableBucketItemFields",
		Method:             "GET",
		PathPattern:        "/items/fields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAvailableBucketItemFieldsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAvailableBucketItemFieldsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAvailableBucketItemFields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetItems gets all items

Get items across all buckets. The returned items will include only items from buckets for which the user is authorized. If the user is not authorized to any buckets, an empty list will be returned.
*/
func (a *Client) GetItems(params *GetItemsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetItemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetItemsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getItems",
		Method:             "GET",
		PathPattern:        "/items",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetItemsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetItemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getItems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetItemsInBucket gets bucket items

Gets the items located in the given bucket.
*/
func (a *Client) GetItemsInBucket(params *GetItemsInBucketParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetItemsInBucketOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetItemsInBucketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getItemsInBucket",
		Method:             "GET",
		PathPattern:        "/items/{bucketId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetItemsInBucketReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetItemsInBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getItemsInBucket: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

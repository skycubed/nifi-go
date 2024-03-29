// Code generated by go-swagger; DO NOT EDIT.

package extensions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new extensions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for extensions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetExtensions(params *GetExtensionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionsOK, error)

	GetExtensionsProvidingServiceAPI(params *GetExtensionsProvidingServiceAPIParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionsProvidingServiceAPIOK, error)

	GetTags(params *GetTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTagsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	GetExtensions gets all extensions

	Gets the metadata for all extensions that match the filter params and are part of bundles located in buckets the current user is authorized for. If the user is not authorized to any buckets, an empty result set will be returned.

NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.
*/
func (a *Client) GetExtensions(params *GetExtensionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExtensionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExtensions",
		Method:             "GET",
		PathPattern:        "/extensions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetExtensionsReader{formats: a.formats},
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
	success, ok := result.(*GetExtensionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExtensions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetExtensionsProvidingServiceAPI gets extensions providing service API

	Gets the metadata for extensions that provide the specified API and are part of bundles located in buckets the current user is authorized for. If the user is not authorized to any buckets, an empty result set will be returned.

NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.
*/
func (a *Client) GetExtensionsProvidingServiceAPI(params *GetExtensionsProvidingServiceAPIParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionsProvidingServiceAPIOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExtensionsProvidingServiceAPIParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExtensionsProvidingServiceAPI",
		Method:             "GET",
		PathPattern:        "/extensions/provided-service-api",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetExtensionsProvidingServiceAPIReader{formats: a.formats},
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
	success, ok := result.(*GetExtensionsProvidingServiceAPIOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExtensionsProvidingServiceAPI: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetTags gets extension tags

	Gets all the extension tags known to this NiFi Registry instance, along with the number of extensions that have the given tag.

NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.
*/
func (a *Client) GetTags(params *GetTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTags",
		Method:             "GET",
		PathPattern:        "/extensions/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTagsReader{formats: a.formats},
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
	success, ok := result.(*GetTagsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

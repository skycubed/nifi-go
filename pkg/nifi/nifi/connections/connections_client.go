// Code generated by go-swagger; DO NOT EDIT.

package connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new connections API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for connections API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteConnection(params *DeleteConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteConnectionOK, error)

	GetConnection(params *GetConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetConnectionOK, error)

	UpdateConnection(params *UpdateConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateConnectionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteConnection deletes a connection
*/
func (a *Client) DeleteConnection(params *DeleteConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteConnection",
		Method:             "DELETE",
		PathPattern:        "/connections/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteConnectionReader{formats: a.formats},
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
	success, ok := result.(*DeleteConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetConnection gets a connection
*/
func (a *Client) GetConnection(params *GetConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getConnection",
		Method:             "GET",
		PathPattern:        "/connections/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateConnection updates a connection
*/
func (a *Client) UpdateConnection(params *UpdateConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateConnection",
		Method:             "PUT",
		PathPattern:        "/connections/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateConnectionReader{formats: a.formats},
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
	success, ok := result.(*UpdateConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

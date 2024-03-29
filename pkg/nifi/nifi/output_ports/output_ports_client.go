// Code generated by go-swagger; DO NOT EDIT.

package output_ports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new output ports API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for output ports API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetOutputPort(params *GetOutputPortParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOutputPortOK, error)

	RemoveOutputPort(params *RemoveOutputPortParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveOutputPortOK, error)

	UpdateOutputPort(params *UpdateOutputPortParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOutputPortOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetOutputPort gets an output port
*/
func (a *Client) GetOutputPort(params *GetOutputPortParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOutputPortOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOutputPortParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOutputPort",
		Method:             "GET",
		PathPattern:        "/output-ports/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOutputPortReader{formats: a.formats},
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
	success, ok := result.(*GetOutputPortOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOutputPort: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RemoveOutputPort deletes an output port
*/
func (a *Client) RemoveOutputPort(params *RemoveOutputPortParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveOutputPortOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveOutputPortParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeOutputPort",
		Method:             "DELETE",
		PathPattern:        "/output-ports/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveOutputPortReader{formats: a.formats},
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
	success, ok := result.(*RemoveOutputPortOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeOutputPort: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateOutputPort updates an output port
*/
func (a *Client) UpdateOutputPort(params *UpdateOutputPortParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOutputPortOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOutputPortParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateOutputPort",
		Method:             "PUT",
		PathPattern:        "/output-ports/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateOutputPortReader{formats: a.formats},
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
	success, ok := result.(*UpdateOutputPortOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateOutputPort: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

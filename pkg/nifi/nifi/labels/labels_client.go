// Code generated by go-swagger; DO NOT EDIT.

package labels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new labels API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for labels API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetLabel(params *GetLabelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLabelOK, error)

	RemoveLabel(params *RemoveLabelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveLabelOK, error)

	UpdateLabel(params *UpdateLabelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateLabelOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetLabel gets a label
*/
func (a *Client) GetLabel(params *GetLabelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLabelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLabelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getLabel",
		Method:             "GET",
		PathPattern:        "/labels/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLabelReader{formats: a.formats},
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
	success, ok := result.(*GetLabelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLabel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemoveLabel deletes a label
*/
func (a *Client) RemoveLabel(params *RemoveLabelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveLabelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveLabelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeLabel",
		Method:             "DELETE",
		PathPattern:        "/labels/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveLabelReader{formats: a.formats},
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
	success, ok := result.(*RemoveLabelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeLabel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateLabel updates a label
*/
func (a *Client) UpdateLabel(params *UpdateLabelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateLabelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateLabelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateLabel",
		Method:             "PUT",
		PathPattern:        "/labels/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateLabelReader{formats: a.formats},
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
	success, ok := result.(*UpdateLabelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateLabel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
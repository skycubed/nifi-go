// Code generated by go-swagger; DO NOT EDIT.

package controller_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new controller services API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for controller services API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetControllerService(params *GetControllerServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetControllerServiceOK, error)

	GetControllerServiceReferences(params *GetControllerServiceReferencesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetControllerServiceReferencesOK, error)

	RemoveControllerService(params *RemoveControllerServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveControllerServiceOK, error)

	UpdateControllerService(params *UpdateControllerServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateControllerServiceOK, error)

	UpdateControllerServiceReferences(params *UpdateControllerServiceReferencesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateControllerServiceReferencesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetControllerService gets a controller service
*/
func (a *Client) GetControllerService(params *GetControllerServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetControllerServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetControllerServiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getControllerService",
		Method:             "GET",
		PathPattern:        "/controller-services/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetControllerServiceReader{formats: a.formats},
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
	success, ok := result.(*GetControllerServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getControllerService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetControllerServiceReferences gets a controller service
*/
func (a *Client) GetControllerServiceReferences(params *GetControllerServiceReferencesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetControllerServiceReferencesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetControllerServiceReferencesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getControllerServiceReferences",
		Method:             "GET",
		PathPattern:        "/controller-services/{id}/references",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetControllerServiceReferencesReader{formats: a.formats},
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
	success, ok := result.(*GetControllerServiceReferencesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getControllerServiceReferences: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemoveControllerService deletes a controller service
*/
func (a *Client) RemoveControllerService(params *RemoveControllerServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveControllerServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveControllerServiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeControllerService",
		Method:             "DELETE",
		PathPattern:        "/controller-services/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveControllerServiceReader{formats: a.formats},
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
	success, ok := result.(*RemoveControllerServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeControllerService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateControllerService updates a controller service
*/
func (a *Client) UpdateControllerService(params *UpdateControllerServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateControllerServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateControllerServiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateControllerService",
		Method:             "PUT",
		PathPattern:        "/controller-services/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateControllerServiceReader{formats: a.formats},
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
	success, ok := result.(*UpdateControllerServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateControllerService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateControllerServiceReferences updates a controller services references
*/
func (a *Client) UpdateControllerServiceReferences(params *UpdateControllerServiceReferencesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateControllerServiceReferencesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateControllerServiceReferencesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateControllerServiceReferences",
		Method:             "PUT",
		PathPattern:        "/controller-services/{id}/references",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateControllerServiceReferencesReader{formats: a.formats},
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
	success, ok := result.(*UpdateControllerServiceReferencesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateControllerServiceReferences: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
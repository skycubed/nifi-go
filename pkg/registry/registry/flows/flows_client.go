// Code generated by go-swagger; DO NOT EDIT.

package flows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new flows API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for flows API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAvailableFlowFields(params *GetAvailableFlowFieldsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAvailableFlowFieldsOK, error)

	GlobalGetFlow(params *GlobalGetFlowParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GlobalGetFlowOK, error)

	GlobalGetFlowVersion(params *GlobalGetFlowVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GlobalGetFlowVersionOK, error)

	GlobalGetFlowVersions(params *GlobalGetFlowVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GlobalGetFlowVersionsOK, error)

	GlobalGetLatestFlowVersion(params *GlobalGetLatestFlowVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GlobalGetLatestFlowVersionOK, error)

	GlobalGetLatestFlowVersionMetadata(params *GlobalGetLatestFlowVersionMetadataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GlobalGetLatestFlowVersionMetadataOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetAvailableFlowFields gets flow fields

Retrieves the flow field names that can be used for searching or sorting on flows.
*/
func (a *Client) GetAvailableFlowFields(params *GetAvailableFlowFieldsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAvailableFlowFieldsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAvailableFlowFieldsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAvailableFlowFields",
		Method:             "GET",
		PathPattern:        "/flows/fields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAvailableFlowFieldsReader{formats: a.formats},
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
	success, ok := result.(*GetAvailableFlowFieldsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAvailableFlowFields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GlobalGetFlow gets flow

Gets a flow by id.
*/
func (a *Client) GlobalGetFlow(params *GlobalGetFlowParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GlobalGetFlowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGlobalGetFlowParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "globalGetFlow",
		Method:             "GET",
		PathPattern:        "/flows/{flowId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GlobalGetFlowReader{formats: a.formats},
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
	success, ok := result.(*GlobalGetFlowOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for globalGetFlow: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GlobalGetFlowVersion gets flow version

Gets the given version of a flow, including metadata and flow content.
*/
func (a *Client) GlobalGetFlowVersion(params *GlobalGetFlowVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GlobalGetFlowVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGlobalGetFlowVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "globalGetFlowVersion",
		Method:             "GET",
		PathPattern:        "/flows/{flowId}/versions/{versionNumber}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GlobalGetFlowVersionReader{formats: a.formats},
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
	success, ok := result.(*GlobalGetFlowVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for globalGetFlowVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GlobalGetFlowVersions gets flow versions

Gets summary information for all versions of a given flow. Versions are ordered newest->oldest.
*/
func (a *Client) GlobalGetFlowVersions(params *GlobalGetFlowVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GlobalGetFlowVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGlobalGetFlowVersionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "globalGetFlowVersions",
		Method:             "GET",
		PathPattern:        "/flows/{flowId}/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GlobalGetFlowVersionsReader{formats: a.formats},
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
	success, ok := result.(*GlobalGetFlowVersionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for globalGetFlowVersions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GlobalGetLatestFlowVersion gets latest flow version

Gets the latest version of a flow, including metadata and flow content.
*/
func (a *Client) GlobalGetLatestFlowVersion(params *GlobalGetLatestFlowVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GlobalGetLatestFlowVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGlobalGetLatestFlowVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "globalGetLatestFlowVersion",
		Method:             "GET",
		PathPattern:        "/flows/{flowId}/versions/latest",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GlobalGetLatestFlowVersionReader{formats: a.formats},
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
	success, ok := result.(*GlobalGetLatestFlowVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for globalGetLatestFlowVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GlobalGetLatestFlowVersionMetadata gets latest flow version metadata

Gets the metadata for the latest version of a flow.
*/
func (a *Client) GlobalGetLatestFlowVersionMetadata(params *GlobalGetLatestFlowVersionMetadataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GlobalGetLatestFlowVersionMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGlobalGetLatestFlowVersionMetadataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "globalGetLatestFlowVersionMetadata",
		Method:             "GET",
		PathPattern:        "/flows/{flowId}/versions/latest/metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GlobalGetLatestFlowVersionMetadataReader{formats: a.formats},
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
	success, ok := result.(*GlobalGetLatestFlowVersionMetadataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for globalGetLatestFlowVersionMetadata: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

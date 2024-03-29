// Code generated by go-swagger; DO NOT EDIT.

package bucket_bundles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new bucket bundles API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for bucket bundles API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateExtensionBundleVersion(params *CreateExtensionBundleVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateExtensionBundleVersionOK, error)

	GetExtensionBundles(params *GetExtensionBundlesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionBundlesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	CreateExtensionBundleVersion creates extension bundle version

	Creates a version of an extension bundle by uploading a binary artifact. If an extension bundle already exists in the given bucket with the same group id and artifact id as that of the bundle being uploaded, then it will be added as a new version to the existing bundle. If an extension bundle does not already exist in the given bucket with the same group id and artifact id, then a new extension bundle will be created and this version will be added to the new bundle. Client's may optionally supply a SHA-256 in hex format through the multi-part form field 'sha256'. If supplied, then this value will be compared against the SHA-256 computed by the server, and the bundle will be rejected if the values do not match. If not supplied, the bundle will be accepted, but will be marked to indicate that the client did not supply a SHA-256 during creation.

NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.
*/
func (a *Client) CreateExtensionBundleVersion(params *CreateExtensionBundleVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateExtensionBundleVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateExtensionBundleVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createExtensionBundleVersion",
		Method:             "POST",
		PathPattern:        "/buckets/{bucketId}/bundles/{bundleType}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateExtensionBundleVersionReader{formats: a.formats},
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
	success, ok := result.(*CreateExtensionBundleVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createExtensionBundleVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetExtensionBundles gets extension bundles by bucket

NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.
*/
func (a *Client) GetExtensionBundles(params *GetExtensionBundlesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionBundlesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExtensionBundlesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExtensionBundles",
		Method:             "GET",
		PathPattern:        "/buckets/{bucketId}/bundles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetExtensionBundlesReader{formats: a.formats},
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
	success, ok := result.(*GetExtensionBundlesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExtensionBundles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

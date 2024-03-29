// Code generated by go-swagger; DO NOT EDIT.

package extension_repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new extension repository API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for extension repository API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetExtensionRepoArtifacts(params *GetExtensionRepoArtifactsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionRepoArtifactsOK, error)

	GetExtensionRepoBuckets(params *GetExtensionRepoBucketsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionRepoBucketsOK, error)

	GetExtensionRepoGroups(params *GetExtensionRepoGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionRepoGroupsOK, error)

	GetExtensionRepoVersion(params *GetExtensionRepoVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionRepoVersionOK, error)

	GetExtensionRepoVersionContent(params *GetExtensionRepoVersionContentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionRepoVersionContentOK, error)

	GetExtensionRepoVersionExtension(params *GetExtensionRepoVersionExtensionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionRepoVersionExtensionOK, error)

	GetExtensionRepoVersionExtensionAdditionalDetailsDocs(params *GetExtensionRepoVersionExtensionAdditionalDetailsDocsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionRepoVersionExtensionAdditionalDetailsDocsOK, error)

	GetExtensionRepoVersionExtensionDocs(params *GetExtensionRepoVersionExtensionDocsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionRepoVersionExtensionDocsOK, error)

	GetExtensionRepoVersionExtensions(params *GetExtensionRepoVersionExtensionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionRepoVersionExtensionsOK, error)

	GetExtensionRepoVersionSha256(params *GetExtensionRepoVersionSha256Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionRepoVersionSha256OK, error)

	GetExtensionRepoVersions(params *GetExtensionRepoVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionRepoVersionsOK, error)

	GetGlobalExtensionRepoVersionSha256(params *GetGlobalExtensionRepoVersionSha256Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGlobalExtensionRepoVersionSha256OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	GetExtensionRepoArtifacts gets extension repo artifacts

	Gets the artifacts in the extension repository in the given bucket and group.

NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.
*/
func (a *Client) GetExtensionRepoArtifacts(params *GetExtensionRepoArtifactsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionRepoArtifactsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExtensionRepoArtifactsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExtensionRepoArtifacts",
		Method:             "GET",
		PathPattern:        "/extension-repository/{bucketName}/{groupId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetExtensionRepoArtifactsReader{formats: a.formats},
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
	success, ok := result.(*GetExtensionRepoArtifactsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExtensionRepoArtifacts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetExtensionRepoBuckets gets extension repo buckets

	Gets the names of the buckets the current user is authorized for in order to browse the repo by bucket.

NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.
*/
func (a *Client) GetExtensionRepoBuckets(params *GetExtensionRepoBucketsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionRepoBucketsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExtensionRepoBucketsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExtensionRepoBuckets",
		Method:             "GET",
		PathPattern:        "/extension-repository",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetExtensionRepoBucketsReader{formats: a.formats},
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
	success, ok := result.(*GetExtensionRepoBucketsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExtensionRepoBuckets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetExtensionRepoGroups gets extension repo groups

	Gets the groups in the extension repository in the given bucket.

NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.
*/
func (a *Client) GetExtensionRepoGroups(params *GetExtensionRepoGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionRepoGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExtensionRepoGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExtensionRepoGroups",
		Method:             "GET",
		PathPattern:        "/extension-repository/{bucketName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetExtensionRepoGroupsReader{formats: a.formats},
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
	success, ok := result.(*GetExtensionRepoGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExtensionRepoGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetExtensionRepoVersion gets extension repo version

	Gets information about the version in the given bucket, group, and artifact.

NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.
*/
func (a *Client) GetExtensionRepoVersion(params *GetExtensionRepoVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionRepoVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExtensionRepoVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExtensionRepoVersion",
		Method:             "GET",
		PathPattern:        "/extension-repository/{bucketName}/{groupId}/{artifactId}/{version}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetExtensionRepoVersionReader{formats: a.formats},
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
	success, ok := result.(*GetExtensionRepoVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExtensionRepoVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetExtensionRepoVersionContent gets extension repo version content

	Gets the binary content of the bundle with the given bucket, group, artifact, and version.

NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.
*/
func (a *Client) GetExtensionRepoVersionContent(params *GetExtensionRepoVersionContentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionRepoVersionContentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExtensionRepoVersionContentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExtensionRepoVersionContent",
		Method:             "GET",
		PathPattern:        "/extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/content",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetExtensionRepoVersionContentReader{formats: a.formats},
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
	success, ok := result.(*GetExtensionRepoVersionContentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExtensionRepoVersionContent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetExtensionRepoVersionExtension gets extension repo extension

	Gets information about the extension with the given name in the given bucket, group, artifact, and version.

NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.
*/
func (a *Client) GetExtensionRepoVersionExtension(params *GetExtensionRepoVersionExtensionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionRepoVersionExtensionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExtensionRepoVersionExtensionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExtensionRepoVersionExtension",
		Method:             "GET",
		PathPattern:        "/extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetExtensionRepoVersionExtensionReader{formats: a.formats},
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
	success, ok := result.(*GetExtensionRepoVersionExtensionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExtensionRepoVersionExtension: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetExtensionRepoVersionExtensionAdditionalDetailsDocs gets extension repo extension details

	Gets the additional details documentation for the extension with the given name in the given bucket, group, artifact, and version.

NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.
*/
func (a *Client) GetExtensionRepoVersionExtensionAdditionalDetailsDocs(params *GetExtensionRepoVersionExtensionAdditionalDetailsDocsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionRepoVersionExtensionAdditionalDetailsDocsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExtensionRepoVersionExtensionAdditionalDetailsDocsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExtensionRepoVersionExtensionAdditionalDetailsDocs",
		Method:             "GET",
		PathPattern:        "/extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs/additional-details",
		ProducesMediaTypes: []string{"text/html"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetExtensionRepoVersionExtensionAdditionalDetailsDocsReader{formats: a.formats},
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
	success, ok := result.(*GetExtensionRepoVersionExtensionAdditionalDetailsDocsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExtensionRepoVersionExtensionAdditionalDetailsDocs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetExtensionRepoVersionExtensionDocs gets extension repo extension docs

	Gets the documentation for the extension with the given name in the given bucket, group, artifact, and version.

NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.
*/
func (a *Client) GetExtensionRepoVersionExtensionDocs(params *GetExtensionRepoVersionExtensionDocsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionRepoVersionExtensionDocsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExtensionRepoVersionExtensionDocsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExtensionRepoVersionExtensionDocs",
		Method:             "GET",
		PathPattern:        "/extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs",
		ProducesMediaTypes: []string{"text/html"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetExtensionRepoVersionExtensionDocsReader{formats: a.formats},
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
	success, ok := result.(*GetExtensionRepoVersionExtensionDocsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExtensionRepoVersionExtensionDocs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetExtensionRepoVersionExtensions gets extension repo extensions

	Gets information about the extensions in the given bucket, group, artifact, and version.

NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.
*/
func (a *Client) GetExtensionRepoVersionExtensions(params *GetExtensionRepoVersionExtensionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionRepoVersionExtensionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExtensionRepoVersionExtensionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExtensionRepoVersionExtensions",
		Method:             "GET",
		PathPattern:        "/extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetExtensionRepoVersionExtensionsReader{formats: a.formats},
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
	success, ok := result.(*GetExtensionRepoVersionExtensionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExtensionRepoVersionExtensions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetExtensionRepoVersionSha256 gets extension repo version checksum

	Gets the hex representation of the SHA-256 digest for the binary content of the bundle with the given bucket, group, artifact, and version.

NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.
*/
func (a *Client) GetExtensionRepoVersionSha256(params *GetExtensionRepoVersionSha256Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionRepoVersionSha256OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExtensionRepoVersionSha256Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExtensionRepoVersionSha256",
		Method:             "GET",
		PathPattern:        "/extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/sha256",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetExtensionRepoVersionSha256Reader{formats: a.formats},
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
	success, ok := result.(*GetExtensionRepoVersionSha256OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExtensionRepoVersionSha256: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetExtensionRepoVersions gets extension repo versions

	Gets the versions in the extension repository for the given bucket, group, and artifact.

NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.
*/
func (a *Client) GetExtensionRepoVersions(params *GetExtensionRepoVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionRepoVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExtensionRepoVersionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExtensionRepoVersions",
		Method:             "GET",
		PathPattern:        "/extension-repository/{bucketName}/{groupId}/{artifactId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetExtensionRepoVersionsReader{formats: a.formats},
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
	success, ok := result.(*GetExtensionRepoVersionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExtensionRepoVersions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetGlobalExtensionRepoVersionSha256 gets global extension repo version checksum

	Gets the hex representation of the SHA-256 digest for the binary content with the given bucket, group, artifact, and version. Since the same group-artifact-version can exist in multiple buckets, this will return the checksum of the first one returned. This will be consistent since the checksum must be the same when existing in multiple buckets.

NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.
*/
func (a *Client) GetGlobalExtensionRepoVersionSha256(params *GetGlobalExtensionRepoVersionSha256Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGlobalExtensionRepoVersionSha256OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGlobalExtensionRepoVersionSha256Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGlobalExtensionRepoVersionSha256",
		Method:             "GET",
		PathPattern:        "/extension-repository/{groupId}/{artifactId}/{version}/sha256",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetGlobalExtensionRepoVersionSha256Reader{formats: a.formats},
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
	success, ok := result.(*GetGlobalExtensionRepoVersionSha256OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGlobalExtensionRepoVersionSha256: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

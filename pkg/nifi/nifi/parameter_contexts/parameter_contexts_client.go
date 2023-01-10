// Code generated by go-swagger; DO NOT EDIT.

package parameter_contexts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new parameter contexts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for parameter contexts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateParameterContext(params *CreateParameterContextParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateParameterContextCreated, error)

	DeleteParameterContext(params *DeleteParameterContextParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteParameterContextOK, error)

	DeleteValidationRequest(params *DeleteValidationRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteValidationRequestOK, error)

	GetParameterContext(params *GetParameterContextParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetParameterContextOK, error)

	GetParameterContextUpdate(params *GetParameterContextUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetParameterContextUpdateOK, error)

	GetValidationRequest(params *GetValidationRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetValidationRequestOK, error)

	SubmitParameterContextUpdate(params *SubmitParameterContextUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubmitParameterContextUpdateOK, error)

	SubmitValidationRequest(params *SubmitValidationRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubmitValidationRequestOK, error)

	UpdateParameterContext(params *UpdateParameterContextParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateParameterContextOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateParameterContext creates a parameter context
*/
func (a *Client) CreateParameterContext(params *CreateParameterContextParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateParameterContextCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateParameterContextParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createParameterContext",
		Method:             "POST",
		PathPattern:        "/parameter-contexts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateParameterContextReader{formats: a.formats},
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
	success, ok := result.(*CreateParameterContextCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createParameterContext: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteParameterContext deletes the parameter context with the given ID

Deletes the Parameter Context with the given ID.
*/
func (a *Client) DeleteParameterContext(params *DeleteParameterContextParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteParameterContextOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteParameterContextParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteParameterContext",
		Method:             "DELETE",
		PathPattern:        "/parameter-contexts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteParameterContextReader{formats: a.formats},
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
	success, ok := result.(*DeleteParameterContextOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteParameterContext: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteValidationRequest deletes the validation request with the given ID

Deletes the Validation Request with the given ID. After a request is created via a POST to /nifi-api/validation-contexts, it is expected that the client will properly clean up the request by DELETE'ing it, once the validation process has completed. If the request is deleted before the request completes, then the Validation request will finish the step that it is currently performing and then will cancel any subsequent steps.
*/
func (a *Client) DeleteValidationRequest(params *DeleteValidationRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteValidationRequestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteValidationRequestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteValidationRequest",
		Method:             "DELETE",
		PathPattern:        "/parameter-contexts/{contextId}/validation-requests/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteValidationRequestReader{formats: a.formats},
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
	success, ok := result.(*DeleteValidationRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteValidationRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetParameterContext returns the parameter context with the given ID

Returns the Parameter Context with the given ID.
*/
func (a *Client) GetParameterContext(params *GetParameterContextParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetParameterContextOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetParameterContextParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getParameterContext",
		Method:             "GET",
		PathPattern:        "/parameter-contexts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetParameterContextReader{formats: a.formats},
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
	success, ok := result.(*GetParameterContextOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getParameterContext: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetParameterContextUpdate returns the update request with the given ID

Returns the Update Request with the given ID. Once an Update Request has been created by performing a POST to /nifi-api/parameter-contexts, that request can subsequently be retrieved via this endpoint, and the request that is fetched will contain the updated state, such as percent complete, the current state of the request, and any failures.
*/
func (a *Client) GetParameterContextUpdate(params *GetParameterContextUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetParameterContextUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetParameterContextUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getParameterContextUpdate",
		Method:             "GET",
		PathPattern:        "/parameter-contexts/{contextId}/update-requests/{requestId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetParameterContextUpdateReader{formats: a.formats},
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
	success, ok := result.(*GetParameterContextUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getParameterContextUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetValidationRequest returns the validation request with the given ID

Returns the Validation Request with the given ID. Once a Validation Request has been created by performing a POST to /nifi-api/validation-contexts, that request can subsequently be retrieved via this endpoint, and the request that is fetched will contain the updated state, such as percent complete, the current state of the request, and any failures.
*/
func (a *Client) GetValidationRequest(params *GetValidationRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetValidationRequestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetValidationRequestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getValidationRequest",
		Method:             "GET",
		PathPattern:        "/parameter-contexts/{contextId}/validation-requests/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetValidationRequestReader{formats: a.formats},
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
	success, ok := result.(*GetValidationRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getValidationRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SubmitParameterContextUpdate initiates the update request of a parameter context

This will initiate the process of updating a Parameter Context. Changing the value of a Parameter may require that one or more components be stopped and restarted, so this action may take significantly more time than many other REST API actions. As a result, this endpoint will immediately return a ParameterContextUpdateRequestEntity, and the process of updating the necessary components will occur asynchronously in the background. The client may then periodically poll the status of the request by issuing a GET request to /parameter-contexts/update-requests/{requestId}. Once the request is completed, the client is expected to issue a DELETE request to /parameter-contexts/update-requests/{requestId}.
*/
func (a *Client) SubmitParameterContextUpdate(params *SubmitParameterContextUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubmitParameterContextUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubmitParameterContextUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "submitParameterContextUpdate",
		Method:             "POST",
		PathPattern:        "/parameter-contexts/{contextId}/update-requests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SubmitParameterContextUpdateReader{formats: a.formats},
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
	success, ok := result.(*SubmitParameterContextUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for submitParameterContextUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SubmitValidationRequest initiates a validation request to determine how the validity of components will change if a parameter context were to be updated

This will initiate the process of validating all components whose Process Group is bound to the specified Parameter Context. Performing validation against an arbitrary number of components may be expect and take significantly more time than many other REST API actions. As a result, this endpoint will immediately return a ParameterContextValidationRequestEntity, and the process of validating the necessary components will occur asynchronously in the background. The client may then periodically poll the status of the request by issuing a GET request to /parameter-contexts/validation-requests/{requestId}. Once the request is completed, the client is expected to issue a DELETE request to /parameter-contexts/validation-requests/{requestId}.
*/
func (a *Client) SubmitValidationRequest(params *SubmitValidationRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubmitValidationRequestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubmitValidationRequestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "submitValidationRequest",
		Method:             "POST",
		PathPattern:        "/parameter-contexts/{contextId}/validation-requests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SubmitValidationRequestReader{formats: a.formats},
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
	success, ok := result.(*SubmitValidationRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for submitValidationRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateParameterContext modifies a parameter context

This endpoint will update a Parameter Context to match the provided entity. However, this request will fail if any component is running and is referencing a Parameter in the Parameter Context. Generally, this endpoint is not called directly. Instead, an update request should be submitted by making a POST to the /parameter-contexts/update-requests endpoint. That endpoint will, in turn, call this endpoint.
*/
func (a *Client) UpdateParameterContext(params *UpdateParameterContextParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateParameterContextOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateParameterContextParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateParameterContext",
		Method:             "PUT",
		PathPattern:        "/parameter-contexts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateParameterContextReader{formats: a.formats},
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
	success, ok := result.(*UpdateParameterContextOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateParameterContext: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

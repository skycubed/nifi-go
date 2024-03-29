// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new access API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for access API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateAccessTokenByTryingAllProviders(params *CreateAccessTokenByTryingAllProvidersParams, opts ...ClientOption) (*CreateAccessTokenByTryingAllProvidersOK, error)

	CreateAccessTokenUsingBasicAuthCredentials(params *CreateAccessTokenUsingBasicAuthCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAccessTokenUsingBasicAuthCredentialsOK, error)

	CreateAccessTokenUsingIdentityProviderCredentials(params *CreateAccessTokenUsingIdentityProviderCredentialsParams, opts ...ClientOption) (*CreateAccessTokenUsingIdentityProviderCredentialsOK, error)

	CreateAccessTokenUsingKerberosTicket(params *CreateAccessTokenUsingKerberosTicketParams, opts ...ClientOption) (*CreateAccessTokenUsingKerberosTicketOK, error)

	GetAccessStatus(params *GetAccessStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAccessStatusOK, error)

	GetIdentityProviderUsageInstructions(params *GetIdentityProviderUsageInstructionsParams, opts ...ClientOption) (*GetIdentityProviderUsageInstructionsOK, error)

	LogOut(params *LogOutParams, opts ...ClientOption) (*LogOutOK, error)

	OidcCallback(params *OidcCallbackParams, opts ...ClientOption) error

	OidcExchange(params *OidcExchangeParams, opts ...ClientOption) (*OidcExchangeOK, error)

	OidcLogout(params *OidcLogoutParams, opts ...ClientOption) error

	OidcRequest(params *OidcRequestParams, opts ...ClientOption) error

	TestIdentityProviderRecognizesCredentialsFormat(params *TestIdentityProviderRecognizesCredentialsFormatParams, opts ...ClientOption) (*TestIdentityProviderRecognizesCredentialsFormatOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateAccessTokenByTryingAllProviders creates token trying all providers

Creates a token for accessing the REST API via auto-detected method of verifying client identity claim credentials. The token returned is formatted as a JSON Web Token (JWT). The token is base64 encoded and comprised of three parts. The header, the body, and the signature. The expiration of the token is a contained within the body. The token can be used in the Authorization header in the format 'Authorization: Bearer <token>'.
*/
func (a *Client) CreateAccessTokenByTryingAllProviders(params *CreateAccessTokenByTryingAllProvidersParams, opts ...ClientOption) (*CreateAccessTokenByTryingAllProvidersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAccessTokenByTryingAllProvidersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createAccessTokenByTryingAllProviders",
		Method:             "POST",
		PathPattern:        "/access/token",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateAccessTokenByTryingAllProvidersReader{formats: a.formats},
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
	success, ok := result.(*CreateAccessTokenByTryingAllProvidersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createAccessTokenByTryingAllProviders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateAccessTokenUsingBasicAuthCredentials creates token using basic auth

Creates a token for accessing the REST API via username/password. The user credentials must be passed in standard HTTP Basic Auth format. That is: 'Authorization: Basic <credentials>', where <credentials> is the base64 encoded value of '<username>:<password>'. The token returned is formatted as a JSON Web Token (JWT). The token is base64 encoded and comprised of three parts. The header, the body, and the signature. The expiration of the token is a contained within the body. The token can be used in the Authorization header in the format 'Authorization: Bearer <token>'.
*/
func (a *Client) CreateAccessTokenUsingBasicAuthCredentials(params *CreateAccessTokenUsingBasicAuthCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAccessTokenUsingBasicAuthCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAccessTokenUsingBasicAuthCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createAccessTokenUsingBasicAuthCredentials",
		Method:             "POST",
		PathPattern:        "/access/token/login",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateAccessTokenUsingBasicAuthCredentialsReader{formats: a.formats},
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
	success, ok := result.(*CreateAccessTokenUsingBasicAuthCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createAccessTokenUsingBasicAuthCredentials: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateAccessTokenUsingIdentityProviderCredentials creates token using identity provider

Creates a token for accessing the REST API via a custom identity provider. The user credentials must be passed in a format understood by the custom identity provider, e.g., a third-party auth token in an HTTP header. The exact format of the user credentials expected by the custom identity provider can be discovered by 'GET /access/token/identity-provider/usage'. The token returned is formatted as a JSON Web Token (JWT). The token is base64 encoded and comprised of three parts. The header, the body, and the signature. The expiration of the token is a contained within the body. The token can be used in the Authorization header in the format 'Authorization: Bearer <token>'.
*/
func (a *Client) CreateAccessTokenUsingIdentityProviderCredentials(params *CreateAccessTokenUsingIdentityProviderCredentialsParams, opts ...ClientOption) (*CreateAccessTokenUsingIdentityProviderCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAccessTokenUsingIdentityProviderCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createAccessTokenUsingIdentityProviderCredentials",
		Method:             "POST",
		PathPattern:        "/access/token/identity-provider",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateAccessTokenUsingIdentityProviderCredentialsReader{formats: a.formats},
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
	success, ok := result.(*CreateAccessTokenUsingIdentityProviderCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createAccessTokenUsingIdentityProviderCredentials: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateAccessTokenUsingKerberosTicket creates token using kerberos

Creates a token for accessing the REST API via Kerberos Service Tickets or SPNEGO Tokens (which includes Kerberos Service Tickets). The token returned is formatted as a JSON Web Token (JWT). The token is base64 encoded and comprised of three parts. The header, the body, and the signature. The expiration of the token is a contained within the body. The token can be used in the Authorization header in the format 'Authorization: Bearer <token>'.
*/
func (a *Client) CreateAccessTokenUsingKerberosTicket(params *CreateAccessTokenUsingKerberosTicketParams, opts ...ClientOption) (*CreateAccessTokenUsingKerberosTicketOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAccessTokenUsingKerberosTicketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createAccessTokenUsingKerberosTicket",
		Method:             "POST",
		PathPattern:        "/access/token/kerberos",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateAccessTokenUsingKerberosTicketReader{formats: a.formats},
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
	success, ok := result.(*CreateAccessTokenUsingKerberosTicketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createAccessTokenUsingKerberosTicket: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAccessStatus gets access status

Returns the current client's authenticated identity and permissions to top-level resources
*/
func (a *Client) GetAccessStatus(params *GetAccessStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAccessStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccessStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAccessStatus",
		Method:             "GET",
		PathPattern:        "/access",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAccessStatusReader{formats: a.formats},
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
	success, ok := result.(*GetAccessStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAccessStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetIdentityProviderUsageInstructions gets identity provider usage

Provides a description of how the currently configured identity provider expects credentials to be passed to POST /access/token/identity-provider
*/
func (a *Client) GetIdentityProviderUsageInstructions(params *GetIdentityProviderUsageInstructionsParams, opts ...ClientOption) (*GetIdentityProviderUsageInstructionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIdentityProviderUsageInstructionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getIdentityProviderUsageInstructions",
		Method:             "GET",
		PathPattern:        "/access/token/identity-provider/usage",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetIdentityProviderUsageInstructionsReader{formats: a.formats},
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
	success, ok := result.(*GetIdentityProviderUsageInstructionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getIdentityProviderUsageInstructions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	LogOut performs a logout for other providers that have been issued a j w t

NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.
*/
func (a *Client) LogOut(params *LogOutParams, opts ...ClientOption) (*LogOutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLogOutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "logOut",
		Method:             "DELETE",
		PathPattern:        "/access/logout",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LogOutReader{formats: a.formats},
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
	success, ok := result.(*LogOutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for logOut: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	OidcCallback redirects callback URI for processing the result of the open Id connect login sequence

NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.
*/
func (a *Client) OidcCallback(params *OidcCallbackParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOidcCallbackParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "oidcCallback",
		Method:             "GET",
		PathPattern:        "/access/oidc/callback",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &OidcCallbackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
	OidcExchange retrieves a j w t following a successful login sequence using the configured open Id connect provider

NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.
*/
func (a *Client) OidcExchange(params *OidcExchangeParams, opts ...ClientOption) (*OidcExchangeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOidcExchangeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "oidcExchange",
		Method:             "POST",
		PathPattern:        "/access/oidc/exchange",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &OidcExchangeReader{formats: a.formats},
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
	success, ok := result.(*OidcExchangeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for oidcExchange: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	OidcLogout performs a logout in the open Id provider

NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.
*/
func (a *Client) OidcLogout(params *OidcLogoutParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOidcLogoutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "oidcLogout",
		Method:             "DELETE",
		PathPattern:        "/access/oidc/logout",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &OidcLogoutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
	OidcRequest initiates a request to authenticate through the configured open Id connect provider

NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.
*/
func (a *Client) OidcRequest(params *OidcRequestParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOidcRequestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "oidcRequest",
		Method:             "GET",
		PathPattern:        "/access/oidc/request",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &OidcRequestReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
TestIdentityProviderRecognizesCredentialsFormat tests identity provider

Tests the format of the credentials against this identity provider without preforming authentication on the credentials to validate them. The user credentials should be passed in a format understood by the custom identity provider as defined by 'GET /access/token/identity-provider/usage'.
*/
func (a *Client) TestIdentityProviderRecognizesCredentialsFormat(params *TestIdentityProviderRecognizesCredentialsFormatParams, opts ...ClientOption) (*TestIdentityProviderRecognizesCredentialsFormatOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestIdentityProviderRecognizesCredentialsFormatParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "testIdentityProviderRecognizesCredentialsFormat",
		Method:             "POST",
		PathPattern:        "/access/token/identity-provider/test",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TestIdentityProviderRecognizesCredentialsFormatReader{formats: a.formats},
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
	success, ok := result.(*TestIdentityProviderRecognizesCredentialsFormatOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for testIdentityProviderRecognizesCredentialsFormat: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

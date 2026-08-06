package registry

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// WithBearerToken configures every generated request to use HTTP Bearer
// authentication. The token is validated when the client is constructed.
func WithBearerToken(token string) ClientOption {
	return func(client *Client) error {
		if token == "" || token != strings.TrimSpace(token) || strings.ContainsAny(token, "\r\n") {
			return errors.New("registry: bearer token must be non-empty and contain no surrounding whitespace or newlines")
		}
		return WithRequestEditorFn(func(_ context.Context, request *http.Request) error {
			request.Header.Set("Authorization", "Bearer "+token)
			return nil
		})(client)
	}
}

// JSONResponseBody returns a typed successful JSON payload from a generated
// response. Registry releases differ in whether successful payloads are
// declared as an explicit 2xx response or as the OpenAPI default response.
func JSONResponseBody[T any](response any) *T {
	if typed, ok := response.(interface{ GetJSON200() *T }); ok && typed.GetJSON200() != nil {
		return typed.GetJSON200()
	}
	if typed, ok := response.(interface{ GetJSON201() *T }); ok && typed.GetJSON201() != nil {
		return typed.GetJSON201()
	}
	if typed, ok := response.(interface{ GetJSON202() *T }); ok && typed.GetJSON202() != nil {
		return typed.GetJSON202()
	}
	if typed, ok := response.(interface{ GetJSONDefault() *T }); ok {
		return typed.GetJSONDefault()
	}
	return nil
}

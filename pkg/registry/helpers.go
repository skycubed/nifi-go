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

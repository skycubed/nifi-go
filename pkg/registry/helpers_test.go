package registry

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

type doerFunc func(*http.Request) (*http.Response, error)

func (f doerFunc) Do(request *http.Request) (*http.Response, error) {
	return f(request)
}

func TestAuthenticatedVersionRequest(t *testing.T) {
	t.Parallel()
	doer := doerFunc(func(r *http.Request) (*http.Response, error) {
		if got, want := r.URL.Path, "/nifi-registry-api/about"; got != want {
			t.Errorf("path = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer registry.jwt.token"; got != want {
			t.Errorf("Authorization = %q, want %q", got, want)
		}
		return &http.Response{
			Status:     "200 OK",
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(strings.NewReader(fmt.Sprintf(`{"registryAboutVersion":%q}`, TargetRegistryVersion))),
			Request:    r,
		}, nil
	})

	client, err := NewClientWithResponses("https://example.test/nifi-registry-api", WithBearerToken("registry.jwt.token"), WithHTTPClient(doer))
	if err != nil {
		t.Fatal(err)
	}
	response, err := client.GetVersionWithResponse(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode() != http.StatusOK || response.JSONDefault == nil || response.JSONDefault.RegistryAboutVersion == nil {
		t.Fatalf("unexpected response: status=%d body=%s", response.StatusCode(), response.Body)
	}
	if got, want := *response.JSONDefault.RegistryAboutVersion, TargetRegistryVersion; got != want {
		t.Errorf("version = %q, want %q", got, want)
	}
}

func TestWithBearerTokenRejectsInvalidValue(t *testing.T) {
	t.Parallel()
	if _, err := NewClientWithResponses("https://example.test/nifi-registry-api", WithBearerToken("bad\ntoken")); err == nil {
		t.Fatal("expected invalid token to be rejected")
	}
}

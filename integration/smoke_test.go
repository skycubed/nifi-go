//go:build integration

package integration_test

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/skycubed/nifi-go/v2/pkg/nifi"
	"github.com/skycubed/nifi-go/v2/pkg/registry"
)

func TestNiFiAbout(t *testing.T) {
	baseURL := os.Getenv("NIFI_BASE_URL")
	token := os.Getenv("NIFI_TOKEN")
	if baseURL == "" || token == "" {
		t.Skip("NIFI_BASE_URL and NIFI_TOKEN are required")
	}

	certificatePath := os.Getenv("NIFI_CA_FILE")
	certificate, err := os.ReadFile(certificatePath)
	if err != nil {
		t.Fatalf("read NIFI_CA_FILE: %v", err)
	}
	roots := x509.NewCertPool()
	if !roots.AppendCertsFromPEM(certificate) {
		t.Fatal("NIFI_CA_FILE contained no certificates")
	}
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.TLSClientConfig = &tls.Config{MinVersion: tls.VersionTLS12, RootCAs: roots}
	httpClient := &http.Client{Transport: transport}
	client, err := nifi.NewClientWithResponses(baseURL, nifi.WithHTTPClient(httpClient), nifi.WithBearerToken(token))
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	response, err := client.GetAboutInfoWithResponse(ctx)
	if err != nil {
		t.Fatal(err)
	}
	entity := nifiAboutEntity(response)
	if response.StatusCode() != http.StatusOK || entity == nil || entity.About == nil || entity.About.Version == nil {
		t.Fatalf("unexpected response: status=%d body=%s", response.StatusCode(), response.Body)
	}
	if got, want := *entity.About.Version, nifi.TargetNiFiVersion; got != want {
		t.Fatalf("server version = %q, binding version = %q", got, want)
	}
}

func nifiAboutEntity(response *nifi.GetAboutInfoResponse) *nifi.AboutEntity {
	if typed, ok := any(response).(interface{ GetJSON200() *nifi.AboutEntity }); ok {
		return typed.GetJSON200()
	}
	if typed, ok := any(response).(interface{ GetJSONDefault() *nifi.AboutEntity }); ok {
		return typed.GetJSONDefault()
	}
	return nil
}

func TestRegistryAbout(t *testing.T) {
	baseURL := os.Getenv("REGISTRY_BASE_URL")
	if baseURL == "" {
		t.Skip("REGISTRY_BASE_URL is required")
	}

	options := []registry.ClientOption{registry.WithHTTPClient(http.DefaultClient)}
	if token := os.Getenv("REGISTRY_TOKEN"); token != "" {
		options = append(options, registry.WithBearerToken(token))
	}
	client, err := registry.NewClientWithResponses(baseURL, options...)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	response, err := client.GetVersionWithResponse(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode() != http.StatusOK || response.JSONDefault == nil || response.JSONDefault.RegistryAboutVersion == nil {
		t.Fatalf("unexpected response: status=%d body=%s", response.StatusCode(), response.Body)
	}
	if got, want := *response.JSONDefault.RegistryAboutVersion, registry.TargetRegistryVersion; got != want {
		t.Fatalf("server version = %q, binding version = %q", got, want)
	}
}

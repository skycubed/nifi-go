package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestNormalizeProducesValidDeterministicSpec(t *testing.T) {
	t.Parallel()
	raw := []byte(`{
  "openapi": "3.0.1",
  "info": {"title": "test", "version": "2.11.0"},
  "components": {"securitySchemes": {"JWT": {"type": "http", "scheme": "Bearer"}}},
  "paths": {"/about": {"get": {
    "operationId": "getAbout",
    "security": [{"Read - /flow": []}, {"JWT": []}],
    "responses": {"200": {"content": {"application/json": {"schema": {"type": "object"}}}}}
  }}}
}`)

	first, err := normalize(raw, "2.11.0")
	if err != nil {
		t.Fatal(err)
	}
	second, err := normalize(raw, "2.11.0")
	if err != nil {
		t.Fatal(err)
	}
	if string(first) != string(second) {
		t.Fatal("normalization is not deterministic")
	}
	if err := validate(first, "2.11.0"); err != nil {
		t.Fatalf("normalized spec is invalid: %v", err)
	}

	var doc map[string]any
	if err := json.Unmarshal(first, &doc); err != nil {
		t.Fatal(err)
	}
	components, _ := object(doc["components"])
	schemes, _ := object(components["securitySchemes"])
	jwt, _ := object(schemes["JWT"])
	if got, want := jwt["scheme"], "bearer"; got != want {
		t.Errorf("scheme = %v, want %v", got, want)
	}
	paths, _ := object(doc["paths"])
	pathItem, _ := object(paths["/about"])
	operation, _ := object(pathItem["get"])
	if got := operation["x-nifi-authorizations"].([]any); len(got) != 1 || got[0] != "Read - /flow" {
		t.Errorf("authorizations = %#v", got)
	}
	responses, _ := object(operation["responses"])
	response, _ := object(responses["200"])
	if got, want := response["description"], "HTTP 200 response"; got != want {
		t.Errorf("description = %v, want %v", got, want)
	}
}

func TestNormalizeRejectsVersionMismatch(t *testing.T) {
	t.Parallel()
	_, err := normalize([]byte(`{"info":{"version":"2.10.0"},"paths":{}}`), "2.11.0")
	if err == nil || !strings.Contains(err.Error(), "does not match") {
		t.Fatalf("expected version mismatch, got %v", err)
	}
}

func TestNormalizeRejectsDuplicateOperationID(t *testing.T) {
	t.Parallel()
	raw := []byte(`{
  "info": {"version": "2.11.0"},
  "paths": {
    "/one": {"get": {"operationId": "duplicate"}},
    "/two": {"post": {"operationId": "duplicate"}}
  }
}`)
	_, err := normalize(raw, "2.11.0")
	if err == nil || !strings.Contains(err.Error(), "duplicate operationId") {
		t.Fatalf("expected duplicate operationId, got %v", err)
	}
}

func TestNormalizeRejectsMissingOperationID(t *testing.T) {
	t.Parallel()
	_, err := normalize([]byte(`{
  "info": {"version": "2.11.0"},
  "paths": {"/one": {"get": {"responses": {"204": {}}}}}
}`), "2.11.0")
	if err == nil || !strings.Contains(err.Error(), "has no operationId") {
		t.Fatalf("expected missing operationId, got %v", err)
	}
}

func TestValidateArtifactSourceRejectsUntrustedURL(t *testing.T) {
	t.Parallel()
	if err := validateArtifactSource("nifi", "2.11.0", "https://example.test/nifi.war"); err == nil {
		t.Fatal("expected untrusted artifact URL to be rejected")
	}
}

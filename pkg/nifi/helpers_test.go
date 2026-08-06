package nifi

import (
	"context"
	"encoding/json"
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

func TestWithBearerToken(t *testing.T) {
	t.Parallel()

	doer := doerFunc(func(r *http.Request) (*http.Response, error) {
		if got, want := r.URL.Path, "/nifi-api/flow/about"; got != want {
			t.Errorf("path = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test.jwt.token"; got != want {
			t.Errorf("Authorization = %q, want %q", got, want)
		}
		return &http.Response{
			Status:     "200 OK",
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(strings.NewReader(fmt.Sprintf(`{"about":{"version":%q}}`, TargetNiFiVersion))),
			Request:    r,
		}, nil
	})

	client, err := NewClientWithResponses("https://example.test/nifi-api", WithBearerToken("test.jwt.token"), WithHTTPClient(doer))
	if err != nil {
		t.Fatal(err)
	}
	response, err := client.GetAboutInfoWithResponse(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode() != http.StatusOK {
		t.Fatalf("status = %d, body = %s", response.StatusCode(), response.Body)
	}
	entity := JSONResponseBody[AboutEntity](response)
	if entity == nil || entity.About == nil || entity.About.Version == nil {
		t.Fatalf("missing typed response: %#v", response)
	}
	if got, want := *entity.About.Version, TargetNiFiVersion; got != want {
		t.Errorf("version = %q, want %q", got, want)
	}
}

func TestWithBearerTokenRejectsInvalidValues(t *testing.T) {
	t.Parallel()
	for _, token := range []string{"", " token", "token ", "token\rvalue", "token\nvalue"} {
		t.Run(strings.ReplaceAll(token, "\n", "newline"), func(t *testing.T) {
			if _, err := NewClientWithResponses("https://example.test/nifi-api", WithBearerToken(token)); err == nil {
				t.Fatalf("expected token %q to be rejected", token)
			}
		})
	}
}

func TestTypedResponsePreservesNonSuccessBody(t *testing.T) {
	t.Parallel()
	doer := doerFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			Status:     "409 Conflict",
			StatusCode: http.StatusConflict,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader("NiFi is not in the appropriate state")),
			Request:    r,
		}, nil
	})

	client, err := NewClientWithResponses("https://example.test/nifi-api", WithHTTPClient(doer))
	if err != nil {
		t.Fatal(err)
	}
	response, err := client.GetAboutInfoWithResponse(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if got, want := response.StatusCode(), http.StatusConflict; got != want {
		t.Errorf("status = %d, want %d", got, want)
	}
	if got, want := string(response.Body), "NiFi is not in the appropriate state"; got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}

func TestRequestConstructionBuildsPathAndQuery(t *testing.T) {
	t.Parallel()
	req, err := NewGetProcessorRequest("https://example.test/nifi-api/", "processor-id")
	if err != nil {
		t.Fatal(err)
	}
	if got, want := req.URL.Path, "/nifi-api/processors/processor-id"; got != want {
		t.Errorf("path = %q, want %q", got, want)
	}

	nodewise := true
	req, err = NewGetCountersRequest("https://example.test/nifi-api/", &GetCountersParams{Nodewise: &nodewise})
	if err != nil {
		t.Fatal(err)
	}
	if got, want := req.URL.Query().Get("nodewise"), "true"; got != want {
		t.Errorf("nodewise = %q, want %q", got, want)
	}
}

func TestRequestConstructionSerializesJSON(t *testing.T) {
	t.Parallel()
	version := int64(7)
	name := "Updated processor"
	body := ProcessorEntity{
		Revision:  &RevisionDTO{Version: &version},
		Component: &ProcessorDTO{Name: &name},
	}
	request, err := NewUpdateProcessorRequest("https://example.test/nifi-api/", "processor-id", body)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := request.Header.Get("Content-Type"), "application/json"; got != want {
		t.Errorf("Content-Type = %q, want %q", got, want)
	}
	defer request.Body.Close()
	var decoded ProcessorEntity
	if err := json.NewDecoder(request.Body).Decode(&decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.Revision == nil || decoded.Revision.Version == nil || *decoded.Revision.Version != version {
		t.Fatalf("revision was not serialized: %#v", decoded.Revision)
	}
	if decoded.Component == nil || decoded.Component.Name == nil || *decoded.Component.Name != name {
		t.Fatalf("component was not serialized: %#v", decoded.Component)
	}
}

func TestRawClientLeavesBodyForStreamingCaller(t *testing.T) {
	t.Parallel()
	doer := doerFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			Status:     "200 OK",
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader("streamed response")),
			Request:    r,
		}, nil
	})

	client, err := NewClient("https://example.test/nifi-api", WithHTTPClient(doer))
	if err != nil {
		t.Fatal(err)
	}
	response, err := client.GetAboutInfo(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(body), "streamed response"; got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}

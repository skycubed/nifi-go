# Apache NiFi Go bindings

Generated, version-matched Go clients for the Apache NiFi and Apache NiFi
Registry REST APIs. Every release is generated from the OpenAPI document in an
official Apache NiFi release artifact and is intended for one exact NiFi
version.

## Compatibility and installation

Use the binding tag matching the server. Cross-version compatibility is not
assumed or guaranteed.

| NiFi / Registry | Branch | Binding tag |
| --- | --- | --- |
| 2.0.0 | `nifi-2.0.0` | `v2.0.0-bindings.2` |
| 2.1.0 | `nifi-2.1.0` | `v2.1.0-bindings.2` |
| 2.2.0 | `nifi-2.2.0` | `v2.2.0-bindings.2` |
| 2.3.0 | `nifi-2.3.0` | `v2.3.0-bindings.2` |
| 2.4.0 | `nifi-2.4.0` | `v2.4.0-bindings.2` |
| 2.5.0 | `nifi-2.5.0` | `v2.5.0-bindings.2` |
| 2.6.0 | `nifi-2.6.0` | `v2.6.0-bindings.2` |
| 2.7.0 | `nifi-2.7.0` | `v2.7.0-bindings.2` |
| 2.7.1 | `nifi-2.7.1` | `v2.7.1-bindings.2` |
| 2.7.2 | `nifi-2.7.2` | `v2.7.2-bindings.2` |
| 2.8.0 | `nifi-2.8.0` | `v2.8.0-bindings.2` |
| 2.9.0 | `nifi-2.9.0` | `v2.9.0-bindings.2` |
| 2.10.0 | `nifi-2.10.0` | `v2.10.0-bindings.2` |
| 2.11.0 | `nifi-2.11.0` | `v2.11.0-bindings.2` |

For example:

```sh
go get github.com/skycubed/nifi-go/v2@v2.11.0-bindings.2
```

Legacy `*-beta.*` tags remain available but are not rewritten or promoted as
current releases. New binding-only fixes increment the final binding revision,
for example `v2.11.0-bindings.2`.

The minimum supported Go version is 1.25.

## Client usage

The response-aware client parses documented response bodies while retaining the
status, headers, and original bytes:

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/skycubed/nifi-go/v2/pkg/nifi"
)

func main() {
	client, err := nifi.NewClientWithResponses(
		"https://nifi.example.com/nifi-api",
		nifi.WithBearerToken("eyJ..."),
	)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.GetAboutInfoWithResponse(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	about := nifi.JSONResponseBody[nifi.AboutEntity](response)
	if response.StatusCode() != 200 || about == nil {
		log.Fatalf("NiFi returned %s: %s", response.Status(), response.Body)
	}
	fmt.Println(*about.About.Version)
}
```

The server argument must include the API root: `/nifi-api` for NiFi and
`/nifi-registry-api` for Registry. `WithBearerToken` rejects empty tokens,
surrounding whitespace, and header-injection characters.

Mutating endpoints accept generated entity types. NiFi uses optimistic
revisions, so read the current entity, preserve its revision, make the desired
change, and submit that entity to the matching update method.

```go
current, err := client.GetProcessorWithResponse(ctx, processorID)
if err != nil {
	return err
}
currentEntity := nifi.JSONResponseBody[nifi.ProcessorEntity](current)
if currentEntity == nil || currentEntity.Component == nil {
	return fmt.Errorf("read processor failed: %s: %s", current.Status(), current.Body)
}
newName := "Renamed processor"
currentEntity.Component.Name = &newName
updated, err := client.UpdateProcessorWithResponse(ctx, processorID, *currentEntity)
if err != nil {
	return err
}
if nifi.JSONResponseBody[nifi.ProcessorEntity](updated) == nil {
	return fmt.Errorf("update failed: %s: %s", updated.Status(), updated.Body)
}
```

For downloads and other large responses, use the raw `Client`, stream from
`response.Body`, and close it. The `ClientWithResponses` helpers intentionally
read and close response bodies in order to parse them; use their retained
`Body` byte slice when diagnosing a non-2xx response.

```go
client, err := nifi.NewClient("https://nifi.example.com/nifi-api",
	nifi.WithBearerToken("eyJ..."))
if err != nil {
	return err
}
response, err := client.GetAboutInfo(ctx)
if err != nil {
	return err
}
defer response.Body.Close()
```

### Custom trust roots and mTLS

Pass a configured `http.Client` through `WithHTTPClient`. Keep certificate
verification enabled and load the deployment CA explicitly:

```go
roots := x509.NewCertPool()
if !roots.AppendCertsFromPEM(caPEM) {
	return errors.New("invalid NiFi CA certificate")
}
transport := http.DefaultTransport.(*http.Transport).Clone()
transport.TLSClientConfig = &tls.Config{
	MinVersion:   tls.VersionTLS12,
	RootCAs:      roots,
	Certificates: []tls.Certificate{clientCertificate}, // omit when not using mTLS
}
httpClient := &http.Client{Transport: transport}

client, err := nifi.NewClientWithResponses(
	"https://nifi.example.com/nifi-api",
	nifi.WithHTTPClient(httpClient),
	nifi.WithBearerToken(token),
)
```

No helper disables TLS verification or imposes a global timeout. Callers should
use context deadlines appropriate for each operation, especially for large
uploads and downloads.

## NiFi Registry deprecation

The `pkg/registry` client is generated through NiFi Registry 2.11.0 for existing
deployments. Apache deprecated NiFi Registry in February 2026 and plans to
remove it in NiFi 3. New deployments should use a Git-based Flow Registry
Client.

## Reproducible generation

`source.lock.json` records exact Maven Central WAR URLs and SHA-512 digests.
The generation tool verifies each artifact, extracts the upstream document,
preserves it under `openapi/raw`, applies narrowly defined OpenAPI repairs under
`openapi/normalized`, validates it, and invokes pinned `oapi-codegen` 2.8.0.

```sh
make spec              # download, verify, normalize, validate
make generate          # regenerate split model/client files
make verify-generated  # regenerate and require a clean diff
make test
make test-race
make vet
make lint
make vuln
```

Generated `*.gen.go` files must not be edited. Hand-written helpers and tests
live beside them and are never removed by generation. See
[`docs/adr/0001-openapi-generator.md`](docs/adr/0001-openapi-generator.md)
for the generator decision and [`docs/releasing.md`](docs/releasing.md) for the
branch, validation, and immutable-tag process.

The generated method and model API is preserved within a binding-only revision.
Any incompatible public API change requires explicit review and documentation;
the API-diff gate rejects an accidental break when the target NiFi version has
not changed. A new upstream NiFi version may necessarily add or change generated
API, and its diff is reported for review.

## Security and support

See [`SECURITY.md`](SECURITY.md). The bindings secure HTTP requests only as
configured by the caller; using a newer client does not remediate a vulnerable
NiFi server. Run a supported, patched NiFi release.

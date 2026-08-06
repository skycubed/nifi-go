# ADR 0001: Generate clients with oapi-codegen

## Status

Accepted, August 2026.

## Context

Apache NiFi publishes OpenAPI 3.0 documents inside the released NiFi and NiFi
Registry web API WARs. The documents contain two recurring issues: authorization
prose is encoded as undeclared security requirement names, and many responses
omit the required description. Recent NiFi releases also spell the HTTP bearer
scheme with non-canonical casing.

The previous workflow cloned and built the full NiFi source tree, patched Java
annotations, installed the latest generator at build time, and embedded the
specification in a single generated Go file. That process was slow,
non-reproducible, and added parser dependencies to every client.

## Decision

Use pinned `oapi-codegen` v2.8.0 after a deterministic, validated normalization
step. Generate models and client code into separate files and do not embed the
specification. Keep the source and normalized documents for provenance and
review.

`oapi-codegen` successfully generated and compiled the NiFi and Registry 2.1.0
and 2.11.0 documents during evaluation. It retains the existing standard-library
HTTP client shape and response-aware methods, minimizing migration for current
users.

## Alternatives

- **ogen v1.23.0** generates validation and optimized serialization, but it
  rejected NiFi's undeclared security requirements and then rejected binary and
  wildcard content shapes. Supporting every operation would require additional
  semantic repairs and a breaking public API migration.
- **OpenAPI Generator v7.24.0** has broad language support and a mature release
  process, but adds a Java generation toolchain and produces a substantially
  different, larger Go client without a demonstrated compatibility benefit.
- **go-swagger** does not support OpenAPI 3.x and therefore cannot consume the
  released NiFi documents directly.

## Consequences

Generator, normalizer, and runtime upgrades are explicit reviewed changes.
Generated API changes caused by a new NiFi release are isolated on the matching
version branch. Binding-only revisions use API-diff reporting and must not make
unapproved incompatible changes.

# Release process

The exact-version branch and immutable tag are the compatibility contract. A
binding release is prepared and published in separate, reviewable operations.
The automated discovery workflow may open an issue for a new stable Apache NiFi
release, but never publishes code or tags.

## Prepare

1. Verify the Apache release announcement and run **Prepare or publish a binding
   release** with `action=prepare`, the exact release version, and revision `1`.
2. The workflow confirms the matching upstream release tag exists, downloads
   the NiFi and Registry web API WARs from their exact Maven Central coordinates,
   creates `source.lock.json`, verifies and normalizes both specifications, and
   regenerates the clients.
3. It creates the `nifi-X.Y.Z` branch if necessary and opens a draft generated
   pull request against it. Existing exact-version branches are never replaced.
4. Review the raw and normalized spec diff, generated public API report, and all
   required CI checks. Run the real-server smoke job for the exact version.
5. Merge the generated pull request. For the newest upstream release, merge the
   same approved snapshot to `main`, which remains the canonical tooling and
   documentation branch and the latest generated release.

For a binding-only correction, prepare revision `N+1` from the existing
exact-version branch. The target version in the lock does not change, and the
API-diff gate rejects incompatible public API changes unless the project
explicitly approves and documents an exception.

## Publish

Run the same workflow with `action=publish` only after the exact-version branch
has passed required checks and its real-server smoke test. The `release`
environment should require maintainer approval. The workflow regenerates,
tests, runs static analysis and `govulncheck`, then creates an annotated
`vX.Y.Z-bindings.N` tag. It fails if the tag already exists and never moves or
rewrites a tag.

Legacy branches and beta tags are historical records. Do not force-push them,
delete them, or reuse their names.

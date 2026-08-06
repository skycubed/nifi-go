# Security policy

## Supported releases

The current `v2.X.Y-bindings.N` revision on each `nifi-2.X.Y` branch is
supported for binding defects. Legacy beta tags are retained but do not receive
routine fixes. Apache NiFi server vulnerabilities and server lifecycle are
governed by the Apache NiFi project; a binding update cannot remediate the
server itself.

This module supports the two current Go release families, with Go 1.25 as the
minimum. CI uses patched Go toolchains and rejects reachable vulnerabilities
reported by `govulncheck`.

## Reporting a vulnerability

Do not open a public issue for a suspected vulnerability. Use GitHub's private
vulnerability reporting feature for this repository:

https://github.com/skycubed/nifi-go/security/advisories/new

Include the binding tag, target NiFi version, affected operation, reproduction,
and impact. Do not include live credentials, JWTs, private keys, or production
URLs.

## Client security expectations

- Keep TLS certificate verification enabled and provide trusted roots or mTLS
  credentials through a custom `http.Client`.
- Prefer short-lived bearer tokens and contexts with operation-appropriate
  deadlines.
- Do not log authorization headers, token response bodies, or sensitive NiFi
  entity fields.
- Select the exact binding version matching the target NiFi server.

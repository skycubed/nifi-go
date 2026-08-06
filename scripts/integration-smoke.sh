#!/bin/sh
set -eu

version=${1:?usage: integration-smoke.sh NIFI_VERSION}
nifi_container="nifi-bindings-${version}"
registry_container="nifi-registry-bindings-${version}"
username=integration-admin
password='Integration-Only-Password-2026'
work_dir=${RUNNER_TEMP:-/tmp}/nifi-bindings-${version}
certificate=${work_dir}/nifi-ca.pem

mkdir -p "${work_dir}"

cleanup() {
    docker rm -f "${nifi_container}" "${registry_container}" >/dev/null 2>&1 || true
}
trap cleanup EXIT INT TERM
cleanup

docker run --detach --name "${nifi_container}" --hostname localhost \
    --publish 8443:8443 \
    --env NIFI_WEB_HTTPS_HOST=0.0.0.0 \
    --env NIFI_WEB_PROXY_HOST=localhost:8443 \
    --env SINGLE_USER_CREDENTIALS_USERNAME="${username}" \
    --env SINGLE_USER_CREDENTIALS_PASSWORD="${password}" \
    "apache/nifi:${version}" >/dev/null

docker run --detach --name "${registry_container}" --hostname localhost \
    --publish 18080:18080 \
    --env NIFI_REGISTRY_WEB_HTTP_HOST=0.0.0.0 \
    "apache/nifi-registry:${version}" >/dev/null

attempt=0
while [ "${attempt}" -lt 90 ]; do
    if openssl s_client -connect localhost:8443 -servername localhost -showcerts </dev/null 2>/dev/null \
        | openssl x509 -outform PEM >"${certificate}" 2>/dev/null \
        && curl --fail --silent --show-error --cacert "${certificate}" \
            https://localhost:8443/nifi-api/authentication/configuration >/dev/null 2>&1 \
        && curl --fail --silent --show-error \
            http://localhost:18080/nifi-registry-api/about >/dev/null 2>&1; then
        break
    fi
    attempt=$((attempt + 1))
    sleep 10
done

if [ "${attempt}" -eq 90 ]; then
    docker logs "${nifi_container}" || true
    docker logs "${registry_container}" || true
    echo "NiFi services did not become ready" >&2
    exit 1
fi

token=$(curl --fail --silent --show-error --cacert "${certificate}" \
    --header 'Content-Type: application/x-www-form-urlencoded' \
    --data-urlencode "username=${username}" \
    --data-urlencode "password=${password}" \
    https://localhost:8443/nifi-api/access/token)

NIFI_BASE_URL=https://localhost:8443/nifi-api \
NIFI_TOKEN="${token}" \
REGISTRY_BASE_URL=http://localhost:18080/nifi-registry-api \
NIFI_CA_FILE="${certificate}" \
go test -tags=integration -count=1 ./integration/...

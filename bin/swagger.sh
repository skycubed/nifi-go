#!/bin/bash -x

set -euo pipefail

paths201=(
  "/access/kerberos"
  "/access/token"
  "/controller/controller-services"
  "/controller/registry-clients"
  "/controller/reporting-tasks"
  "/parameter-contexts"
  "/policies"
  "/process-groups/{id}/connections"
  "/process-groups/{id}/controller-services"
  "/process-groups/{id}/funnels"
  "/process-groups/{id}/input-ports"
  "/process-groups/{id}/labels"
  "/process-groups/{id}/output-ports"
  "/process-groups/{id}/process-groups"
  "/process-groups/{id}/process-groups/import"
  "/process-groups/{id}/processors"
  "/process-groups/{id}/remote-process-groups"
  "/process-groups/{id}/snippet-instance"
  "/process-groups/{id}/template-instance"
  "/process-groups/{id}/templates"
  "/process-groups/{id}/templates/import"
  "/provenance"
  "/provenance-events/replays"
  "/provenance/lineage"
  "/snippets"
  "/tenants/user-groups"
  "/tenants/users"
)

declare -A modelFieldMap
modelFieldMap=(
  [RevisionDTO.1]="version"
)

sf=$1
for path in "${paths201[@]}"
do
  jq --arg p "$path" '(.paths[$p].post.responses |= (. + {"201":."200"}|del(."200")))' "${sf}" > "${sf}".tmp
  mv "${sf}".tmp "${sf}"
done

for model in "${!modelFieldMap[@]}"; do
  m=${model%.*}
  f=${modelFieldMap[$model]}
  jq --arg m "$m" --arg f "$f" '(.definitions[$m].properties[$f] |= (. + {"x-omitempty":false}))' "${sf}" > "${sf}".tmp
  mv "${sf}".tmp "${sf}"
done

jq '. += {"securityDefinitions" : {
              "Authorization" : {
                "description" : "NiFi Auth Token (JWT)",
                "type" : "apiKey",
                "name" : "Authorization",
                "in" : "header"
              },
              "BasicAuth" : {
                "description" : "HTTP Basic Auth",
                "type" : "basic"
              }
          }}' "${sf}" > "${sf}".tmp

mv "${sf}".tmp "${sf}"

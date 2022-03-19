#!/bin/bash

set -euo pipefail

files=(
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

sf=$1
for path in "${files[@]}"
do
  jq --arg p "$path" '(.paths[$p].post.responses |= (. + {"201":."200"}|del(."200")))' "${sf}" > "${sf}".tmp
  mv  "${sf}".tmp "${sf}"
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

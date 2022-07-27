#!/bin/bash -x

set -euo pipefail

declare -A modelFieldMap
modelFieldMap=(
  [RevisionInfo.1]="version"
)

sf=$1
for model in "${!modelFieldMap[@]}"; do
  m=${model%.*}
  f=${modelFieldMap[$model]}
  jq --arg m "$m" --arg f "$f" '(.definitions[$m].properties[$f] |= (. + {"x-omitempty":false}))' "${sf}" > "${sf}".tmp
  mv "${sf}".tmp "${sf}"
done


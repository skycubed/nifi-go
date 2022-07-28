#!/bin/bash -x

set -euo pipefail

declare -A modelFieldMap
modelFieldMap=(
  [RevisionInfo.1]="version"
  [VersionedProcessor.1]="runDurationMillis"
  [VersionedProcessor.2]="annotationData"
  [VersionedProcessor.3]="style"
  [VersionedProcessor.4]="comments"
)

sf=$1
for model in "${!modelFieldMap[@]}"; do
  m=${model%.*}
  f=${modelFieldMap[$model]}
  jq --arg m "$m" --arg f "$f" '(.definitions[$m].properties[$f] |= (. + {"x-omitempty":false}))' "${sf}" > "${sf}".tmp
  mv "${sf}".tmp "${sf}"
done

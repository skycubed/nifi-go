#!/bin/bash -x

set -euo pipefail

declare -A modelFieldMapOmitEmptyFalse
modelFieldMapOmitEmptyFalse=(
  [RevisionDTO.1]="version"
  [VersionedProcessor.1]="runDurationMillis"
  [VersionedProcessor.2]="annotationData"
  [VersionedProcessor.3]="style"
  [VersionedProcessor.4]="comments"
)

declare -A modelFieldMapOmitEmptyTrue
modelFieldMapOmitEmptyTrue=(
  [VersionedProcessor.1]="properties"
)

sf=$1
for model in "${!modelFieldMapOmitEmptyFalse[@]}"; do
  m=${model%.*}
  f=${modelFieldMapOmitEmptyFalse[$model]}
  jq --arg m "$m" --arg f "$f" '(.definitions[$m].properties[$f] |= (. + {"x-omitempty":false}))' "${sf}" > "${sf}".tmp
  mv "${sf}".tmp "${sf}"
done

for model in "${!modelFieldMapOmitEmptyTrue[@]}"; do
  m=${model%.*}
  f=${modelFieldMapOmitEmptyTrue[$model]}
  jq --arg m "$m" --arg f "$f" '(.definitions[$m].properties[$f] |= (. + {"x-omitempty":true}))' "${sf}" > "${sf}".tmp
  mv "${sf}".tmp "${sf}"
done

SHELL := /bin/sh

.DEFAULT_GOAL := test

.PHONY: spec
spec:
	cd tools && go run ./cmd/specsync -lock ../source.lock.json -cache ../.cache/artifacts -output ../openapi

.PHONY: generate
generate: spec
	cd tools && go run ./cmd/versiongen -lock ../source.lock.json -root ..
	cd tools && go tool oapi-codegen --config ../config/oapi/nifi-types.yaml ../openapi/normalized/nifi.json
	cd tools && go tool oapi-codegen --config ../config/oapi/nifi-client.yaml ../openapi/normalized/nifi.json
	cd tools && go tool oapi-codegen --config ../config/oapi/registry-types.yaml ../openapi/normalized/registry.json
	cd tools && go tool oapi-codegen --config ../config/oapi/registry-client.yaml ../openapi/normalized/registry.json

.PHONY: verify-generated
verify-generated: generate
	git diff --exit-code -- openapi pkg/nifi/types.gen.go pkg/nifi/client.gen.go pkg/nifi/version.gen.go pkg/registry/types.gen.go pkg/registry/client.gen.go pkg/registry/version.gen.go

.PHONY: test
test:
	go test ./...
	cd tools && go test ./...

.PHONY: test-race
test-race:
	go test -race ./...

.PHONY: vet
vet:
	go vet ./...
	cd tools && go vet ./...

.PHONY: tools-bin
tools-bin:
	mkdir -p .cache/bin
	cd tools && go build -o ../.cache/bin/staticcheck honnef.co/go/tools/cmd/staticcheck
	cd tools && go build -o ../.cache/bin/govulncheck golang.org/x/vuln/cmd/govulncheck
	cd tools && go build -o ../.cache/bin/apidiff golang.org/x/exp/cmd/apidiff

.PHONY: lint
lint: tools-bin
	.cache/bin/staticcheck ./...
	cd tools && ../.cache/bin/staticcheck ./...

.PHONY: vuln
vuln: tools-bin
	.cache/bin/govulncheck ./...

.PHONY: integration
integration:
	go test -tags=integration ./integration/...

.PHONY: clean
clean:
	rm -rf target

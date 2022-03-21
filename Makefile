TRG := target
NIFI_VER := 1.15.3

.PHONY: swagger
swagger:
	-docker rm -f swagger
	mkdir -p $(TRG)/swagger/{nifi,registry}
	docker build -f Dockerfile.swagger --build-arg NIFI_VER=$(NIFI_VER) -t swagger .
	docker create -ti --name swagger swagger bash
	docker cp swagger:/nifi/swagger.json     $(TRG)/swagger/nifi
	docker cp swagger:/registry/swagger.json $(TRG)/swagger/registry
	docker rm -f swagger
	rm -rf swagger
	mv $(TRG)/swagger swagger

.PHONY: generate
generate:
	-docker rm -f generate
	docker build -f Dockerfile.generate -t generate .
	docker create -ti --name generate generate sh
	-rm -rf pkg/nifi pkg/registry
	docker cp generate:/go/src/github.com/skycubed/nifi-go/pkg/nifi pkg
	docker cp generate:/go/src/github.com/skycubed/nifi-go/pkg/registry pkg
	docker rm -f generate

clean:
	rm -rf $(TRG)
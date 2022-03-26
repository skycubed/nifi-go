NIFI_VER := 1.15.2

.PHONY: generate
generate: swagger
	-docker rm -f generate
	docker build -f Dockerfile.generate -t generate .
	docker create -ti --name generate generate sh
	-rm -rf pkg/{nifi,registry}
	docker cp generate:/go/src/github.com/skycubed/nifi-go/pkg/nifi     pkg
	docker cp generate:/go/src/github.com/skycubed/nifi-go/pkg/registry pkg
	docker rm -f generate

.PHONY: swagger
swagger:
	-docker rm -f swagger
	mkdir -p target/swagger/{nifi,registry}
	docker build -f Dockerfile.swagger --build-arg NIFI_VER=$(NIFI_VER) -t swagger .
	docker create -ti --name swagger swagger bash
	docker cp swagger:/nifi/swagger.json     target/swagger/nifi
	docker cp swagger:/registry/swagger.json target/swagger/registry
	docker rm -f swagger
	rm -rf swagger
	mv target/swagger swagger

clean:
	rm -rf target
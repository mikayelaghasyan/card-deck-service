TARGET_DIR=target
BINARY_NAME=card-deck-service

build:
	go build -o ${TARGET_DIR}/${BINARY_NAME} cmd/main/card-deck-service.go

run: build
	./${TARGET_DIR}/${BINARY_NAME}

unit-test: build
	go test ./pkg/...

integration-test: build
	go test ./cmd/integration-tests/...

test: test-unit test-integration

clean:
	go clean
	rm -rf ${TARGET_DIR}

generate:
	oapi-codegen -package api -generate types pkg/api/openapi.yml > pkg/api/types.gen.go
	oapi-codegen -package api -generate server pkg/api/openapi.yml > pkg/api/server.gen.go

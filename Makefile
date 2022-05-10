TARGET_DIR=target
BINARY_NAME=card-deck-service

build:
	go build -o ${TARGET_DIR}/${BINARY_NAME} cmd/main/card-deck-service.go

run: build
	./${TARGET_DIR}/${BINARY_NAME}

test-ig: build
	go test ./cmd/integration-tests/...

clean:
	go clean
	rm -rf ${TARGET_DIR}

generate:
	oapi-codegen -package api -generate types pkg/api/openapi.yml > pkg/api/types.gen.go
	oapi-codegen -package api -generate server pkg/api/openapi.yml > pkg/api/server.gen.go

TARGET_DIR=target
BINARY_NAME=card-deck-service

build:
	go build -o ${TARGET_DIR}/${BINARY_NAME} cmd/card-deck-service.go

run: build
	./${TARGET_DIR}/${BINARY_NAME}

clean:
	go clean
	rm -rf ${TARGET_DIR}

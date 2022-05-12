# Card Deck service

## Table of Contents
- [General](#general)
- [Implementation](#implementation)
- [Makefile](#makefile)
- [Exposed Endpoints](#exposed-endpoints)
- [Not Implemented](#not-implemented)

## General
Card deck service is a demo service for managing playing card decks.
The service is provided as a set of REST API endpoints.

## Implementation
- The service is implemented in Go language.
- The API is first described in OpenAPI format, and then the code for data types and server interface is generated from that OpenAPI definition with [oapi-codegen](https://github.com/deepmap/oapi-codegen) library.
- The service was developed with TDD approach, both on integration and unit levels. (You can navigate through commits to see the progress of TDD cycles.)
- The data is stored in-memory, which means it will not be persisted over app restarts. However, the architecture allows to easily add other implementations as well (e.g. database), by just implementing Repository interface.

## Makefile
Makefile is provided with a few targets:
- `make build` - builds the project and places the executable in the `./target` directory
- `make run` - runs the app
- `make unit-test` - runs all unit tests
- `make integration-test` - runs all integration tests
- `make test` - runs both unit and integration tests
- `make clean` - cleans the project
- `make generate` - generates data types and server interface from the OpenAPI definition

## Exposed Endpoints
The server runs on port `1323` and exposes the following endpoints:
- `/swagger/index.html` - Swagger UI with the API definitions and tools for sending requests to the endpoints
- `/api` - API endpoints base path

## Not Implemented
There are few things which were initially planned to be implemented, but actually not implemented:
- SQL DB repository
- configuration
- dockerization
- e2e (end-to-end) testing, i.e. testing directly REST API endpoints with HTTP requests

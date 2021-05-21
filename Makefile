.PHONY: install
install: # Install on system for development
	go install -ldflags="-X main.baseURL=https://vax-availability-api-staging.azurewebsites.net"

.PHONY: build
build: # Build on system for development
	go build --ldflags="-X main.baseURL=https://vax-availability-api-staging.azurewebsites.net"

.PHONY: fmt
fmt: # Run all formatting
	go fmt ./...

.PHONY: test
test: # Run all tests
	go test ./...

.PHONY: api-codegen
api-codegen:
	oapi-codegen -config=./.oapi-codegen.yaml https://vax-availability-api.azurewebsites.net/openapi.json

.PHONY: docs
docs: # Generate documentation at docs/
	go run ./scripts/gen-docs/main.go

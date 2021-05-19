.PHONY: install
install: # Install on system
	go install

.PHONY: build
build: # Build code
	go build

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

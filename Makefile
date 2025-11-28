.PHONY: generate generate.proto build run

PODMAN ?= podman
GIT_COMMIT=$(shell git rev-list -1 HEAD --abbrev-commit)

BINARY_NAME=agent
BINARY_PATH=bin/$(BINARY_NAME)
MAIN_PATH=./main.go

generate:
	@echo "Generating code..."
	go generate ./...
	@echo "Code generation complete."

# Generate protobuf code using buf in container
generate.proto:
	@echo "Generating protobuf code with buf in container..."
	$(PODMAN) run --rm \
		-v $(CURDIR)/api/v2/:/workspace \
		-w /workspace \
		bufbuild/buf:latest \
		generate
	@echo "Protobuf generation complete."

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	go build -ldflags="-X main.sha=${GIT_COMMIT}" -o $(BINARY_PATH) $(MAIN_PATH)
	@echo "Build complete: $(BINARY_PATH)"

run:
	$(BINARY_PATH) run

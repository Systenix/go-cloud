# Variables
CLI_NAME := go-cloud-cli
CLI_DIR := cmd
BIN_DIR := bin
PKG_LIST := $(shell go list ./... | grep -v /vendor/)

# Default target
all: build

# Build the CLI
build:
	@echo "Building the CLI..."
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(CLI_NAME) $(CLI_DIR)/main.go
	@echo "CLI built successfully."

# Run tests
test:
	@echo "Running tests..."
	go test -v $(PKG_LIST)
	@echo "Tests completed."

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf $(BIN_DIR)
	rm -f $(BIN_DIR)/$(CLI_NAME)
	@echo "Cleanup completed."

# Run linter
lint:
	@echo "Running linter..."
	golangci-lint run ./...
	@echo "Linting completed."

# Format code
fmt:
	@echo "Formatting code..."
	go fmt $(PKG_LIST)
	@echo "Formatting completed."

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy
	@echo "Dependencies installed."

# Help
help:
	@echo "Available targets:"
	@echo "  make build     - Build the CLI application"
	@echo "  make test      - Run tests"
	@echo "  make clean     - Clean build artifacts"
	@echo "  make lint      - Run linter"
	@echo "  make fmt       - Format code"
	@echo "  make deps      - Install dependencies"
	@echo "  make all       - Build the CLI (default target)"

.PHONY: all build test clean lint fmt deps help
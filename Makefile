# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOLINT=golangci-lint
GOFMT=gofmt
BINARY_NAME=app
BINARY_UNIX=$(BINARY_NAME)_unix

# Build info
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GIT_COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Linker flags
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT)"

.PHONY: all build clean test coverage deps fmt lint vet staticcheck gosec pre-commit help

## help: Show this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## all: Run all checks and build
all: deps fmt lint vet staticcheck gosec test coverage build

## build: Build the binary
build:
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) -v ./cmd/app

## build-linux: Build for Linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BINARY_UNIX) -v ./cmd/app

## clean: Clean build artifacts
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -f coverage.out
	rm -f coverage.html

## test: Run tests
test:
	$(GOTEST) -v ./...

## test-race: Run tests with race detector
test-race:
	$(GOTEST) -race -v ./...

## coverage: Run tests with coverage
coverage:
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

## deps: Download dependencies
deps:
	$(GOMOD) download
	$(GOMOD) tidy

## fmt: Format Go code
fmt:
	$(GOFMT) -s -w .

## lint: Run golangci-lint
lint:
	$(GOLINT) run

## vet: Run go vet
vet:
	$(GOCMD) vet ./...

## staticcheck: Run staticcheck
staticcheck:
	staticcheck ./...

## gosec: Run gosec security scanner
gosec:
	gosec ./...

## pre-commit: Run all pre-commit checks (for Go Report Card A+)
pre-commit: fmt lint vet staticcheck gosec test
	@echo "All pre-commit checks passed! ✅"

## install-tools: Install development tools
install-tools:
	@echo "Installing development tools..."
	$(GOCMD) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GOCMD) install honnef.co/go/tools/cmd/staticcheck@latest
	$(GOCMD) install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest
	@echo "Tools installed successfully! ✅"

## docker-build: Build Docker image
docker-build:
	docker build -t $(BINARY_NAME) .

## run: Run the application
run: build
	./$(BINARY_NAME)

## dev: Run in development mode with auto-reload (requires air)
dev:
	air

## install-air: Install air for development auto-reload
install-air:
	$(GOCMD) install github.com/cosmtrek/air@latest
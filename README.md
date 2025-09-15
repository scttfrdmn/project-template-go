# Go Project Template

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/scttfrdmn/project-name)](https://goreportcard.com/report/github.com/scttfrdmn/project-name)
[![Go Version](https://img.shields.io/github/go-mod/go-version/scttfrdmn/project-name)](https://golang.org/)
[![Build Status](https://img.shields.io/github/actions/workflow/status/scttfrdmn/project-name/ci.yml?branch=main)](https://github.com/scttfrdmn/project-name/actions)
[![Coverage Status](https://coveralls.io/repos/github/scttfrdmn/project-name/badge.svg?branch=main)](https://coveralls.io/github/scttfrdmn/project-name?branch=main)
[![GoDoc](https://godoc.org/github.com/scttfrdmn/project-name?status.svg)](https://godoc.org/github.com/scttfrdmn/project-name)
[![Docker Pulls](https://img.shields.io/docker/pulls/scttfrdmn/project-name)](https://hub.docker.com/r/scttfrdmn/project-name)
[![Release](https://img.shields.io/github/release/scttfrdmn/project-name.svg)](https://github.com/scttfrdmn/project-name/releases/latest)
[![GitHub issues](https://img.shields.io/github/issues/scttfrdmn/project-template-go)](https://github.com/scttfrdmn/project-template-go/issues)
[![GitHub forks](https://img.shields.io/github/forks/scttfrdmn/project-template-go)](https://github.com/scttfrdmn/project-template-go/network)
[![GitHub stars](https://img.shields.io/github/stars/scttfrdmn/project-template-go)](https://github.com/scttfrdmn/project-template-go/stargazers)
[![GitHub contributors](https://img.shields.io/github/contributors/scttfrdmn/project-template-go)](https://github.com/scttfrdmn/project-template-go/graphs/contributors)

A comprehensive Go project template following best practices for September 2025, with pre-commit hooks configured to maintain Go Report Card A+ grade.

## Features

- **Go 1.23** with modern project structure (`cmd/`, `internal/`, `pkg/`)
- **Pre-commit hooks** configured for Go Report Card A+ grade
- **Comprehensive tooling**: golangci-lint, staticcheck, gosec, go vet
- **Docker support** with multi-stage builds and security best practices
- **Makefile** with common Go development tasks
- **Air integration** for hot reloading during development
- **GitHub workflows** for CI/CD (coming soon)
- **Comprehensive documentation** (CONTRIBUTING, CODE_OF_CONDUCT, SECURITY)
- **GitHub templates** for issues and pull requests
- **Dependabot configuration** for automated dependency updates
- **EditorConfig** for consistent code formatting
- **MIT License** template
- **Changelog** following Keep a Changelog format
- **Ko-fi funding** integration

## Installation

```bash
# Clone the repository
git clone git@github.com:scttfrdmn/project-template-go.git
cd project-template-go

# Install development tools and pre-commit hooks
./scripts/install-hooks.sh

# Update go.mod with your project name
sed -i '' 's|github.com/scttfrdmn/project-name|github.com/yourusername/yourproject|g' go.mod

# Update import paths in Go files
find . -name '*.go' -exec sed -i '' 's|github.com/scttfrdmn/project-name|github.com/yourusername/yourproject|g' {} +

# Install dependencies
make deps

# Verify everything works
make all
```

## Quick Start

```bash
# Build and run the application
make build
./app

# Or run in development mode with hot reloading
make dev

# Run all quality checks (for Go Report Card A+)
make pre-commit

# Run tests with coverage
make coverage

# Build Docker image
make docker-build
```

## Development

### Prerequisites

- Go 1.23 or higher
- Make
- Docker (optional)
- pre-commit (for git hooks)

### Available Make Targets

```bash
make help           # Show all available targets
make build          # Build the application
make test           # Run tests
make coverage       # Run tests with coverage report
make lint           # Run golangci-lint
make fmt            # Format code
make pre-commit     # Run all pre-commit checks
make install-tools  # Install development tools
```

## What's Included

```
project-template-go/
├── .github/
│   ├── dependabot.yml          # Automated dependency updates
│   ├── FUNDING.yml             # Ko-fi sponsorship integration
│   ├── ISSUE_TEMPLATE.md       # Bug report template
│   └── PULL_REQUEST_TEMPLATE.md # PR checklist template
├── cmd/
│   └── app/
│       └── main.go             # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go           # Configuration management
│   ├── handler/
│   │   └── handler.go          # HTTP handlers
│   └── service/
│       └── service.go          # Business logic
├── pkg/                        # Public packages (optional)
├── scripts/
│   └── install-hooks.sh        # Setup script for pre-commit hooks
├── docs/
│   ├── CODE_OF_CONDUCT.md      # Community guidelines
│   ├── CONTRIBUTING.md         # Contribution guidelines
│   └── SECURITY.md             # Security policy
├── tests/
│   ├── unit/                   # Unit tests
│   ├── integration/            # Integration tests
│   ├── e2e/                    # End-to-end tests
│   ├── fixtures/               # Test data
│   └── README.md               # Testing guidelines
├── .air.toml                   # Air configuration for hot reloading
├── .editorconfig               # Editor configuration
├── .gitignore                  # Git ignore rules
├── .golangci.yml               # golangci-lint configuration
├── .pre-commit-config.yaml     # Pre-commit hooks configuration
├── CHANGELOG.md                # Version history
├── Dockerfile                  # Multi-stage Docker build
├── LICENSE                     # MIT License
├── Makefile                    # Build automation
├── go.mod                      # Go module definition
└── README.md                   # This file
```

## Go Report Card A+ Grade

This template is configured to maintain a Go Report Card A+ grade through:

- **Pre-commit hooks** that run automatically before each commit
- **golangci-lint** with comprehensive linting rules
- **staticcheck** for advanced static analysis
- **gosec** for security vulnerability scanning
- **go vet** for suspicious constructs
- **Automatic code formatting** with `gofmt`
- **Test execution** to ensure functionality

### Setting Up Pre-commit Hooks

```bash
# Run the setup script (installs tools and hooks)
./scripts/install-hooks.sh

# Or manually install pre-commit and run
pre-commit install
```

### Manual Quality Checks

```bash
# Run all quality checks
make pre-commit

# Run individual checks
make fmt          # Format code
make lint         # Run linter
make vet          # Run go vet
make staticcheck  # Run static analysis
make gosec        # Run security check
make test         # Run tests
```

## Contributing

See [CONTRIBUTING.md](docs/CONTRIBUTING.md) for detailed guidelines.

## Versioning

This project uses [Semantic Versioning](https://semver.org/). See [CHANGELOG.md](CHANGELOG.md) for details.

## License

This project is licensed under the [MIT License](LICENSE) - see the LICENSE file for details.

## API Endpoints

The template includes a basic HTTP server with the following endpoints:

- `GET /health` - Health check endpoint
- `GET /info` - Service information
- `POST /api/process` - Example data processing endpoint

## Environment Variables

- `PORT` - Server port (default: 8080)
- `HOST` - Server host (default: 0.0.0.0)
- `ENVIRONMENT` - Environment name (default: development)
- `LOG_LEVEL` - Log level (default: info)
- `DEBUG` - Enable debug mode (default: false)

## Acknowledgments

- Go community for excellent tooling
- golangci-lint, staticcheck, and gosec maintainers
- pre-commit framework contributors
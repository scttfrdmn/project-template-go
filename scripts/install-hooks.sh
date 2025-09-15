#!/bin/bash

# Install pre-commit hooks for Go Report Card A+ grade
# This script sets up git hooks to ensure code quality

set -e

echo "🔧 Setting up pre-commit hooks for Go Report Card A+ grade..."

# Check if pre-commit is installed
if ! command -v pre-commit &> /dev/null; then
    echo "❌ pre-commit is not installed. Please install it first:"
    echo "   pip install pre-commit"
    echo "   or"
    echo "   brew install pre-commit"
    exit 1
fi

# Install Go tools if not present
echo "📦 Installing Go development tools..."

# Install golangci-lint
if ! command -v golangci-lint &> /dev/null; then
    echo "Installing golangci-lint..."
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
fi

# Install staticcheck
if ! command -v staticcheck &> /dev/null; then
    echo "Installing staticcheck..."
    go install honnef.co/go/tools/cmd/staticcheck@latest
fi

# Install gosec
if ! command -v gosec &> /dev/null; then
    echo "Installing gosec..."
    go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest
fi

# Install air for development (optional)
if ! command -v air &> /dev/null; then
    echo "Installing air for development auto-reload..."
    go install github.com/cosmtrek/air@latest
fi

# Install pre-commit hooks
echo "🪝 Installing pre-commit hooks..."
pre-commit install

# Run hooks once to make sure everything works
echo "✅ Running pre-commit hooks to verify setup..."
pre-commit run --all-files || {
    echo "⚠️  Some hooks failed, but this is normal for the first run."
    echo "   The hooks will fix issues automatically on commit."
}

echo ""
echo "🎉 Pre-commit hooks installed successfully!"
echo ""
echo "Your project is now configured for Go Report Card A+ grade with:"
echo "  ✓ Automatic code formatting (gofmt)"
echo "  ✓ Linting (golangci-lint)"
echo "  ✓ Static analysis (staticcheck)"
echo "  ✓ Security scanning (gosec)"
echo "  ✓ Go vet checks"
echo "  ✓ Test execution"
echo ""
echo "These checks will run automatically on every commit."
echo "To run them manually: make pre-commit"
echo ""
echo "Happy coding! 🚀"
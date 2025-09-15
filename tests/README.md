# Tests

This directory contains test files for the project.

## Structure

```
tests/
├── unit/          # Unit tests
├── integration/   # Integration tests
├── e2e/           # End-to-end tests
└── fixtures/      # Test data and fixtures
```

## Running Tests

```bash
# Run all tests
npm test

# Run specific test types
npm run test:unit
npm run test:integration
npm run test:e2e

# Run tests with coverage
npm run test:coverage

# Watch mode for development
npm run test:watch
```

## Writing Tests

- Place unit tests close to the code they test
- Use descriptive test names that explain what is being tested
- Follow the Arrange-Act-Assert pattern
- Mock external dependencies in unit tests
- Use fixtures for consistent test data

## Coverage

Aim for high test coverage, but focus on testing critical paths and edge cases rather than achieving 100% coverage for its own sake.
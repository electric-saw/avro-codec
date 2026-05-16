# Contributing to avro-codec

Thank you for your interest in contributing to avro-codec! This document provides guidelines and information for contributors.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Making Changes](#making-changes)
- [Pull Request Process](#pull-request-process)
- [Coding Standards](#coding-standards)
- [Testing](#testing)
- [Commit Messages](#commit-messages)

## Code of Conduct

This project adheres to the [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## Getting Started

1. **Fork the repository** on GitHub
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/<your-username>/avro-codec.git
   cd avro-codec
   ```
3. **Add the upstream remote**:
   ```bash
   git remote add upstream https://github.com/Mario-Valente/avro-codec.git
   ```

## Development Setup

### Prerequisites

- Go 1.24 or later
- Make
- [golangci-lint](https://golangci-lint.run/welcome/install/) v2.6.2 or later

### Setup

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Verify your setup by running tests:
   ```bash
   make test
   ```

3. Run the linter:
   ```bash
   make lint
   ```

## Making Changes

1. **Create a new branch** from `main`:
   ```bash
   git checkout -b feature/your-feature-name
   # or
   git checkout -b fix/your-bug-fix
   ```

2. **Make your changes** following our [coding standards](#coding-standards)

3. **Write or update tests** as needed

4. **Run the full CI locally**:
   ```bash
   make ci
   ```

5. **Commit your changes** following our [commit message guidelines](#commit-messages)

## Pull Request Process

1. **Update documentation** if you're changing functionality
2. **Ensure all tests pass** and linting is clean
3. **Fill out the PR template** completely
4. **Link any related issues** using keywords like `Fixes #123`
5. **Request review** from maintainers
6. **Address review feedback** promptly

### PR Checklist

Before submitting your PR, ensure:

- [ ] Tests pass locally (`make test`)
- [ ] Linter passes (`make lint`)
- [ ] New code has appropriate test coverage
- [ ] Documentation is updated if needed
- [ ] Commit messages follow conventions
- [ ] PR description explains the changes clearly

## Coding Standards

This project uses `golangci-lint` to enforce coding standards. The configuration is in `.golangci.yml`.

### General Guidelines

- Follow standard Go conventions and idioms
- Keep functions focused and reasonably sized
- Use meaningful variable and function names
- Add comments for non-obvious logic
- Handle errors explicitly
- Avoid global state when possible

### Code Formatting

Format your code before committing:

```bash
make fmt
```

## Testing

### Running Tests

```bash
# Run all tests
make test

# Run tests with race detector
go test -race ./...

# Run tests for a specific package
go test ./registry/...

# Run a specific test
go test -run TestSchemaParser ./...
```

### Writing Tests

- Place tests in `*_test.go` files alongside the code
- Use table-driven tests where appropriate
- Use `github.com/stretchr/testify` for assertions
- Aim for meaningful coverage, not just high percentages
- Include both positive and negative test cases

### Test Data

- Place test fixtures in the `testdata/` directory
- Use descriptive file names for test schemas

## Commit Messages

We follow [Conventional Commits](https://www.conventionalcommits.org/) specification.

### Format

```
<type>(<scope>): <description>

[optional body]

[optional footer(s)]
```

### Types

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation only changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code change that neither fixes a bug nor adds a feature
- `perf`: Performance improvement
- `test`: Adding or updating tests
- `chore`: Maintenance tasks, dependencies, etc.

### Examples

```
feat(schema): add support for logical type duration

fix(decoder): handle nil pointer in union decoding

docs: update README with new API examples

chore(deps): bump golangci-lint to v2.6.2
```

### Breaking Changes

For breaking changes, add `BREAKING CHANGE:` in the footer:

```
feat(api)!: change Marshal signature to return error

BREAKING CHANGE: Marshal now returns ([]byte, error) instead of []byte
```

## Questions?

If you have questions, feel free to:

- Open a [Discussion](https://github.com/Mario-Valente/avro-codec/discussions)
- Check existing [Issues](https://github.com/Mario-Valente/avro-codec/issues)

Thank you for contributing!

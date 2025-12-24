# Go Essential Tools

## Overview

Essential Go tools for development and code quality.

| Tool | Purpose |
|------|---------|
| `go fmt` | Code formatting |
| `go vet` | Static analysis |
| `go doc` | Documentation |

## go fmt

Automatic code formatting tool that enforces consistent code style across your project.

### Features
- Automatic code formatting
- Enforces consistent code style
- Formats code automatically
- Built into Go toolchain

### Usage

```bash
# Format single file
go fmt main.go

# Format entire package
go fmt ./...

# Check formatting without changes
gofmt -d .
```

## go vet

Static analysis tool that detects suspicious constructs and potential bugs.

### Features
- Detects suspicious constructs
- Finds potential bugs
- Runs before compilation
- Built into Go toolchain

### Usage

```bash
# Analyze current package
go vet

# Analyze specific file
go vet main.go

# Analyze all packages
go vet ./...
```

## go doc

Documentation viewer for packages, functions, and types.

### Features
- View package documentation
- Search functions and types
- Quick command-line access
- Built into Go toolchain

### Usage

```bash
# Package documentation
go doc fmt

# Function documentation
go doc fmt.Printf

# Current package
go doc
```

### Example

```bash
$ go doc fmt
package fmt // import "fmt"

func Printf(format string, a ...interface{}) (n int, err error)
    Printf formats according to a format specifier and writes to standard output.
```

## Best Practices

### Automate
Use in CI/CD pipelines for consistent code quality.

### Regular Usage
Make these tools part of your daily development workflow.

### Pre-commit Checks
Run tools before committing changes.

### Documentation
Write clear comments for `go doc` extraction.

### Development Workflow

```
Code → go fmt → go vet → go doc → Commit
```

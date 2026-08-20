# Hash Calculator

![CI](https://github.com/Qyroxen/Hash-Calculator/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Hash-Calculator/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Hash-Calculator?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Hash-Calculator)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Hash-Calculator)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Hash-Calculator?style=social)](https://github.com/Qyroxen/Hash-Calculator/stargazers)

## What is it?

Hash Calculator is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Hash-Calculator.git
cd Hash-Calculator
go build -o hashcalculator .

# Run
./hashcalculator --help
```

## CLI Usage

```bash
# Basic usage
./hashcalculator

# With flags
./hashcalculator --verbose --output json

# Get help
./hashcalculator --help
```

## Examples

```bash
# Example 1
./hashcalculator example1

# Example 2
./hashcalculator example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o hashcalculator .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Hash-Calculator/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Hash-Calculator?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Hash-Calculator/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Hash-Calculator?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Hash-Calculator/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Hash-Calculator" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Hash-Calculator/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Hash-Calculator" alt="Pull Requests">
  </a>
</p>

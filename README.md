# 🔐 Hash Calculator

![Go](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=flat-square&logo=go&logoColor=white)
![Version](https://img.shields.io/badge/Version-v3.0.0-00ADD8?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)
![PRs](https://img.shields.io/badge/PRs-Welcome-brightgreen?style=flat-square)

> Security tool by [AetherCodeHQ](https://github.com/AetherCodeHQ)

`security` `cryptography` `cli` `golang` `crypto`

---

## What is Hash-Calculator?

**Hash-Calculator** is a security-focused tool that analyzes and validates code, configurations, or data for vulnerabilities.

## Features

- ✅ Cryptographic operations
- 🚀 **Zero dependencies** — only Go standard library
- 📦 **Single binary** — compile and run anywhere
- 🔄 **Offline capable** — no internet required

## Installation

```bash
# Clone
git clone https://github.com/AetherCodeHQ/Hash-Calculator.git
cd Hash-Calculator

# Build
go build -o hash-calculator .

# Run
./hash-calculator <file> [md5|sha256|sha512]
```

### Or directly with `go run`:
```bash
go run main.go <file> [md5|sha256|sha512]
```

## Usage

```bash
# Basic usage
./hash-calculator <file> [md5|sha256|sha512]

# With flags
./hash-calculator <file> [md5|sha256|sha512] value <file> [md5|sha256|sha512]
```

### Example Output

```
$ ./hash-calculator <file> [md5|sha256|sha512]
<file> [md5|sha256|sha512]
md5:    %s\n
sha256: %s\n
```

## Project Structure

```
Hash-Calculator/
  main.go          # Entry point (59 lines)
  go.mod            # Go module definition
  go.sum            # Dependency checksums
  README.md         # This file
  LICENSE           # MIT License
  CHANGELOG.md      # Version history
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) for details.

---

Built with ❤️ by [AetherCodeHQ](https://github.com/AetherCodeHQ)

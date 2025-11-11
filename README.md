# Custom TCP/IP Stack

Welcome to the **custom-tcp-ip-stack** repository! This project is a full-featured, custom implementation of the TCP/IP stack written in Go. It is designed as an educational resource, a foundation for embedded or userspace network development, or as a platform for experimentation with network protocols.

[![Go](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)
  
---

## Table of Contents

- [Project Overview](#project-overview)
- [Features](#features)
- [Directory Structure](#directory-structure)
- [Getting Started](#getting-started)
- [Documentation](#documentation)
    - [Architecture](docs/ARCHITECTURE.md)
    - [TCP Implementation Details](docs/TCP.md)
    - [IP Layer Details](docs/IP.md)
    - [Contributing](docs/CONTRIBUTING.md)
- [Usage](#usage)
- [FAQ](docs/FAQ.md)

---

## Project Overview

This repository builds a network stack from scratch, implementing the core layers of the OSI model required for modern networking. Its main focus is on clean, idiomatic Go code and extensibility. 

---

## Features

- **Custom TCP and IP Implementation**  
  Not a wrapper over platform sockets—core protocols are implemented from first principles.
- **Modular Architecture**  
  Each protocol layer and logic is separated for clarity and extensibility.
- **Educational**  
  Richly documented, with [detailed architecture](docs/ARCHITECTURE.md) and protocol-specific write-ups.

---

## Directory Structure

| Path              | Description                                |
|-------------------|--------------------------------------------|
| `main.go`         | Entry point; sample usage/demo app.         |
| `internal/`       | Core internal logic and shared utilities.   |
| `stack/`          | The core TCP/IP stack logic lives here.     |
| `custom-tcp-ip-stack` | Compiled output.                        |
| `go.mod`/`go.sum` | Go modules dependency management.           |
| `dump`            | Placeholder (e.g., for debugging output).   |
| `docs/`           | Project documentation (see below).          |

---

## Getting Started

1. **Prerequisites:**  
   - [Go 1.18+](https://golang.org/dl/)

2. **Clone the repo:**
   ```sh
   git clone https://github.com/toastsandwich/custom-tcp-ip-stack.git
   cd custom-tcp-ip-stack
   ```

3. **Build and run:**  
   ```sh
   go build -o custom-tcp-ip-stack
   ./custom-tcp-ip-stack
   ```

See [Usage](#usage) and [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md) for more details.

---

## Documentation

### High-Level Design

- [Architecture Overview](docs/ARCHITECTURE.md)
- [TCP Protocol Implementation](docs/TCP.md)
- [IP Layer Details](docs/IP.md)
- [Contributing Guide](docs/CONTRIBUTING.md)
- [Frequently Asked Questions](docs/FAQ.md)

---

## Usage

After building, run the demo via `./custom-tcp-ip-stack`.  
Use the library or stack modules as a starting point for your custom network projects.

---

## Contributing

We welcome contributions! Please see [CONTRIBUTING](docs/CONTRIBUTING.md).

---

## License

This project is intended for educational and personal use.

---

## Contact

Maintainer: [toastsandwich](https://github.com/toastsandwich)
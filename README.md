# Kongo CLI

[![Release](https://img.shields.io/github/release/fabiorphp/kongo-cli.svg?style=flat-square)](https://github.com/fabiorphp/kongo-cli/releases/latest)
[![Build Status](https://img.shields.io/travis/fabiorphp/kongo-cli/master.svg?style=flat-square)](https://travis-ci.org/fabiorphp/kongo-cli)
[![Coverage Status](https://img.shields.io/coveralls/fabiorphp/kongo-cli/master.svg?style=flat-square)](https://coveralls.io/github/fabiorphp/kongo-cli?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/fabiorphp/kongo-cli?style=flat-square)](https://goreportcard.com/report/github.com/fabiorphp/kongo-cli)
[![GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](https://github.com/fabiorphp/kongo-cli/blob/master/LICENSE)

Kongo CLI is an open source command-line application for managing Kong instances.

## Instalation

### Using go get
If you have [Go](https://golang.org) installed:

```sh
$ go get github.com/fabiorphp/kongo-cli/cmd/kongo
```

### Manually
Download your preferred flavor from the [releases page](https://github.com/fabiorphp/kongo-cli/releases) and install manually.

## Development

### Requirements

- Install [Go](https://golang.org)
- Install [go dep](https://github.com/golang/dep)

### Makefile
```sh
// Build a beta version of kongo
$ make build

// Clean up
$ make clean

// Creates folders and download dependencies
$ make configure

//Run tests and generates html coverage file
make cover

// Download project dependencies
make depend

// Format all go files
make fmt

//Run linters
make lint

// Run tests
make test
```

## License

This project is released under the MIT licence. See [LICENSE](https://github.com/fabiorphp/kongo-cli/blob/master/LICENSE) for more details.

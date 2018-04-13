# Kongo CLI

[![Build Status](https://img.shields.io/travis/fabiorphp/kongo-cli/master.svg?style=flat-square)](https://travis-ci.org/fabiorphp/kongo-cli)
[![Coverage Status](https://img.shields.io/coveralls/fabiorphp/kongo-cli/master.svg?style=flat-square)](https://coveralls.io/github/fabiorphp/kongo-cli?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/fabiorphp/kongo-cli?style=flat-square)](https://goreportcard.com/report/github.com/fabiorphp/kongo-cli)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](https://github.com/fabiorphp/kongo-cli/blob/master/LICENSE)

Manage Kong instances by CLI.

## Development

### Requirements

- Install [go dep](https://github.com/golang/dep)

### Run tests
```sh
// tests
$ make test

// test with coverage
$ make test-coverage

// clean-up
$ make clean

// configure (download dependencies and run docker containers)
$ make configure
```

## License

This project is released under the MIT licence. See [LICENSE](https://github.com/fabiorphp/kongo-cli/blob/master/LICENSE) for more details.

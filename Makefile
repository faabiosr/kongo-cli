APP_VERSION=0.0.1
LDFLAGS=--ldflags '-X main.version=${APP_VERSION} -extldflags "-static" -w -s'
OS=$(shell uname -s)

.DEFAULT_GOAL := build

# Build a beta version of kongo
build:
	CGO_ENABLED=0 go build -v -a ${LDFLAGS} -o ./build/kongo github.com/fabiorphp/kongo-cli/cmd/kongo
.PHONY: build

# Clean up
clean:
	@rm -fR ./build/ ./vendor/ ./cover.*
.PHONY: clean

# Creates folders and download dependencies
configure:
	@mkdir -p ./build
	dep ensure -v
.PHONY: configure

# Run tests and generates html coverage file
cover: test
	go tool cover -html=./build/cover.out -o ./build/cover.html
.PHONY: cover

depend:
	go get -u gopkg.in/alecthomas/gometalinter.v2
	gometalinter.v2 --install
	go get -u github.com/golang/dep/...
.PHONY: depend

# Format all go files
fmt:
	gofmt -s -w -l $(shell go list -f {{.Dir}} ./... | grep -v /vendor/)
.PHONY: fmt

# Run linters
lint:
	gometalinter.v2 --vendor ./...
.PHONY: lint

# Run tests
test:
	go test -v -race -coverprofile=./build/cover.out $(shell go list ./... | grep -v /vendor/)
.PHONY: test

VERSION := 0.0.13
COMMIT := $(shell git rev-parse --short HEAD)
LDFLAGS := -ldflags "-w -s -X main.version=${VERSION} -X main.commit=${COMMIT}"

.PHONY: all
all: mackerel-plugin-bitcoin-core
	docker build .

.PHONY: build
build: mackerel-plugin-bitcoin-core

mackerel-plugin-bitcoin-core: mackerel-plugin-bitcoin.go Makefile
	go build $(LDFLAGS) -o mackerel-plugin-bitcoin

.PHONY: linux
linux: mackerel-plugin-bitcoin.go
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o mackerel-plugin-bitcoin-core

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -rf mackerel-plugin-bitcoin-core

.PHONY: tag
tag:
	git tag v${VERSION}
	git push origin v${VERSION}
	git push origin main

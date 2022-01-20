VERSION := 0.0.10
COMMIT := $(shell git rev-parse --short HEAD)
LDFLAGS := -ldflags "-w -s -X main.version=${VERSION} -X main.commit=${COMMIT}"

all: mackerel-plugin-bitcoin-core
	docker build .

build: mackerel-plugin-bitcoin-core

mackerel-plugin-bitcoin-core: mackerel-plugin-bitcoin.go Makefile
	go build $(LDFLAGS) -o mackerel-plugin-bitcoin

linux: mackerel-plugin-bitcoin.go
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o mackerel-plugin-bitcoin-core

fmt:
	go fmt ./...

check:
	go test ./...

clean:
	rm -rf mackerel-plugin-bitcoin-core

tag:
	git tag v${VERSION}
	git push origin v${VERSION}
	git push origin main

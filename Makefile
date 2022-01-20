VERSION := 0.0.6
COMMIT := $(shell git rev-parse --short HEAD)
LDFLAGS := -ldflags "-w -s -X main.version=${VERSION} -X main.commit=${COMMIT}"

mackerel-plugin-bitcoin-core: mackerel-plugin-bitcoin-core.go Makefile
	go build $(LDFLAGS) -o mackerel-plugin-bitcoin-core

linux: mackerel-plugin-bitcoin-core.go
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

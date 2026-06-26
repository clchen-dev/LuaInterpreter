GO ?= go
BINARY ?= bin/luago

.PHONY: build test vet fmt ci clean

build:
	$(GO) build -trimpath -o $(BINARY) ./cmd/luago

test:
	$(GO) test -race -cover ./...

vet:
	$(GO) vet ./...

fmt:
	$(GO) fmt ./...

ci: vet test build

clean:
	rm -rf bin dist coverage.out

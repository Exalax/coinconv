LOCAL_BIN=$(CURDIR)/bin
LINTER_VERSION=v1.45.2

build:
	go build -o $(CURDIR)/bin/coinconv $(CURDIR)/cmd/

lint: $(LOCAL_BIN)/golangci-lint
	$(LOCAL_BIN)/golangci-lint run ./...

test:
	go test ./...

$(LOCAL_BIN)/golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(LINTER_VERSION)

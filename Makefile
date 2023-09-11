golangcilint_version := "1.54.2"

default: test

.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...

.PHONY: lint
lint:
	@(which golangci-lint &&  [[ "$$(golangci-lint --version | awk '{print $$4}')" == "$(golangcilint_version)" ]] ) || go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(golangcilint_version)
	@golangci-lint run

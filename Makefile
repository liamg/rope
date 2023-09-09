default: test

.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...


.PHONY: lint test

test:
	go test -v -race -cover ./...

lint:
	golangci-lint run

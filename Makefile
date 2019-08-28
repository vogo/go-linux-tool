format:
		goimports -w -l .
		go fmt ./...

check:
		golangci-lint run --disable=unused

test:
		go test

all: format check test
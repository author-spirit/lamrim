.PHONY: build run test fmt lint tidy

BINARY := bin/lamrim

build:
	go build -o $(BINARY) ./cmd/lamrim

run:
	go run ./cmd/lamrim -workflow first

serve:
	go run ./cmd/lamrim -serve

test:
	go test ./...

fmt:
	gofmt -s -w .

lint:
	golangci-lint run ./...

tidy:
	go mod tidy

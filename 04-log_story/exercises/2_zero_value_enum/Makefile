run:
	@go run .

build:
	@go build -o bin/logger main.go

test:
	@go test ./...

cover:
	@go test ./... -cover

lint:
	@golangci-lint run ./...

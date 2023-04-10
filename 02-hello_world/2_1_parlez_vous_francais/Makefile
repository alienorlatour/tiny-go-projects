run:
	@go run .

build:
	@go build -o bin/helloworld main.go

test:
	@go test ./...

cover:
	@go test ./... -cover

test_french:
	@go test ./... -run French

lint:
	@golangci-lint run ./...

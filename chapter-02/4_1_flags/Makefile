run:
	@go run .

build:
	@go build -o bin/helloworld main.go

test:
	@go test ./...

cover:
	@go test ./... -cover

lint:
	@golangci-lint run ./...

%: ## anything
	@go run . -lang=$@

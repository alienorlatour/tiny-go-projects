run:
	@go run -race .

build:
	@go build -o bin/ .

cover:
	@go test -race ./... -cover

lint:
	@golangci-lint run ./...

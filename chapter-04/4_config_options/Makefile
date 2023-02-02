run:
	@go run .

build:
	@go build -o bin/gordle main.go

test:
	@go test ./...

cover:
	@go test ./... -cover

benchmark:
	@go test ./... -run=^$ bench=. --benchmem

lint:
	@golangci-lint run ./...

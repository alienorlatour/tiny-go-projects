run:
	@go run -race cmd/main.go

build:
	@go build -o bin/ cmd/main.go

test:
	@go test -race ./...

cover:
	@go test -race ./... -cover

benchmark:
	@go test ./... -run=^$ bench=. --benchmem

lint:
	@golangci-lint run ./...

generate:
	@go generate ./...

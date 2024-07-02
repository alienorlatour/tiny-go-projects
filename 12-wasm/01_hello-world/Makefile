run:
	@go run -race main.go

build:
	@GOOS=js GOARCH=wasm go build -o main.wasm

test:
	@go test -race ./...

cover:
	@go test -race ./... -cover

benchmark:
	@go test ./... -run=^$ bench=. --benchmem

lint:
	@golangci-lint run ./...

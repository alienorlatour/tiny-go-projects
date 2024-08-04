build:
	@GOOS=js GOARCH=wasm go build -o main.wasm main.go

lint:
	@golangci-lint run ./...

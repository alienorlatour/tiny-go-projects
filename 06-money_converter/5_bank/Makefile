run:
	@go run . -from USD -to EUR 10.89

build:
	@go build -o bin/ .

test:
	@go test ./...

cover:
	@go test ./... -cover

benchmark:
	@go test ./... -run=^$ bench=. --benchmem

lint:
	@golangci-lint run ./...

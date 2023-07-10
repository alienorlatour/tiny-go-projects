run:
	go run .

build:
	go build -o bin/cache .

test:
	go test -race ./...

cover:
	go test ./... -cover

benchmark:
	go test ./... -run=^$ bench=. --benchmem

lint:
	golangci-lint run ./...

run:
	go run main.go

build:
	go build -o bin/change main.go

test:
	go test ./...

cover:
	go test ./... -cover

benchmark:
	go test ./... -run=^$ bench=. --benchmem

lint:
	 golangci-lint run ./...

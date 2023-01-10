run:
	go run main.go -from USD -to EUR 10.89

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

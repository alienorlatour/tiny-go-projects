run:
	@go run -race cmd/habits-server/main.go

build:
	@go build -o bin/ cmd/habits-server/main.go

test:
	@go test -race -test.short ./...

cover:
	@go test -race -test.short ./... -cover

benchmark:
	@go test ./... -run=^$ bench=. --benchmem

lint:
	@golangci-lint run ./...

generate:
	@protoc -I=api/proto/ --go_out=api/ --go_opt=paths=source_relative --go-grpc_out=api/ --go-grpc_opt=paths=source_relative api/proto/*.proto

integration:
	@go test -race --cover ./...

cover:
	go test --cover ./...

build-01:
	go build -o bin/chapter01 chapter-01/main.go

build: build-01

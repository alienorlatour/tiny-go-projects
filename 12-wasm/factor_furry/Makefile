wasm:
	cd ./cmd/wasm && \
	GOOS=js GOARCH=wasm go build -o  ../../assets/output.wasm

run_server:
	go run ./cmd/server/main.go
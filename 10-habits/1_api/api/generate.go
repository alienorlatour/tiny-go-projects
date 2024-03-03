// Package api contains the definitions of the habit service's endpoints.
package api

// This allows for the compilation of the API of the habit service.
// Since go generate doesn't expand filenames the same way bash does, we don't directly call protoc.
// Instead, we ask bash to run a command (the real protoc command), which allows us to delegate
//go:generate bash -c "protoc -I=proto/ --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto"

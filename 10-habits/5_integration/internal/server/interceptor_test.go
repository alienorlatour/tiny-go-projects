package server

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func Test_timerInterceptor(t *testing.T) {
	info := &grpc.UnaryServerInfo{FullMethod: "TestingFunc"}
	handler := func(ctx context.Context, req any) (any, error) {
		return "123", nil
	}

	var output bytes.Buffer

	interceptor := timerInterceptor(&output)

	_, _ = interceptor(context.Background(), "request", info, handler)
	assert.Contains(t, output.String(), "Time in TestingFunc: ")
}

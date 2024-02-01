package server

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"learngo-pockets/habits/internal/log"
)

func Test_timerInterceptor(t *testing.T) {
	info := &grpc.UnaryServerInfo{FullMethod: "TestingFunc"}
	handler := func(ctx context.Context, req any) (any, error) {
		return "123", nil
	}

	bfr := bytes.NewBuffer([]byte{})
	lgr := log.New(bfr)

	// Use the t variable for logging
	interceptor := timerInterceptor(lgr)

	_, _ = interceptor(context.Background(), "request", info, handler)
	assert.Contains(t, bfr.String(), "time in TestingFunc: ")
}

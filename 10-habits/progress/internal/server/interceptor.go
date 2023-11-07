package server

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

// timerIntercept is a unary interceptor that logs the time spent in each call.
func timerIntercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	defer fmt.Printf("Time in %s: %s\n", info.FullMethod, time.Since(start))

	resp, err := handler(ctx, req)

	return resp, err
}

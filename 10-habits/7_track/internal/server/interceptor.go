package server

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

func timerInterceptor(lgr Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		defer func() {
			lgr.Logf("time in %s: %s\n", info.FullMethod, time.Since(start))
		}()

		return handler(ctx, req)
	}
}

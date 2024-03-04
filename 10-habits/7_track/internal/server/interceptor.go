package server

import (
	"context"
	"fmt"
	"io"
	"time"

	"google.golang.org/grpc"
)

func timerInterceptor(writer io.Writer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		defer func() {
			_, _ = fmt.Fprintf(writer, "Time in %s: %s\n", info.FullMethod, time.Since(start))
		}()

		return handler(ctx, req)
	}
}

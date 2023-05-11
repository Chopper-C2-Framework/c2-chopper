package interceptor

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthInterceptor struct{}

func (i *AuthInterceptor) UnaryServerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Printf("Interceptor called for service %v with metadata %v\n", info.FullMethod, md)
	} else {
		fmt.Printf("Interceptor called for service %v without metadata\n", info.FullMethod)
	}
	return handler(ctx, req)
}

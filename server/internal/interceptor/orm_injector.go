package interceptor

import (
	"context"
	"fmt"

	orm "github.com/chopper-c2-framework/c2-chopper/server/domain"
	"google.golang.org/grpc"
)

type ORMInjectorInterceptor struct {
	DbConnection *orm.ORMConnection
}

func (i *ORMInjectorInterceptor) UnaryServerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	fmt.Println("ORMInjectorInterceptor executed.")
	ctx = context.WithValue(ctx, "Db", i.DbConnection)
	return handler(ctx, req)
}

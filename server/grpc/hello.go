package grpc

import (
	context "context"
	"fmt"

	"github.com/chopper-c2-framework/c2-chopper/grpc/proto"
)

type HelloService struct {
	proto.UnimplementedHelloServiceServer
}

func (s *HelloService) HelloWorld(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {

	return &proto.HelloResponse{
		Message: fmt.Sprintf("Hello World %s", in.Name),
	}, nil

}

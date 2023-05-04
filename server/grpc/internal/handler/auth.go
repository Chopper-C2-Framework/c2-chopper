package handler

import (
	context "context"

	"github.com/chopper-c2-framework/c2-chopper/proto"
)

// We define a server struct that implements the server interface. 🥳🥳🥳
type AuthServer struct {
	proto.UnimplementedAuthServer
}

// We implement the SayHello method of the server interface. 🥳🥳🥳
func (s *AuthServer) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "Hello, " + in.GetName()}, nil
}

package rest

import "github.com/chopper-c2-framework/c2-chopper/grpc/proto"

// Client IMPLEMENTATION OF THE GRPC CLIENT
type Client struct {
	AgentService proto.AgentServiceClient
}

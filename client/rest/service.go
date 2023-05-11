package rest

import "github.com/chopper-c2-framework/c2-chopper/grpc/proto"

// IMPLEMENTATION OF THE GRPC CLIENT
type Client struct {
	listenerService proto.ListenerServiceClient
}

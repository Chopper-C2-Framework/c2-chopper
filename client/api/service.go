package api

import "github.com/chopper-c2-framework/c2-chopper/proto"

// IMPLEMENTATION OF THE GRPC CLIENT
type Client struct {
	listenerService proto.ListenerServiceClient
}

package server

type IServerManager interface {
	IgRPCServer
	IgRPCServerHTTPGateway
}

type Manager struct {
	gRPCServer
	gRPCServerHTTPGateway
}

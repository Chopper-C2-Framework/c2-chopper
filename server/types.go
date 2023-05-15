package server

type IServerManager interface {
	IgRPCServer
	IgRPCServerHTTPGateway
}

type ServerManager struct {
	gRPCServer
	gRPCServerHTTPGateway
}

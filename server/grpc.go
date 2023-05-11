package server

import (
	"fmt"
	"log"

	Cfg "github.com/chopper-c2-framework/c2-chopper/core/config"
	"github.com/chopper-c2-framework/c2-chopper/grpc/proto"
	_ "github.com/chopper-c2-framework/c2-chopper/grpc/proto"
	handler "github.com/chopper-c2-framework/c2-chopper/server/grpc"

	"crypto/tls"

	"net"

	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
)

type IgRPCServer interface {
	NewgRPCServer(config *Cfg.Config) error
}

type gRPCServer struct {
	server *grpc.Server
}

func loadTLSCredentials(certFile string, keyFile string) (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	tlsCfg := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(tlsCfg), nil
}

func (server_m *gRPCServer) NewgRPCServer(config *Cfg.Config) error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Host, config.ServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("[+] Created listener on port", config.ServerPort)

	if config.UseTLS {
		tlsCredentials, err := loadTLSCredentials(config.ServerCert, config.ServerCertKey)
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}
		fmt.Println("[+] Loaded certificates.")
		server_m.server = grpc.NewServer(grpc.Creds(tlsCredentials))
	} else {
		server_m.server = grpc.NewServer()
	}

	proto.RegisterAuthServiceServer(server_m.server, &handler.AuthService{})
	proto.RegisterListenerServiceServer(server_m.server, &handler.ListenerService{})
	proto.RegisterTeamServiceServer(server_m.server, &handler.TeamService{})
	proto.RegisterPluginServiceServer(server_m.server, &handler.PluginService{})
	if err := server_m.server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil
}

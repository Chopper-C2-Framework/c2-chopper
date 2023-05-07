package server

import (
	"fmt"
	"log"

	Cfg "github.com/chopper-c2-framework/c2-chopper/core/config"

	"crypto/tls"

	"github.com/chopper-c2-framework/c2-chopper/proto"

	"github.com/chopper-c2-framework/c2-chopper/server/internal/handler"

	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type IgRPCServer interface {
	NewgRPCServer(config *Cfg.Config) error
}

type gRPCServer struct {
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

func (server_m gRPCServer) NewgRPCServer(config *Cfg.Config) error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Host, config.ServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("[+] Created listener on port", config.ServerPort)

	var s *grpc.Server
	if config.UseTLS {
		tlsCredentials, err := loadTLSCredentials(config.ServerCert, config.ServerCertKey)
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}
		fmt.Println("[+] Loaded certificates.")
		s = grpc.NewServer(grpc.Creds(tlsCredentials))
	} else {
		s = grpc.NewServer()
	}

	proto.RegisterAuthServiceServer(s, &handler.AuthService{})
	proto.RegisterAgentServiceServer(s, &handler.AgentService{})
	proto.RegisterTeamServiceServer(s, &handler.TeamService{})
	proto.RegisterPluginServiceServer(s, &handler.PluginService{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil
}

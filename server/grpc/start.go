package grpc

import (
	"fmt"
	"log"

	Cfg "github.com/chopper-c2-framework/c2-chopper/core/config"

	"crypto/tls"

	pb "github.com/chopper-c2-framework/c2-chopper/proto"
	"github.com/chopper-c2-framework/c2-chopper/server/grpc/internal/handler"

	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

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

func NewgRPCServer(config Cfg.Config) {
	tlsCredentials, err := loadTLSCredentials(config.ServerCert, config.ServerCertKey)
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	fmt.Println("[+] Loaded certificates.", config.ServerPort)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.ServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("[+] Created listener on port", config.ServerPort)

	s := grpc.NewServer(grpc.Creds(tlsCredentials))

	pb.RegisterAuthServer(s, &handler.AuthServer{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

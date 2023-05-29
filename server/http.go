package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/chopper-c2-framework/c2-chopper/core/config"
	gw "github.com/chopper-c2-framework/c2-chopper/grpc/proto"
)

type IgRPCServerHTTPGateway interface {
	// NewgRPCServerHTTPGateway TODO This function will be launched thro a go routine, and no return is expected from now on
	// We will gracefully terminate it when the main thread is done!
	NewgRPCServerHTTPGateway(*config.Config) error
}

type gRPCServerHTTPGateway struct {
}

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9002", "gRPC server endpoint")
)

func handleSvcRegError(err error) {
	if err != nil {
		log.Panicln("Error while registering service: ", err)
	}
}

func (g *gRPCServerHTTPGateway) NewgRPCServerHTTPGateway(config *config.Config) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := gw.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	handleSvcRegError(err)

	err = gw.RegisterTeamServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	handleSvcRegError(err)

	err = gw.RegisterTrackingServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	handleSvcRegError(err)

	err = gw.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	handleSvcRegError(err)

	err = gw.RegisterPluginServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	handleSvcRegError(err)

	err = gw.RegisterTaskServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	handleSvcRegError(err)

	err = gw.RegisterAgentServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	handleSvcRegError(err)

	fmt.Printf("[+] HTTP Gateway on on %d\n", config.ServerHTTPPort)

	corsConfig := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	server := &http.Server{
		Handler: corsConfig.Handler(mux),
		Addr:    fmt.Sprintf(":%d", config.ServerHTTPPort),
	}

	err = server.ListenAndServe()

	if err != nil {
		log.Fatalf("failed to serve: %v\n", err)
		return err
	}

	fmt.Printf("HTTP server started on port %d", config.ServerHTTPPort)

	return nil
}

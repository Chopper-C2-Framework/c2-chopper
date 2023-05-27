package server

import (
	"fmt"
	"log"

	"github.com/chopper-c2-framework/c2-chopper/core"
	Cfg "github.com/chopper-c2-framework/c2-chopper/core/config"
	"github.com/chopper-c2-framework/c2-chopper/grpc/proto"
	handler "github.com/chopper-c2-framework/c2-chopper/server/grpc"
	"github.com/chopper-c2-framework/c2-chopper/server/internal/interceptor"

	"crypto/tls"

	"net"

	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"

	orm "github.com/chopper-c2-framework/c2-chopper/core/domain"
	"github.com/chopper-c2-framework/c2-chopper/core/plugins"
)

type IgRPCServer interface {
	// NewgRPCServer TODO This function will be launched through a go routine, and no return is expected from now on
	// we need to handle error case and inform the main thread
	// > we need to make sure the grpc gateway is only open when this succeeds
	// we will gracefully terminate it when the main thread is done
	NewgRPCServer(
		config *Cfg.Config,
		pluginManager *plugins.PluginManager,
	) error
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

func (Server *gRPCServer) NewgRPCServer(
	frameworkConfig *Cfg.Config,
	pluginManager *plugins.PluginManager,
) error {
	Agent, err := net.Listen("tcp", fmt.Sprintf("%s:%d", frameworkConfig.Host, frameworkConfig.ServergRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("[+] Created Agent on port", frameworkConfig.ServergRPCPort)

	dbConnection, _ := orm.CreateDB(frameworkConfig)
	coreServices := core.InitServices(dbConnection, *frameworkConfig)

	AuthInterceptor := interceptor.AuthInterceptor{
		AuthService: coreServices.AuthService,
		AccessibleRoles: map[string][]string{
			"/TeamService/UpdateTeam": {"User"},
			"/TeamService/CreateTeam": {"User"},
			"/TeamService/DeleteTeam": {"User"},
		},
	}

	UnaryInterceptors := grpc.ChainUnaryInterceptor(
		AuthInterceptor.IsAuthenticatedInterceptor,
	)

	if frameworkConfig.UseTLS {
		tlsCredentials, err := loadTLSCredentials(frameworkConfig.ServerCert, frameworkConfig.ServerCertKey)
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}
		fmt.Println("[+] Loaded certificates.")
		Server.server = grpc.NewServer(
			grpc.Creds(tlsCredentials),
			UnaryInterceptors,
		)
	} else {
		Server.server = grpc.NewServer(
			UnaryInterceptors,
		)
	}

	proto.RegisterAuthServiceServer(Server.server, &handler.AuthService{
		AuthService: coreServices.AuthService,
	})
	proto.RegisterAgentServiceServer(Server.server, &handler.AgentService{
		AgentService: coreServices.AgentService,
	})

	proto.RegisterTeamServiceServer(Server.server, &handler.TeamService{
		TeamService: coreServices.TeamService,
	})
	proto.RegisterPluginServiceServer(Server.server, &handler.PluginService{
		PluginManager: pluginManager,
	})

	proto.RegisterTaskServiceServer(Server.server, &handler.TaskService{
		TaskService:  coreServices.TaskService,
		AgentService: coreServices.AgentService,
	})
	proto.RegisterProfileServiceServer(Server.server, &handler.ProfileService{})
	proto.RegisterTrackingServiceServer(Server.server, &handler.TrackingService{})
	proto.RegisterHelloServiceServer(Server.server, &handler.HelloService{})

	if err := Server.server.Serve(Agent); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("gRPC server started on port", frameworkConfig.ServergRPCPort)
	return nil
}

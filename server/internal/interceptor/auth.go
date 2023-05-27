package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/chopper-c2-framework/c2-chopper/core/services"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthInterceptor struct {
	AuthService     services.IAuthService
	AccessibleRoles map[string][]string
}

type ContextUser struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Id       string `json:"id"`
}

func (i *AuthInterceptor) IsAuthenticatedInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Printf("Interceptor called for service %v with metadata %v\n", info.FullMethod, md)
	} else {
		fmt.Printf("Interceptor called for service %v without metadata\n", info.FullMethod)
	}

	user, err := i.authorize(ctx, info.FullMethod)
	if err != nil {
		return nil, err
	}

	ctx = metadata.AppendToOutgoingContext(ctx, "userId", user.Id)

	return handler(ctx, req)
}

func (i *AuthInterceptor) authorize(ctx context.Context, method string) (*ContextUser, error) {
	logrus.Println("AuthInterceptor.authorize", method)
	accessibleRoles, ok := i.AccessibleRoles[method]
	if !ok {
		// everyone can access
		return nil, nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	claims, err := i.AuthService.ParseToken(accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	for _, role := range accessibleRoles {
		if role == claims.Role {
			return &ContextUser{Id: claims.Subject, Username: claims.Username, Role: claims.Role}, nil
		}
	}

	return nil, status.Error(codes.PermissionDenied, "no permission to access this RPC")
}

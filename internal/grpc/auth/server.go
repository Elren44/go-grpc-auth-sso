package authgrpc

import (
	"context"

	ssov1 "github.com/Elren44/go-grpc-auth-proto/gen/go/sso"
	"google.golang.org/grpc"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	// Implement login logic here
	return &ssov1.LoginResponse{}, nil
}

func (s *serverAPI) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	// Implement register logic here
	return &ssov1.RegisterResponse{}, nil
}

func (s *serverAPI) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	// Implement is admin logic here
	return &ssov1.IsAdminResponse{}, nil
}

package grpcserver

import (
	// "context"

	"github.com/maximus969/users-app/internal/app/services"
	usersv1 "github.com/maximus969/users-app/protos/gen/go/users"
	"google.golang.org/grpc"
)

func RegisterGRPCServer(gRPC *grpc.Server, userService *services.UserService) {
	adapter := NewGRPCServerAdapter(userService)
	usersv1.RegisterUsersServer(gRPC, adapter)
}

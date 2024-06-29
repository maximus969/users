package grpcserver

import (
	"context"

	"github.com/google/uuid"
	"github.com/maximus969/users-app/internal/app/domain"
	"github.com/maximus969/users-app/internal/app/services"
	usersv1 "github.com/maximus969/users-app/protos/gen/go/users"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GRPCServerAdapter struct {
	usersv1.UnimplementedUsersServer
	userService *services.UserService
}

func NewGRPCServerAdapter(userService *services.UserService) *GRPCServerAdapter {
	return &GRPCServerAdapter{
		userService: userService,
	}
}

func (s *GRPCServerAdapter) CreateUser(ctx context.Context, req *usersv1.UserRequest) (*usersv1.UserResponse, error) {
	newUser, err := domain.NewUser(domain.NewUserData{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Age:       uint(req.Age),
	})

	if err != nil {
		return nil, err
	}

	createdUser, err := s.userService.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return &usersv1.UserResponse{
		Id:        createdUser.Id().String(),
		Firstname: createdUser.Firstname(),
		Lastname:  createdUser.Lastname(),
		Email:     createdUser.Email(),
		Age:       int64(createdUser.Age()),
		Created:   timestamppb.New(createdUser.Created()),
	}, nil
}

// DeleteUser implements usersv1.UsersServer.
func (s *GRPCServerAdapter) DeleteUser(ctx context.Context, req *usersv1.DeleteUserRequest) (*usersv1.DeleteUserResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	err = s.userService.DeleteUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return &usersv1.DeleteUserResponse{}, nil
}

// GetUserById implements usersv1.UsersServer.
func (s *GRPCServerAdapter) GetUserById(ctx context.Context, req *usersv1.GetUserByIdRequest) (*usersv1.UserResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	user, err := s.userService.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &usersv1.UserResponse{
		Id:        user.Id().String(),
		Firstname: user.Firstname(),
		Lastname:  user.Lastname(),
		Email:     user.Email(),
		Age:       int64(user.Age()),
		Created:   timestamppb.New(user.Created()),
	}, nil
}

// GetUsers implements usersv1.UsersServer.
func (s *GRPCServerAdapter) GetUsers(ctx context.Context, req *usersv1.GetUsersRequest) (*usersv1.UsersResponse, error) {
	users, err := s.userService.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	var userResponses []*usersv1.UserResponse
	for _, user := range users {
		userResponse := &usersv1.UserResponse{
			Id:        user.Id().String(),
			Firstname: user.Firstname(),
			Lastname:  user.Lastname(),
			Email:     user.Email(),
			Age:       int64(user.Age()),
			Created:   timestamppb.New(user.Created()),
		}
		userResponses = append(userResponses, userResponse)
	}

	return &usersv1.UsersResponse{
		Users: userResponses,
	}, nil
}

// UpdateUser implements usersv1.UsersServer.
func (s *GRPCServerAdapter) UpdateUser(ctx context.Context, req *usersv1.UserRequest) (*usersv1.UserResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	user, err := domain.NewUser(domain.NewUserData{
		Id:        id,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Age:       uint(req.Age),
	})

	if err != nil {
		return nil, err
	}

	updatedUser, err := s.userService.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &usersv1.UserResponse{
		Id:        updatedUser.Id().String(),
		Firstname: updatedUser.Firstname(),
		Lastname:  updatedUser.Lastname(),
		Email:     updatedUser.Email(),
		Age:       int64(updatedUser.Age()),
		Created:   timestamppb.New(updatedUser.Created()),
	}, nil
}

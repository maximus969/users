package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/maximus969/users-app/internal/app/domain"
)

// UserService is a user service
type UserService struct {
	repo UserRepository
}

// NewUserService creates a new user service
func NewUserService(repo UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (s UserService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	return s.repo.CreateUser(ctx, user)
}

func (s UserService) GetUserById(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return s.repo.GetUserById(ctx, id)
}

func (s UserService) GetUsers(ctx context.Context) ([]domain.User, error) {
	return s.repo.GetUsers(ctx)
}

func (s UserService) UpdateUser(ctx context.Context, user domain.User) (domain.User, error) {
	return s.repo.UpdateUser(ctx, user)
}

func (s UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteUser(ctx, id)
}

//go:generate mockery

package httpserver

import (
	"context"

	"github.com/google/uuid"
	"github.com/maximus969/users-app/internal/app/domain"
)

// UserService is a user service
type UserService interface {
	GetUserById(ctx context.Context, id uuid.UUID) (domain.User, error)
	GetUsers(ctx context.Context) ([]domain.User, error)
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	UpdateUser(ctx context.Context, user domain.User) (domain.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

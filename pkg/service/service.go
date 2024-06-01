package service

import (	
	"github.com/maximus969/users-app"
	"github.com/maximus969/users-app/pkg/repository"
)

type Users interface {
	Create(newUser users.User) error
	GetAllUsers() ([]users.User, error)
	GetUser(id string) (users.User, error)
	DeleteUser(id string) error
}

type Service struct {
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Users:      NewUsersService(repos.Users),
	}
}
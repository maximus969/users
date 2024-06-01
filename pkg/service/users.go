package service

import (
	"github.com/maximus969/users-app"
	"github.com/maximus969/users-app/pkg/repository"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) Create(newUser users.User) error {
	return s.repo.Create(newUser)
}
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

func (s *UsersService) GetAllUsers() ([]users.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UsersService) GetUser(id string) (users.User, error) {
	return s.repo.GetUser(id)
}

func (s *UsersService) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}
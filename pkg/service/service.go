package service

import "github.com/maximus969/users-app/pkg/repository"

type Users interface {
	//
}

type Service struct {
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
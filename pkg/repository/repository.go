package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/maximus969/users-app"
)

type Users interface {
	Create(newUser users.User) error
	GetAllUsers() ([]users.User, error)
	GetUser(id string) (users.User, error)
	DeleteUser(id string) error
}

type Repository struct {
	Users
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Users:      NewUsersPostgres(db),
	}
}
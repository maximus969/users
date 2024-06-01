package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/maximus969/users-app"
)

type Users interface {
	Create(newUser users.User) error
}

type Repository struct {
	Users
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Users:      NewUsersPostgres(db),
	}
}
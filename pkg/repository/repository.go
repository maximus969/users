package repository

import "github.com/jmoiron/sqlx"

type Users interface {
	//
}

type Repository struct {
	Users
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
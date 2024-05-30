package repository

type Users interface {
	//
}

type Repository struct {
	Users
}

func NewRepository() *Repository {
	return &Repository{}
}
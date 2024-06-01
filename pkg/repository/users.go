package repository

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/maximus969/users-app"
)

type UsersPostgres struct {
	db *sqlx.DB
}

func NewUsersPostgres(db *sqlx.DB) *UsersPostgres {
	return &UsersPostgres{db: db}
}

func (r *UsersPostgres) Create(user users.User) error {
	user.Id = uuid.New()
	user.Created = time.Now()
	
	query := fmt.Sprintf(`INSERT INTO %s (id, firstname, lastname, email, age, created) values ($1, $2, $3, $4, $5, $6) RETURNING id`, usersTable)
	
	row := r.db.QueryRow(query, user.Id, user.Firstname, user.Lastname, user.Email, user.Age, user.Created )

	var returnedId uuid.UUID
	
	err := row.Scan(&returnedId)
	
	if err != nil {
		fmt.Println("Error saving to Postgres:", err)
		return err
	}	
	
	return nil
}

func (r *UsersPostgres) GetAllUsers() ([]users.User, error) {
	var allUsers []users.User

	query := fmt.Sprintf(`SELECT * FROM %s`, usersTable)

	err := r.db.Select(&allUsers, query)

	return allUsers, err
}

func (r *UsersPostgres) GetUser(id string) (users.User, error) {
	var user users.User

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`,
		usersTable)

	err := r.db.Get(&user, query, id)

	return user, err
}
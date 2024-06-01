package repository

import (
	"fmt"
	"reflect"
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

func (r *UsersPostgres) UpdateUser(id string, updatedUser users.UserUpdate) error {
	userUpdates := StructToMap(updatedUser)

	delete(userUpdates, "id")
	delete(userUpdates, "created")

	if len(userUpdates) == 0 {
		return nil
	}

	query := `UPDATE users SET `

	args := []interface{}{}
	i := 1
	for key, value := range userUpdates {
		query += fmt.Sprintf("%s = $%d, ", key, i)
		args = append(args, value)
		i++
	}

	query = query[:len(query)-2] // Remove the trailing comma and space
	query += fmt.Sprintf(" WHERE id = $%d", i)
	
	args = append(args, id)

	_, err := r.db.Exec(query, args...)

	return err
}

func (r *UsersPostgres) DeleteUser(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1",
		usersTable)
	_, err := r.db.Exec(query, id)

	return err
}

func StructToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	objValue := reflect.ValueOf(obj)
	objType := reflect.TypeOf(obj)

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		fieldValue := objValue.Field(i)

		fieldName := field.Name

		if !fieldValue.IsNil() {
			result[fieldName] = fieldValue.Elem().Interface()
		}
	}

	return result
}

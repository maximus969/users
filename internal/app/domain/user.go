package domain

import (
	"time"

	"github.com/google/uuid"
)

// User is a domain User.
type User struct {
	id        uuid.UUID 
	firstname string    
	lastname  string    
	email     string    
	age       uint
	created   time.Time
}

type NewUserData struct {
	Id        uuid.UUID
	Firstname string   
	Lastname  string   
	Email     string   
	Age       uint	
}

// NewUser creates a new user.
func NewUser(data NewUserData) (User, error) {
	return User{
		id:       data.Id,
		firstname: data.Firstname,
		lastname: data.Lastname,
		email:    data.Email,
		age:    data.Age,		
	}, nil
}

// ID returns the user ID.
func (u User) Id() uuid.UUID {
	return u.id
}

// Username returns the user firstname.
func (u User) Firstname() string {
	return u.firstname
}

// Password returns the user lastname.
func (u User) Lastname() string {
	return u.lastname
}

// Password returns the user email.
func (u User) Email() string {
	return u.email
}

// Password returns the user age.
func (u User) Age() uint {
	return u.age
}

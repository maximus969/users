package users

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname" binding:"required"`
	Lastname  string    `json:"lastname" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Age       uint      `json:"age" binding:"required"`
	Created   time.Time `json:"created"`
}

type UserUpdate struct {
	Id        *uuid.UUID `json:"id,omitempty" db:"id"`
	Firstname *string    `json:"firstname,omitempty" db:"firstname"`
	Lastname  *string    `json:"lastname,omitempty" db:"lastname"`
	Email     *string    `json:"email,omitempty" db:"email"`
	Age       *uint      `json:"age,omitempty" db:"age"`
	Created   *time.Time `json:"created,omitempty" db:"created"`
}

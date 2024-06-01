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

package httpserver

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/maximus969/users-app/internal/app/domain"
)

type UserRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Age       uint   `json:"age"`
}

func (u *UserRequest) Validate() error {
	if u.Firstname == "" {
		return fmt.Errorf("%w: firstname", domain.ErrRequired)
	}
	if u.Lastname == "" {
		return fmt.Errorf("%w: lastname", domain.ErrRequired)
	}
	if u.Email == "" {
		return fmt.Errorf("%w: email", domain.ErrRequired)
	}
	if u.Age <= 0 {
		return fmt.Errorf("%w: age", domain.ErrNegative)
	}

	return nil
}

type UserResponse struct {
	Id        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Age       uint      `json:"age"`
	Created   time.Time `json:"created"`
}

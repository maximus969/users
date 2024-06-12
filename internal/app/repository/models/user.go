package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// User is a domain user.
type User struct {
	bun.BaseModel `bun:"table:users"`
	Id            uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()"`
	Firstname     string
	Lastname      string
	Email         string
	Age           uint
	Created       time.Time `bun:",nullzero"`
}

package domain

import (
	"context"
	"github.com/gofrs/uuid"
)

type User struct {
	id           uuid.UUID
	name         string
	email        string
	passwordHash []byte
}

func (u *User) ID() uuid.UUID        { return u.id }
func (u *User) Name() string         { return u.name }
func (u *User) PasswordHash() []byte { return u.passwordHash }
func (u *User) Email() string {
	return u.email
}

func NewUser(name string, email string, hash []byte) *User {
	return &User{
		id:           uuid.Must(uuid.NewV7()),
		name:         name,
		email:        email,
		passwordHash: hash,
	}
}

type UserRepository interface {
	Save(ctx context.Context, user *User) error
	FindByName(ctx context.Context, name string) (*User, error)
}

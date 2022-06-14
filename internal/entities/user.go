package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(name, email, username, password string) *User {
	return &User{
		UUID:      uuid.New(),
		Name:      name,
		Email:     email,
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type UserService interface {
	Register(name, email, username, password string) (*User, error)
}

type UserRepository interface {
	Create(name, email, username, password string) (*User, error)
	GetByUsername(username string) (*User, error)
}

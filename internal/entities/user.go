package entities

import "time"

type User struct {
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserService interface {
	Create(name, email, username, password string) (*User, error)
}

type UserRepository interface {
	Create(name, email, username, password string) (*User, error)
}

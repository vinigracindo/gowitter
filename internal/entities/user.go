package entities

type User struct {
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserService interface {
	Create(name, email, username, password string) (*User, error)
}

type UserRepository interface {
	Create(name, email, username, password string) (*User, error)
}

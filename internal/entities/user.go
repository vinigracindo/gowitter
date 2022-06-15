package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID      uuid.UUID
	Name      string
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Following []*User
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
		Following: []*User{},
	}
}

func (u *User) Follow(userToBeFollowed *User) error {
	if u.UUID == userToBeFollowed.UUID {
		return ErrCannotFollowYourself
	}

	if u.AlreadyFollowing(userToBeFollowed) {
		return ErrAlreadyFollowing
	}

	u.Following = append(u.Following, userToBeFollowed)
	return nil
}

func (u *User) AlreadyFollowing(user *User) bool {
	for _, follower := range u.Following {
		if follower.UUID == user.UUID {
			return true
		}
	}
	return false
}

type UserService interface {
	Register(name, email, username, password string) (*User, error)
	Follow(userUUID, userToBeFollowedUUID uuid.UUID) error
}

type UserRepository interface {
	Create(name, email, username, password string) (*User, error)
	Update(user *User) error
	FindByUUID(uuid uuid.UUID) (*User, error)
}

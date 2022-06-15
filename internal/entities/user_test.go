package entities_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/gowitter/internal/entities"
)

func TestUserEntity_NewUser(t *testing.T) {
	user := entities.NewUser("John Doe", "john@doe.com", "johndoe", "secretpass")

	t.Run("user cannot be nil", func(t *testing.T) {
		assert.NotNil(t, user)
	})

	t.Run("user uuid cannot be nil. entities.NewUser must generate a random uuid.", func(t *testing.T) {
		assert.NotNil(t, user.UUID)
	})

	t.Run("CreatedAt and UpdatedAt must be set on entities.NewUser", func(t *testing.T) {
		assert.NotZero(t, user.CreatedAt)
		assert.NotZero(t, user.UpdatedAt)
	})
}

func TestUserEntity_Follow(t *testing.T) {
	user := entities.NewUser("John Doe", "john@doe.com", "johndoe", "secretpass")
	userToBeFollowed := entities.NewUser("Patrick", "patrick@gmail.com", "patrick", "secretpass")

	user.Follow(userToBeFollowed)

	t.Run("check if user is following userToBeFollowed", func(t *testing.T) {
		assert.Equal(t, user.Following[0], userToBeFollowed)
	})

	t.Run("user cannot follow yourself", func(t *testing.T) {
		err := user.Follow(user)
		assert.Equal(t, entities.ErrCannotFollowYourself, err)
	})

	t.Run("user cannot follow a user that is already following", func(t *testing.T) {
		err := user.Follow(userToBeFollowed)
		assert.Equal(t, entities.ErrAlreadyFollowing, err)
	})
}

func TestUserEntity_AlreadyFollowing(t *testing.T) {
	user := entities.NewUser("John Doe", "john@doe.com", "johndoe", "secretpass")
	userToBeFollowed := entities.NewUser("Patrick", "patrick@gmail.com", "patrick", "secretpass")

	t.Run("user cannot get follower that is not following", func(t *testing.T) {
		assert.False(t, user.AlreadyFollowing(userToBeFollowed))
	})

	user.Follow(userToBeFollowed)

	t.Run("user can get follower that is following", func(t *testing.T) {
		assert.True(t, user.AlreadyFollowing(userToBeFollowed))
	})

}

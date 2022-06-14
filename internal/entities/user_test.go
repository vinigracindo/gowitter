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

package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/gowitter/internal/entities"
	"github.com/vinigracindo/gowitter/internal/entities/mocks"
	"github.com/vinigracindo/gowitter/internal/services"
)

func TestUserService(t *testing.T) {
	t.Run("UserService must implement UserService interface", func(t *testing.T) {
		assert.Implements(t, (*entities.UserService)(nil), services.NewUserService(nil))
	})
}

func TestUserService_Register(t *testing.T) {

	t.Run("should create a user", func(t *testing.T) {
		mockRepo := mocks.UserRepository{}
		mockRepo.On("GetByUsername", "username").Return(nil, nil)
		mockRepo.On("Create", "name", "email", "username", "password").Return(&entities.User{}, nil)
		service := services.NewUserService(&mockRepo)

		user, err := service.Register("name", "email", "username", "password")
		assert.NoError(t, err)
		assert.NotNil(t, user)
	})

	t.Run("should return error when create user failed", func(t *testing.T) {
		mockRepo := mocks.UserRepository{}
		mockRepo.On("GetByUsername", "username").Return(nil, nil)
		mockRepo.On("Create", "name", "email", "username", "password").Return(nil, entities.ErrUserAlreadyExists)
		service := services.NewUserService(&mockRepo)

		user, err := service.Register("name", "email", "username", "password")
		assert.NotNil(t, err)
		assert.Nil(t, user)
	})

}

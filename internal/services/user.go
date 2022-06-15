package services

import (
	"github.com/google/uuid"
	"github.com/vinigracindo/gowitter/internal/entities"
)

type service struct {
	repo entities.UserRepository
}

func NewUserService(repo entities.UserRepository) entities.UserService {
	return &service{repo: repo}
}

func (s *service) Register(name, email, username, password string) (*entities.User, error) {
	user, err := s.repo.Create(name, email, username, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) Follow(userUUID, userToBeFollowedUUID uuid.UUID) error {
	user, err := s.repo.FindByUUID(userUUID)
	if err != nil {
		return err
	}

	userToBeFollowed, err := s.repo.FindByUUID(userToBeFollowedUUID)
	if err != nil {
		return err
	}

	err = user.Follow(userToBeFollowed)
	if err != nil {
		return err
	}

	err = s.repo.Update(user)
	if err != nil {
		return err
	}

	return nil
}

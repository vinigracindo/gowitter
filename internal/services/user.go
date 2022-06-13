package services

import "github.com/vinigracindo/gowitter/internal/entities"

type service struct {
	repo entities.UserRepository
}

func NewUserService(repo entities.UserRepository) entities.UserService {
	return &service{repo: repo}
}

func (s *service) Create(name, email, username, password string) (*entities.User, error) {
	user, err := s.repo.Create(name, email, username, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

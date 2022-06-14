package services

import "github.com/vinigracindo/gowitter/internal/entities"

type service struct {
	repo entities.UserRepository
}

func NewUserService(repo entities.UserRepository) entities.UserService {
	return &service{repo: repo}
}

func (s *service) Register(name, email, username, password string) (*entities.User, error) {
	userAlreadyExist, _ := s.repo.GetByUsername(username)

	if userAlreadyExist != nil {
		return nil, ErrUsernameMustBeUnique
	}

	user, err := s.repo.Create(name, email, username, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

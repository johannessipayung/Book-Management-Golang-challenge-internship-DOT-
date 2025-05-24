package service

import (
	"challengeGO/model"
	"challengeGO/repository"
	"errors"
)

type UserService interface {
	Register(user *model.User) error
	Login(email, password string) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}

func (s *userService) Register(user *model.User) error {
	return s.repo.Create(user)
}

func (s *userService) Login(email, password string) (*model.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil || user.Password != password {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

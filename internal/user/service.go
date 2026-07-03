package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(user *User) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Register(user *User) error {

	existing, _ := s.repo.GetByEmail(user.Email)

	if existing != nil {
		return errors.New("email already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hash)

	return s.repo.Create(user)
}
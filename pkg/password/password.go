package password

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrPasswordMismatch = errors.New("invalid email or password")
)

// Service defines password-related operations.
type Service interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, plainPassword string) error
}

// service implements Service.
type service struct{}

// New creates a new password service.
func New() Service {
	return &service{}
}

// Hash converts a plain password into a bcrypt hash.
func (s *service) Hash(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// Compare checks whether the plain password matches the stored hash.
func (s *service) Compare(hashedPassword, plainPassword string) error {

	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(plainPassword),
	)

	if err != nil {
		return ErrPasswordMismatch
	}

	return nil
}
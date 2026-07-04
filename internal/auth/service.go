package auth

import (
	"errors"
	"time"

	"example.com/hr-emp-mgmt/internal/refresh"
	"example.com/hr-emp-mgmt/pkg/jwt"
	"example.com/hr-emp-mgmt/pkg/token"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(req LoginRequest) (*LoginResponse, error) {

	// Step 1
	userData, err := s.userRepo.GetByEmail(req.Email)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Step 2
	err = bcrypt.CompareHashAndPassword(
		[]byte(userData.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Step 3
	accessToken, err := jwt.GenerateAccessToken(
		userData.ID,
		userData.Role,
	)

	if err != nil {
		return nil, err
	}

	// Step 4
	refreshToken, err := token.GenerateRefreshToken()

	if err != nil {
		return nil, err
	}

	// Step 5
	refresh := refresh.RefreshToken{
		UserID: userData.ID,
		Token: refreshToken,
		ExpiresAt: time.Now().Add(
			7 * 24 * time.Hour,
		),
	}

	err = s.refreshRepo.Create(&refresh)

	if err != nil {
		return nil, err
	}

	// Step 6
	response := &LoginResponse{

		AccessToken: accessToken,

		RefreshToken: refreshToken,

		User: UserResponse{
			ID:        userData.ID,
			FirstName: userData.FirstName,
			LastName:  userData.LastName,
			Email:     userData.Email,
			Role:      userData.Role,
		},
	}

	return response, nil
}

// package auth

// import (
// 	"example.com/hr-emp-mgmt/internal/refresh"
// 	"example.com/hr-emp-mgmt/internal/user"
// )

// type Service interface {
// 	Login(req LoginRequest) (*LoginResponse, error)
// 	Register(req RegisterRequest) error
// }

// type service struct {
// 	userRepo    user.Repository
// 	refreshRepo refresh.Repository
// }

// func NewService(
// 	userRepo user.Repository,
// 	refreshRepo refresh.Repository,
// ) Service {
// 	return &service{
// 		userRepo:    userRepo,
// 		refreshRepo: refreshRepo,
// 	}
// }
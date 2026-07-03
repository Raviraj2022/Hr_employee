package user

import (
	"example.com/hr-emp-mgmt/config"

	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error
	GetByEmail(email string) (*User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository() Repository {
	return &repository{
		db: config.DB,
	}
}

func (r *repository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *repository) GetByEmail(email string) (*User, error) {
	var user User

	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
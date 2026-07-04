package refresh

import (
	"example.com/hr-emp-mgmt/config"

	"gorm.io/gorm"
)

type Repository interface {
	Create(token *RefreshToken) error
	GetByToken(token string) (*RefreshToken, error)
	Delete(token string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository() Repository {
	return &repository{
		db: config.DB,
	}
}

func (r *repository) Create(token *RefreshToken) error {
	return r.db.Create(token).Error
}

func (r *repository) GetByToken(token string) (*RefreshToken, error) {

	var refresh RefreshToken

	err := r.db.Where("token = ?", token).First(&refresh).Error

	if err != nil {
		return nil, err
	}

	return &refresh, nil
}

func (r *repository) Delete(token string) error {
	return r.db.Where("token = ?", token).Delete(&RefreshToken{}).Error
}
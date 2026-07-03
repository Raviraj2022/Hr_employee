package user

import "time"

type User struct {
	ID         string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FirstName  string    `gorm:"not null"`
	LastName   string
	Email      string    `gorm:"uniqueIndex;not null"`
	Password   string    `gorm:"not null"`
	Role       string    `gorm:"default:employee"`
	IsActive   bool      `gorm:"default:true"`
	IsVerified bool      `gorm:"default:false"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (User) TableName() string {
	return "users"
}
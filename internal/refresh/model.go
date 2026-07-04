package refresh

import "time"

type RefreshToken struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    string    `gorm:"type:uuid"`
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
}

func (RefreshToken) TableName() string {
	return "refresh_tokens"
}
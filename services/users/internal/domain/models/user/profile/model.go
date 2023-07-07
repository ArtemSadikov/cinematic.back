package profile

import "github.com/google/uuid"

type Profile struct {
	UserID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email    string    `gorm:"not null;unique"`
	Username string    `gorm:"not null"`
}

func (p *Profile) TableName() string {
	return "users_profiles"
}

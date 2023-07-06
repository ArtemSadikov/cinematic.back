package auth

import (
	"github.com/google/uuid"
	"time"
)

type UserAuth struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	UserID    string    `gorm:"unique"`
	Password  string
	Token     string
	CreatedAt time.Time
}

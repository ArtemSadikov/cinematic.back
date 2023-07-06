package user

import (
	"cinematic.back/services/users/internal/domain/models/user/profile"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID      uuid.UUID       `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Profile profile.Profile `gorm:"foreignKey:UserId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

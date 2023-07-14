package user

import (
	"cinematic.back/services/users/internal/domain/models/auth"
	"cinematic.back/services/users/internal/domain/models/user/profile"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID      uuid.UUID       `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Profile profile.Profile `gorm:"foreignKey:UserID;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Auth    *auth.UserAuth  `gorm:"foreignKey:UserID;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;nullable"`

	CreatedAt time.Time `gorm:"<-create"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (u *User) IsNull() bool {
	return u.ID == uuid.Nil
}

package token

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type AuthToken struct {
	AuthID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Token         string    `gorm:"not null;type:character varying(255)"`
	TokenID       uuid.UUID `gorm:"not null;type:uuid;default:uuid_generate_v4()"`
	LastRefreshAt time.Time `gorm:"default:now();not null;autoUpdateTime:false"`
}

func (a AuthToken) BeforeSave(db *gorm.DB) error {
	db.Statement.SetColumn("LastRefreshAt", time.Now())
	return nil
}

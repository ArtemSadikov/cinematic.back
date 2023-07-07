package auth

import (
	"cinematic.back/pkg/crypto"
	"cinematic.back/services/users/internal/domain/models/auth/token"
	"github.com/google/uuid"
	"time"
)

type UserAuth struct {
	ID           uuid.UUID        `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID       uuid.UUID        `gorm:"unique;not null"`
	Password     string           `gorm:"not null;type:character varying(255)"`
	PasswordID   uuid.UUID        `gorm:"not null;type:uuid;default:uuid_generate_v4()"`
	RegisteredAt time.Time        `gorm:"default:now();autoUpdateTime:false;not null"`
	Token        *token.AuthToken `gorm:"nullable;foreignKey:AuthID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:id"`
}

func (m *UserAuth) TableName() string {
	return "users_auths"
}

func (m *UserAuth) IsNull() bool {
	return m.ID == uuid.Nil
}

func (m *UserAuth) SetToken(data string, tokenId uuid.UUID) {
	if m.Token == nil {
		m.Token = &token.AuthToken{}
	}

	m.Token.Token = crypto.HashToken(data)
	m.Token.TokenID = tokenId
}

func (m *UserAuth) ChangePassword(pswd string) {
	if err := crypto.ComparePassword(pswd, m.Password); err == nil {
		return
	}

	m.Password = crypto.HashPassword(pswd)
	m.PasswordID = uuid.New()
}

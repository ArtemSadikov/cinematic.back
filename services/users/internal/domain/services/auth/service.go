package auth

import (
	"cinematic.back/services/users/internal/domain/models/auth"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
}

func (s service) Save(ctx context.Context, auth *auth.UserAuth) error {
	return s.db.WithContext(ctx).
		Session(&gorm.Session{FullSaveAssociations: true}).
		Save(&auth).Error
}

func (s service) FindUserAuthByUserID(ctx context.Context, userID uuid.UUID) (*auth.UserAuth, error) {
	var result auth.UserAuth

	err := s.db.WithContext(ctx).
		Preload("Token").
		Where(auth.UserAuth{UserID: userID}).
		First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, err
}

func NewService(db *gorm.DB) *service {
	return &service{db: db}
}

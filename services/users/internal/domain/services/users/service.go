package users

import (
	"cinematic.back/services/users/internal/domain/models/user"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type service struct {
	db *gorm.DB
}

func New(db *gorm.DB) *service {
	return &service{db: db}
}

func (s *service) Save(ctx context.Context, user *user.User) error {
	return s.db.WithContext(ctx).
		Session(&gorm.Session{FullSaveAssociations: true}).
		Save(user).Error
}

func (s *service) FindByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	result := user.User{ID: id}

	err := s.db.WithContext(ctx).
		Preload(clause.Associations).
		First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *service) DeleteByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	result, err := s.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := s.db.WithContext(ctx).Delete(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

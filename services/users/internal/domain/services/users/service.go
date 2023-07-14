package users

import (
	"cinematic.back/services/users/internal/domain/models/user"
	"cinematic.back/services/users/internal/domain/models/user/profile"
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type service struct {
	db *gorm.DB
}

func (s *service) FindUsersByIds(ctx context.Context, ids ...uuid.UUID) ([]*user.User, error) {
	var users []*user.User

	if err := s.db.WithContext(ctx).Where(ids).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
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

func (s *service) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	var result user.User

	err := s.db.WithContext(ctx).
		InnerJoins("Profile", s.db.Where(&profile.Profile{Email: email})).
		First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &result, nil
		}
		return nil, err
	}

	return &result, nil
}

func New(db *gorm.DB) *service {
	return &service{db: db}
}

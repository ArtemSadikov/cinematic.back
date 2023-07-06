package auth

import (
	"context"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
}

func (s service) AuthByCredentials(ctx context.Context, email, password string) error {

	//TODO implement me
	panic("implement me")
}

func (s service) AuthByAccessToken(ctx context.Context, token string) error {
	//TODO implement me
	panic("implement me")
}

func New(db *gorm.DB) *service {
	return &service{db: db}
}

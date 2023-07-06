package services

import (
	"cinematic.back/services/users/internal/domain/models/user"
	"context"
	"github.com/google/uuid"
)

type UserService interface {
	Save(ctx context.Context, user *user.User) error
	FindByID(ctx context.Context, id uuid.UUID) (*user.User, error)
	DeleteByID(ctx context.Context, id uuid.UUID) (*user.User, error)
}

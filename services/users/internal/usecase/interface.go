package usecase

import (
	"cinematic.back/services/users/internal/domain/models/user"
	"context"
	"github.com/google/uuid"
)

type UsersUseCases interface {
	Create(ctx context.Context, data *user.User) (*user.User, error)
	UpdateByID(ctx context.Context, id uuid.UUID, data *user.User) (*user.User, error)
	DeleteUserByID(ctx context.Context, id uuid.UUID) (*user.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*user.User, error)
}

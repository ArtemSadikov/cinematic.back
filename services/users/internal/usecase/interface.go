package usecase

import (
	"cinematic.back/services/users/internal/domain/models/user"
	"cinematic.back/services/users/internal/domain/services/token"
	"context"
	"github.com/google/uuid"
)

type UsersUseCases interface {
	Create(ctx context.Context, data *user.User) (*user.User, error)
	UpdateByID(ctx context.Context, id uuid.UUID, data *user.User) (*user.User, error)
	DeleteUserByID(ctx context.Context, id uuid.UUID) (*user.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*user.User, error)
	FindUsersByIds(ctx context.Context, ids ...uuid.UUID) ([]*user.User, error)
}

type AuthUseCases interface {
	Register(ctx context.Context, email, username, password string) (*token.Tokens, error)
	AuthByCredentials(ctx context.Context, email, password string) (*token.Tokens, error)
	AuthByAccessToken(ctx context.Context, token string) (*user.User, error)
	RefreshToken(ctx context.Context, token string) (*token.Tokens, error)
	ChangePassword(ctx context.Context, userId uuid.UUID, password string) (*token.Tokens, error)
}

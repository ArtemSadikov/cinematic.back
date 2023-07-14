package services

import (
	"cinematic.back/services/users/internal/domain/models/auth"
	"cinematic.back/services/users/internal/domain/models/user"
	tokenService "cinematic.back/services/users/internal/domain/services/token"
	"context"
	"github.com/google/uuid"
)

type UserService interface {
	Save(ctx context.Context, user *user.User) error
	FindByID(ctx context.Context, id uuid.UUID) (*user.User, error)
	DeleteByID(ctx context.Context, id uuid.UUID) (*user.User, error)
	FindByEmail(ctx context.Context, email string) (*user.User, error)
	FindUsersByIds(ctx context.Context, ids ...uuid.UUID) ([]*user.User, error)
}

type AuthService interface {
	Save(ctx context.Context, auth *auth.UserAuth) error
	FindUserAuthByUserID(ctx context.Context, userID uuid.UUID) (*auth.UserAuth, error)
}

type TokenService interface {
	GenerateTokens(userId, passwordId, tokenId string) (tokenService.Tokens, error)
	ValidateAccessToken(token string) (*tokenService.AccessJWTClaims, error)
	ValidateRefreshToken(token string) (*tokenService.RefreshJWTClaims, error)
}

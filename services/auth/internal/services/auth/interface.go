package auth

import "context"

type Service interface {
	CreateUserAuthData() error
	AuthByCredentials(ctx context.Context, email, password string) error
	AuthByAccessToken(ctx context.Context, token string) error
}

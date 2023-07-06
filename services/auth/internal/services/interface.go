package services

import "context"

type AuthService interface {
	AuthByCredentials(ctx context.Context, email, password string) error
	AuthByAccessToken(ctx context.Context, token string) error
}

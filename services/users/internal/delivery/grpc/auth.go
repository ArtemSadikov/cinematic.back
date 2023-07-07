package grpc

import (
	"cinematic.back/services/users/internal/delivery/grpc/interface/auth"
	"context"
)

func (s *Server) AuthByCredentials(ctx context.Context, in *auth.AuthByCredentialsRequest) (*auth.AuthByCredentialsResponse, error) {
	tokens, err := s.authUC.AuthByCredentials(ctx, in.Credentials.Email, in.Credentials.Password)
	if err != nil {
		return nil, err
	}

	resp := &auth.AuthByCredentialsResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}

	return resp, err
}

func (s *Server) AuthByAccessToken(ctx context.Context, in *auth.AuthByAccessTokenRequest) (*auth.OkResponse, error) {
	if err := s.authUC.AuthByAccessToken(ctx, in.Token); err != nil {
		return nil, err
	}

	return &auth.OkResponse{Message: "OK"}, nil
}

func (s *Server) Register(ctx context.Context, in *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	creds := in.GetCredentials()
	tokens, err := s.authUC.Register(ctx, creds.Email, in.Username, creds.Password)
	if err != nil {
		return nil, err
	}

	resp := &auth.RegisterResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}

	return resp, nil
}

func (s *Server) RefreshToken(ctx context.Context, in *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
	tokens, err := s.authUC.RefreshToken(ctx, in.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	resp := &auth.RefreshTokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}

	return resp, err
}

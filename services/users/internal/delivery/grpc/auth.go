package grpc

import (
	"cinematic.back/api/users/pb"
	"cinematic.back/services/users/internal/delivery/grpc/mappers"
	"context"
)

func (s *Server) AuthByCredentials(ctx context.Context, in *pb.AuthByCredentialsRequest) (*pb.AuthByCredentialsResponse, error) {
	tokens, err := s.authUC.AuthByCredentials(ctx, in.Credentials.Email, in.Credentials.Password)
	if err != nil {
		return nil, err
	}

	resp := &pb.AuthByCredentialsResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}

	return resp, err
}

func (s *Server) AuthByAccessToken(ctx context.Context, in *pb.AuthByAccessTokenRequest) (*pb.AuthByAccessTokenResponse, error) {
	user, err := s.authUC.AuthByAccessToken(ctx, in.Token)
	if err != nil {
		return nil, err
	}

	return &pb.AuthByAccessTokenResponse{
		User: mappers.MakeUserFromModel(*user),
	}, nil
}

func (s *Server) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	creds := in.GetCredentials()
	tokens, err := s.authUC.Register(ctx, creds.Email, in.Username, creds.Password)
	if err != nil {
		return nil, err
	}

	resp := &pb.RegisterResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}

	return resp, nil
}

func (s *Server) RefreshToken(ctx context.Context, in *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	tokens, err := s.authUC.RefreshToken(ctx, in.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	resp := &pb.RefreshTokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}

	return resp, err
}

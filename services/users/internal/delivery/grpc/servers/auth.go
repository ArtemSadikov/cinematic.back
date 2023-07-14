package servers

import (
	"cinematic.back/api/users/pb"
	"cinematic.back/services/users/internal/delivery/grpc/mappers"
	"cinematic.back/services/users/internal/usecase"
	"context"
	"github.com/google/uuid"
)

type AuthServer struct {
	authUC usecase.AuthUseCases
}

func (s *AuthServer) AuthByCredentials(ctx context.Context, in *pb.AuthByCredentialsRequest) (*pb.AuthByCredentialsResponse, error) {
	tokens, err := s.authUC.AuthByCredentials(ctx, in.Credentials.Email, in.Credentials.Password)
	if err != nil {
		return nil, err
	}

	resp := &pb.AuthByCredentialsResponse{
		Tokens: mappers.MakeTokens(tokens),
	}

	return resp, err
}

func (s *AuthServer) AuthByAccessToken(ctx context.Context, in *pb.AuthByAccessTokenRequest) (*pb.AuthByAccessTokenResponse, error) {
	user, err := s.authUC.AuthByAccessToken(ctx, in.Token)
	if err != nil {
		return nil, err
	}

	return &pb.AuthByAccessTokenResponse{
		User: mappers.MakeUserFromModel(*user),
	}, nil
}

func (s *AuthServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	creds := in.GetCredentials()
	tokens, err := s.authUC.Register(ctx, creds.Email, in.Username, creds.Password)
	if err != nil {
		return nil, err
	}

	resp := &pb.RegisterResponse{
		Tokens: mappers.MakeTokens(tokens),
	}

	return resp, nil
}

func (s *AuthServer) RefreshToken(ctx context.Context, in *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	tokens, err := s.authUC.RefreshToken(ctx, in.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	resp := &pb.RefreshTokenResponse{
		Tokens: mappers.MakeTokens(tokens),
	}

	return resp, err
}

func (s *AuthServer) ChangePassword(ctx context.Context, in *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	tokens, err := s.authUC.ChangePassword(ctx, uuid.MustParse(in.UserId), in.Password)
	if err != nil {
		return nil, err
	}

	resp := &pb.ChangePasswordResponse{
		Tokens: mappers.MakeTokens(tokens),
	}

	return resp, nil
}

func NewAuthServer(authUC usecase.AuthUseCases) *AuthServer {
	return &AuthServer{authUC: authUC}
}

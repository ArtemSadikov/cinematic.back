package grpc

import (
	auth "cinematic.back/services/auth/internal/delivery/grpc/interface"
	"context"
)

func (s *Server) AuthByCredentials(ctx context.Context, in *auth.AuthByCredentialsRequest) (*auth.OkResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) AuthByAccessToken(ctx context.Context, in *auth.AuthByAccessTokenRequest) (*auth.OkResponse, error) {
	//TODO implement me
	panic("implement me")
}

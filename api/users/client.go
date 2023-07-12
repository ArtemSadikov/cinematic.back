package users

import (
	"cinematic.back/api/users/pb"
	"cinematic.back/pkg/grpc"
	"context"
)

type Client interface {
	AuthByAccessToken(ctx context.Context, in *pb.AuthByAccessTokenRequest) (*pb.AuthByAccessTokenResponse, error)
	AuthByCredentials(ctx context.Context, in *pb.AuthByCredentialsRequest) (*pb.AuthByCredentialsResponse, error)
	Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error)
	RefreshToken(ctx context.Context, in *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error)
	FindUserById(ctx context.Context, in *pb.FindUserByIdRequest) (*pb.FindUserByIdResponse, error)
}

type client struct {
	uClient pb.UsersServiceClient
	aClient pb.AuthServiceClient
}

func (c client) FindUserById(ctx context.Context, in *pb.FindUserByIdRequest) (*pb.FindUserByIdResponse, error) {
	return c.uClient.FindUserById(ctx, in)
}

func (c client) AuthByAccessToken(ctx context.Context, in *pb.AuthByAccessTokenRequest) (*pb.AuthByAccessTokenResponse, error) {
	return c.aClient.AuthByAccessToken(ctx, in)
}

func (c client) AuthByCredentials(ctx context.Context, in *pb.AuthByCredentialsRequest) (*pb.AuthByCredentialsResponse, error) {
	return c.aClient.AuthByCredentials(ctx, in)
}

func (c client) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return c.aClient.Register(ctx, in)
}

func (c client) RefreshToken(ctx context.Context, in *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	return c.aClient.RefreshToken(ctx, in)
}

func NewClient(host string) Client {
	grpcClient := grpc.NewClient(host)
	conn := grpcClient.OpenConn(context.Background())

	uClient := pb.NewUsersServiceClient(conn)
	aClient := pb.NewAuthServiceClient(conn)

	return &client{uClient, aClient}
}

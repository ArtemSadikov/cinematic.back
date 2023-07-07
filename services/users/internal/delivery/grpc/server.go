package grpc

import (
	"cinematic.back/services/users/internal/usecase"
)

type Server struct {
	usersUC usecase.UsersUseCases
	authUC  usecase.AuthUseCases
}

func New(
	usersUC usecase.UsersUseCases,
	authUC usecase.AuthUseCases,
) *Server {
	return &Server{
		usersUC: usersUC,
		authUC:  authUC,
	}
}

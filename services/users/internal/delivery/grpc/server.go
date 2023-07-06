package grpc

import (
	"cinematic.back/services/users/internal/usecase"
)

type Server struct {
	uUseCase usecase.UsersUseCases
}

func New(service usecase.UsersUseCases) *Server {
	return &Server{uUseCase: service}
}

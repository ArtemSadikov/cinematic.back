package main

import (
	"cinematic.back/pkg/provider/database"
	"cinematic.back/services/users/internal/delivery/grpc"
	"cinematic.back/services/users/internal/delivery/grpc/interface/auth"
	"cinematic.back/services/users/internal/delivery/grpc/interface/users"
	container2 "cinematic.back/services/users/internal/infrastructure/container"
	"cinematic.back/services/users/internal/usecase"
	"context"
	grpc2 "google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	container, err := container2.New()
	if err != nil {
		log.Fatalf("%e", err)
	}

	err = container.Invoke(func(
		provider database.Provider,
		uUseCase usecase.UsersUseCases,
		aUseCase usecase.AuthUseCases,
	) error {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if err := provider.Migrate(ctx); err != nil {
			return err
		}

		var errCh chan error

		server := grpc.New(uUseCase, aUseCase)
		listener, err := net.Listen("tcp", "localhost:3000")
		if err != nil {
			log.Fatalf("%e", err)
		}

		grpcServer := grpc2.NewServer()

		users.RegisterUsersServiceServer(grpcServer, server)
		auth.RegisterAuthServiceServer(grpcServer, server)

		go func() {
			log.Println("server is listening on 3000 port")
			err := grpcServer.Serve(listener)
			errCh <- err
		}()

		return <-errCh
	})
	if err != nil {
		log.Fatalf("%e", err)
	}
}

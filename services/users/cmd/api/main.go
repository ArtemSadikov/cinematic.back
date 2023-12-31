package main

import (
	"cinematic.back/api/users/pb"
	"cinematic.back/pkg/provider/database"
	"cinematic.back/services/users/internal/delivery/grpc/servers"
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

	err = container.Provide(func(
		aUseCase usecase.AuthUseCases,
	) *servers.AuthServer {
		return servers.NewAuthServer(aUseCase)
	})
	if err != nil {
		log.Fatalf("%e", err)
	}

	err = container.Provide(func(
		uUseCase usecase.UsersUseCases,
	) *servers.UserServer {
		return servers.NewUserServer(uUseCase)
	})
	if err != nil {
		log.Fatalf("%e", err)
	}

	err = container.Invoke(func(
		provider database.Provider,
		aServer *servers.AuthServer,
		uServer *servers.UserServer,
	) error {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if err := provider.Migrate(ctx); err != nil {
			return err
		}

		var errCh chan error

		listener, err := net.Listen("tcp", "localhost:3001")
		if err != nil {
			log.Fatalf("%e", err)
		}

		grpcServer := grpc2.NewServer()

		pb.RegisterUsersServiceServer(grpcServer, uServer)
		pb.RegisterAuthServiceServer(grpcServer, aServer)

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

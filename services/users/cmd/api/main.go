package main

import (
	"cinematic.back/pkg/provider/database"
	"cinematic.back/services/users/internal/delivery/grpc"
	users2 "cinematic.back/services/users/internal/delivery/grpc/interface"
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

	err = container.Invoke(func(provider database.Provider, service usecase.UsersUseCases) error {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if err := provider.Migrate(ctx); err != nil {
			return err
		}

		var errCh chan error

		server := grpc.New(service)
		listener, err := net.Listen("tcp", "localhost:3000")
		if err != nil {
			log.Fatalf("%e", err)
		}

		grpcServer := grpc2.NewServer()
		users2.RegisterUsersServiceServer(grpcServer, server)
		go func() {
			log.Println("server is listening on 3000 port")
			err := grpcServer.Serve(listener)
			errCh <- err
		}()

		if err := <-errCh; err != nil {
			log.Fatalf("%e", err)
			return err
		}

		return nil
	})
	if err != nil {
		log.Fatalf("%e", err)
	}
}

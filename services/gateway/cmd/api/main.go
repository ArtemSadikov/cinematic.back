package main

import (
	"cinematic.back/api/users"
	"cinematic.back/services/gateway/internal/api/graphql"
	"cinematic.back/services/gateway/internal/infrastructure/container"
	"cinematic.back/services/gateway/internal/services/user"
	"log"
)

func main() {
	c, err := container.New()
	if err != nil {
		log.Fatal(err)
	}

	err = c.Provide(func(uClient users.Client, uService user.Service) *graphql.Server {
		return graphql.NewServer(uClient, uService)
	})
	if err != nil {
		log.Fatal(err)
	}

	err = c.Invoke(func(server *graphql.Server) {
		var errs chan error

		go func() {
			log.Println("Gateway server started at localhost:3000")
			errs <- server.Start()
		}()

		log.Println(<-errs)
	})
	if err != nil {
		log.Fatal(err)
	}
}

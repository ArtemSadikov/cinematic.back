package container

import (
	"cinematic.back/api/users"
	"cinematic.back/services/gateway/internal/api/graphql/loaders/user"
	uService "cinematic.back/services/gateway/internal/services/user"
	"go.uber.org/dig"
)

func New() (*dig.Container, error) {
	container := dig.New()

	var err error
	err = container.Provide(func() users.Client {
		return users.NewClient("localhost:3001")
	})
	if err != nil {
		return nil, err
	}

	err = container.Provide(func() uService.Service {
		return uService.NewService()
	})
	if err != nil {
		return nil, err
	}

	err = container.Provide(func(uClient users.Client, uService uService.Service) *user.Loader {
		return user.NewLoader(uClient, uService)
	})
	if err != nil {
		return nil, err
	}

	return container, nil
}

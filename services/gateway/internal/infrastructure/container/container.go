package container

import (
	"cinematic.back/api/users"
	"cinematic.back/services/gateway/internal/services/user"
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

	err = container.Provide(func() user.Service {
		return user.NewService()
	})
	if err != nil {
		return nil, err
	}

	return container, nil
}

package container

import (
	"cinematic.back/pkg/postgres"
	"cinematic.back/pkg/provider/database"
	"cinematic.back/services/users/internal/domain/services"
	"cinematic.back/services/users/internal/domain/services/users"
	postgres2 "cinematic.back/services/users/internal/infrastructure/providers/postgres"
	"cinematic.back/services/users/internal/usecase"
	uUseCase "cinematic.back/services/users/internal/usecase/users"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func New(opts ...dig.Option) (*dig.Container, error) {
	c := dig.New(opts...)

	err := c.Provide(func() *postgres.Postgres {
		return postgres.New(
			"localhost",
			5432,
			"cinematic",
			"password",
			"cinematic_users",
			false,
			"pgx",
		)
	})
	if err != nil {
		return nil, err
	}

	err = c.Provide(func(pg *postgres.Postgres) database.Provider {
		return postgres2.NewProvider(pg)
	})
	if err != nil {
		return nil, err
	}

	err = c.Provide(func(provider database.Provider) (*gorm.DB, error) {
		if err := provider.Connect(); err != nil {
			return nil, err
		}

		return provider.DB(), nil
	})
	if err != nil {
		return nil, err
	}

	err = c.Provide(func(db *gorm.DB) services.UserService {
		return users.New(db)
	})
	if err != nil {
		return nil, err
	}

	err = c.Provide(func(uServices services.UserService) usecase.UsersUseCases {
		return uUseCase.NewUseCase(uServices)
	})

	return c, err
}

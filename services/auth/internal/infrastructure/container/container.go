package container

import (
	"cinematic.back/pkg/postgres"
	pgPr "cinematic.back/services/auth/internal/infrastructure/providers/database/postgres"
	"go.uber.org/dig"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"time"
)

func New(opts ...dig.Option) (*dig.Container, error) {
	container := dig.New(opts...)

	err := container.Provide(func() *postgres.Postgres {
		return postgres.New(
			"localhost",
			5432,
			"cinematic_auth",
			"password",
			"cinematic_auth",
			false,
			"pgx",
		)
	})
	if err != nil {
		return nil, err
	}

	err = container.Provide(func(pg *postgres.Postgres) *pgPr.Provider {
		return pgPr.New(pg)
	})
	if err != nil {
		return nil, err
	}

	err = container.Provide(func(provider *pgPr.Provider) (*gorm.DB, error) {
		db, err := provider.Connect()
		if err != nil {
			return nil, err
		}

		return db, err
	})
	if err != nil {
		return nil, err
	}

	err = container.Invoke(func(provider *pgPr.Provider) error {
		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
		defer cancel()

		return provider.Migrate(ctx)
	})
	if err != nil {
		return nil, err
	}

	return container, nil
}

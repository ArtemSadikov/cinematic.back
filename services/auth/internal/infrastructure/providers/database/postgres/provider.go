package postgres

import (
	"cinematic.back/pkg/postgres"
	"cinematic.back/services/auth/internal/domain/auth"
	"context"
	pgd "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Provider struct {
	pgConn postgres.DBConnection
	db     *gorm.DB
}

func New(pg *postgres.Postgres) *Provider {
	conn := postgres.DBConnection{Postgres: pg}

	return &Provider{pgConn: conn}
}

func (p *Provider) Connect() (*gorm.DB, error) {
	dbConnection, err := p.pgConn.GetDBConnection()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(pgd.New(pgd.Config{Conn: dbConnection}))
	if err != nil {
		return nil, err
	}
	return db, err
}

func (p *Provider) Migrate(ctx context.Context) error {
	db, err := p.db.DB()
	if err != nil {
		return err
	}

	if _, err := db.QueryContext(ctx, "CREATE EXTENTION IF NOT EXISTS \"uuid-ossp\";"); err != nil {
		return err
	}

	migrator := p.db.Migrator()
	if err := migrator.AutoMigrate(auth.UserAuth{}); err != nil {
		return err
	}
	return nil
}

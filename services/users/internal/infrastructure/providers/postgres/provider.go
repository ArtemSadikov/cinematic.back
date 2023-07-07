package postgres

import (
	"cinematic.back/pkg/postgres"
	"cinematic.back/services/users/internal/domain/models/auth"
	"cinematic.back/services/users/internal/domain/models/auth/token"
	"cinematic.back/services/users/internal/domain/models/user"
	"cinematic.back/services/users/internal/domain/models/user/profile"
	"context"
	postgres2 "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Provider struct {
	dbConn *postgres.DBConnection
	db     *gorm.DB
}

func NewProvider(pg *postgres.Postgres) *Provider {
	dbConn := &postgres.DBConnection{Postgres: pg}
	return &Provider{dbConn: dbConn}
}

func (p *Provider) Connect() error {
	db, err := p.dbConn.GetDBConnection()
	if err != nil {
		return err
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      false,
			Colorful:                  true,
		},
	)

	p.db, err = gorm.Open(
		postgres2.New(postgres2.Config{Conn: db}),
		&gorm.Config{
			Logger: newLogger,
		},
	)
	return err
}

func (p *Provider) Migrate(ctx context.Context) error {
	db, err := p.db.DB()
	if err != nil {
		return err
	}

	if _, err := db.QueryContext(ctx, "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"); err != nil {
		return err
	}

	migrator := p.db.Migrator()
	err = migrator.AutoMigrate(
		&user.User{},
		&profile.Profile{},
		&auth.UserAuth{},
		&token.AuthToken{},
	)
	if err != nil {
		return err
	}
	return nil
}

func (p *Provider) DB() *gorm.DB {
	return p.db
}

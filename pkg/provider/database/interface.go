package database

import (
	"context"
	"gorm.io/gorm"
)

type Provider interface {
	Connect() error
	Migrate(ctx context.Context) error
	DB() *gorm.DB
}

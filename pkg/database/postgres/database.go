package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
	"users/pkg/config"
	"users/pkg/database"
)

type postgresDatabase struct {
	db *gorm.DB
}

var (
	once       sync.Once
	onceError  error
	dbInstance *postgresDatabase
)

func NewDatabase(cfg config.Storage) (database.Database, error) {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
			cfg.GetHost(),
			cfg.GetUser(),
			cfg.GetPassword(),
			cfg.GetDatabase(),
			cfg.GetPort(),
			"disable",
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			onceError = err

			return
		}

		dbInstance = &postgresDatabase{db: db}
	})

	if onceError != nil {
		return nil, onceError
	}

	return dbInstance, nil
}

func (p *postgresDatabase) GetInstance() *gorm.DB {
	return p.db
}

func (p *postgresDatabase) Migrate(dst ...interface{}) error {
	return p.db.AutoMigrate(dst...)
}

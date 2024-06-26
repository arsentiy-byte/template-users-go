package app

import (
	"users/internal/config"
	"users/pkg/database"
	"users/pkg/database/postgres"
)

type provider struct {
	cfg *config.Config
	db  database.Database
}

func newProvider(cfg *config.Config) *provider {
	return &provider{
		cfg: cfg,
	}
}

func (p *provider) Database() database.Database {
	if p.db == nil {
		db, err := postgres.NewDatabase(p.cfg.Storage)
		if err != nil {
			panic(err)
		}

		p.db = db
	}

	return p.db
}

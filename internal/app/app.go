package app

import (
	"context"
	"fmt"
	"users/internal/config"
	"users/internal/entities"
)

type App struct {
	cfg      *config.Config
	provider *provider
}

func New(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	fmt.Println(a.provider.Database())

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.migrateDatabase,
	}

	for _, init := range inits {
		if err := init(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	a.cfg = cfg

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.provider = newProvider(a.cfg)

	return nil
}

func (a *App) migrateDatabase(_ context.Context) error {
	return a.provider.Database().Migrate(&entities.User{})
}

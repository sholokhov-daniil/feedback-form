package repository

import (
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/sholokhov-daniil/feedback-form/internal/config"
)

var (
	instance *Container
	once 	 sync.Once
)

type Container struct {
	Database *sqlx.DB
	Config config.Config
}

func ServiceContainer() *Container {
	once.Do(func()  {
		instance = &Container{};
	})

	return instance
}
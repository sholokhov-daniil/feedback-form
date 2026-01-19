package repository

import (
	"sync"

	"gorm.io/gorm"
	"github.com/sholokhov-daniil/feedback-form/internal/config"
)

var (
	instance *Container
	once 	 sync.Once
)

type Container struct {
	Database *gorm.DB
	Config config.Config
}

func ServiceContainer() *Container {
	once.Do(func()  {
		instance = &Container{};
	})

	return instance
}
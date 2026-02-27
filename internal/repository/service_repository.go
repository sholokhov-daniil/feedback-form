package repository

import (
	"sync"

	"github.com/sholokhov-daniil/feedback-form/internal/config"
	"gorm.io/gorm"
)

var (
	instance *Container
	once     sync.Once
)

type Container struct {
	Database *gorm.DB
	Config   config.Config
}

func ServiceContainer() *Container {
	once.Do(func() {
		instance = &Container{}
	})

	return instance
}

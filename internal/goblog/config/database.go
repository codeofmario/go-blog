package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(settings *Settings) *gorm.DB {
	db, err := gorm.Open(postgres.Open(settings.PostgresDsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect with database.")
	}

	return db
}

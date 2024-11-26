package database

import (
	"go_carwash/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConn(config *config.Env) *gorm.DB {

	db, err := gorm.Open(postgres.Open(config.DBSOURCE), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	return db
}

package db

import (
	"github.com/laidingqing/sokr/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	db, dbErr := gorm.Open(postgres.Open(cfg.Database.DSN()), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate()

	return db, dbErr
}

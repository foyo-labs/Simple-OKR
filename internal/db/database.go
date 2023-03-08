package db

import (
	"github.com/laidingqing/sokr/internal/config"
	"github.com/laidingqing/sokr/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	db, dbErr := gorm.Open(postgres.Open(cfg.Database.DSN()), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(
		new(entity.User),
		new(entity.Group),
		new(entity.UserGroup),
		new(entity.Objective),
		new(entity.KeyResult),
		new(entity.UserObjective),
		new(entity.GroupObjective),
	)

	return db, dbErr
}

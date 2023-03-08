package repository

import (
	"context"

	"github.com/laidingqing/sokr/internal/entity"
	"github.com/laidingqing/sokr/internal/schema"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IUserObjectiveRepository interface {
	Add(ctx context.Context, item schema.UserObjective) error
}

type UserObjectiveRepository struct {
	DB *gorm.DB
}

var _ IUserObjectiveRepository = &UserObjectiveRepository{}

func NewUserObjectiveRepository(dbConn *gorm.DB) IUserObjectiveRepository {
	usoRep := UserObjectiveRepository{DB: dbConn}
	return &usoRep
}

func (a *UserObjectiveRepository) Add(ctx context.Context, item schema.UserObjective) error {
	sitem := entity.SchemaUserObjective(item)
	db := entity.GetUserObjectiveDB(ctx, a.DB)
	result := db.Create(sitem.ToUserObjective())
	return errors.WithStack(result.Error)
}

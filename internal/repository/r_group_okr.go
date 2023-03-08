package repository

import (
	"context"

	"github.com/laidingqing/sokr/internal/entity"
	"github.com/laidingqing/sokr/internal/schema"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IGroupObjectiveRepository interface {
	Add(ctx context.Context, item schema.GroupObjective) error
}

type GroupObjectiveRepository struct {
	DB *gorm.DB
}

var _ IGroupObjectiveRepository = &GroupObjectiveRepository{}

func NewGroupObjectiveRepository(dbConn *gorm.DB) IGroupObjectiveRepository {
	groupOkrRep := GroupObjectiveRepository{DB: dbConn}
	return &groupOkrRep
}

func (a *GroupObjectiveRepository) Add(ctx context.Context, item schema.GroupObjective) error {
	sitem := entity.SchemaGroupObjective(item)
	db := entity.GetGroupObjectiveDB(ctx, a.DB)
	result := db.Create(sitem.ToGroupObjective())
	return errors.WithStack(result.Error)
}

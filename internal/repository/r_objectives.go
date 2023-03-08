package repository

import (
	"context"
	"time"

	"github.com/laidingqing/sokr/internal/entity"
	"github.com/laidingqing/sokr/internal/schema"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IObjectiveRepository interface {
	Add(ctx context.Context, item schema.Objective) error
}

type ObjectiveRepository struct {
	DB *gorm.DB
}

var _ IObjectiveRepository = &ObjectiveRepository{}

func NewObjectiveRepository(dbConn *gorm.DB) IObjectiveRepository {
	objectiveRep := ObjectiveRepository{DB: dbConn}
	return &objectiveRep
}

func (a *ObjectiveRepository) Add(ctx context.Context, objective schema.Objective) error {
	objective.Created = uint64(time.Now().UnixNano())
	objective.Actived = schema.ObjectiveActived
	sitem := entity.SchemaObjective(objective)
	db := entity.GetObjectiveDB(ctx, a.DB)
	result := db.Create(sitem.ToObjective())
	return errors.WithStack(result.Error)
}

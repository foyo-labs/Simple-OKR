package repository

import (
	"context"

	"github.com/laidingqing/sokr/internal/entity"
	"github.com/laidingqing/sokr/internal/schema"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ICycleRepository interface {
	Add(ctx context.Context, item schema.Cycle) error
	Find(ctx context.Context, query schema.CycleQueryParam) ([]*schema.Cycle, error)
}

type CycleRepository struct {
	DB *gorm.DB
}

var _ ICycleRepository = &CycleRepository{}

func NewCycleRepository(dbConn *gorm.DB) ICycleRepository {
	companyRep := CycleRepository{DB: dbConn}
	return &companyRep
}

func (a *CycleRepository) Add(ctx context.Context, item schema.Cycle) error {
	sitem := entity.SchemaCycle(item)
	db := entity.GetCycleDB(ctx, a.DB)
	result := db.Create(sitem.ToCycle())
	return errors.WithStack(result.Error)
}

func (a *CycleRepository) Find(ctx context.Context, param schema.CycleQueryParam) ([]*schema.Cycle, error) {
	var items entity.Cycles
	db := entity.GetCycleDB(ctx, a.DB)
	result := db.Where("id=?", param.ID).Find(&items)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}
	return items.ToSchemaCycles(), nil
}

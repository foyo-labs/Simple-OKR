package repository

import (
	"context"

	"github.com/laidingqing/sokr/internal/entity"
	"github.com/laidingqing/sokr/internal/schema"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IKeyResultRepository interface {
	Add(ctx context.Context, items schema.KeyResults) error
}

type KeyResultRepository struct {
	DB *gorm.DB
}

var _ IKeyResultRepository = &KeyResultRepository{}

func NewKeyResultRepository(dbConn *gorm.DB) IKeyResultRepository {
	krRep := KeyResultRepository{DB: dbConn}
	return &krRep
}

func (a *KeyResultRepository) Add(ctx context.Context, items schema.KeyResults) error {
	db := entity.GetKeyResultDB(ctx, a.DB)
	sitems := entity.SchemaKeyResults(items)
	result := db.CreateInBatches(sitems.ToKeyResults(), 50)
	return errors.WithStack(result.Error)
}

package repository

import (
	"github.com/laidingqing/sokr/internal/entity"
	"gorm.io/gorm"
)

type IKeyResultRepository interface {
	Add(kr []entity.KeyResult) error
}

type KeyResultRepository struct {
	DbConn *gorm.DB
}

var _ IKeyResultRepository = &KeyResultRepository{}

func NewKeyResultRepository(dbConn *gorm.DB) IKeyResultRepository {
	krRep := KeyResultRepository{DbConn: dbConn}
	return &krRep
}

func (rep *KeyResultRepository) Add(krs []entity.KeyResult) error {

	return nil
}

package repository

import (
	"github.com/laidingqing/sokr/internal/entity"
	"gorm.io/gorm"
)

type ICycleRepository interface {
	Add(cycle entity.Cycle) error
}

type CycleRepository struct {
	DbConn *gorm.DB
}

var _ ICycleRepository = &CycleRepository{}

func NewCycleRepository(dbConn *gorm.DB) ICycleRepository {
	companyRep := CycleRepository{DbConn: dbConn}
	return &companyRep
}

func (rep *CycleRepository) Add(cycle entity.Cycle) error {

	return nil
}

package repository

import (
	"github.com/laidingqing/sokr/internal/entity"
	"gorm.io/gorm"
)

type IObjectiveRepository interface {
	Add(account entity.Objective) error
}

type ObjectiveRepository struct {
	DbConn *gorm.DB
}

var _ IObjectiveRepository = &ObjectiveRepository{}

func NewObjectiveRepository(dbConn *gorm.DB) IObjectiveRepository {
	objectiveRep := ObjectiveRepository{DbConn: dbConn}
	return &objectiveRep
}

func (rep *ObjectiveRepository) Add(objective entity.Objective) error {

	return nil
}

package repository

import (
	"github.com/laidingqing/sokr/internal/entities"
	"gorm.io/gorm"
)

type IObjectiveRepository interface {
	Add(account entities.Objective) error
}

type ObjectiveRepository struct {
	DbConn *gorm.DB
}

var _ IObjectiveRepository = &ObjectiveRepository{}

func NewObjectiveRepository(dbConn *gorm.DB) IObjectiveRepository {
	objectiveRep := ObjectiveRepository{DbConn: dbConn}
	return &objectiveRep
}

func (rep *ObjectiveRepository) Add(objective entities.Objective) error {

	return nil
}

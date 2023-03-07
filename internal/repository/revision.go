package repository

import (
	"github.com/laidingqing/sokr/internal/entity"
	"gorm.io/gorm"
)

type IRevisionRepository interface {
	Add(account entity.Revision) error
}

type RevisionRepository struct {
	DbConn *gorm.DB
}

var _ IRevisionRepository = &RevisionRepository{}

func NewRevisionRepository(dbConn *gorm.DB) IRevisionRepository {
	revisionRep := RevisionRepository{DbConn: dbConn}
	return &revisionRep
}

func (rep *RevisionRepository) Add(objective entity.Revision) error {

	return nil
}

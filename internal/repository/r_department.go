package repository

import (
	"github.com/laidingqing/sokr/internal/entity"
	"gorm.io/gorm"
)

type IDepartmentRepository interface {
	Add(department entity.Department) error
}

type DepartmentRepository struct {
	DbConn *gorm.DB
}

var _ IDepartmentRepository = &DepartmentRepository{}

func NewDepartmentRepository(dbConn *gorm.DB) IDepartmentRepository {
	departmentRep := DepartmentRepository{DbConn: dbConn}
	return &departmentRep
}

func (rep *DepartmentRepository) Add(department entity.Department) error {

	return nil
}

package repository

import (
	"github.com/laidingqing/sokr/internal/entity"
	"gorm.io/gorm"
)

type ICompanyRepository interface {
	Add(company entity.Company) error
	Get(id string) (*entity.Company, error)
}

type CompanyRepository struct {
	DbConn *gorm.DB
}

var _ ICompanyRepository = &CompanyRepository{}

func NewCompanyRepository(dbConn *gorm.DB) ICompanyRepository {
	companyRep := CompanyRepository{DbConn: dbConn}
	return &companyRep
}

func (rep *CompanyRepository) Add(obj entity.Company) error {

	return nil
}

func (rep *CompanyRepository) Get(id string) (*entity.Company, error) {

	return nil, nil
}

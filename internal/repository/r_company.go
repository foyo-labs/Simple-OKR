package repository

import (
	"context"
	"time"

	"github.com/laidingqing/sokr/internal/entity"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ICompanyRepository interface {
	Add(ctx context.Context, company entity.Company) error
	Get(ctx context.Context, id string) (*entity.Company, error)
}

type CompanyRepository struct {
	DB *gorm.DB
}

var _ ICompanyRepository = &CompanyRepository{}

func NewCompanyRepository(dbConn *gorm.DB) ICompanyRepository {
	companyRep := CompanyRepository{DB: dbConn}
	return &companyRep
}

func (a *CompanyRepository) Add(ctx context.Context, item entity.Company) error {
	sitem := entity.SchemaCompany(item)
	sitem.Created = uint64(time.Now().UnixNano())
	db := entity.GetCompanyDB(ctx, a.DB)
	result := db.Create(sitem.ToCompany())
	return errors.WithStack(result.Error)
}

func (a *CompanyRepository) Get(ctx context.Context, id string) (*entity.Company, error) {

	return nil, nil
}

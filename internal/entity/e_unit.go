package entity

import (
	"context"

	"github.com/laidingqing/sokr/internal/schema"
	"github.com/laidingqing/sokr/pkg/util/structure"
	"gorm.io/gorm"
)

type Department struct {
	ID       int32  `gorm:"column:id;not null;primaryKey"`
	Name     string `gorm:"column:name;not null"`
	ParentID int32  `gorm:"column:parent_id;not null"`
	Created  uint64 `gorm:"column:created;"`
}

func GetCompanyDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(Company))
}

type Company struct {
	ID      int32  `gorm:"column:id;not null;primaryKey"`
	Name    string `gorm:"column:name;not null"`
	Abbr    string `gorm:"column:abbr;"`
	Created uint64 `gorm:"column:created;"`
}

func (a *Company) ToSchemaCompany() *schema.Company {
	item := new(schema.Company)
	structure.Copy(a, item)
	return item
}

type SchemaCompany schema.Company

func (a SchemaCompany) ToCompany() *Company {
	item := new(Company)
	structure.Copy(a, item)
	return item
}

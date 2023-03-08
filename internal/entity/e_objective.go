package entity

import (
	"context"

	"github.com/laidingqing/sokr/internal/schema"
	"github.com/laidingqing/sokr/pkg/util/structure"
	"gorm.io/gorm"
)

func GetObjectiveDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(Objective))
}

// Objective 目标
type Objective struct {
	ID          int32                  `gorm:"column:id;not null;primaryKey"`
	Name        string                 `gorm:"column:name;not null"`
	Description string                 `gorm:"column:discription"`
	Actived     schema.ObjectiveStatus `gorm:"column:actived;not null"`
	Sequence    int32                  `gorm:"column:sequence"`
	ParentID    int32                  `gorm:"column:parent_id"`
	Created     uint64                 `gorm:"column:created"`
}

func (a Objective) ToSchemaObjective() *schema.Objective {
	item := new(schema.Objective)
	structure.Copy(a, item)
	return item
}

type SchemaObjective schema.Objective

func (a SchemaObjective) ToObjective() *Objective {
	item := new(Objective)
	structure.Copy(a, item)
	return item
}

package entity

import (
	"context"

	"github.com/laidingqing/sokr/internal/schema"
	"github.com/laidingqing/sokr/pkg/util/structure"
	"gorm.io/gorm"
)

func GetKeyResultDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(KeyResult))
}

type KeyResult struct {
	ID           string  `gorm:"column:id;not null;primaryKey"`
	ObjectiveID  string  `gorm:"column:objective_id;not null"`
	Name         string  `gorm:"column:name;not null"`
	StartValue   float64 `gorm:"column:start_value;not null;default:0;"`
	TargetValue  float64 `gorm:"column:target_value;not null;default:0;"`
	CurrentValue float64 `gorm:"column:current_value;not null;default:0;"`
	Sequence     int32   `gorm:"column:sequence;not null"`
	Created      uint64  `gorm:"column:created;index;"`
}

type SchemaKeyResult schema.KeyResult

func (a SchemaKeyResult) ToKeyResult() *KeyResult {
	item := new(KeyResult)
	structure.Copy(a, item)
	return item
}

type SchemaKeyResults []*schema.KeyResult

func (a SchemaKeyResults) ToKeyResults() []*KeyResult {
	list := make([]*KeyResult, len(a))
	for i, item := range a {
		sitem := SchemaKeyResult(*item)
		structure.Copy(item, sitem)
		list[i] = sitem.ToKeyResult()
	}
	return list
}

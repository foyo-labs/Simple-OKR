package entity

import (
	"context"

	"github.com/laidingqing/sokr/internal/schema"
	"github.com/laidingqing/sokr/pkg/util/structure"
	"gorm.io/gorm"
)

func GetCycleDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(Cycle))
}

// Cycle OKR 周期
// 逻辑控制团队及下属周期一致，不重叠。
type Cycle struct {
	ID      string `gorm:"column:id;not null;primaryKey"`
	Name    string `gorm:"column:name;not null"`
	StartAt uint64 `gorm:"column:start_at;not null"`
	EndAt   uint64 `gorm:"column:end_at;not null"`
}

func (a Cycle) ToSchemaCycle() *schema.Cycle {
	item := new(schema.Cycle)
	structure.Copy(a, item)
	return item
}

type SchemaCycle schema.Cycle

func (a SchemaCycle) ToCycle() *Cycle {
	item := new(Cycle)
	structure.Copy(a, item)
	return item
}

type Cycles []*Cycle

func (a Cycles) ToSchemaCycles() []*schema.Cycle {
	list := make([]*schema.Cycle, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaCycle()
	}
	return list
}

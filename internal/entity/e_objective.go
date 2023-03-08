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
	ID          string                 `gorm:"column:id;not null;primaryKey"`
	Name        string                 `gorm:"column:name;not null"`
	Description string                 `gorm:"column:discription"`
	Actived     schema.ObjectiveStatus `gorm:"column:actived;not null"`
	Sequence    int32                  `gorm:"column:sequence"`
	ParentID    int32                  `gorm:"column:parent_id"`
	Created     uint64                 `gorm:"column:created"`
	Updated     uint64                 `gorm:"column:updated"`
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

func GetUserObjectiveDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(UserObjective))
}

// UserObjective 成员OKR关联
type UserObjective struct {
	ID          string `gorm:"column:id;not null;primaryKey"`
	UserID      string `gorm:"column:user_id;not null"`
	ObjectiveID string `gorm:"column:objective_id;not null"`
	CycleID     string `gorm:"column:cycle_id;not null"`
}

func (a UserObjective) ToSchemaUserObjective() *schema.UserObjective {
	item := new(schema.UserObjective)
	structure.Copy(a, item)
	return item
}

type SchemaUserObjective schema.UserObjective

func (a SchemaUserObjective) ToUserObjective() *UserObjective {
	item := new(UserObjective)
	structure.Copy(a, item)
	return item
}

func GetGroupObjectiveDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(GroupObjective))
}

// GroupObjective 团队OKR关联
type GroupObjective struct {
	ID          string `gorm:"column:id;not null;primaryKey"`
	GroupID     string `gorm:"column:group_id;not null"`
	ObjectiveID string `gorm:"column:objective_id;not null"`
	CycleID     string `gorm:"column:cycle_id;not null"`
}

func (a GroupObjective) ToSchemaGroupObjective() *schema.GroupObjective {
	item := new(schema.GroupObjective)
	structure.Copy(a, item)
	return item
}

type SchemaGroupObjective schema.GroupObjective

func (a SchemaGroupObjective) ToGroupObjective() *GroupObjective {
	item := new(GroupObjective)
	structure.Copy(a, item)
	return item
}

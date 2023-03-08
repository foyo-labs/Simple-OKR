package entity

import (
	"context"

	"github.com/laidingqing/sokr/internal/schema"
	"github.com/laidingqing/sokr/pkg/util/structure"
	"gorm.io/gorm"
)

func GetGroupDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(Group))
}

type Group struct {
	ID           string `gorm:"column:id;not null;primaryKey"`
	Name         string `gorm:"column:name;not null"`
	ParentID     string `gorm:"column:parent_id;not null"`
	LevelNum     string `gorm:"column:level_num;size:20;index;"` //分级设置：LIKE　0.1.3
	OrderNum     int64  `gorm:"column:order_num;index;"`         //排序
	CurrentLevel uint64 `gorm:"column:curr_lev;index;"`          //当前级别,第一级为0，依次增长
	Created      uint64 `gorm:"column:created;"`
}

func (a Group) ToSchemaGroup() *schema.Group {
	item := new(schema.Group)
	structure.Copy(a, item)
	return item
}

type SchemaGroup schema.Group

func (a SchemaGroup) ToGroup() *Group {
	item := new(Group)
	structure.Copy(a, item)
	return item
}

type Groups []*Group

func (a Groups) ToSchemaGroups() []*schema.Group {
	list := make([]*schema.Group, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaGroup()
	}
	return list
}

func GetUserGroupDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(UserGroup))
}

type UserGroup struct {
	UserID  string `gorm:"column:user_id;not null;primaryKey"`
	GroupID string `gorm:"column:group_id;not null;primaryKey"`
}

type SchemaUserGroup schema.UserGroup

func (a SchemaUserGroup) ToUserGroup() *UserGroup {
	item := new(UserGroup)
	structure.Copy(a, item)
	return item
}

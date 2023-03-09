package entity

import (
	"context"

	"github.com/laidingqing/sokr/internal/schema"
	"github.com/laidingqing/sokr/pkg/util/structure"
	"gorm.io/gorm"
)

func GetUserDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(User))
}

type User struct {
	ID       string      `gorm:"column:id;not null;primaryKey"`
	Email    string      `gorm:"column:email;not null"`
	Password string      `gorm:"column:password;not null"`
	Name     string      `gorm:"column:name;"`
	Status   int         `gorm:"column:status;index;default:0;not null;"` // 状态(1:启用 2:停用)
	Role     schema.Role `gorm:"column:role;default:1;"`                  //角色, 0
	Created  uint64      `gorm:"column:created;not null"`
}

// ToSchemaUser 转换为用户对象
func (a User) ToSchemaUser() *schema.User {
	item := new(schema.User)
	structure.Copy(a, item)
	return item
}

// SchemaUser 用户对象
type SchemaUser schema.User

// ToUser 转换为用户实体
func (a SchemaUser) ToUser() *User {
	item := new(User)
	structure.Copy(a, item)
	return item
}

type Users []*User

func (a Users) ToSchemaUsers() []*schema.User {
	list := make([]*schema.User, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaUser()
	}
	return list
}

type AdminUser struct {
	ID      int32  `gorm:"column:id;not null;primaryKey"`
	Created uint64 `gorm:"column:created;not null"`
}

type AuditorUser struct {
	ID int32 `gorm:"column:id;not null;primaryKey"`
}

func GetUserSettingsDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(UserSettings))
}

type UserSettings struct {
	ID           int32 `gorm:"column:id;not null;primaryKey"`
	UserID       int32 `gorm:"column:user_id;not null"`
	CompanyID    int32 `gorm:"column:company_id;not null"`
	DepartmentID int32 `gorm:"column:department_id;not null"`
}

package repository

import (
	"context"

	"github.com/laidingqing/sokr/internal/entity"
	"github.com/laidingqing/sokr/internal/schema"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IUserGroupRepository interface {
	Create(ctx context.Context, item schema.UserGroup) error
	Update(ctx context.Context, item schema.UserGroup) error
}

type UserGroupRepository struct {
	DB *gorm.DB
}

var _ IUserGroupRepository = &UserGroupRepository{}

func NewUserGroupRepository(dbConn *gorm.DB) IUserGroupRepository {
	userGroupRep := UserGroupRepository{DB: dbConn}
	return &userGroupRep
}

func (a *UserGroupRepository) Create(ctx context.Context, item schema.UserGroup) error {
	sitem := entity.SchemaUserGroup(item)
	db := entity.GetUserGroupDB(ctx, a.DB)
	result := db.Create(sitem.ToUserGroup())
	return errors.WithStack(result.Error)
}

func (a *UserGroupRepository) Update(ctx context.Context, item schema.UserGroup) error {
	sitem := entity.SchemaUserGroup(item)
	db := entity.GetUserGroupDB(ctx, a.DB)
	result := db.Where("user_id=?", item.UserID).Updates(sitem)
	return errors.WithStack(result.Error)
}

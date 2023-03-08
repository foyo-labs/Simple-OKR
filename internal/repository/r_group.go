package repository

import (
	"context"
	"time"

	"github.com/laidingqing/sokr/internal/entity"
	"github.com/laidingqing/sokr/internal/schema"
	"github.com/laidingqing/sokr/pkg/logger"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IGroupRepository interface {
	Create(ctx context.Context, group schema.Group) error
	Get(ctx context.Context, id string) (*schema.Group, error)
}

type GroupRepository struct {
	DB *gorm.DB
}

var _ IGroupRepository = &GroupRepository{}

func NewGroupRepository(dbConn *gorm.DB) IGroupRepository {
	groupRep := GroupRepository{DB: dbConn}
	return &groupRep
}

func (a *GroupRepository) Create(ctx context.Context, item schema.Group) error {
	sitem := entity.SchemaGroup(item)
	sitem.Created = uint64(time.Now().UnixNano())
	db := entity.GetGroupDB(ctx, a.DB)
	result := db.Create(sitem.ToGroup())
	return errors.WithStack(result.Error)
}

func (a *GroupRepository) Get(ctx context.Context, id string) (*schema.Group, error) {
	var item entity.Group
	db := entity.GetGroupDB(ctx, a.DB)
	result := db.Where("id=?", id).First(&item)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}
	logger.Infof("found group: %s", item.ID)
	return item.ToSchemaGroup(), nil
}

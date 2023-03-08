package service

import (
	"context"
	"fmt"

	"github.com/laidingqing/sokr/internal/repository"
	"github.com/laidingqing/sokr/internal/schema"
	"github.com/laidingqing/sokr/pkg/errors"
	"github.com/laidingqing/sokr/pkg/uuid"
)

type groupService struct {
	IGroupRepository     repository.IGroupRepository
	IUserGroupRepository repository.IUserGroupRepository
	Trans                repository.Trans
}

type IGroupService interface {
	Create(ctx context.Context, group schema.Group) (*schema.IDResult, error)
}

func NewGroupService(
	groupRepository repository.IGroupRepository,
	userGroupRepository repository.IUserGroupRepository,
	trans repository.Trans,
) IGroupService {
	return &groupService{IGroupRepository: groupRepository, IUserGroupRepository: userGroupRepository, Trans: trans}
}

func (a *groupService) Create(ctx context.Context, item schema.Group) (*schema.IDResult, error) {
	item.ID = uuid.NextID()
	currentLevel := uint64(0)
	levelNum := "0"
	parentID := "0"

	if item.ParentID != "" {
		group, err := a.IGroupRepository.Get(ctx, item.ParentID)
		if err != nil {
			return nil, errors.ErrBadRequest
		}
		if group != nil {
			currentLevel = group.CurrentLevel + 1
			levelNum = fmt.Sprintf("%s.%d", group.LevelNum, currentLevel)
		}
		parentID = item.ParentID
	}

	item.ParentID = parentID
	item.CurrentLevel = currentLevel
	item.LevelNum = levelNum

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		err := a.IGroupRepository.Create(ctx, item)
		if err != nil {
			return err
		}
		userGroup := schema.UserGroup{
			UserID:  item.UserID,
			GroupID: item.ID,
		}

		if item.ParentID == "" {
			err = a.IUserGroupRepository.Create(ctx, userGroup)
		}

		return err
	})

	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

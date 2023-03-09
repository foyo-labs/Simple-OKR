package service

import (
	"context"

	"github.com/laidingqing/sokr/internal/repository"
	"github.com/laidingqing/sokr/internal/schema"
	"github.com/laidingqing/sokr/pkg/errors"
	"github.com/laidingqing/sokr/pkg/util/hash"
	"github.com/laidingqing/sokr/pkg/util/jwt"
	"github.com/laidingqing/sokr/pkg/uuid"
)

type userService struct {
	IUserRepository      repository.IUserRepository
	IGroupRepository     repository.IGroupRepository
	IUserGroupRepository repository.IUserGroupRepository
	Trans                repository.Trans
}

type IUserService interface {
	Verify(ctx context.Context, email string, password string) (*schema.User, error)
	Create(ctx context.Context, user schema.User) (*schema.IDResult, error)
	GenerateToken(ctx context.Context, userID string) (*schema.LoginTokenInfo, error)
}

func NewUserService(
	userRepository repository.IUserRepository,
	userGroupRepo repository.IUserGroupRepository,
	groupRepo repository.IGroupRepository,
	trans repository.Trans,
) IUserService {
	return &userService{
		IUserRepository:      userRepository,
		IGroupRepository:     groupRepo,
		IUserGroupRepository: userGroupRepo,
		Trans:                trans,
	}
}

func (a *userService) GenerateToken(ctx context.Context, userID string) (*schema.LoginTokenInfo, error) {
	token, err := jwt.GenerateToken(userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	group, err := a.IGroupRepository.QueryUserGroup(ctx, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	item := &schema.LoginTokenInfo{
		AccessToken: token,
		UserID:      userID,
		UserInfo: schema.UserInfo{
			GroupID:   group.ID,
			GroupName: group.Name,
		},
	}
	return item, nil
}

func (a *userService) Verify(ctx context.Context, email string, password string) (*schema.User, error) {
	result, err := a.IUserRepository.Query(ctx, schema.UserQueryParam{
		Email: email,
	})
	if err != nil {
		return nil, err
	} else if len(result.Data) == 0 {
		return nil, errors.ErrInvalidUserName
	}

	item := result.Data[0]
	if item.Password != hash.SHA1String(password) {
		return nil, errors.ErrInvalidPassword
	} else if item.Status != 1 {
		return nil, errors.ErrUserDisable
	}

	return item, nil
}

func (a *userService) Create(ctx context.Context, item schema.User) (*schema.IDResult, error) {
	err := a.checkEmail(ctx, item)
	if err != nil {
		return nil, err
	}

	groupID := ""

	if item.GroupID != "" {
		group, err := a.checkGroup(ctx, item)
		if err != nil {
			return nil, err
		}
		groupID = group.ID
	}

	item.Password = hash.SHA1String(item.Password)
	item.ID = uuid.NextID()
	item.Status = 1
	item.Role = schema.LocalUser
	err = a.Trans.Exec(ctx, func(ctx context.Context) error {
		item.ID = uuid.NextID()
		err = a.IUserRepository.Create(ctx, item)
		if err != nil {
			return err
		}
		err = a.IUserGroupRepository.Create(ctx, schema.UserGroup{
			UserID:  item.ID,
			GroupID: groupID,
		})
		return err
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

func (a *userService) checkGroup(ctx context.Context, item schema.User) (*schema.Group, error) {
	result, err := a.IGroupRepository.Find(ctx, schema.GroupQueryParam{
		ID: item.GroupID,
	})
	if err != nil {
		return nil, err
	} else if len(result) == 0 {
		return nil, errors.New400Response("分组不存在")
	}
	return result[0], nil
}

func (a *userService) checkEmail(ctx context.Context, item schema.User) error {
	result, err := a.IUserRepository.Query(ctx, schema.UserQueryParam{
		PaginationParam: schema.PaginationParam{OnlyCount: true},
		Email:           item.Email,
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.New400Response("用户名已经存在")
	}
	return nil
}

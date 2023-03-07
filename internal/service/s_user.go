package service

import (
	"context"

	"github.com/laidingqing/sokr/internal/repository"
	"github.com/laidingqing/sokr/internal/schema"
	"github.com/laidingqing/sokr/pkg/errors"
	"github.com/laidingqing/sokr/pkg/util/hash"
)

type userService struct {
	IUserRepository repository.IUserRepository
}

type IUserService interface {
	Verify(ctx context.Context, email string, password string) (*schema.User, error)
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &userService{IUserRepository: userRepository}
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

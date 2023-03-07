package repository

import (
	"context"

	"github.com/laidingqing/sokr/internal/entity"
	"github.com/laidingqing/sokr/internal/schema"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(ctx context.Context, company schema.User) error
	Get(ctx context.Context, id string) (*schema.User, error)
	Query(ctx context.Context, query schema.UserQueryParam) (*schema.UserQueryResult, error)
}

type UserRepository struct {
	DB *gorm.DB
}

var _ IUserRepository = &UserRepository{}

func NewUserRepository(dbConn *gorm.DB) IUserRepository {
	userRep := UserRepository{DB: dbConn}
	return &userRep
}

func (a *UserRepository) Create(ctx context.Context, item schema.User) error {
	sitem := entity.SchemaUser(item)
	db := entity.GetUserDB(ctx, a.DB)
	result := db.Create(sitem.ToUser())
	return errors.WithStack(result.Error)
}

func (a *UserRepository) Get(ctx context.Context, id string) (*schema.User, error) {

	return nil, nil
}

func (a *UserRepository) Query(ctx context.Context, params schema.UserQueryParam) (*schema.UserQueryResult, error) {
	db := entity.GetUserDB(ctx, a.DB)

	if v := params.Email; v != "" {
		db = db.Where("email=?", v)
	}
	if v := params.Status; v > 0 {
		db = db.Where("status=?", v)
	}

	var list entity.Users

	pr, err := WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.UserQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaUsers(),
	}
	return qr, nil
}

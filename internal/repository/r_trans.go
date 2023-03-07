package repository

import (
	"context"

	"github.com/laidingqing/sokr/internal/contextx"
	"gorm.io/gorm"
)

// Trans 事务管理
type Trans struct {
	DB *gorm.DB
}

func NewTrans(dbConn *gorm.DB) Trans {
	trans := Trans{DB: dbConn}
	return trans
}

// Exec 执行事务
func (a *Trans) Exec(ctx context.Context, fn func(context.Context) error) error {
	if _, ok := contextx.FromTrans(ctx); ok {
		return fn(ctx)
	}

	return a.DB.Transaction(func(db *gorm.DB) error {
		return fn(contextx.NewTrans(ctx, db))
	})
}

package entity

import (
	"context"

	"github.com/laidingqing/sokr/internal/contextx"
	"gorm.io/gorm"
)

// GetDB ...
func GetDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	trans, ok := contextx.FromTrans(ctx)
	if ok && !contextx.FromNoTrans(ctx) {
		db, ok := trans.(*gorm.DB)
		if ok {
			if contextx.FromTransLock(ctx) {
				// Only Postgres and MySQL.
				db = db.Set("gorm:query_option", "FOR UPDATE")
			}
			return db
		}
	}
	return defDB
}

func GetDBWithModel(ctx context.Context, defDB *gorm.DB, m interface{}) *gorm.DB {
	return GetDB(ctx, defDB).Model(m)
}

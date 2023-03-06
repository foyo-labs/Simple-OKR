//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/laidingqing/sokr/internal/api"
	"github.com/laidingqing/sokr/internal/repository"
	"github.com/laidingqing/sokr/internal/service"
	"gorm.io/gorm"
)

func InitObjectiveAPI(db *gorm.DB) api.ObjectivesAPI {
	wire.Build(repository.NewObjectiveRepository, service.NewObjectiveService, api.NewObjectivesAPI)
	return api.ObjectivesAPI{}
}

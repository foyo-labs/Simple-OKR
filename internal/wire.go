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

func InitUserAPI(db *gorm.DB) api.UserAPI {
	wire.Build(repository.NewUserRepository, repository.NewTrans, service.NewUserService, api.NewUserAPI)
	return api.UserAPI{}
}

func InitUnitAPI(db *gorm.DB) api.UnitAPI {
	wire.Build(repository.NewCompanyRepository, repository.NewTrans, service.NewUnitService, api.NewUnitAPI)
	return api.UnitAPI{}
}

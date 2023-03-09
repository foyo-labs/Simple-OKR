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
	wire.Build(
		repository.NewObjectiveRepository,
		repository.NewKeyResultRepository,
		repository.NewGroupObjectiveRepository,
		repository.NewUserObjectiveRepository,
		repository.NewCycleRepository,
		repository.NewTrans,
		service.NewObjectiveService,
		api.NewObjectivesAPI,
	)
	return api.ObjectivesAPI{}
}

func InitUserAPI(db *gorm.DB) api.UserAPI {
	wire.Build(
		repository.NewUserRepository,
		repository.NewUserGroupRepository,
		repository.NewGroupRepository,
		repository.NewTrans,
		service.NewUserService,
		service.NewGroupService,
		api.NewUserAPI,
	)
	return api.UserAPI{}
}

func InitGroupAPI(db *gorm.DB) api.GroupAPI {
	wire.Build(repository.NewGroupRepository, repository.NewUserGroupRepository, repository.NewTrans, service.NewGroupService, api.NewGroupAPI)
	return api.GroupAPI{}
}

func InitCycleAPI(db *gorm.DB) api.CycleAPI {
	wire.Build(
		repository.NewCycleRepository,
		repository.NewTrans,
		service.NewCycleService,
		api.NewCycleAPI,
	)
	return api.CycleAPI{}
}

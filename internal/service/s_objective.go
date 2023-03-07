package service

import "github.com/laidingqing/sokr/internal/repository"

type objectiveService struct {
	IObjectiveRepository repository.IObjectiveRepository
}

type IObjectiveService interface {
}

func NewObjectiveService(objectiveRepository repository.IObjectiveRepository) IObjectiveService {
	return &objectiveService{IObjectiveRepository: objectiveRepository}
}

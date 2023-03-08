package service

import (
	"context"

	"github.com/laidingqing/sokr/internal/repository"
	"github.com/laidingqing/sokr/internal/schema"
	"github.com/laidingqing/sokr/pkg/uuid"
)

type cycleService struct {
	ICycleRepository repository.ICycleRepository
	Trans            repository.Trans
}

type ICycleService interface {
	Create(ctx context.Context, group schema.Cycle) (*schema.IDResult, error)
}

func NewCycleService(
	cycleRepository repository.ICycleRepository,
	trans repository.Trans,
) ICycleService {
	return &cycleService{
		ICycleRepository: cycleRepository,
		Trans:            trans,
	}
}

func (a *cycleService) Create(ctx context.Context, item schema.Cycle) (*schema.IDResult, error) {
	item.ID = uuid.NextID()
	err := a.ICycleRepository.Add(ctx, item)
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

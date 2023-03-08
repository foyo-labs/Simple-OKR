package service

import (
	"context"
	"time"

	"github.com/laidingqing/sokr/internal/repository"
	"github.com/laidingqing/sokr/internal/schema"
	"github.com/laidingqing/sokr/pkg/uuid"
)

type objectiveService struct {
	IObjectiveRepository repository.IObjectiveRepository
	IKeyResultRepository repository.IKeyResultRepository
	Trans                repository.Trans
}

// IObjectiveService 目标服务接口
type IObjectiveService interface {
	Create(ctx context.Context, item schema.Objective) (*schema.IDResult, error)
}

func NewObjectiveService(
	objectiveRepository repository.IObjectiveRepository,
	keyResultRepository repository.IKeyResultRepository,
	trans repository.Trans,
) IObjectiveService {
	return &objectiveService{IObjectiveRepository: objectiveRepository, IKeyResultRepository: keyResultRepository, Trans: trans}
}

func (a *objectiveService) Create(ctx context.Context, item schema.Objective) (*schema.IDResult, error) {
	item.ID = uuid.NextID()
	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		err := a.IObjectiveRepository.Add(ctx, item)
		if err != nil {
			return err
		}

		var tKeyResults []*schema.KeyResult
		for _, v := range item.KeyResults {
			v.ObjectiveID = item.ID
			v.ID = uuid.NextID()
			v.Created = uint64(time.Now().UnixNano())
			tKeyResults = append(tKeyResults, v)
		}

		return a.IKeyResultRepository.Add(ctx, tKeyResults)
	})

	if err != nil {
		return nil, err
	}
	return schema.NewIDResult(item.ID), nil
}

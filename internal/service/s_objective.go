package service

import (
	"context"
	"time"

	"github.com/laidingqing/sokr/internal/repository"
	"github.com/laidingqing/sokr/internal/schema"
	"github.com/laidingqing/sokr/pkg/uuid"
)

type objectiveService struct {
	IObjectiveRepo      repository.IObjectiveRepository
	IKeyResultRepo      repository.IKeyResultRepository
	IGroupObjectiveRepo repository.IGroupObjectiveRepository
	IUserObjectiveRepo  repository.IUserObjectiveRepository
	Trans               repository.Trans
}

// IObjectiveService 目标服务接口
type IObjectiveService interface {
	Create(ctx context.Context, item schema.Objective) (*schema.IDResult, error)
}

func NewObjectiveService(
	objectiveRepository repository.IObjectiveRepository,
	keyResultRepository repository.IKeyResultRepository,
	groupObjectiveRepository repository.IGroupObjectiveRepository,
	userObjectiveRepository repository.IUserObjectiveRepository,
	trans repository.Trans,
) IObjectiveService {
	return &objectiveService{
		IObjectiveRepo:      objectiveRepository,
		IKeyResultRepo:      keyResultRepository,
		IUserObjectiveRepo:  userObjectiveRepository,
		IGroupObjectiveRepo: groupObjectiveRepository,
		Trans:               trans,
	}
}

func (a *objectiveService) Create(ctx context.Context, item schema.Objective) (*schema.IDResult, error) {
	item.ID = uuid.NextID()
	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		err := a.IObjectiveRepo.Add(ctx, item)
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
		err = a.IKeyResultRepo.Add(ctx, tKeyResults)

		if item.ObjectiveType == schema.GroupObjectiveType {
			groupObjectID := uuid.NextID()
			if err = a.IGroupObjectiveRepo.Add(ctx, schema.GroupObjective{
				ID:          groupObjectID,
				GroupID:     item.GroupID,
				ObjectiveID: item.ID,
				CycleID:     item.CycleID,
			}); err != nil {
				return err
			}
		}

		if item.ObjectiveType == schema.UserObjectiveType {
			groupObjectID := uuid.NextID()
			if err = a.IUserObjectiveRepo.Add(ctx, schema.UserObjective{
				ID:          groupObjectID,
				UserID:      item.UserID,
				ObjectiveID: item.ID,
				CycleID:     item.CycleID,
			}); err != nil {
				return err
			}
		}

		return err
	})

	if err != nil {
		return nil, err
	}
	return schema.NewIDResult(item.ID), nil
}

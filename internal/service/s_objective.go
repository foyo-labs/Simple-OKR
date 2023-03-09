package service

import (
	"context"
	"time"

	"github.com/laidingqing/sokr/internal/repository"
	"github.com/laidingqing/sokr/internal/schema"
	"github.com/laidingqing/sokr/pkg/errors"
	"github.com/laidingqing/sokr/pkg/logger"
	"github.com/laidingqing/sokr/pkg/uuid"
)

type objectiveService struct {
	IObjectiveRepo      repository.IObjectiveRepository
	IKeyResultRepo      repository.IKeyResultRepository
	IGroupObjectiveRepo repository.IGroupObjectiveRepository
	IUserObjectiveRepo  repository.IUserObjectiveRepository
	ICycleRepository    repository.ICycleRepository
	Trans               repository.Trans
}

// IObjectiveService 目标服务接口
type IObjectiveService interface {
	Create(ctx context.Context, item schema.Objective) (*schema.IDResult, error)
	Query(ctx context.Context, item schema.ObjectiveQueryParam) ([]*schema.Objective, error)
}

func NewObjectiveService(
	objectiveRepository repository.IObjectiveRepository,
	keyResultRepository repository.IKeyResultRepository,
	groupObjectiveRepository repository.IGroupObjectiveRepository,
	userObjectiveRepository repository.IUserObjectiveRepository,
	cycleRepository repository.ICycleRepository,
	trans repository.Trans,
) IObjectiveService {
	return &objectiveService{
		IObjectiveRepo:      objectiveRepository,
		IKeyResultRepo:      keyResultRepository,
		IUserObjectiveRepo:  userObjectiveRepository,
		IGroupObjectiveRepo: groupObjectiveRepository,
		ICycleRepository:    cycleRepository,
		Trans:               trans,
	}
}

func (a *objectiveService) Create(ctx context.Context, item schema.Objective) (*schema.IDResult, error) {
	err := a.checkCycle(ctx, schema.CycleQueryParam{
		ID: item.CycleID,
	})
	if err != nil {
		return nil, err
	}

	item.ID = uuid.NextID()
	err = a.Trans.Exec(ctx, func(ctx context.Context) error {
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

// Query 查询Objective
// 包含成员与组的O
func (a *objectiveService) Query(ctx context.Context, item schema.ObjectiveQueryParam) ([]*schema.Objective, error) {

	return nil, nil
}

// Private func

func (a *objectiveService) checkCycle(ctx context.Context, query schema.CycleQueryParam) error {
	result, err := a.ICycleRepository.Find(ctx, query)
	logger.Infof("%v", result)
	if err != nil {
		return err
	} else if len(result) == 0 {
		return errors.New400Response("OKR周期未发现")
	}
	return nil
}

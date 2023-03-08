package service

import (
	"context"

	"github.com/laidingqing/sokr/internal/repository"
	"github.com/laidingqing/sokr/internal/schema"
)

type keyresultService struct {
	IKeyResultRepository repository.IKeyResultRepository
}

// IKeyResultService 关键结果服务接口
type IKeyResultService interface {
	Create(ctx context.Context, items []schema.KeyResult) ([]*schema.IDResult, error)
}

func NewKeyResultService(keyresultRepository repository.IKeyResultRepository) IKeyResultService {
	return &keyresultService{IKeyResultRepository: keyresultRepository}
}

func (a *keyresultService) Create(ctx context.Context, items []schema.KeyResult) ([]*schema.IDResult, error) {

	return nil, nil
}

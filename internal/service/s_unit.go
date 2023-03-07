package service

import "github.com/laidingqing/sokr/internal/repository"

type unitService struct {
	ICompanyRepository repository.ICompanyRepository
	Trans              repository.Trans
}

type IUnitService interface {
}

func NewUnitService(companyRepository repository.ICompanyRepository, trans repository.Trans) IUnitService {
	return &unitService{ICompanyRepository: companyRepository, Trans: trans}
}

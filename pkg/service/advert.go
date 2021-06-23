package service

import (
	//"github.com/p12s/avito-advertising-http-api"
	common "github.com/p12s/avito-advertising-http-api"
	"github.com/p12s/avito-advertising-http-api/pkg/repository"
)

type AdvertService struct {
	repo repository.Advert
}

func NewAdvertService(repo repository.Advert) *AdvertService {
	return &AdvertService{repo: repo}
}

func (a *AdvertService) GetByOrder(params common.AdvertSortOrderParams, advertCount int) ([]common.AdvertsListItem, error) {
	return a.repo.GetByOrder(params, advertCount)
}

func (a *AdvertService) GetById() error {
	return a.repo.GetById()
}

func (a *AdvertService) Create() error {
	return a.repo.Create()
}

func (a *AdvertService) Get(advertId int, params common.AdvertFieldParams) (common.AdvertWithPhoto, error) {
	return a.repo.Get(advertId, params)
}

func (a *AdvertService) Update() error {
	return a.repo.Update()
}

func (a *AdvertService) Delete() error {
	return a.repo.Delete()
}

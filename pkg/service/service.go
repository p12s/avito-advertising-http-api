package service

import (
	common "github.com/p12s/avito-advertising-http-api"
	"github.com/p12s/avito-advertising-http-api/pkg/repository"
)

type Advert interface {
	Get() error
	Create() error //int,
	GetByOrder(params common.AdvertSortOrderParams, advertCount int) ([]common.AdvertWithPhoto, error)
	GetById() error //common.Advert,
	Delete() error
	Update() error
}

type Service struct {
	Advert
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Advert: NewAdvertService(repos.Advert),
	}
}

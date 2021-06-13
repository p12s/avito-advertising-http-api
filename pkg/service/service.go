package service

import (
	"github.com/p12s/avito-advertising-http-api/pkg/repository"
)

type Advert interface {
	Create() error     //int,
	GetByOrder() error //[]common.Advert,
	GetById() error    //common.Advert,
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

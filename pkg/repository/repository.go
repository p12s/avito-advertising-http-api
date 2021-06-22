package repository

import (
	"github.com/jmoiron/sqlx"
	common "github.com/p12s/avito-advertising-http-api"
)

type Advert interface {
	GetByOrder(params common.AdvertSortOrderParams, advertCount int) ([]common.AdvertWithPhoto, error)
	GetById() error
	Create() error
	Get() error
	Update() error
	Delete() error
}

type Repository struct {
	Advert
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Advert: NewAdvertPostgres(db),
	}
}

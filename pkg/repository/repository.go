package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/p12s/avito-advertising-http-api"
)

type Advert interface {
	GetByOrder() (error)
	Create() (error)
	Get() (error)
	Update() (error)
	Delete() (error)
}

type Repository struct {
	Advert
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Advert: NewAdvertPostgres(db),
	}
}

package repository

import (
	"github.com/jmoiron/sqlx"
	//"github.com/p12s/avito-advertising-http-api"
)

type AdvertPostgres struct {
	db *sqlx.DB
}

func NewAdvertPostgres(db *sqlx.DB) *AdvertPostgres {
	return &AdvertPostgres{db: db}
}

func (a *AdvertPostgres) Create() error {
	return nil
}

func (a *AdvertPostgres) GetById() error { //common.Advert, error
	return nil
}

func (a *AdvertPostgres) Delete() error {
	return nil
}

func (a *AdvertPostgres) Update() error {
	return nil
}

package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	common "github.com/p12s/avito-advertising-http-api"
	//"github.com/p12s/avito-advertising-http-api"
)

type AdvertPostgres struct {
	db *sqlx.DB
}

func NewAdvertPostgres(db *sqlx.DB) *AdvertPostgres {
	return &AdvertPostgres{db: db}
}

func (a *AdvertPostgres) Get() error {
	return nil
}

func (a *AdvertPostgres) Create() error {
	return nil
}

func (a *AdvertPostgres) GetByOrder(params common.AdvertSortOrderParams, advertCount int) ([]common.AdvertWithPhoto, error) {
	priceOrder := params.Price.String()
	createdAtOrder := params.CreatedAt.String()
	offset := params.Offset // 	TODO проверить на (> 0 && < MAX)

	var items []common.AdvertWithPhoto
	query := fmt.Sprintf(`SELECT DISTINCT ON (a.price, a.created_at) a.title, a.price, ph.url AS photo_url FROM %s a 
		LEFT JOIN %s ph ON a.id = ph.id_advert
		ORDER BY a.price %s, a.created_at %s LIMIT %d OFFSET %d`,
		advertTable, photoTable, priceOrder, createdAtOrder, advertCount, offset)

	if err := a.db.Select(&items, query); err != nil {
		return nil, err
	}

	return items, nil
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

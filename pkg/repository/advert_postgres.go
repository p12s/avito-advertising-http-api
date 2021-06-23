package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	common "github.com/p12s/avito-advertising-http-api"
	"time"
	//"github.com/p12s/avito-advertising-http-api"
)

type AdvertPostgres struct {
	db *sqlx.DB
}

func NewAdvertPostgres(db *sqlx.DB) *AdvertPostgres {
	return &AdvertPostgres{db: db}
}

func (a *AdvertPostgres) Get(advertId int, params common.AdvertFieldParams) (common.AdvertWithPhoto, error) {
	// advert
	queryPart := ""
	if params.Fields.Description {
		queryPart = ", a.description"
	}
	var item common.AdvertWithPhoto
	query := fmt.Sprintf(`SELECT a.title, a.price %s FROM %s a LEFT JOIN %s ph ON a.id = ph.id_advert WHERE a.id = %d`,
		queryPart, advertTable, photoTable, advertId)

	if err := a.db.Get(&item, query); err != nil {
		return common.AdvertWithPhoto{}, err
	}

	// photo
	queryPart = "AND p.is_general IS TRUE"
	if params.Fields.AllPhoto {
		queryPart = "ORDER BY p.is_general DESC, p.id ASC"
	}
	query = fmt.Sprintf(`SELECT p.id, p.id_advert, p.url, p.is_general FROM %s p WHERE p.id_advert = %d %s`,
		photoTable, advertId, queryPart)
	if err := a.db.Select(&item.Photos, query); err != nil {
		return common.AdvertWithPhoto{}, err
	}

	return item, nil
}

func (a *AdvertPostgres) Create(advert common.AdvertWithPhoto) (int, error) {
	tx, err := a.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	query := fmt.Sprintf("INSERT INTO %s (title, description, price, created_at) values ($1, $2, $3, $4) RETURNING id", advertTable)
	row := tx.QueryRow(query, advert.Title, advert.Description, advert.Price, time.Now())
	err = row.Scan(&itemId)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	values := ""
	for i, val := range advert.Photos {
		isGeneralStr := "FALSE"
		if val.IsGeneral {
			isGeneralStr = "TRUE"
		}
		values += fmt.Sprintf("(%d, '%s', %s)", itemId, val.Url, isGeneralStr)
		if i < len(advert.Photos)-1 {
			values += ","
		}
	}
	createPhotosQuery := fmt.Sprintf("INSERT INTO %s (id_advert, url, is_general) values %s", photoTable, values)
	_, err = tx.Exec(createPhotosQuery)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (a *AdvertPostgres) GetByOrder(params common.AdvertSortOrderParams, advertCount int) ([]common.AdvertsListItem, error) {
	priceOrder := params.Price.String()
	createdAtOrder := params.CreatedAt.String()
	offset := params.Offset // 	TODO проверить на (> 0 && < MAX)

	var items []common.AdvertsListItem
	query := fmt.Sprintf(`SELECT DISTINCT ON (a.price, a.created_at) a.title, a.price, ph.url AS photo_url FROM %s a 
		LEFT JOIN %s ph ON a.id = ph.id_advert
		WHERE ph.is_general IS TRUE
		ORDER BY a.price %s, a.created_at %s LIMIT %d OFFSET %d`,
		advertTable, photoTable, priceOrder, createdAtOrder, advertCount, offset)
	if err := a.db.Select(&items, query); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return items, nil
}

func (a *AdvertPostgres) GetById() error {
	return nil
}

func (a *AdvertPostgres) Delete() error {
	return nil
}

func (a *AdvertPostgres) Update() error {
	return nil
}

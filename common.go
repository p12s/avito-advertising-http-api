package common

import (
	"time"
)

type Photo struct {
	Id       int    `json:"id" db:"id"`
	IdAdvert int    `json:"id_advert" db:"id_advert"`
	Url      string `json:"url" db:"url"`
}

type Advert struct {
	Id          int       `json:"-" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Price       float32   `json:"price" db:"price"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type SortOrder int

const (
	Asc SortOrder = iota
	Desc
)

func (c SortOrder) String() string {
	sortOrderString := ""
	if c == Desc {
		sortOrderString = "DESC"
	} else {
		sortOrderString = "ASC"
	}
	return sortOrderString
}

type AdvertSortOrderParams struct {
	Price     SortOrder `json:"price"`
	CreatedAt SortOrder `json:"created_at"`
	Offset    int       `json:"offset"`
}

type AdvertWithPhoto struct {
	Id          int       `json:"-" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"-" db:"description"`
	Price       float32   `json:"price" db:"price"`
	CreatedAt   time.Time `json:"-" db:"created_at"`
	PhotoUrl    string    `json:"photo_url" db:"photo_url"`
}

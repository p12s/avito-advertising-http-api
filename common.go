package common

import (
	"errors"
	"time"
)

type Photo struct {
	Id        int    `json:"id" db:"id"`
	IdAdvert  int    `json:"id_advert" db:"id_advert"`
	Url       string `json:"url" db:"url"`
	IsGeneral bool   `json:"is_general" db:"is_general"`
}

// Advert - рекламное объявление
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

// AdvertSortOrderParams - передаваемые клиентом поля для сортировки и пагинации
type AdvertSortOrderParams struct {
	Price     SortOrder `json:"price"`      // сортировка по цене
	CreatedAt SortOrder `json:"created_at"` // по дате создания
	Offset    int       `json:"offset"`     // смещение от 0
}

// AdvertsListItem - элемент возвращаемого клиенту списка объявлений
type AdvertsListItem struct {
	Id          int       `json:"-" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description,omitempty" db:"description"`
	Price       float32   `json:"price" db:"price"`
	CreatedAt   time.Time `json:"-" db:"created_at"`
	PhotoUrl    string    `json:"photo_url,omitempty" db:"photo_url"` // адрес главного фото
}

// AdvertWithPhoto - возвращаемое клиенту 1 объявление
type AdvertWithPhoto struct {
	Id          int       `json:"-" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description,omitempty" db:"description"`
	Price       float32   `json:"price" db:"price"`
	CreatedAt   time.Time `json:"-" db:"created_at"`
	Photos      []Photo   `json:"photos,omitempty" db:"photos"` // массив фото
}

// Validate - не больше 3 ссылок на объявление, название не больше 200 символов, описание не больше 1000 символов
func (a AdvertWithPhoto) Validate() error {
	if len(a.Title) > 200 {
		return errors.New("title cannot be longer than 200 characters")
	}
	if len(a.Description) > 1000 {
		return errors.New("description cannot be longer than 1000 characters")
	}
	if len(a.Photos) > 3 {
		return errors.New("photos count cannot be more than 3")
	}
	return nil
}

type Fields struct {
	Description bool `json:"description" db:"description"` // включить в ответ описание, необязательное поле
	AllPhoto    bool `json:"all_photo" db:"all_photo"`     // включить в ответ все фото, а не только главное, необязательное поле
}
// AdvertFieldParams - доплнительные передаваемые в запросе клиента поля
type AdvertFieldParams struct {
	Fields Fields
}

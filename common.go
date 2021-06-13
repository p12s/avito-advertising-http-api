package common

type Advert struct {
	Id           int    `json:"-"`
	Title        string `json:"name"`
	Description  string `json:"description"`
	CreationDate string `json:"creation_date"`
}

type Photo struct {
	Id       int    `json:"id"`
	IdAdvert int    `json:"id_advert"`
	Url      string `json:"url"`
}

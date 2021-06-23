package handler

import (
	"github.com/gin-gonic/gin"
	common "github.com/p12s/avito-advertising-http-api"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

type getAllAdvertsResponse struct {
	Data []common.Advert `json:"data"` // todo  возможно надо добавить фото-поля?
}

// getByOrder - получение списка объявлений
// доступна сортировка по цене/дате создания и пагинация по 10 шт
func (h *Handler) getByOrder(c *gin.Context) {
	var input common.AdvertSortOrderParams
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	adverts, err := h.services.Advert.GetByOrder(input, viper.GetInt("settings.pageAdvertCount"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, adverts)
}

// create - создание объявления
// принимает поля: название, описание, цена, несколько ссылок на фотографии (сами фото не загружаем)
// Возвращает ID созданного объявления и код результата (ошибка или успех)
func (h *Handler) create(c *gin.Context) {
	var input common.AdvertWithPhoto
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	advertId, err := h.services.Advert.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": advertId,
	})
}

// get - получение конкретного объявления
// обязательные поля в ответе клиенту: название объявления, цена, ссылка на главное фото;
// опциональные поля (можно запросить, передав параметр fields): описание, ссылки на все фото.
func (h *Handler) get(c *gin.Context) {
	advertId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid advert id param")
		return
	}

	var input common.AdvertFieldParams
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	advert, err := h.services.Advert.Get(advertId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, advert)
}

func (h *Handler) update(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, map[string]interface{}{
		"id": 444,
	})
}

func (h *Handler) delete(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, map[string]interface{}{
		"id": 555,
	})
}

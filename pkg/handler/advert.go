package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	common "github.com/p12s/avito-advertising-http-api"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

type getAllAdvertsResponse struct {
	Data []common.Advert `json:"data"` // todo  возможно надо добавить фото-поля?
}

// getByOrder - получение списка объявлений (доступна сортировка по цене/дате создания и пагинация по 10 шт)
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

func (h *Handler) create(c *gin.Context) {

	c.JSON(http.StatusNotImplemented, map[string]interface{}{
		"id": 222,
	})
}

// Обязательные поля в ответе: название объявления, цена, ссылка на главное фото;
// Опциональные поля (можно запросить, передав параметр fields): описание, ссылки на все фото.
func (h *Handler) get(c *gin.Context) {
	advertId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid advert id param")
		return
	}
	fmt.Println(advertId)

	var input common.AdvertFieldParams
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(input)

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

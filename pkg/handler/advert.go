package handler

import (
	"github.com/gin-gonic/gin"
	common "github.com/p12s/avito-advertising-http-api"
	"github.com/spf13/viper"
	"net/http"
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

	c.JSON(http.StatusNotImplemented, map[string]interface{}{
		"id": 333,
	})
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

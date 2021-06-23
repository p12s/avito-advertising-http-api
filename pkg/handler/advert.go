package handler

import (
	"github.com/gin-gonic/gin"
	common "github.com/p12s/avito-advertising-http-api"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

// @Summary GetByOrder
// @Tags getByOrder
// @Description getting ordered advert
// @ID get-by-order
// @Accept  json
// @Produce  json
// @Param input body common.AdvertSortOrderParams true "sort info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/advt/ [get]
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

// @Summary Create
// @Tags create
// @Description create advert
// @ID create
// @Accept  json
// @Produce  json
// @Param input body common.AdvertWithPhoto true "advert info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/advt/ [post]
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

// @Summary Get
// @Tags get
// @Description get advert
// @ID get
// @Accept  json
// @Produce  json
// @Param input body common.AdvertFieldParams true "field getting info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/advt/:id/item/ [get]
func (h *Handler) get(c *gin.Context) { // получение конкретного объявления
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

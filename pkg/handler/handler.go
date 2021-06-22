package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/p12s/avito-advertising-http-api/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		advt := api.Group("/advt")
		{
			advt.GET("/", h.getByOrder)
			advt.POST("/", h.create)

			item := advt.Group(":id/item")
			{
				item.GET("/", h.get)
				item.PUT("/", h.update)
				item.DELETE("/", h.delete)
			}
		}
	}

	return router
}

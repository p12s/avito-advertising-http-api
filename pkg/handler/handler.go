package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/p12s/avito-advertising-http-api/pkg/service"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/p12s/avito-advertising-http-api/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

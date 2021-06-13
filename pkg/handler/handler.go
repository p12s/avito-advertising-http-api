package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		advt := api.Group("/advt")
		{
			advt.GET("/", h.getByOrder)
			advt.POST("/", h.create)

			item := advt.Group(":/id")
			{
				item.GET("/", h.get)
				item.PUT("/", h.update)
				item.DELETE("/", h.delete)
			}
		}
	}

	return router
}

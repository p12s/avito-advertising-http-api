package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getByOrder(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, map[string]interface{}{
		"id": 111,
	})
}

func (h *Handler) create(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, map[string]interface{}{
		"id": 222,
	})
}

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

package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) checkHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"health": "ok"})
}

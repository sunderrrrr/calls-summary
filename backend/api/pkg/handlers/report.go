package handlers

import (
	"api/pkg/utils/logger"
	"api/pkg/utils/responser"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) reportOfCall(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		responser.NewErrorResponse(c, http.StatusBadRequest, "file not found")
		logger.Log.Errorf("error getting file from form: %v", err)
		return
	}
	defer file.Close()
}

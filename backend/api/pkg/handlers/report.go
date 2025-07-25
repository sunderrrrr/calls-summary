package handlers

import (
	"api/pkg/utils/logger"
	"api/pkg/utils/responser"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) reportOfCall(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		responser.NewErrorResponse(c, http.StatusBadRequest, "file not found")
		logger.Log.Errorf("error getting file from form: %v", err)
		return
	}
	logger.Log.Info("processing report call")
	defer file.Close()
	logger.Log.Infof("file name: %s", fileHeader.Filename)
	resp, err := h.service.Report.ReportCall(file, fileHeader.Filename)
	logger.Log.Info("report call finished")
	if err != nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, "failed to process report")
		logger.Log.Errorf("error processing report: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": resp})
}

package handlers

import (
	"api/models"
	"api/pkg/utils/logger"
	"api/pkg/utils/responser"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) analyzeCall(c *gin.Context) { // Возвращает id созданного анализа
	userId, err := getUserId(c)
	if err != nil {
		responser.NewErrorResponse(c, http.StatusUnauthorized, "unauthorized")
	}
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		responser.NewErrorResponse(c, http.StatusBadRequest, "file not found")
		logger.Log.Errorf("error getting file from form: %v", err)
		return
	}
	logger.Log.Debugln("processing report call")
	defer file.Close()
	logger.Log.Infof("file name: %s", fileHeader.Filename)
	resp, err := h.service.Analysis.AnalyzeCall(userId, file, fileHeader.Filename)
	logger.Log.Debugln("report call finished")
	if err != nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, "failed to process report")
		logger.Log.Errorf("error processing report: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": resp})
}

func (h *Handler) AddMessage(c *gin.Context) {
	id, err := getUserId(c)
	analysisId := c.Param("analysisId")
	if err != nil {
		responser.NewErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		return
	}
	var input models.ChatMessage
	if err := c.ShouldBindJSON(&input); err != nil {
		responser.NewErrorResponse(c, http.StatusBadRequest, "field validation fail")
		logger.Log.Errorf("failed to bind JSON: %v", err)
		return
	}
	if err := h.service.Analysis.SendMessageToChat(analysisId, id, input); err != nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, "failed to send message")
		logger.Log.Errorf("failed to send message: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "message sent successfully"})

}

func (h *Handler) GetChatHistory(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		responser.NewErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		return
	}
	analysisId := c.Param("analysisId")
	messages, err := h.service.Analysis.GetChatHistory(analysisId, userId)
	if err != nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, "failed to get chat history")
		logger.Log.Errorf("failed to get chat history: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

func (h *Handler) GetAllAnalyses(c *gin.Context) { //Получение списка в формате название:анализ

}

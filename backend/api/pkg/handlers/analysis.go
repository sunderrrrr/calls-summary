package handlers

import (
	"api/models"
	"api/pkg/utils/logger"
	"api/pkg/utils/responser"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) analyzeCall(c *gin.Context) {
	// Обрабатывает загрузку файла звонка, передает его в сервис для анализа
	// и возвращает ID созданного анализа
	// Использует сервисный метод Analysis.AnalyzeCall
	// Возвращает HTTP 401, если пользователь не авторизован
	// Возвращает HTTP 400, если файл не найден
	// Возвращает HTTP 500, если произошла ошибка обработки
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
	// Добавляет сообщение в чат-бота для конкретного анализа
	// Использует сервисный метод Analysis.SendMessageToChat
	// Проверяет авторизацию пользователя и валидирует входные данные
	// Возвращает HTTP 401, если пользователь не авторизован
	// Возвращает HTTP 400, если входные данные некорректны
	// Возвращает HTTP 500, если произошла ошибка при отправке сообщения
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
	// Получает историю сообщений чата для конкретного анализа
	// Использует сервисный метод Analysis.GetChatHistory
	// Возвращает HTTP 401, если пользователь не авторизован
	// Возвращает HTTP 500, если произошла ошибка при получении данных
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

func (h *Handler) GetAllAnalyses(c *gin.Context) {

}

package service

import (
	"api/models"
	"api/pkg/utils/httpClient"
	"api/pkg/utils/logger"
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
)

func ReportCall(file io.Reader, filename string) (models.AnalysisResponse, error) {
	var buffer bytes.Buffer
	var response models.AnalysisResponse
	writer := multipart.NewWriter(&buffer)

	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		logger.Log.Errorf("create form file: %v", err)
		return models.AnalysisResponse{}, err
	}

	if _, err := io.Copy(part, file); err != nil {
		logger.Log.Errorf("copy file: %v", err)
		return models.AnalysisResponse{}, err
	}

	if err := writer.Close(); err != nil {
		logger.Log.Errorf("close writer: %v", err)
		return models.AnalysisResponse{}, err
	}

	req, err := http.NewRequest("POST", "http://localhost:8090/call-analysis", &buffer)
	if err != nil {
		logger.Log.Errorf("new request: %v", err)
		return models.AnalysisResponse{}, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := httpClient.DefaultClient.Do(req)
	if err != nil {
		logger.Log.Errorf("send request: %v", err)
		return models.AnalysisResponse{}, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		logger.Log.Errorf("decode response: %v", err)
		return models.AnalysisResponse{}, err
	}
	return response, nil
}

func AskLLM(message []models.ChatMessage) (models.ChatMessage, error) {
	var response struct {
		Message string `json:"message"`
	}
	marsh, err := json.Marshal(message)
	if err != nil {
		logger.Log.Errorf("marshal message err: %v", err)
		return models.ChatMessage{}, err
	}
	req, err := http.NewRequest("POST", "http://localhost:8090/chat-response", bytes.NewBuffer(marsh))
	if err != nil {
		logger.Log.Errorf("new request: %v", err)
		return models.ChatMessage{}, err
	}

	resp, err := httpClient.DefaultClient.Do(req)
	if err != nil {
		logger.Log.Errorf("send request: %v", err)
		return models.ChatMessage{}, err
	}
	defer resp.Body.Close()
	// Читаем тело для диагностики
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Errorf("read response body: %v", err)
		return models.ChatMessage{}, err
	}
	logger.Log.Infof("response body: %s", string(bodyBytes))

	// Теперь пробуем декодировать тело, для этого создаем новый reader
	if err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&response); err != nil {
		logger.Log.Errorf("decode response: %v", err)
		return models.ChatMessage{}, err
	}

	return models.ChatMessage{
		Message: response.Message,
		Sender:  "bot",
	}, nil
}

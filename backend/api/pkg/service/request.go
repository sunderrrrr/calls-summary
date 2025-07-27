package service

import (
	"api/models"
	"api/pkg/utils/httpClient"
	"api/pkg/utils/logger"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
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

	if err := writer.WriteField("model", "gpt-4o-mini-transcribe"); err != nil {
		logger.Log.Errorf("write field: %v", err)
		return models.AnalysisResponse{}, err
	}

	if err := writer.Close(); err != nil {
		logger.Log.Errorf("close writer: %v", err)
		return models.AnalysisResponse{}, err
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return models.AnalysisResponse{}, errors.New("OPENAI_API_KEY not set")
	}

	req, err := http.NewRequest("POST", "https://openai.api.proxyapi.ru/v1/audio/transcriptions", &buffer)
	if err != nil {
		logger.Log.Errorf("new request: %v", err)
		return models.AnalysisResponse{}, err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
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

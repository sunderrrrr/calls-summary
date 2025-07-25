package service

import (
	"api/pkg/utils/httpClient"
	"api/pkg/utils/logger"
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type ReportService struct{}

func NewReportService() *ReportService {
	return &ReportService{}
}

func (s *ReportService) ReportCall(file io.Reader, filename string) (string, error) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		logger.Log.Errorf("create form file: %v", err)
		return "", err
	}

	if _, err := io.Copy(part, file); err != nil {
		logger.Log.Errorf("copy file: %v", err)
		return "", err
	}

	if err := writer.WriteField("model", "gpt-4o-mini-transcribe"); err != nil {
		logger.Log.Errorf("write field: %v", err)
		return "", err
	}

	if err := writer.Close(); err != nil {
		logger.Log.Errorf("close writer: %v", err)
		return "", err
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", errors.New("OPENAI_API_KEY not set")
	}

	req, err := http.NewRequest("POST", "https://openai.api.proxyapi.ru/v1/audio/transcriptions", &buffer)
	if err != nil {
		logger.Log.Errorf("new request: %v", err)
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := httpClient.DefaultClient.Do(req)
	if err != nil {
		logger.Log.Errorf("send request: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Errorf("read response: %v", err)
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		logger.Log.Errorf("OpenAI API error (%d): %s", resp.StatusCode, string(body))
		return "", errors.New("OpenAI API error: " + string(body))
	}

	return string(body), nil
}

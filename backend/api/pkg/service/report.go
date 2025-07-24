package service

import (
	"api/pkg/utils/httpClient"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type ReportService struct {
}

func NewReportService() *ReportService {
	return &ReportService{}
}

func (s *ReportService) reportCall(file io.Reader, filename string) (string, error) {
	var buffer bytes.Buffer
	w := multipart.NewWriter(&buffer)

	part, err := w.CreateFormFile("file", filename)
	if err != nil {
		return "", err
	}
	if _, err := io.Copy(part, file); err != nil {
		return "", err
	}
	w.Close()
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/analyze", os.Getenv("FASTAPI_URL")), &buffer)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := httpClient.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

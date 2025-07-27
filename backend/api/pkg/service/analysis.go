package service

import (
	"api/models"
	"api/pkg/repository"
	"errors"
	"io"
)

type AnalysisService struct {
	repo repository.Analysis
}

func NewAnalysisService(repo repository.Analysis) *AnalysisService {
	return &AnalysisService{repo: repo}
}

func (s *AnalysisService) AnalyzeCall(userId int, file io.Reader, filename string) (string, error) {
	analysis, err := ReportCall(file, filename)
	if err != nil {
		return "", errors.New("analysis failed: " + err.Error())
	}
	id, err := s.repo.CreateAnalysis(userId, analysis)
	if err != nil {
		return "", errors.New("analysis add to db failed: " + err.Error())
	}
	return id, nil
}

func (s *AnalysisService) SendMessageToChat(analysisId string, userId int, message models.ChatMessage) error {
	if message.Sender != "user" && message.Sender != "bot" {
		return errors.New("invalid sender")
	}
	err := s.repo.AddChatMessage(analysisId, userId, message.Sender, message.Message)
	if err != nil {
		return errors.New("failed to send message: " + err.Error())
	}
	return nil
}

func (s *AnalysisService) GetChatHistory(analysisId string, userId int) ([]models.ChatMessage, error) {
	messages, err := s.repo.GetAnalysisChatHistory(analysisId, userId)
	if err != nil {
		return nil, errors.New("failed to get chat history: " + err.Error())
	}
	return messages, nil
}

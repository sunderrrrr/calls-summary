package service

import (
	"api/models"
	"api/pkg/repository"
	"api/pkg/utils/logger"
	"errors"
	"io"
)

type AnalysisService struct {
	repo repository.Analysis
}

func NewAnalysisService(repo repository.Analysis) *AnalysisService {
	return &AnalysisService{repo: repo}
}

/*
Обработка звонка, получение файла и имени файла, отправка на fast api сервис для анализа звонка
Функционал чат-бота по конкретному анализу
Данные возвращаются и принимаются в handlers/analysis.go
*/

func (s *AnalysisService) AnalyzeCall(userId int, file io.Reader, filename string) (string, error) {
	// Отправляет файл звонка на внешний сервис для анализа.
	// Сохраняет результат анализа в базу данных через репозиторий.
	// Возвращает ID созданного анализа.
	analysis, err := ReportCall(file, filename)
	if err != nil {
		return "", errors.New("analysis failed: " + err.Error())
	}
	id, err := s.repo.CreateAnalysis(userId, analysis)
	if err != nil {
		return "", errors.New("analysis add to db failed: " + err.Error())
	}
	if err = s.repo.AddChatMessage(id, userId, "bot", analysis.Analysis); err != nil {
		return "", errors.New("failed to add initial bot message: " + err.Error())
	}
	return id, nil
}

func (s *AnalysisService) SendMessageToChat(analysisId string, userId int, message models.ChatMessage) error {
	// Добавляет сообщение в чат-бота.
	// Если сообщение отправлено пользователем, запрашивает ответ у LLM
	// и добавляет его в чат.
	// Использует методы репозитория для работы с базой данных.
	if message.Sender != "user" && message.Sender != "bot" {
		return errors.New("invalid sender")
	}
	err := s.repo.AddChatMessage(analysisId, userId, message.Sender, message.Message)
	if err != nil {
		return errors.New("failed to send message: " + err.Error())
	}

	if message.Sender == "user" {
		history, err := s.GetChatHistory(analysisId, userId)
		if err != nil {
			return errors.New("failed to get chat history: " + err.Error())
		}
		newMsg, err := AskLLM(history)
		if err != nil {
			return errors.New("failed to get response from LLM: " + err.Error())
		}
		err = s.repo.AddChatMessage(analysisId, userId, "bot", newMsg.Message)
		if err != nil {
			return errors.New("failed to send bot message: " + err.Error())
		}
		logger.Log.Debugln("bot message: %s", newMsg.Message)
	}
	return nil
}

func (s *AnalysisService) GetChatHistory(analysisId string, userId int) ([]models.ChatMessage, error) {
	// Получает историю сообщений чата из базы данных через репозиторий.
	messages, err := s.repo.GetAnalysisChatHistory(analysisId, userId)
	if err != nil {
		return nil, errors.New("failed to get chat history: " + err.Error())
	}
	return messages, nil
}

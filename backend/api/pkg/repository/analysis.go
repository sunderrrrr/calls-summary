package repository

import (
	"api/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AnalysisRepository struct {
	db *sqlx.DB
}

func NewAnalysisRepository(db *sqlx.DB) *AnalysisRepository {
	return &AnalysisRepository{
		db: db,
	}
}

func (r *AnalysisRepository) CreateAnalysis(userId int, analysis models.AnalysisResponse) (string, error) {
	var analysisId string
	query := fmt.Sprintf("INSERT INTO analyses (user_id, title, report) VALUES ($1, $2, $3) RETURNING id")
	err := r.db.QueryRow(query, userId, analysis.Title, analysis.Analysis).Scan(&analysisId)
	if err != nil {
		return "", fmt.Errorf("error creating analysis: %w", err)
	}
	return analysisId, nil
}

func (r *AnalysisRepository) GetAllAnalysis(id string) ([]models.Analysis, error) {
	var analysis []models.Analysis
	query := fmt.Sprintf("SELECT id, user_id, title, report, created_at FROM analyses WHERE user_id=$1 ORDER BY created_at DESC")
	if err := r.db.Select(&analysis, query, id); err != nil {
		return nil, fmt.Errorf("error getting analysis: %w", err)
	}
	return analysis, nil
}

func (r *AnalysisRepository) GetAnalysisChatHistory(analysisId string, userId int) ([]models.ChatMessage, error) {
	var messages []models.ChatMessage
	query := fmt.Sprintf(" SELECT m.id, m.analysis_id, m.sender, m.message, m.created_at FROM chat_messages m JOIN analyses a ON m.analysis_id = a.id WHERE m.analysis_id = $1 AND a.user_id = $2 ORDER BY m.created_at ASC")
	if err := r.db.Select(&messages, query, analysisId, userId); err != nil {
		return nil, fmt.Errorf("error getting analysis chat history: %w", err)
	}
	return messages, nil
}

func (r *AnalysisRepository) AddChatMessage(analysisId string, userId int, sender, message string) error {
	query := fmt.Sprintf("INSERT INTO chat_messages (analysis_id, sender, message) SELECT $1, $2, $3 FROM analyses WHERE id=$1 AND user_id=$4")
	_, err := r.db.Exec(query, analysisId, sender, message, userId)
	if err != nil {
		return fmt.Errorf("error adding chat message: %w", err)
	}
	return nil
}

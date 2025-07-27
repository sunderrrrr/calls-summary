package models

import "time"

type Analysis struct {
	ID        string    `db:"id" json:"id"`
	UserID    int       `db:"user_id" json:"user_id"`
	Title     string    `db:"title" json:"title"`
	Report    string    `db:"report" json:"report"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
type ChatMessage struct {
	ID         string    `db:"id" json:"id"`
	AnalysisID string    `db:"analysis_id" json:"analysis_id"`
	Sender     string    `db:"sender" json:"sender"` // "user" или "bot"
	Message    string    `db:"message" json:"message"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}

type AnalyzeListItem struct {
	ID    string `db:"id" json:"id"`
	Title string `db:"title" json:"title" binding:"required"`
}

type AnalysisResponse struct {
	Title    string `json:"title"`
	Analysis string `json:"analysis"`
}

type AnalysisList struct {
	_ []AnalyzeListItem `json:"items"`
}

package repository

import (
	"Uptique/internal/database/models"
	"database/sql"
	"time"
)

func CreateWebsite(db *sql.DB, url string, siteType string) (*models.WebsiteModel, error) {
	query := `
		INSERT INTO websites (url, type, created_at, updated_at)
		VALUES (?, ?, ?, ?)
		RETURNING id, url, type, created_at, updated_at
	`
	now := time.Now()
	row := db.QueryRow(query, url, siteType, now, now)
	var w models.WebsiteModel
	err := row.Scan(&w.ID, &w.URL, &w.Type, &w.CreatedAt, &w.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &w, err
}

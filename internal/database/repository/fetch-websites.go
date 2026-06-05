package repository

import (
	"Uptique/internal/database/models"
	"database/sql"
)

func FetchWebsites(db *sql.DB) ([]models.WebsiteModel, error) {
	query := `
	SELECT id, url, type, created_at, updated_at FROM websites
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var websites []models.WebsiteModel
	for rows.Next() {
		var w models.WebsiteModel
		err := rows.Scan(&w.ID, &w.URL, &w.Type, &w.CreatedAt, &w.UpdatedAt)
		if err != nil {
			return nil, err
		}
		websites = append(websites, w)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return websites, err
}

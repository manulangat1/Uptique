package repository

import (
	"Uptique/internal/database/models"
	"database/sql"

	// "net/url"
	"time"
)

func CreateStatusMonitor(db *sql.DB, id int, status string) (*models.StatusMonitor, error) {
	query := `
	INSERT INTO statuses (website_id, status, created_at)
	VALUES (?,?,?)
	RETURNING id, website_id, status, created_at
	`
	now := time.Now()
	row := db.QueryRow(query, id, status, now)

	var sm models.StatusMonitor
	err := row.Scan(&sm.ID, &sm.WebsiteID, &sm.Status, &sm.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &sm, nil
}

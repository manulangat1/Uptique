package repository

import (
	"Uptique/internal/database/models"
	"database/sql"
)

func FetchAllStatus(db *sql.DB) ([]models.StatusMonitor, error) {
	query := `
	SELECT * FROM statuses
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var status []models.StatusMonitor
	for rows.Next() {
		var s models.StatusMonitor
		err := rows.Scan(&s.ID, &s.WebsiteID, &s.Status, &s.CreatedAt)
		if err != nil {
			return nil, err
		}
		status = append(status, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return status, err
}

package models

import "time"

type WebsiteModel struct {
	ID        int64     `db:"id"`
	URL       string    `db:"url"`
	Type      string    `db:"type"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	status []StatusMonitor `db:"-"`
}

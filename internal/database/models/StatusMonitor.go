package models

import "time"

type StatusMonitor struct {
	ID        int64     `db:"id"`
	WebsiteID int64     `db:"website_id"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`

	// belongs to one website
	Website *WebsiteModel `db:"-"`
}

CREATE TABLE IF NOT EXISTS statuses (
    id         INTEGER PRIMARY KEY,
    website_id INTEGER NOT NULL,
    status     TEXT    NOT NULL CHECK(status IN ('up', 'down', 'degraded')),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (website_id) REFERENCES websites(id) ON DELETE CASCADE
);

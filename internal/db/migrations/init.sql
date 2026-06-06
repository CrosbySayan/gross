CREATE TABLE IF NOT EXISTS channels (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		slug TEXT NOT NULL UNIQUE
	);

CREATE TABLE IF NOT EXISTS articles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    external_id TEXT NOT NULL UNIQUE,
    title TEXT NOT NULL,
    url TEXT NOT NULL,
    summary TEXT,
    published_at INTEGER NOT NULL,
    published_by TEXT,
    channel_id INTEGER,
    FOREIGN KEY(channel_id) REFERENCES channels(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_articles_published_at
ON articles(published_at);

CREATE INDEX IF NOT EXISTS idx_articles_channel_id
ON articles(channel_id);

CREATE UNIQUE INDEX IF NOT EXISTS idx_articles_external_id
ON articles(external_id);

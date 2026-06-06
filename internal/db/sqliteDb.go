package db

import (
	"database/sql"
	"fmt"
	"os"

	rss "github.com/CrosbySayan/gross/internal/types"

	_ "modernc.org/sqlite"
)

// SqliteDbManager is an impl of DataBaseManager, allowing for the storage of Rss Articles and channels
type SqliteDbManager struct {
	DB *sql.DB
}

func NewDataBaseManager(filename string) (DataBaseManager, error) {
	db, err := sql.Open("sqlite", filename)
	db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return nil, err
	}

	manager := &SqliteDbManager{
		DB: db,
	}

	if err := manager.initSchema(); err != nil {
		_ = db.Close()
		return nil, err
	}

	return manager, nil
}

func (s *SqliteDbManager) initSchema() error {
	sqlBytes, err := os.ReadFile("internal/db/migrations/init.sql")
	if err != nil {
		return err
	}

	_, err = s.DB.Exec(string(sqlBytes))
	if err != nil {
		return err
	}
	rows, err := s.DB.Query(`
		SELECT name
		FROM sqlite_master
		WHERE type='table'
	`)
	for rows.Next() {
		var name string
		rows.Scan(&name)
		fmt.Println(name)
	}
	return err
}

func (s *SqliteDbManager) AddChannel(channel rss.Channel) (int64, error) {
	res, err := s.DB.Exec(`
	INSERT INTO channels (name, slug)
	VALUES (?, ?)
	`, channel.Name, channel.Slug)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId() // Gives the id of the entry
}

func (s *SqliteDbManager) GetChannelBySlug(slug string) (rss.Channel, error) {
	var channel rss.Channel
	err := s.DB.QueryRow(`SELECT id, name, slug
	FROM channels
	WHERE slug = ?`, slug).Scan(
		&channel.ID,
		&channel.Name,
		&channel.Slug,
	)
	if err != nil {
		return rss.Channel{}, err
	}

	return channel, err
}

func (s *SqliteDbManager) GetChannelByID(id int64) (rss.Channel, error) {
	var channel rss.Channel

	err := s.DB.QueryRow(`
	SELECT id, name, slug
	FROM channels
	WHERE id = ?`, id).Scan(
		&channel.ID,
		&channel.Name,
		&channel.Slug,
	)
	if err != nil {
		return rss.Channel{}, err
	}

	return channel, nil
}

func (s *SqliteDbManager) ListChannels() ([]rss.Channel, error) {
	rows, err := s.DB.Query(`SELECT * FROM channels`)
	if err != nil {
		return nil, nil
	}
	defer rows.Close()

	var channels []rss.Channel
	// rows.Next() acts like a linked list the curr can be treated as a single query
	for rows.Next() {
		var c rss.Channel
		if err := rows.Scan(&c.ID, &c.Name, &c.Slug); err != nil {
			return nil, err
		}

		channels = append(channels, c)
	}
	return channels, nil
}

func (s *SqliteDbManager) UpsertArticle(article rss.Article) (int64, error) {
	res, err := s.DB.Exec(`
	INSERT INTO articles (
		external_id,
		title,
		url,
		summary,
		published_at,
		published_by,
		channel_id
	)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	ON CONFLICT(external_id) DO UPDATE SET
		channel_id = excluded.channel_id,
		title = excluded.title,
		url = excluded.url,
		summary = excluded.summary,
		published_at = excluded.published_at,
		published_by = excluded.published_by
	`,
		article.ExternalID,
		article.Title,
		article.URL,
		article.Summary,
		article.PublishedAt,
		article.PublishedBy,
		article.ChannelID,
	)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (s *SqliteDbManager) RemoveArticle(id int64) error {
	_, err := s.DB.Exec(`
	DELETE FROM articles
	WHERE id = ?
	`, id)
	return err
}

func (s *SqliteDbManager) GetAllArticlesFromChannel(channelID int64) ([]rss.Article, error) {
	rows, err := s.DB.Query(`
		SELECT id, external_id, title, url, summary, published_at, published_by
		FROM articles
		WHERE channel_id = ?
		ORDER BY published_at DESC
	`, channelID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []rss.Article

	for rows.Next() {
		var a rss.Article

		err := rows.Scan(
			&a.ID,
			&a.ExternalID,
			&a.Title,
			&a.URL,
			&a.Summary,
			&a.PublishedAt,
			&a.PublishedBy,
		)
		if err != nil {
			return nil, err
		}

		articles = append(articles, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

func (s *SqliteDbManager) GetArticleByExternalID(externalID string) (rss.Article, error) {
	var article rss.Article

	err := s.DB.QueryRow(`SELECT title, external_id, id, url, summary, published_at, published_by
	FROM articles WHERE external_id = ?`, externalID).Scan(&article.Title, &article.ExternalID, &article.ID, &article.URL, &article.Summary, &article.PublishedAt, &article.PublishedBy)
	if err != nil {
		return rss.Article{}, err
	}
	return article, nil
}

func (s *SqliteDbManager) Close() error {
	return s.DB.Close()
}

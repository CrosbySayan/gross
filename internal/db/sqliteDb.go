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

func (s *SqliteDbManager) AddChannel(channel rss.Channel) (int64, error) { return 0, nil }
func (s *SqliteDbManager) GetChannelBySlug(slug string) (rss.Channel, error) {
	return rss.Channel{}, nil
}
func (s *SqliteDbManager) GetChannelByID(id int64) (rss.Channel, error) { return rss.Channel{}, nil }
func (s *SqliteDbManager) ListChannels() ([]rss.Channel, error)         { return nil, nil }

func (s *SqliteDbManager) GetArticleByExternalID(externalID string) (rss.Article, error) {
	return rss.Article{}, nil
}
func (s *SqliteDbManager) UpsertArticle(article rss.Article) (dbID int64, err error) { return 0, nil }
func (s *SqliteDbManager) RemoveArticle(id int64) error                              { return nil }
func (s *SqliteDbManager) SearchArticle(q rss.ArticleQuery) ([]rss.Article, error)   { return nil, nil }

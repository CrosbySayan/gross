package db

import rss "github.com/CrosbySayan/gross/internal/types"

type DataBaseManager interface {
	AddChannel(channel rss.Channel) (int64, error)
	GetChannelBySlug(slug string) (rss.Channel, error)
	GetChannelByID(id int64) (rss.Channel, error)
	ListChannels() ([]rss.Channel, error)

	GetAllArticlesFromChannel(channelID int64) ([]rss.Article, error)

	GetArticleByExternalID(externalID string) (rss.Article, error) // Tries to find an article with some GUID
	UpsertArticle(article rss.Article) (dbID int64, err error)     // Either updates an existing article or adds a new article
	RemoveArticle(id int64) error                                  // Removes an article from the db
	Close() error
}

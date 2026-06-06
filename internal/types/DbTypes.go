package types

import "time"

type Article struct {
	FeedId      int64
	Title       string
	URL         string
	Summary     string
	PublishedAt time.Time
	PublishedBy string
	ExternalId  string
}

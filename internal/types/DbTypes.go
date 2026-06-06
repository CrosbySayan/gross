package types

import "time"

type Article struct {
	ID         int64
	ExternalID string // GUID

	Title       string
	URL         string
	Summary     string
	PublishedAt time.Time
	PublishedBy string
}

type Channel struct {
	ID   int64  // Db ID
	Name string // External Name
	Slug string // Internal Name
}

type ArticleQuery struct {
	// For text searching
	Text string

	// Structure Filters
	ChannelIDs []int64

	PublishedAfter  *time.Time
	PublishedBefore *time.Time
	// Result Control
	Limit int

	SortBy    SortField
	SortOrder SortOrder
}

type SortField string

const (
	SortByPublishedAt SortField = "published_at"
	SortByRelevance   SortField = "relevance"
)

type SortOrder string

const (
	Asc  SortOrder = "asc"
	Desc SortOrder = "desc"
)

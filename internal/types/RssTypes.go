package types

import (
	"regexp"
	"strings"
	"time"
)

type Rss struct {
	Channel RssChannel `xml:"channel"`
}

type RssChannel struct {
	Title string    `xml:"title"`
	Link  string    `xml:"link"`
	Desc  string    `xml:"description"`
	Lang  string    `xml:"language"`
	Items []RssItem `xml:"item"`
}

type RssItem struct {
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	Desc    string `xml:"description"`
	PubDate string `xml:"pubDate"`
	GUID    string `xml:"guid"`
}

func (r RssChannel) ToDBChannel() *Channel {
	return &Channel{
		Name: r.Title,
		Slug: Slugify(r.Title),
	}
}

func (r RssItem) ToArticle(publisher string) (*Article, error) {
	publishedAt, err := parsePubDate(r.PubDate)
	if err != nil {
		return nil, err
	}

	return &Article{
		ExternalID:  r.GUID,
		Title:       r.Title,
		URL:         r.Link,
		Summary:     r.Desc,
		PublishedAt: publishedAt,
		PublishedBy: publisher,
	}, nil
}

func (r Rss) Compile() (*Channel, []Article) {
	channel := r.Channel.ToDBChannel()
	var articles []Article
	for _, rssArticle := range r.Channel.Items {
		article, err := rssArticle.ToArticle(channel.Slug)
		if err != nil {
			panic(err)
		}

		articles = append(articles, *article)
	}
	return channel, articles
}

// Helpers
var nonAlnum = regexp.MustCompile(`[^a-z0-9]+`)

func Slugify(name string) string {
	name = strings.ToLower(name)
	name = strings.TrimSpace(name)
	name = nonAlnum.ReplaceAllString(name, "-")
	name = strings.Trim(name, "-")
	return name
}

func parsePubDate(s string) (time.Time, error) {
	layouts := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC3339,
		time.RFC822,
		time.RFC822Z,
	}

	var err error
	for _, layout := range layouts {
		var t time.Time
		t, err = time.Parse(layout, s)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, err
}

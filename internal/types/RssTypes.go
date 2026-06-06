package types

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

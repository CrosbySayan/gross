package types

type RssChannel struct {
	title string `xml:"title"`
	link  string `xml:"link"`
	desc  string `xml:"description"`
	lang  string `xml:"language"`
}

type RssItem struct {
	title   string `xml:"title"`
	link    string `xml:"link"`
	desc    string `xml:"description"`
	pubDate string `xml:"pubDate"`
	guid    string `xml:"guid"`
}


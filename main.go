package main

import (
	"github.com/CrosbySayan/gross/internal/server"
)

func main() {
	rssMan := server.CreateSubManager("session.json")

	rssMan.Start()
	defer rssMan.Close()

	rssMan.Subscribe("Reuters", "https://ir.thomsonreuters.com/rss/news-releases.xml?items=15")
	_, err := rssMan.Fetch("Reuters")
	if err != nil {
		panic(err)
	}

	// fmt.Printf("Title: %s\nLink: %s\nDesc: %s\n", channel.Title, channel.Link, channel.Desc)
}

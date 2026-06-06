package main

import (
	"fmt"

	"github.com/CrosbySayan/gross/internal/db"
	"github.com/CrosbySayan/gross/internal/server"
)

func main() {
	rssMan := server.CreateSubManager("session.json")

	rssMan.Start()
	defer rssMan.Close()

	rssMan.Subscribe("Reuters", "https://ir.thomsonreuters.com/rss/news-releases.xml?items=15")
	rssData, err := rssMan.Fetch("Reuters")
	if err != nil {
		panic(err)
	}
	channel, articles := rssData.Compile()
	// fmt.Printf("Title: %s\nLink: %s\nDesc: %s\n", channel.Title, channel.Link, channel.Desc)
	sqlDB, err := db.NewDataBaseManager("new_db.sql")
	if err != nil {
		panic(err)
	}

	id, err := sqlDB.AddChannel(*channel)

	fmt.Println("Channel ID:", id)

	savedChannel, err := sqlDB.GetChannelBySlug(channel.Slug)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Saved Channel: %+v\n", savedChannel)

	channelID := savedChannel.ID

	for _, article := range articles {
		article.ChannelID = channelID

		id, err := sqlDB.UpsertArticle(article)
		if err != nil {
			panic(err)
		}
		fmt.Println("Article ID:", id)
	}

	articles, err = sqlDB.GetAllArticlesFromChannel(savedChannel.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println("Articles count:", len(articles))
	for i, a := range articles {
		fmt.Printf("%d: %s (%s)\n", i, a.Title, a.PublishedAt)
	}
}

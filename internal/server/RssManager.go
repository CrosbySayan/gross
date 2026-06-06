package server

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/CrosbySayan/gross/internal/types"
)

type RssSubManager struct {
	session string
	subs    map[string]string
}

// Creation
func CreateSubManager(session string) SubManager {
	return &RssSubManager{
		session: session,
		subs:    make(map[string]string),
	}
}

// Contract Completetion

func (r *RssSubManager) Start() {
	// Read from session file
	file, err := os.Open(r.session)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		panic(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&r.subs)
	if err != nil {
		panic(err)
	}
}

func (r *RssSubManager) Close() {
	// Write to Session Manager
	file, err := os.Create(r.session)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(r.subs)
}

func (r *RssSubManager) Subscribe(name string, url string) error {
	// TODO: Add test to validate URL
	r.subs[name] = url
	return nil
}

func (r *RssSubManager) Unsubscribe(name string) error {
	_, found := r.subs[name]

	if !found {
		return errors.New("key given is not currently a subscribed Rss")
	}
	delete(r.subs, name)
	return nil
}

// Fetch returns an RssChannel representing parsed XML from a particular subscribed Rss Feed.
func (r *RssSubManager) Fetch(name string) (*types.Rss, error) {
	RssFeedURL, found := r.subs[name]

	if !found {
		return &types.Rss{}, errors.New("key given is not currently a subscribed Rss")
	}

	resp, err := http.Get(RssFeedURL)
	if err != nil {
		return &types.Rss{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &types.Rss{}, err
	}

	var rss types.Rss
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		return &types.Rss{}, err
	}

	fmt.Printf("Title: %s\nLink: %s\nDesc: %s\n", rss.Channel.Title, rss.Channel.Link, rss.Channel.Desc)
	return &rss, nil
}

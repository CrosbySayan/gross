package server

import (
	"encoding/xml"
	"net/http"

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
}

func (r *RssSubManager) Close() {
	// Write to Session Manager
}

func (r *RssSubManager) Subscribe(name string, url string) error {
	return nil
}

func (r *RssSubManager) Unsubscribe(name string) error {
	return nil
}

func (r *RssSubManager) Fetch(name string) (types.RssChannel, []types.RssItem, error) {
	RssFeedURL := r.subs[name]
}

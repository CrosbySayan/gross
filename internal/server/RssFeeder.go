package server

import (
	"github.com/CrosbySayan/gross/internal/types"
)

type SubManager interface {
	Start()
	Close()
	// Add or Remove Rss from State
	Subscribe(name string, url string) error
	Unsubscribe(name string) error

	Fetch(name string) (types.RssChannel, []types.RssItem, error)
}

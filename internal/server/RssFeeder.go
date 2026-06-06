package server

type RssManager interface {
	// Add or Remove Rss from State
	Subscribe(name string, url string) error
	Unsubscribe(name string) error

	Fetch(name string)
}

package frontier

import (
	"github.com/imthaghost/crawler/pkg/queue"
	"net/url"
)

// frontier
type frontier struct {
	queue queue.Queue // FIFO (First in First Out)
}

// Service ...
type Service interface {
	Search(u url.URL) (url.URL, error)
}

// NewService instantiates a tree data structure that contains
// all the URLs that are to be crawled. Calling the crawl
// frontier will return the next URL that we need to be crawled.
func NewService() Service {
	return &frontier{
		queue: queue.NewQueue(),
		}
}

// Search ...
func  (f *frontier) Search(u url.URL) (url.URL, error) {
	return url.URL{}, nil
}
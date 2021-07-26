package frontier

import (
	"net/url"

	"github.com/imthaghost/crawler/pkg/queue"
)

// frontier
type frontier struct {
	qService queue.Service // FIFO (First in First Out)
}

// NewFrontier instantiates a tree data structure that contains
// all the URLs that are to be crawled. Calling the crawl
// frontier will return the next URL that we need to be crawled.
func NewFrontier() {

}

type Service interface {
	// We can crawl sub URL’s by BFS (breadth first search)
	BFS(url url.URL) (url.URL, error)
	// starting from root page and traverse inside. BFS so that we can keep adding the URL’s to the queue as
}

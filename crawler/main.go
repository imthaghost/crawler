package main

import (
	"github.com/imthaghost/learning/crawler/fetcher"
	"github.com/imthaghost/learning/crawler/renderer"
)

func main() {
	someurl := "golang.org"

	fetcher.Ip_address(someurl)
	renderer.Render(someurl)
}

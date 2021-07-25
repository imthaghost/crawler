package main

import (
	"fmt"

	"github.com/imthaghost/learning/crawler/fetcher"
	"github.com/imthaghost/learning/crawler/renderer"
)

func main() {
	someurl := "https://golang.org"

	ip, err := fetcher.Ip_address(someurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(ip)
	renderer.Render(someurl)
}

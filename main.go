package main

import (
	"fmt"

	"github.com/imthaghost/crawler/fetcher"
	"github.com/imthaghost/crawler/renderer"
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

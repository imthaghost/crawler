package main

import (
	"fmt"
	"github.com/imthaghost/crawler/pkg/queue"

	"github.com/imthaghost/crawler/pkg/fetcher"
	"github.com/imthaghost/crawler/pkg/renderer"
)

func main() {
	url := "https://golang.org"

	ip, err := fetcher.Ip_address(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(ip)
	renderer.Render(url)

	q := queue.NewQueue()
	err = q.Enqueue(url)
	if err != nil {
		panic(err)
	}
	u, err := q.Dequeue()
	if err != nil {
		panic(err)
	}
	fmt.Println(u)

}

package main

import (
	"fmt"
	"github.com/imthaghost/crawler/pkg/fetcher"
)

func main() {
	url := "https://github.com"

	fetcherService := fetcher.NewFetcher()

	ip, err :=  fetcherService.GetIP(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(ip)


}

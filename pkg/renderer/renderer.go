package renderer

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Render(url string) {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", html)
}

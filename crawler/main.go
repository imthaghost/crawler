package crawler

import "github.com/imthaghost/learning/crawler/frontier"

func main() {
	// start passing values into dfs
	someurl := "https://google.com"
	frontier.DFS(someurl)
}

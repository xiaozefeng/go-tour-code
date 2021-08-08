package concurrency

import "testing"

func TestWebCrawl(t *testing.T) {
	Crawl("https://golang.org/", 4, fetcher)
}

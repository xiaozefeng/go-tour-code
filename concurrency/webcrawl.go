package concurrency

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and a slice of URLS found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl use fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// todo: Fetch URLS in parallel
	// todo: Don't fetch the same URL twice
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s, %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
}

type FakeFetcher map[string]*FakeResult

func (f FakeFetcher) Fetch(url string) (body string, urls []string, err error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not foud: %s", url)
}

type FakeResult struct {
	body string
	urls []string
}

var fetcher = FakeFetcher{
	"https://golang.org/": &FakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &FakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &FakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &FakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
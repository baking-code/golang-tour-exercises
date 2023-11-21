package main

import (
	"exercises/web-crawler/utils"
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var urlStore = utils.MakeMap()

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, exitSignal chan bool) {
	if depth <= 0 {
		// depth reached, signal end
		exitSignal <- true
		return
	}
	body, urls, err := fetcher.Fetch(url)
	urlStore.Set(url)
	if err != nil {
		fmt.Println(err)
		exitSignal <- true
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	exit := make(chan bool)
	// count how many crawlers we're starting
	counter := 0
	for _, u := range urls {
		// only crawl url if we haven't seen it yet
		if urlStore.Get(u) == "" {
			counter++
			go Crawl(u, depth-1, fetcher, exit)
		}
	}
	for i := 0; i < counter; i++ {
		// wait for all started crawlers (in current scope) to return
		<-exit
	}
	// return from this crawler
	exitSignal <- true
	return
}

func main() {
	exit := make(chan bool)
	go Crawl("https://golang.org/", 4, fetcher, exit)
	// wait for above crawler to finish
	<-exit
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

// Write a Go program that implements a simple concurrent web crawler.
// The crawler should start from a given URL and visit links on that page, continuing to a specified depth.
// Use Go routines to crawl multiple pages concurrently, and ensure that each page is only visited once.

// Hint: Use a map to track visited URLs and channels to coordinate the crawling process. Use sync.Mutex to manage access to shared resources.
type Crawler struct {
	visited map[string]bool
	mu      sync.Mutex
	wg      sync.WaitGroup
}

func (c *Crawler) crawl(url string, depth int, ch chan string) {
	defer c.wg.Done()
	if depth <= 0 {
		return
	}

	c.mu.Lock()
	if c.visited[url] {
		c.mu.Unlock()
		return
	}
	c.visited[url] = true
	c.mu.Unlock()

	links := fetchLinks(url)
	for _, link := range links {
		ch <- link
		c.wg.Add(1)
		go c.crawl(link, depth-1, ch)
	}
}

func fetchLinks(url string) []string {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	var links []string
	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}
		t := z.Token()
		if t.Type == html.StartTagToken && t.Data == "a" {
			for _, a := range t.Attr {
				if a.Key == "href" && strings.HasPrefix(a.Val, "http") {
					links = append(links, a.Val)
				}
			}
		}
	}
	return links
}

func main() {
	startURL := "https://example.com"
	maxDepth := 2

	c := &Crawler{visited: make(map[string]bool)}
	ch := make(chan string)

	c.wg.Add(1)
	go c.crawl(startURL, maxDepth, ch)

	go func() {
		for link := range ch {
			fmt.Println("Discovered:", link)
		}
	}()

	c.wg.Wait()
	close(ch)
	fmt.Println("Crawling complete.")
}

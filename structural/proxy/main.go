package main

import (
	"fmt"
)

func main() {
	origin := NewOriginServer()
	proxy := NewCachingProxy(origin)
	proxy.Block("https://blocked.example")

	urls := []string{
		"https://example.com/home",
		"https://example.com/home", // cache hit
		"https://example.com/about",
		"https://blocked.example",
	}

	for _, u := range urls {
		content, err := proxy.Fetch(u)
		if err != nil {
			fmt.Println("Fetch error:", err)
			continue
		}
		fmt.Printf("Fetched: %s => %s\n", u, content)
	}

	fmt.Printf("\nCache entries: %d\n", len(proxy.cache))
}

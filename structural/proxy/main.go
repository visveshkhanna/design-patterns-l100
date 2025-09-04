package main

import (
	"fmt"

	"design-patterns/structural/proxy/origin"
	"design-patterns/structural/proxy/proxy"
)

func main() {
	originServer := origin.NewServer()

	cachingProxy := proxy.NewCachingProxy(originServer)

	cachingProxy.Block("https://blocked.example")

	urls := []string{
		"https://example.com/home",
		"https://example.com/home",
		"https://example.com/about",
		"https://blocked.example",
	}

	fmt.Println("=== Proxy Pattern Demo ===")
	for _, url := range urls {
		content, err := cachingProxy.Fetch(url)
		if err != nil {
			fmt.Printf("❌ Error fetching %s: %v\n", url, err)
			continue
		}
		fmt.Printf("✅ Fetched: %s => %s\n", url, content)
	}

	fmt.Printf("\n📊 Final cache size: %d entries\n", cachingProxy.CacheSize())
}

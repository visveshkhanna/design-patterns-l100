package proxy

import (
	"fmt"

	"design-patterns/structural/proxy/interfaces"
)

type CachingProxy struct {
	origin    interfaces.IFetcher
	cache     map[string]string
	blocklist map[string]bool
}

func NewCachingProxy(origin interfaces.IFetcher) *CachingProxy {
	return &CachingProxy{
		origin:    origin,
		cache:     make(map[string]string),
		blocklist: make(map[string]bool),
	}
}

func (p *CachingProxy) Block(url string) {
	p.blocklist[url] = true
}

func (p *CachingProxy) Fetch(url string) (string, error) {
	if p.blocklist[url] {
		return "", fmt.Errorf("access blocked for %s", url)
	}

	if content, ok := p.cache[url]; ok {
		fmt.Printf("Cache hit for %s\n", url)
		return content, nil
	}

	content, err := p.origin.Fetch(url)
	if err != nil {
		return "", err
	}

	p.cache[url] = content
	fmt.Printf("Cached new content for %s\n", url)
	return content, nil
}

func (p *CachingProxy) CacheSize() int {
	return len(p.cache)
}

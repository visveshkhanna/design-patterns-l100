package main

import "fmt"

// CachingProxy controls access to the origin and adds caching and blocking.
type CachingProxy struct {
	origin    Fetcher
	cache     map[string]string
	blocklist map[string]bool
}

func NewCachingProxy(origin Fetcher) *CachingProxy {
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
		return content, nil
	}
	content, err := p.origin.Fetch(url)
	if err != nil {
		return "", err
	}
	p.cache[url] = content
	return content, nil
}

package main

// Fetcher defines the subject interface for fetching resources.
type Fetcher interface {
	Fetch(url string) (string, error)
}

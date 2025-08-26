package main

import "fmt"

// OriginServer is the real subject that performs the actual fetch.
type OriginServer struct{}

func NewOriginServer() *OriginServer {
	return &OriginServer{}
}

func (s *OriginServer) Fetch(url string) (string, error) {
	return fmt.Sprintf("content from origin for %s", url), nil
}

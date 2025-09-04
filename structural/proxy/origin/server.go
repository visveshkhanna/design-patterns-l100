package origin

import (
	"fmt"

	"design-patterns/structural/proxy/interfaces"
)

type Server struct{}

func NewServer() interfaces.IFetcher {
	return &Server{}
}

func (s *Server) Fetch(url string) (string, error) {
	return fmt.Sprintf("content from origin for %s", url), nil
}

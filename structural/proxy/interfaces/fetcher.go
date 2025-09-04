package interfaces

type IFetcher interface {
	Fetch(url string) (string, error)
}

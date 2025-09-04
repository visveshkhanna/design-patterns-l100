package interfaces

type ICard interface {
	Render(title string, body string) string
}

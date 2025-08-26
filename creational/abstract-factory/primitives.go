package main

type Button interface {
	Render() string
}

type Card interface {
	Render(title string, body string) string
}

type ThemeFactory interface {
	MakeButton() Button
	MakeCard() Card
}

// Light
type LightButton struct{}

func (LightButton) Render() string { return "[Light Button]" }

type LightCard struct{}

func (LightCard) Render(title string, body string) string {
	return "LightCard(" + title + "): " + body
}

type LightTheme struct{}

func (LightTheme) MakeButton() Button { return LightButton{} }
func (LightTheme) MakeCard() Card     { return LightCard{} }

// Dark
type DarkButton struct{}

func (DarkButton) Render() string { return "[Dark Button]" }

type DarkCard struct{}

func (DarkCard) Render(title string, body string) string {
	return "DarkCard(" + title + "): " + body
}

type DarkTheme struct{}

func (DarkTheme) MakeButton() Button { return DarkButton{} }
func (DarkTheme) MakeCard() Card     { return DarkCard{} }

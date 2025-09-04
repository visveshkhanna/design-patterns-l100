package themes

import "design-patterns/creational/abstract-factory/interfaces"

type DarkButton struct{}

func (DarkButton) Render() string {
	return "[Dark Button]"
}

type DarkCard struct{}

func (DarkCard) Render(title string, body string) string {
	return "DarkCard(" + title + "): " + body
}

type DarkTheme struct{}

func (DarkTheme) MakeButton() interfaces.IButton {
	return DarkButton{}
}

func (DarkTheme) MakeCard() interfaces.ICard {
	return DarkCard{}
}

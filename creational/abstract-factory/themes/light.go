package themes

import "design-patterns/creational/abstract-factory/interfaces"

type LightButton struct{}

func (LightButton) Render() string {
	return "[Light Button]"
}

type LightCard struct{}

func (LightCard) Render(title string, body string) string {
	return "LightCard(" + title + "): " + body
}

type LightTheme struct{}

func (LightTheme) MakeButton() interfaces.IButton {
	return LightButton{}
}

func (LightTheme) MakeCard() interfaces.ICard {
	return LightCard{}
}

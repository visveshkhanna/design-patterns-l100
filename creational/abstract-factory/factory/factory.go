package factory

import (
	"design-patterns/creational/abstract-factory/interfaces"
	"design-patterns/creational/abstract-factory/themes"
	"fmt"
)

func GetThemeFactory(theme string) (interfaces.IThemeFactory, error) {
	switch theme {
	case "light":
		return &themes.LightTheme{}, nil
	case "dark":
		return &themes.DarkTheme{}, nil
	default:
		return nil, fmt.Errorf("unknown theme: %s", theme)
	}
}

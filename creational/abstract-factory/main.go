package main

import (
	"fmt"
)

func getThemeFactory(theme string) (ThemeFactory, error) {
	if theme == "light" {
		return &LightTheme{}, nil
	}
	if theme == "dark" {
		return &DarkTheme{}, nil
	}
	return nil, fmt.Errorf("unknown theme: %s", theme)
}

func main() {
	light, _ := getThemeFactory("light")
	dark, _ := getThemeFactory("dark")

	render := func(name string, f ThemeFactory) {
		btn := f.MakeButton()
		card := f.MakeCard()
		fmt.Println("--", name, "--")
		fmt.Println(btn.Render())
		fmt.Println(card.Render("Welcome", "Minimal themed UI"))
	}

	render("Light", light)
	render("Dark", dark)
}

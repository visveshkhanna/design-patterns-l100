package main

import (
	"design-patterns/creational/abstract-factory/factory"
	"design-patterns/creational/abstract-factory/interfaces"
	"fmt"
)

func main() {

	light, _ := factory.GetThemeFactory("light")
	dark, _ := factory.GetThemeFactory("dark")

	render := func(name string, f interfaces.IThemeFactory) {
		btn := f.MakeButton()
		card := f.MakeCard()
		fmt.Println("--", name, "--")
		fmt.Println(btn.Render())
		fmt.Println(card.Render("Welcome", "Minimal themed UI"))
	}

	render("Light", light)
	render("Dark", dark)
}

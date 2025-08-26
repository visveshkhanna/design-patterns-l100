package main

import "fmt"

type Mediator interface {
	notify(sender string, event string)
}

type Dialog struct {
	button *Button
	input  *Input
}

func (d *Dialog) notify(sender string, event string) {
	if sender == "button" && event == "click" {
		fmt.Println("Dialog: Button clicked, enabling input")
		d.input.enable()
	}
	if sender == "input" && event == "change" {
		fmt.Println("Dialog: Input changed, disabling button")
		d.button.disable()
	}
}

package main

import "fmt"

type Light struct{ isOn bool }

func (l *Light) on() {
	l.isOn = true
	fmt.Println("Light is ON")
}

func (l *Light) off() {
	l.isOn = false
	fmt.Println("Light is OFF")
}

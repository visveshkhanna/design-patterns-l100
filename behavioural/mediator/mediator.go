package main

import "fmt"

type Mediator interface {
	notify(sender string, event string)
}

type Greenhouse struct {
	fan    *Fan
	heater *Heater
	sensor *Sensor
}

func (g *Greenhouse) notify(sender string, event string) {
	if sender == "sensor" && event == "hot" {
		fmt.Println("Greenhouse: too hot, turning fan on, heater off")
		g.fan.on()
		g.heater.off()
		return
	}
	if sender == "sensor" && event == "cold" {
		fmt.Println("Greenhouse: too cold, turning heater on, fan off")
		g.heater.on()
		g.fan.off()
		return
	}
	if sender == "sensor" && event == "ok" {
		fmt.Println("Greenhouse: conditions ok, turning everything off")
		g.heater.off()
		g.fan.off()
		return
	}
}

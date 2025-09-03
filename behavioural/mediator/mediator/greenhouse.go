package mediator

import (
	"design-patterns/behavioural/mediator/components"
	"design-patterns/behavioural/mediator/interfaces"
	"fmt"
)

type Greenhouse struct {
	fan    *components.Fan
	heater *components.Heater
	sensor *components.Sensor
}

var _ interfaces.Mediator = (*Greenhouse)(nil)

func NewGreenhouse() *Greenhouse {
	gh := &Greenhouse{}

	gh.fan = components.NewFan(gh)
	gh.heater = components.NewHeater(gh)
	gh.sensor = components.NewSensor(gh)

	return gh
}

func (g *Greenhouse) Notify(sender string, event string) {
	switch {
	case sender == "sensor" && event == "hot":
		fmt.Println("Greenhouse: too hot, turning fan on, heater off")
		g.fan.On()
		g.heater.Off()
	case sender == "sensor" && event == "cold":
		fmt.Println("Greenhouse: too cold, turning heater on, fan off")
		g.heater.On()
		g.fan.Off()
	case sender == "sensor" && event == "ok":
		fmt.Println("Greenhouse: conditions ok, turning everything off")
		g.heater.Off()
		g.fan.Off()
	}
}

func (g *Greenhouse) GetSensor() *components.Sensor {
	return g.sensor
}

package main

import "design-patterns/behavioural/mediator/mediator"

func main() {

	gh := mediator.NewGreenhouse()

	sensor := gh.GetSensor()

	sensor.TempHot()
	sensor.TempCold()
	sensor.TempOk()
}

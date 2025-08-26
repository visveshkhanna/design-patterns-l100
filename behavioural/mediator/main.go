package main

func main() {
	gh := &Greenhouse{}
	fan := &Fan{mediator: gh}
	heater := &Heater{mediator: gh}
	sensor := &Sensor{mediator: gh}
	gh.fan = fan
	gh.heater = heater
	gh.sensor = sensor

	sensor.tempHot()
	sensor.tempCold()
	sensor.tempOk()
}

package main

type Fan struct {
	mediator Mediator
	onState  bool
}

type Heater struct {
	mediator Mediator
	onState  bool
}

type Sensor struct{ mediator Mediator }

func (f *Fan) on()  { f.onState = true }
func (f *Fan) off() { f.onState = false }

func (h *Heater) on()  { h.onState = true }
func (h *Heater) off() { h.onState = false }

func (s *Sensor) tempHot()  { s.mediator.notify("sensor", "hot") }
func (s *Sensor) tempCold() { s.mediator.notify("sensor", "cold") }
func (s *Sensor) tempOk()   { s.mediator.notify("sensor", "ok") }

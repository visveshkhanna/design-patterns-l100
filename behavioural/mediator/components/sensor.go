package components

import "design-patterns/behavioural/mediator/interfaces"

type Sensor struct {
	mediator interfaces.IMediator
}

func NewSensor(m interfaces.IMediator) *Sensor {
	return &Sensor{
		mediator: m,
	}
}

func (s *Sensor) TempHot() {
	s.mediator.Notify("sensor", "hot")
}

func (s *Sensor) TempCold() {
	s.mediator.Notify("sensor", "cold")
}

func (s *Sensor) TempOk() {
	s.mediator.Notify("sensor", "ok")
}

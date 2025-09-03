package components

import "design-patterns/behavioural/mediator/interfaces"

type Heater struct {
	mediator interfaces.Mediator
	onState  bool
}

func NewHeater(m interfaces.Mediator) *Heater {
	return &Heater{
		mediator: m,
		onState:  false,
	}
}

func (h *Heater) On() {
	h.onState = true
}

func (h *Heater) Off() {
	h.onState = false
}

func (h *Heater) IsOn() bool {
	return h.onState
}

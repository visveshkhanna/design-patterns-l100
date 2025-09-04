package components

import "design-patterns/behavioural/mediator/interfaces"

type Fan struct {
	mediator interfaces.IMediator
	onState  bool
}

func NewFan(m interfaces.IMediator) *Fan {
	return &Fan{
		mediator: m,
		onState:  false,
	}
}

func (f *Fan) On() {
	f.onState = true
}

func (f *Fan) Off() {
	f.onState = false
}

func (f *Fan) IsOn() bool {
	return f.onState
}

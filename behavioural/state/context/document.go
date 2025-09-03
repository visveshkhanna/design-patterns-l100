package context

import "design-patterns/behavioural/state/interfaces"

type Document struct {
	state interfaces.State
	Title string
}

func NewDocument(title string, initialState interfaces.State) *Document {
	return &Document{
		state: initialState,
		Title: title,
	}
}

func (d *Document) SetState(s interfaces.State) {
	d.state = s
}

func (d *Document) Submit() {
	d.state.Submit(d)
}

func (d *Document) Approve() {
	d.state.Approve(d)
}

func (d *Document) Reject() {
	d.state.Reject(d)
}

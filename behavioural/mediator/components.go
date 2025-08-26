package main

type Button struct {
	mediator Mediator
	enabled  bool
}

type Input struct {
	mediator Mediator
	enabled  bool
}

func (b *Button) click() { b.mediator.notify("button", "click") }

func (b *Button) disable() { b.enabled = false }

func (i *Input) change() { i.mediator.notify("input", "change") }

func (i *Input) enable() { i.enabled = true }

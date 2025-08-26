package main

// Command defines the executable action
type Command interface {
	execute()
}

type OnCommand struct{ device *Light }

type OffCommand struct{ device *Light }

func (c *OnCommand) execute() { c.device.on() }

func (c *OffCommand) execute() { c.device.off() }

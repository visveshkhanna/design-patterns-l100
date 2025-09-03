package interfaces

type Mediator interface {
	Notify(sender string, event string)
}

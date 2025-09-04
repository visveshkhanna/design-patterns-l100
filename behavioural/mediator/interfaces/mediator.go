package interfaces

type IMediator interface {
	Notify(sender string, event string)
}

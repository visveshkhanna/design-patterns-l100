package interfaces

type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	NotifyAll()
}

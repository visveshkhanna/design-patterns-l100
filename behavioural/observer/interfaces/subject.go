package interfaces

type ISubject interface {
	Register(observer IObserver)
	Unregister(observer IObserver)
	NotifyAll()
}

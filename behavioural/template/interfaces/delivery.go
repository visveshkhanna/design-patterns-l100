package interfaces

type Delivery interface {
	Prepare()
	Dispatch()
	Report()
}

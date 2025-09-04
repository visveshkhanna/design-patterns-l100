package interfaces

type IDelivery interface {
	Prepare()
	Dispatch()
	Report()
}

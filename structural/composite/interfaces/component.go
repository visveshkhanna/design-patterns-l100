package interfaces

type IComponent interface {
	Price() int
	Print(indent string)
}

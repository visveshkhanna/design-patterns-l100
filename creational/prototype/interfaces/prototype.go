package interfaces

type IPrototype interface {
	Print(indentation string)
	Clone() IPrototype
}

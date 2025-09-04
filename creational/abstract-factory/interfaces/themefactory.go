package interfaces

type IThemeFactory interface {
	MakeButton() IButton
	MakeCard() ICard
}

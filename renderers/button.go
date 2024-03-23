package renderers

type Button struct {
	*VanillaAction
}

func NewButton() *Button{
	return &Button{
		VanillaAction: NewVanillaAction(),
	}
}
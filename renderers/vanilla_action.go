package renderers

type VanillaAction struct {
	*BaseRenderer
}

func NewVanillaAction() *VanillaAction{
	return &VanillaAction{
		BaseRenderer: NewBaseRenderer(),
	}
}


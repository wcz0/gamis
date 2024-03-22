package renderers

type Component struct {
	*BaseRenderer
}

func NewComponent() *Component {
	return &Component{
		BaseRenderer: NewBaseRenderer(),
	}
}

func (c *Component) SetType(value string) *Component {
	c.Set("type", value)
	return c
}
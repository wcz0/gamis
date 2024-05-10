package renderers

type Html struct {
	*BaseRenderer
}

func NewHtml() *Html {
	a := &Html{
		BaseRenderer: NewBaseRenderer(),
	}
	a.Set("type", "html")
	return a
}
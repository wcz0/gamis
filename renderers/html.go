package renderers

type Html struct {
	*Tpl
}

func NewHtml() *Tpl {
	return NewTpl()
}
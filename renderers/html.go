package renderers

type Html struct {
	*Tpl
}

func NewHtml() *Html {
	return &Html{
		Tpl: NewTpl(),
	}
}
package renderers

type Amis struct {}

func NewAmis() *Amis{
	return &Amis{}
}

func (a *Amis) Action() *Action{
	return NewAction()
}

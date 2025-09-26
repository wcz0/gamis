package renderers

type Markdown struct {
	*BaseRenderer
}

func NewMarkdown() *Markdown {
	a := &Markdown{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "markdown")
	return a
}

func (a *Markdown) Set(name string, value interface{}) *Markdown {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Markdown) ClassName(value interface{}) *Markdown {
	a.Set("className", value)
	return a
}

func (a *Markdown) Name(value interface{}) *Markdown {
	a.Set("name", value)
	return a
}

func (a *Markdown) Options(value interface{}) *Markdown {
	a.Set("options", value)
	return a
}

func (a *Markdown) Src(value interface{}) *Markdown {
	a.Set("src", value)
	return a
}

func (a *Markdown) Type(value interface{}) *Markdown {
	a.Set("type", value)
	return a
}

func (a *Markdown) Value(value interface{}) *Markdown {
	a.Set("value", value)
	return a
}

package renderers

type GridNav struct {
	*BaseRenderer
}

func NewGridNav() *GridNav {
	a := &GridNav{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "grid-nav")
	return a
}

func (a *GridNav) Set(name string, value interface{}) *GridNav {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *GridNav) Border(value interface{}) *GridNav {
	a.Set("border", value)
	return a
}

func (a *GridNav) Center(value interface{}) *GridNav {
	a.Set("center", value)
	return a
}

func (a *GridNav) ClassName(value interface{}) *GridNav {
	a.Set("className", value)
	return a
}

func (a *GridNav) ColumnNum(value interface{}) *GridNav {
	a.Set("columnNum", value)
	return a
}

func (a *GridNav) Direction(value interface{}) *GridNav {
	a.Set("direction", value)
	return a
}

func (a *GridNav) Gutter(value interface{}) *GridNav {
	a.Set("gutter", value)
	return a
}

func (a *GridNav) IconRatio(value interface{}) *GridNav {
	a.Set("iconRatio", value)
	return a
}

func (a *GridNav) ItemClassName(value interface{}) *GridNav {
	a.Set("itemClassName", value)
	return a
}

func (a *GridNav) Options(value interface{}) *GridNav {
	a.Set("options", value)
	return a
}

func (a *GridNav) Reverse(value interface{}) *GridNav {
	a.Set("reverse", value)
	return a
}

func (a *GridNav) Source(value interface{}) *GridNav {
	a.Set("source", value)
	return a
}

func (a *GridNav) Square(value interface{}) *GridNav {
	a.Set("square", value)
	return a
}

func (a *GridNav) Type(value interface{}) *GridNav {
	a.Set("type", value)
	return a
}

func (a *GridNav) Value(value interface{}) *GridNav {
	a.Set("value", value)
	return a
}

package renderers

type Property struct {
	*BaseRenderer
}

func NewProperty() *Property {
	a := &Property{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "property")
	return a
}

func (a *Property) Set(name string, value interface{}) *Property {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Property) ClassName(value interface{}) *Property {
	a.Set("className", value)
	return a
}

func (a *Property) Column(value interface{}) *Property {
	a.Set("column", value)
	return a
}

func (a *Property) ContentStyle(value interface{}) *Property {
	a.Set("contentStyle", value)
	return a
}

func (a *Property) Items(value interface{}) *Property {
	a.Set("items", value)
	return a
}

func (a *Property) LabelStyle(value interface{}) *Property {
	a.Set("labelStyle", value)
	return a
}

func (a *Property) Mode(value interface{}) *Property {
	a.Set("mode", value)
	return a
}

func (a *Property) Separator(value interface{}) *Property {
	a.Set("separator", value)
	return a
}

func (a *Property) Source(value interface{}) *Property {
	a.Set("source", value)
	return a
}

func (a *Property) Style(value interface{}) *Property {
	a.Set("style", value)
	return a
}

func (a *Property) Title(value interface{}) *Property {
	a.Set("title", value)
	return a
}

func (a *Property) Type(value interface{}) *Property {
	a.Set("type", value)
	return a
}

package renderers

type WebComponent struct {
	*BaseRenderer
}

func NewWebComponent() *WebComponent {
	a := &WebComponent{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "web-component")
	return a
}

func (a *WebComponent) Set(name string, value interface{}) *WebComponent {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *WebComponent) Body(value interface{}) *WebComponent {
	a.Set("body", value)
	return a
}

func (a *WebComponent) Props(value interface{}) *WebComponent {
	a.Set("props", value)
	return a
}

func (a *WebComponent) Tag(value interface{}) *WebComponent {
	a.Set("tag", value)
	return a
}

func (a *WebComponent) Type(value interface{}) *WebComponent {
	a.Set("type", value)
	return a
}

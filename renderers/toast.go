package renderers

type Toast struct {
	*BaseRenderer
}

func NewToast() *Toast {
	a := &Toast{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *Toast) Set(name string, value interface{}) *Toast {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Toast) Body(value interface{}) *Toast {
	a.Set("body", value)
	return a
}

func (a *Toast) CloseButton(value interface{}) *Toast {
	a.Set("closeButton", value)
	return a
}

func (a *Toast) Items(value interface{}) *Toast {
	a.Set("items", value)
	return a
}

func (a *Toast) Level(value interface{}) *Toast {
	a.Set("level", value)
	return a
}

func (a *Toast) Position(value interface{}) *Toast {
	a.Set("position", value)
	return a
}

func (a *Toast) ShowIcon(value interface{}) *Toast {
	a.Set("showIcon", value)
	return a
}

func (a *Toast) Timeout(value interface{}) *Toast {
	a.Set("timeout", value)
	return a
}

func (a *Toast) Title(value interface{}) *Toast {
	a.Set("title", value)
	return a
}

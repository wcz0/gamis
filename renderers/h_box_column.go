package renderers

type HBoxColumn struct {
	*BaseRenderer
}

func NewHBoxColumn() *HBoxColumn {
	a := &HBoxColumn{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *HBoxColumn) Set(name string, value interface{}) *HBoxColumn {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *HBoxColumn) Body(value interface{}) *HBoxColumn {
	a.Set("body", value)
	return a
}

func (a *HBoxColumn) ColumnClassName(value interface{}) *HBoxColumn {
	a.Set("columnClassName", value)
	return a
}

func (a *HBoxColumn) Height(value interface{}) *HBoxColumn {
	a.Set("height", value)
	return a
}

func (a *HBoxColumn) Horizontal(value interface{}) *HBoxColumn {
	a.Set("horizontal", value)
	return a
}

func (a *HBoxColumn) Mode(value interface{}) *HBoxColumn {
	a.Set("mode", value)
	return a
}

func (a *HBoxColumn) Style(value interface{}) *HBoxColumn {
	a.Set("style", value)
	return a
}

func (a *HBoxColumn) Valign(value interface{}) *HBoxColumn {
	a.Set("valign", value)
	return a
}

func (a *HBoxColumn) Visible(value interface{}) *HBoxColumn {
	a.Set("visible", value)
	return a
}

func (a *HBoxColumn) VisibleOn(value interface{}) *HBoxColumn {
	a.Set("visibleOn", value)
	return a
}

func (a *HBoxColumn) Width(value interface{}) *HBoxColumn {
	a.Set("width", value)
	return a
}

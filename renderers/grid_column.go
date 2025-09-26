package renderers

type GridColumn struct {
	*BaseRenderer
}

func NewGridColumn() *GridColumn {
	a := &GridColumn{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *GridColumn) Set(name string, value interface{}) *GridColumn {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *GridColumn) Body(value interface{}) *GridColumn {
	a.Set("body", value)
	return a
}

func (a *GridColumn) ColumnClassName(value interface{}) *GridColumn {
	a.Set("columnClassName", value)
	return a
}

func (a *GridColumn) Horizontal(value interface{}) *GridColumn {
	a.Set("horizontal", value)
	return a
}

func (a *GridColumn) Id(value interface{}) *GridColumn {
	a.Set("id", value)
	return a
}

func (a *GridColumn) Lg(value interface{}) *GridColumn {
	a.Set("lg", value)
	return a
}

func (a *GridColumn) Md(value interface{}) *GridColumn {
	a.Set("md", value)
	return a
}

func (a *GridColumn) Mode(value interface{}) *GridColumn {
	a.Set("mode", value)
	return a
}

func (a *GridColumn) Sm(value interface{}) *GridColumn {
	a.Set("sm", value)
	return a
}

func (a *GridColumn) Style(value interface{}) *GridColumn {
	a.Set("style", value)
	return a
}

func (a *GridColumn) ThemeCss(value interface{}) *GridColumn {
	a.Set("themeCss", value)
	return a
}

func (a *GridColumn) Valign(value interface{}) *GridColumn {
	a.Set("valign", value)
	return a
}

func (a *GridColumn) WrapperCustomStyle(value interface{}) *GridColumn {
	a.Set("wrapperCustomStyle", value)
	return a
}

func (a *GridColumn) Xs(value interface{}) *GridColumn {
	a.Set("xs", value)
	return a
}

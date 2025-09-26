package renderers

type SearchBox struct {
	*BaseRenderer
}

func NewSearchBox() *SearchBox {
	a := &SearchBox{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "search-box")
	return a
}

func (a *SearchBox) Set(name string, value interface{}) *SearchBox {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *SearchBox) ClassName(value interface{}) *SearchBox {
	a.Set("className", value)
	return a
}

func (a *SearchBox) ClearAndSubmit(value interface{}) *SearchBox {
	a.Set("clearAndSubmit", value)
	return a
}

func (a *SearchBox) Clearable(value interface{}) *SearchBox {
	a.Set("clearable", value)
	return a
}

func (a *SearchBox) Disabled(value interface{}) *SearchBox {
	a.Set("disabled", value)
	return a
}

func (a *SearchBox) DisabledOn(value interface{}) *SearchBox {
	a.Set("disabledOn", value)
	return a
}

func (a *SearchBox) EditorSetting(value interface{}) *SearchBox {
	a.Set("editorSetting", value)
	return a
}

func (a *SearchBox) Enhance(value interface{}) *SearchBox {
	a.Set("enhance", value)
	return a
}

func (a *SearchBox) Hidden(value interface{}) *SearchBox {
	a.Set("hidden", value)
	return a
}

func (a *SearchBox) HiddenOn(value interface{}) *SearchBox {
	a.Set("hiddenOn", value)
	return a
}

func (a *SearchBox) Id(value interface{}) *SearchBox {
	a.Set("id", value)
	return a
}

func (a *SearchBox) Loading(value interface{}) *SearchBox {
	a.Set("loading", value)
	return a
}

func (a *SearchBox) Mini(value interface{}) *SearchBox {
	a.Set("mini", value)
	return a
}

func (a *SearchBox) Name(value interface{}) *SearchBox {
	a.Set("name", value)
	return a
}

func (a *SearchBox) OnEvent(value interface{}) *SearchBox {
	a.Set("onEvent", value)
	return a
}

func (a *SearchBox) Placeholder(value interface{}) *SearchBox {
	a.Set("placeholder", value)
	return a
}

func (a *SearchBox) SearchImediately(value interface{}) *SearchBox {
	a.Set("searchImediately", value)
	return a
}

func (a *SearchBox) Static(value interface{}) *SearchBox {
	a.Set("static", value)
	return a
}

func (a *SearchBox) StaticClassName(value interface{}) *SearchBox {
	a.Set("staticClassName", value)
	return a
}

func (a *SearchBox) StaticInputClassName(value interface{}) *SearchBox {
	a.Set("staticInputClassName", value)
	return a
}

func (a *SearchBox) StaticLabelClassName(value interface{}) *SearchBox {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *SearchBox) StaticOn(value interface{}) *SearchBox {
	a.Set("staticOn", value)
	return a
}

func (a *SearchBox) StaticPlaceholder(value interface{}) *SearchBox {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *SearchBox) StaticSchema(value interface{}) *SearchBox {
	a.Set("staticSchema", value)
	return a
}

func (a *SearchBox) Style(value interface{}) *SearchBox {
	a.Set("style", value)
	return a
}

func (a *SearchBox) TestIdBuilder(value interface{}) *SearchBox {
	a.Set("testIdBuilder", value)
	return a
}

func (a *SearchBox) Testid(value interface{}) *SearchBox {
	a.Set("testid", value)
	return a
}

func (a *SearchBox) Type(value interface{}) *SearchBox {
	a.Set("type", value)
	return a
}

func (a *SearchBox) UseMobileUI(value interface{}) *SearchBox {
	a.Set("useMobileUI", value)
	return a
}

func (a *SearchBox) Visible(value interface{}) *SearchBox {
	a.Set("visible", value)
	return a
}

func (a *SearchBox) VisibleOn(value interface{}) *SearchBox {
	a.Set("visibleOn", value)
	return a
}

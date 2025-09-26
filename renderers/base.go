package renderers

type Base struct {
	*BaseRenderer
}

func NewBase() *Base {
	a := &Base{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "form")
	return a
}

func (a *Base) Set(name string, value interface{}) *Base {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Base) ClassName(value interface{}) *Base {
	a.Set("className", value)
	return a
}

func (a *Base) Disabled(value interface{}) *Base {
	a.Set("disabled", value)
	return a
}

func (a *Base) DisabledOn(value interface{}) *Base {
	a.Set("disabledOn", value)
	return a
}

func (a *Base) EditorSetting(value interface{}) *Base {
	a.Set("editorSetting", value)
	return a
}

func (a *Base) Hidden(value interface{}) *Base {
	a.Set("hidden", value)
	return a
}

func (a *Base) HiddenOn(value interface{}) *Base {
	a.Set("hiddenOn", value)
	return a
}

func (a *Base) Id(value interface{}) *Base {
	a.Set("id", value)
	return a
}

func (a *Base) OnEvent(value interface{}) *Base {
	a.Set("onEvent", value)
	return a
}

func (a *Base) Static(value interface{}) *Base {
	a.Set("static", value)
	return a
}

func (a *Base) StaticClassName(value interface{}) *Base {
	a.Set("staticClassName", value)
	return a
}

func (a *Base) StaticInputClassName(value interface{}) *Base {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Base) StaticLabelClassName(value interface{}) *Base {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Base) StaticOn(value interface{}) *Base {
	a.Set("staticOn", value)
	return a
}

func (a *Base) StaticPlaceholder(value interface{}) *Base {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Base) StaticSchema(value interface{}) *Base {
	a.Set("staticSchema", value)
	return a
}

func (a *Base) Style(value interface{}) *Base {
	a.Set("style", value)
	return a
}

func (a *Base) TestIdBuilder(value interface{}) *Base {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Base) Testid(value interface{}) *Base {
	a.Set("testid", value)
	return a
}

func (a *Base) Type(value interface{}) *Base {
	a.Set("type", value)
	return a
}

func (a *Base) UseMobileUI(value interface{}) *Base {
	a.Set("useMobileUI", value)
	return a
}

func (a *Base) Visible(value interface{}) *Base {
	a.Set("visible", value)
	return a
}

func (a *Base) VisibleOn(value interface{}) *Base {
	a.Set("visibleOn", value)
	return a
}

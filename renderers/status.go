package renderers

type Status struct {
	*BaseRenderer
}

func NewStatus() *Status {
	a := &Status{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "status")
	return a
}

func (a *Status) Set(name string, value interface{}) *Status {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Status) ClassName(value interface{}) *Status {
	a.Set("className", value)
	return a
}

func (a *Status) Disabled(value interface{}) *Status {
	a.Set("disabled", value)
	return a
}

func (a *Status) DisabledOn(value interface{}) *Status {
	a.Set("disabledOn", value)
	return a
}

func (a *Status) EditorSetting(value interface{}) *Status {
	a.Set("editorSetting", value)
	return a
}

func (a *Status) Hidden(value interface{}) *Status {
	a.Set("hidden", value)
	return a
}

func (a *Status) HiddenOn(value interface{}) *Status {
	a.Set("hiddenOn", value)
	return a
}

func (a *Status) Id(value interface{}) *Status {
	a.Set("id", value)
	return a
}

func (a *Status) LabelMap(value interface{}) *Status {
	a.Set("labelMap", value)
	return a
}

func (a *Status) Map(value interface{}) *Status {
	a.Set("map", value)
	return a
}

func (a *Status) OnEvent(value interface{}) *Status {
	a.Set("onEvent", value)
	return a
}

func (a *Status) Placeholder(value interface{}) *Status {
	a.Set("placeholder", value)
	return a
}

func (a *Status) Source(value interface{}) *Status {
	a.Set("source", value)
	return a
}

func (a *Status) Static(value interface{}) *Status {
	a.Set("static", value)
	return a
}

func (a *Status) StaticClassName(value interface{}) *Status {
	a.Set("staticClassName", value)
	return a
}

func (a *Status) StaticInputClassName(value interface{}) *Status {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Status) StaticLabelClassName(value interface{}) *Status {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Status) StaticOn(value interface{}) *Status {
	a.Set("staticOn", value)
	return a
}

func (a *Status) StaticPlaceholder(value interface{}) *Status {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Status) StaticSchema(value interface{}) *Status {
	a.Set("staticSchema", value)
	return a
}

func (a *Status) Style(value interface{}) *Status {
	a.Set("style", value)
	return a
}

func (a *Status) TestIdBuilder(value interface{}) *Status {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Status) Testid(value interface{}) *Status {
	a.Set("testid", value)
	return a
}

func (a *Status) Type(value interface{}) *Status {
	a.Set("type", value)
	return a
}

func (a *Status) UseMobileUI(value interface{}) *Status {
	a.Set("useMobileUI", value)
	return a
}

func (a *Status) Visible(value interface{}) *Status {
	a.Set("visible", value)
	return a
}

func (a *Status) VisibleOn(value interface{}) *Status {
	a.Set("visibleOn", value)
	return a
}

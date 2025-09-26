package renderers

type Date struct {
	*BaseRenderer
}

func NewDate() *Date {
	a := &Date{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "date")
	return a
}

func (a *Date) Set(name string, value interface{}) *Date {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Date) ClassName(value interface{}) *Date {
	a.Set("className", value)
	return a
}

func (a *Date) Disabled(value interface{}) *Date {
	a.Set("disabled", value)
	return a
}

func (a *Date) DisabledOn(value interface{}) *Date {
	a.Set("disabledOn", value)
	return a
}

func (a *Date) DisplayFormat(value interface{}) *Date {
	a.Set("displayFormat", value)
	return a
}

func (a *Date) DisplayTimeZone(value interface{}) *Date {
	a.Set("displayTimeZone", value)
	return a
}

func (a *Date) EditorSetting(value interface{}) *Date {
	a.Set("editorSetting", value)
	return a
}

func (a *Date) Format(value interface{}) *Date {
	a.Set("format", value)
	return a
}

func (a *Date) FromNow(value interface{}) *Date {
	a.Set("fromNow", value)
	return a
}

func (a *Date) Hidden(value interface{}) *Date {
	a.Set("hidden", value)
	return a
}

func (a *Date) HiddenOn(value interface{}) *Date {
	a.Set("hiddenOn", value)
	return a
}

func (a *Date) Id(value interface{}) *Date {
	a.Set("id", value)
	return a
}

func (a *Date) OnEvent(value interface{}) *Date {
	a.Set("onEvent", value)
	return a
}

func (a *Date) Placeholder(value interface{}) *Date {
	a.Set("placeholder", value)
	return a
}

func (a *Date) Static(value interface{}) *Date {
	a.Set("static", value)
	return a
}

func (a *Date) StaticClassName(value interface{}) *Date {
	a.Set("staticClassName", value)
	return a
}

func (a *Date) StaticInputClassName(value interface{}) *Date {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Date) StaticLabelClassName(value interface{}) *Date {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Date) StaticOn(value interface{}) *Date {
	a.Set("staticOn", value)
	return a
}

func (a *Date) StaticPlaceholder(value interface{}) *Date {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Date) StaticSchema(value interface{}) *Date {
	a.Set("staticSchema", value)
	return a
}

func (a *Date) Style(value interface{}) *Date {
	a.Set("style", value)
	return a
}

func (a *Date) TestIdBuilder(value interface{}) *Date {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Date) Testid(value interface{}) *Date {
	a.Set("testid", value)
	return a
}

func (a *Date) Type(value interface{}) *Date {
	a.Set("type", value)
	return a
}

func (a *Date) UpdateFrequency(value interface{}) *Date {
	a.Set("updateFrequency", value)
	return a
}

func (a *Date) UseMobileUI(value interface{}) *Date {
	a.Set("useMobileUI", value)
	return a
}

func (a *Date) ValueFormat(value interface{}) *Date {
	a.Set("valueFormat", value)
	return a
}

func (a *Date) Visible(value interface{}) *Date {
	a.Set("visible", value)
	return a
}

func (a *Date) VisibleOn(value interface{}) *Date {
	a.Set("visibleOn", value)
	return a
}

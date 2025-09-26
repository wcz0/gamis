package renderers

type DateRange struct {
	*BaseRenderer
}

func NewDateRange() *DateRange {
	a := &DateRange{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "date-range")
	return a
}

func (a *DateRange) Set(name string, value interface{}) *DateRange {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *DateRange) ClassName(value interface{}) *DateRange {
	a.Set("className", value)
	return a
}

func (a *DateRange) Connector(value interface{}) *DateRange {
	a.Set("connector", value)
	return a
}

func (a *DateRange) Delimiter(value interface{}) *DateRange {
	a.Set("delimiter", value)
	return a
}

func (a *DateRange) Disabled(value interface{}) *DateRange {
	a.Set("disabled", value)
	return a
}

func (a *DateRange) DisabledOn(value interface{}) *DateRange {
	a.Set("disabledOn", value)
	return a
}

func (a *DateRange) DisplayFormat(value interface{}) *DateRange {
	a.Set("displayFormat", value)
	return a
}

func (a *DateRange) EditorSetting(value interface{}) *DateRange {
	a.Set("editorSetting", value)
	return a
}

func (a *DateRange) Format(value interface{}) *DateRange {
	a.Set("format", value)
	return a
}

func (a *DateRange) Hidden(value interface{}) *DateRange {
	a.Set("hidden", value)
	return a
}

func (a *DateRange) HiddenOn(value interface{}) *DateRange {
	a.Set("hiddenOn", value)
	return a
}

func (a *DateRange) Id(value interface{}) *DateRange {
	a.Set("id", value)
	return a
}

func (a *DateRange) OnEvent(value interface{}) *DateRange {
	a.Set("onEvent", value)
	return a
}

func (a *DateRange) Static(value interface{}) *DateRange {
	a.Set("static", value)
	return a
}

func (a *DateRange) StaticClassName(value interface{}) *DateRange {
	a.Set("staticClassName", value)
	return a
}

func (a *DateRange) StaticInputClassName(value interface{}) *DateRange {
	a.Set("staticInputClassName", value)
	return a
}

func (a *DateRange) StaticLabelClassName(value interface{}) *DateRange {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *DateRange) StaticOn(value interface{}) *DateRange {
	a.Set("staticOn", value)
	return a
}

func (a *DateRange) StaticPlaceholder(value interface{}) *DateRange {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *DateRange) StaticSchema(value interface{}) *DateRange {
	a.Set("staticSchema", value)
	return a
}

func (a *DateRange) Style(value interface{}) *DateRange {
	a.Set("style", value)
	return a
}

func (a *DateRange) TestIdBuilder(value interface{}) *DateRange {
	a.Set("testIdBuilder", value)
	return a
}

func (a *DateRange) Testid(value interface{}) *DateRange {
	a.Set("testid", value)
	return a
}

func (a *DateRange) Type(value interface{}) *DateRange {
	a.Set("type", value)
	return a
}

func (a *DateRange) UseMobileUI(value interface{}) *DateRange {
	a.Set("useMobileUI", value)
	return a
}

func (a *DateRange) ValueFormat(value interface{}) *DateRange {
	a.Set("valueFormat", value)
	return a
}

func (a *DateRange) Visible(value interface{}) *DateRange {
	a.Set("visible", value)
	return a
}

func (a *DateRange) VisibleOn(value interface{}) *DateRange {
	a.Set("visibleOn", value)
	return a
}

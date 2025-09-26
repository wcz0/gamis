package renderers

type Color struct {
	*BaseRenderer
}

func NewColor() *Color {
	a := &Color{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "color")
	return a
}

func (a *Color) Set(name string, value interface{}) *Color {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Color) ClassName(value interface{}) *Color {
	a.Set("className", value)
	return a
}

func (a *Color) DefaultColor(value interface{}) *Color {
	a.Set("defaultColor", value)
	return a
}

func (a *Color) Disabled(value interface{}) *Color {
	a.Set("disabled", value)
	return a
}

func (a *Color) DisabledOn(value interface{}) *Color {
	a.Set("disabledOn", value)
	return a
}

func (a *Color) EditorSetting(value interface{}) *Color {
	a.Set("editorSetting", value)
	return a
}

func (a *Color) Hidden(value interface{}) *Color {
	a.Set("hidden", value)
	return a
}

func (a *Color) HiddenOn(value interface{}) *Color {
	a.Set("hiddenOn", value)
	return a
}

func (a *Color) Id(value interface{}) *Color {
	a.Set("id", value)
	return a
}

func (a *Color) OnEvent(value interface{}) *Color {
	a.Set("onEvent", value)
	return a
}

func (a *Color) ShowValue(value interface{}) *Color {
	a.Set("showValue", value)
	return a
}

func (a *Color) Static(value interface{}) *Color {
	a.Set("static", value)
	return a
}

func (a *Color) StaticClassName(value interface{}) *Color {
	a.Set("staticClassName", value)
	return a
}

func (a *Color) StaticInputClassName(value interface{}) *Color {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Color) StaticLabelClassName(value interface{}) *Color {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Color) StaticOn(value interface{}) *Color {
	a.Set("staticOn", value)
	return a
}

func (a *Color) StaticPlaceholder(value interface{}) *Color {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Color) StaticSchema(value interface{}) *Color {
	a.Set("staticSchema", value)
	return a
}

func (a *Color) Style(value interface{}) *Color {
	a.Set("style", value)
	return a
}

func (a *Color) TestIdBuilder(value interface{}) *Color {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Color) Testid(value interface{}) *Color {
	a.Set("testid", value)
	return a
}

func (a *Color) Type(value interface{}) *Color {
	a.Set("type", value)
	return a
}

func (a *Color) UseMobileUI(value interface{}) *Color {
	a.Set("useMobileUI", value)
	return a
}

func (a *Color) Visible(value interface{}) *Color {
	a.Set("visible", value)
	return a
}

func (a *Color) VisibleOn(value interface{}) *Color {
	a.Set("visibleOn", value)
	return a
}

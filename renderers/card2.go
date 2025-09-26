package renderers

type Card2 struct {
	*BaseRenderer
}

func NewCard2() *Card2 {
	a := &Card2{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "card2")
	return a
}

func (a *Card2) Set(name string, value interface{}) *Card2 {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Card2) Body(value interface{}) *Card2 {
	a.Set("body", value)
	return a
}

func (a *Card2) BodyClassName(value interface{}) *Card2 {
	a.Set("bodyClassName", value)
	return a
}

func (a *Card2) CheckOnItemClick(value interface{}) *Card2 {
	a.Set("checkOnItemClick", value)
	return a
}

func (a *Card2) ClassName(value interface{}) *Card2 {
	a.Set("className", value)
	return a
}

func (a *Card2) Disabled(value interface{}) *Card2 {
	a.Set("disabled", value)
	return a
}

func (a *Card2) DisabledOn(value interface{}) *Card2 {
	a.Set("disabledOn", value)
	return a
}

func (a *Card2) EditorSetting(value interface{}) *Card2 {
	a.Set("editorSetting", value)
	return a
}

func (a *Card2) Hidden(value interface{}) *Card2 {
	a.Set("hidden", value)
	return a
}

func (a *Card2) HiddenOn(value interface{}) *Card2 {
	a.Set("hiddenOn", value)
	return a
}

func (a *Card2) HideCheckToggler(value interface{}) *Card2 {
	a.Set("hideCheckToggler", value)
	return a
}

func (a *Card2) Id(value interface{}) *Card2 {
	a.Set("id", value)
	return a
}

func (a *Card2) OnEvent(value interface{}) *Card2 {
	a.Set("onEvent", value)
	return a
}

func (a *Card2) Static(value interface{}) *Card2 {
	a.Set("static", value)
	return a
}

func (a *Card2) StaticClassName(value interface{}) *Card2 {
	a.Set("staticClassName", value)
	return a
}

func (a *Card2) StaticInputClassName(value interface{}) *Card2 {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Card2) StaticLabelClassName(value interface{}) *Card2 {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Card2) StaticOn(value interface{}) *Card2 {
	a.Set("staticOn", value)
	return a
}

func (a *Card2) StaticPlaceholder(value interface{}) *Card2 {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Card2) StaticSchema(value interface{}) *Card2 {
	a.Set("staticSchema", value)
	return a
}

func (a *Card2) Style(value interface{}) *Card2 {
	a.Set("style", value)
	return a
}

func (a *Card2) TestIdBuilder(value interface{}) *Card2 {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Card2) Testid(value interface{}) *Card2 {
	a.Set("testid", value)
	return a
}

func (a *Card2) Type(value interface{}) *Card2 {
	a.Set("type", value)
	return a
}

func (a *Card2) UseMobileUI(value interface{}) *Card2 {
	a.Set("useMobileUI", value)
	return a
}

func (a *Card2) Visible(value interface{}) *Card2 {
	a.Set("visible", value)
	return a
}

func (a *Card2) VisibleOn(value interface{}) *Card2 {
	a.Set("visibleOn", value)
	return a
}

func (a *Card2) WrapperComponent(value interface{}) *Card2 {
	a.Set("wrapperComponent", value)
	return a
}

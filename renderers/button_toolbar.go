package renderers

type ButtonToolbar struct {
	*BaseRenderer
}

func NewButtonToolbar() *ButtonToolbar {
	a := &ButtonToolbar{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "button-toolbar")
	return a
}

func (a *ButtonToolbar) Set(name string, value interface{}) *ButtonToolbar {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *ButtonToolbar) Buttons(value interface{}) *ButtonToolbar {
	a.Set("buttons", value)
	return a
}

func (a *ButtonToolbar) ClassName(value interface{}) *ButtonToolbar {
	a.Set("className", value)
	return a
}

func (a *ButtonToolbar) Disabled(value interface{}) *ButtonToolbar {
	a.Set("disabled", value)
	return a
}

func (a *ButtonToolbar) DisabledOn(value interface{}) *ButtonToolbar {
	a.Set("disabledOn", value)
	return a
}

func (a *ButtonToolbar) EditorSetting(value interface{}) *ButtonToolbar {
	a.Set("editorSetting", value)
	return a
}

func (a *ButtonToolbar) Hidden(value interface{}) *ButtonToolbar {
	a.Set("hidden", value)
	return a
}

func (a *ButtonToolbar) HiddenOn(value interface{}) *ButtonToolbar {
	a.Set("hiddenOn", value)
	return a
}

func (a *ButtonToolbar) Id(value interface{}) *ButtonToolbar {
	a.Set("id", value)
	return a
}

func (a *ButtonToolbar) OnEvent(value interface{}) *ButtonToolbar {
	a.Set("onEvent", value)
	return a
}

func (a *ButtonToolbar) Static(value interface{}) *ButtonToolbar {
	a.Set("static", value)
	return a
}

func (a *ButtonToolbar) StaticClassName(value interface{}) *ButtonToolbar {
	a.Set("staticClassName", value)
	return a
}

func (a *ButtonToolbar) StaticInputClassName(value interface{}) *ButtonToolbar {
	a.Set("staticInputClassName", value)
	return a
}

func (a *ButtonToolbar) StaticLabelClassName(value interface{}) *ButtonToolbar {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *ButtonToolbar) StaticOn(value interface{}) *ButtonToolbar {
	a.Set("staticOn", value)
	return a
}

func (a *ButtonToolbar) StaticPlaceholder(value interface{}) *ButtonToolbar {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *ButtonToolbar) StaticSchema(value interface{}) *ButtonToolbar {
	a.Set("staticSchema", value)
	return a
}

func (a *ButtonToolbar) Style(value interface{}) *ButtonToolbar {
	a.Set("style", value)
	return a
}

func (a *ButtonToolbar) TestIdBuilder(value interface{}) *ButtonToolbar {
	a.Set("testIdBuilder", value)
	return a
}

func (a *ButtonToolbar) Testid(value interface{}) *ButtonToolbar {
	a.Set("testid", value)
	return a
}

func (a *ButtonToolbar) Type(value interface{}) *ButtonToolbar {
	a.Set("type", value)
	return a
}

func (a *ButtonToolbar) UseMobileUI(value interface{}) *ButtonToolbar {
	a.Set("useMobileUI", value)
	return a
}

func (a *ButtonToolbar) Visible(value interface{}) *ButtonToolbar {
	a.Set("visible", value)
	return a
}

func (a *ButtonToolbar) VisibleOn(value interface{}) *ButtonToolbar {
	a.Set("visibleOn", value)
	return a
}

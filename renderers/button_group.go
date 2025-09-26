package renderers

type ButtonGroup struct {
	*BaseRenderer
}

func NewButtonGroup() *ButtonGroup {
	a := &ButtonGroup{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "button-group")
	return a
}

func (a *ButtonGroup) Set(name string, value interface{}) *ButtonGroup {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *ButtonGroup) BtnActiveClassName(value interface{}) *ButtonGroup {
	a.Set("btnActiveClassName", value)
	return a
}

func (a *ButtonGroup) BtnActiveLevel(value interface{}) *ButtonGroup {
	a.Set("btnActiveLevel", value)
	return a
}

func (a *ButtonGroup) BtnClassName(value interface{}) *ButtonGroup {
	a.Set("btnClassName", value)
	return a
}

func (a *ButtonGroup) BtnLevel(value interface{}) *ButtonGroup {
	a.Set("btnLevel", value)
	return a
}

func (a *ButtonGroup) Buttons(value interface{}) *ButtonGroup {
	a.Set("buttons", value)
	return a
}

func (a *ButtonGroup) ClassName(value interface{}) *ButtonGroup {
	a.Set("className", value)
	return a
}

func (a *ButtonGroup) Disabled(value interface{}) *ButtonGroup {
	a.Set("disabled", value)
	return a
}

func (a *ButtonGroup) DisabledOn(value interface{}) *ButtonGroup {
	a.Set("disabledOn", value)
	return a
}

func (a *ButtonGroup) EditorSetting(value interface{}) *ButtonGroup {
	a.Set("editorSetting", value)
	return a
}

func (a *ButtonGroup) Hidden(value interface{}) *ButtonGroup {
	a.Set("hidden", value)
	return a
}

func (a *ButtonGroup) HiddenOn(value interface{}) *ButtonGroup {
	a.Set("hiddenOn", value)
	return a
}

func (a *ButtonGroup) Id(value interface{}) *ButtonGroup {
	a.Set("id", value)
	return a
}

func (a *ButtonGroup) OnEvent(value interface{}) *ButtonGroup {
	a.Set("onEvent", value)
	return a
}

func (a *ButtonGroup) Size(value interface{}) *ButtonGroup {
	a.Set("size", value)
	return a
}

func (a *ButtonGroup) Static(value interface{}) *ButtonGroup {
	a.Set("static", value)
	return a
}

func (a *ButtonGroup) StaticClassName(value interface{}) *ButtonGroup {
	a.Set("staticClassName", value)
	return a
}

func (a *ButtonGroup) StaticInputClassName(value interface{}) *ButtonGroup {
	a.Set("staticInputClassName", value)
	return a
}

func (a *ButtonGroup) StaticLabelClassName(value interface{}) *ButtonGroup {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *ButtonGroup) StaticOn(value interface{}) *ButtonGroup {
	a.Set("staticOn", value)
	return a
}

func (a *ButtonGroup) StaticPlaceholder(value interface{}) *ButtonGroup {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *ButtonGroup) StaticSchema(value interface{}) *ButtonGroup {
	a.Set("staticSchema", value)
	return a
}

func (a *ButtonGroup) Style(value interface{}) *ButtonGroup {
	a.Set("style", value)
	return a
}

func (a *ButtonGroup) TestIdBuilder(value interface{}) *ButtonGroup {
	a.Set("testIdBuilder", value)
	return a
}

func (a *ButtonGroup) Testid(value interface{}) *ButtonGroup {
	a.Set("testid", value)
	return a
}

func (a *ButtonGroup) Tiled(value interface{}) *ButtonGroup {
	a.Set("tiled", value)
	return a
}

func (a *ButtonGroup) Type(value interface{}) *ButtonGroup {
	a.Set("type", value)
	return a
}

func (a *ButtonGroup) UseMobileUI(value interface{}) *ButtonGroup {
	a.Set("useMobileUI", value)
	return a
}

func (a *ButtonGroup) Vertical(value interface{}) *ButtonGroup {
	a.Set("vertical", value)
	return a
}

func (a *ButtonGroup) Visible(value interface{}) *ButtonGroup {
	a.Set("visible", value)
	return a
}

func (a *ButtonGroup) VisibleOn(value interface{}) *ButtonGroup {
	a.Set("visibleOn", value)
	return a
}

package renderers

type SwitchContainer struct {
	*BaseRenderer
}

func NewSwitchContainer() *SwitchContainer {
	a := &SwitchContainer{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "switch-container")
	return a
}

func (a *SwitchContainer) Set(name string, value interface{}) *SwitchContainer {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *SwitchContainer) ClassName(value interface{}) *SwitchContainer {
	a.Set("className", value)
	return a
}

func (a *SwitchContainer) Disabled(value interface{}) *SwitchContainer {
	a.Set("disabled", value)
	return a
}

func (a *SwitchContainer) DisabledOn(value interface{}) *SwitchContainer {
	a.Set("disabledOn", value)
	return a
}

func (a *SwitchContainer) EditorSetting(value interface{}) *SwitchContainer {
	a.Set("editorSetting", value)
	return a
}

func (a *SwitchContainer) Hidden(value interface{}) *SwitchContainer {
	a.Set("hidden", value)
	return a
}

func (a *SwitchContainer) HiddenOn(value interface{}) *SwitchContainer {
	a.Set("hiddenOn", value)
	return a
}

func (a *SwitchContainer) Id(value interface{}) *SwitchContainer {
	a.Set("id", value)
	return a
}

func (a *SwitchContainer) Items(value interface{}) *SwitchContainer {
	a.Set("items", value)
	return a
}

func (a *SwitchContainer) OnEvent(value interface{}) *SwitchContainer {
	a.Set("onEvent", value)
	return a
}

func (a *SwitchContainer) Static(value interface{}) *SwitchContainer {
	a.Set("static", value)
	return a
}

func (a *SwitchContainer) StaticClassName(value interface{}) *SwitchContainer {
	a.Set("staticClassName", value)
	return a
}

func (a *SwitchContainer) StaticInputClassName(value interface{}) *SwitchContainer {
	a.Set("staticInputClassName", value)
	return a
}

func (a *SwitchContainer) StaticLabelClassName(value interface{}) *SwitchContainer {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *SwitchContainer) StaticOn(value interface{}) *SwitchContainer {
	a.Set("staticOn", value)
	return a
}

func (a *SwitchContainer) StaticPlaceholder(value interface{}) *SwitchContainer {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *SwitchContainer) StaticSchema(value interface{}) *SwitchContainer {
	a.Set("staticSchema", value)
	return a
}

func (a *SwitchContainer) Style(value interface{}) *SwitchContainer {
	a.Set("style", value)
	return a
}

func (a *SwitchContainer) TestIdBuilder(value interface{}) *SwitchContainer {
	a.Set("testIdBuilder", value)
	return a
}

func (a *SwitchContainer) Testid(value interface{}) *SwitchContainer {
	a.Set("testid", value)
	return a
}

func (a *SwitchContainer) Type(value interface{}) *SwitchContainer {
	a.Set("type", value)
	return a
}

func (a *SwitchContainer) UseMobileUI(value interface{}) *SwitchContainer {
	a.Set("useMobileUI", value)
	return a
}

func (a *SwitchContainer) Visible(value interface{}) *SwitchContainer {
	a.Set("visible", value)
	return a
}

func (a *SwitchContainer) VisibleOn(value interface{}) *SwitchContainer {
	a.Set("visibleOn", value)
	return a
}

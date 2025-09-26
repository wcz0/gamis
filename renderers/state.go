package renderers

type State struct {
	*BaseRenderer
}

func NewState() *State {
	a := &State{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *State) Set(name string, value interface{}) *State {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *State) Body(value interface{}) *State {
	a.Set("body", value)
	return a
}

func (a *State) ClassName(value interface{}) *State {
	a.Set("className", value)
	return a
}

func (a *State) Disabled(value interface{}) *State {
	a.Set("disabled", value)
	return a
}

func (a *State) DisabledOn(value interface{}) *State {
	a.Set("disabledOn", value)
	return a
}

func (a *State) EditorSetting(value interface{}) *State {
	a.Set("editorSetting", value)
	return a
}

func (a *State) Hidden(value interface{}) *State {
	a.Set("hidden", value)
	return a
}

func (a *State) HiddenOn(value interface{}) *State {
	a.Set("hiddenOn", value)
	return a
}

func (a *State) Id(value interface{}) *State {
	a.Set("id", value)
	return a
}

func (a *State) OnEvent(value interface{}) *State {
	a.Set("onEvent", value)
	return a
}

func (a *State) Static(value interface{}) *State {
	a.Set("static", value)
	return a
}

func (a *State) StaticClassName(value interface{}) *State {
	a.Set("staticClassName", value)
	return a
}

func (a *State) StaticInputClassName(value interface{}) *State {
	a.Set("staticInputClassName", value)
	return a
}

func (a *State) StaticLabelClassName(value interface{}) *State {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *State) StaticOn(value interface{}) *State {
	a.Set("staticOn", value)
	return a
}

func (a *State) StaticPlaceholder(value interface{}) *State {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *State) StaticSchema(value interface{}) *State {
	a.Set("staticSchema", value)
	return a
}

func (a *State) Style(value interface{}) *State {
	a.Set("style", value)
	return a
}

func (a *State) TestIdBuilder(value interface{}) *State {
	a.Set("testIdBuilder", value)
	return a
}

func (a *State) Testid(value interface{}) *State {
	a.Set("testid", value)
	return a
}

func (a *State) Title(value interface{}) *State {
	a.Set("title", value)
	return a
}

func (a *State) UseMobileUI(value interface{}) *State {
	a.Set("useMobileUI", value)
	return a
}

func (a *State) Visible(value interface{}) *State {
	a.Set("visible", value)
	return a
}

func (a *State) VisibleOn(value interface{}) *State {
	a.Set("visibleOn", value)
	return a
}

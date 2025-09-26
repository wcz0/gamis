package renderers

type Step struct {
	*BaseRenderer
}

func NewStep() *Step {
	a := &Step{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *Step) Set(name string, value interface{}) *Step {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Step) ClassName(value interface{}) *Step {
	a.Set("className", value)
	return a
}

func (a *Step) Description(value interface{}) *Step {
	a.Set("description", value)
	return a
}

func (a *Step) Disabled(value interface{}) *Step {
	a.Set("disabled", value)
	return a
}

func (a *Step) DisabledOn(value interface{}) *Step {
	a.Set("disabledOn", value)
	return a
}

func (a *Step) EditorSetting(value interface{}) *Step {
	a.Set("editorSetting", value)
	return a
}

func (a *Step) Hidden(value interface{}) *Step {
	a.Set("hidden", value)
	return a
}

func (a *Step) HiddenOn(value interface{}) *Step {
	a.Set("hiddenOn", value)
	return a
}

func (a *Step) Icon(value interface{}) *Step {
	a.Set("icon", value)
	return a
}

func (a *Step) Id(value interface{}) *Step {
	a.Set("id", value)
	return a
}

func (a *Step) OnEvent(value interface{}) *Step {
	a.Set("onEvent", value)
	return a
}

func (a *Step) Static(value interface{}) *Step {
	a.Set("static", value)
	return a
}

func (a *Step) StaticClassName(value interface{}) *Step {
	a.Set("staticClassName", value)
	return a
}

func (a *Step) StaticInputClassName(value interface{}) *Step {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Step) StaticLabelClassName(value interface{}) *Step {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Step) StaticOn(value interface{}) *Step {
	a.Set("staticOn", value)
	return a
}

func (a *Step) StaticPlaceholder(value interface{}) *Step {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Step) StaticSchema(value interface{}) *Step {
	a.Set("staticSchema", value)
	return a
}

func (a *Step) Style(value interface{}) *Step {
	a.Set("style", value)
	return a
}

func (a *Step) SubTitle(value interface{}) *Step {
	a.Set("subTitle", value)
	return a
}

func (a *Step) TestIdBuilder(value interface{}) *Step {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Step) Testid(value interface{}) *Step {
	a.Set("testid", value)
	return a
}

func (a *Step) Title(value interface{}) *Step {
	a.Set("title", value)
	return a
}

func (a *Step) UseMobileUI(value interface{}) *Step {
	a.Set("useMobileUI", value)
	return a
}

func (a *Step) Value(value interface{}) *Step {
	a.Set("value", value)
	return a
}

func (a *Step) Visible(value interface{}) *Step {
	a.Set("visible", value)
	return a
}

func (a *Step) VisibleOn(value interface{}) *Step {
	a.Set("visibleOn", value)
	return a
}

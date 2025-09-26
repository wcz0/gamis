package renderers

type Operation struct {
	*BaseRenderer
}

func NewOperation() *Operation {
	a := &Operation{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "operation")
	return a
}

func (a *Operation) Set(name string, value interface{}) *Operation {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Operation) Buttons(value interface{}) *Operation {
	a.Set("buttons", value)
	return a
}

func (a *Operation) ClassName(value interface{}) *Operation {
	a.Set("className", value)
	return a
}

func (a *Operation) Disabled(value interface{}) *Operation {
	a.Set("disabled", value)
	return a
}

func (a *Operation) DisabledOn(value interface{}) *Operation {
	a.Set("disabledOn", value)
	return a
}

func (a *Operation) EditorSetting(value interface{}) *Operation {
	a.Set("editorSetting", value)
	return a
}

func (a *Operation) Fixed(value interface{}) *Operation {
	a.Set("fixed", value)
	return a
}

func (a *Operation) Hidden(value interface{}) *Operation {
	a.Set("hidden", value)
	return a
}

func (a *Operation) HiddenOn(value interface{}) *Operation {
	a.Set("hiddenOn", value)
	return a
}

func (a *Operation) Id(value interface{}) *Operation {
	a.Set("id", value)
	return a
}

func (a *Operation) Label(value interface{}) *Operation {
	a.Set("label", value)
	return a
}

func (a *Operation) OnEvent(value interface{}) *Operation {
	a.Set("onEvent", value)
	return a
}

func (a *Operation) Placeholder(value interface{}) *Operation {
	a.Set("placeholder", value)
	return a
}

func (a *Operation) Static(value interface{}) *Operation {
	a.Set("static", value)
	return a
}

func (a *Operation) StaticClassName(value interface{}) *Operation {
	a.Set("staticClassName", value)
	return a
}

func (a *Operation) StaticInputClassName(value interface{}) *Operation {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Operation) StaticLabelClassName(value interface{}) *Operation {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Operation) StaticOn(value interface{}) *Operation {
	a.Set("staticOn", value)
	return a
}

func (a *Operation) StaticPlaceholder(value interface{}) *Operation {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Operation) StaticSchema(value interface{}) *Operation {
	a.Set("staticSchema", value)
	return a
}

func (a *Operation) Style(value interface{}) *Operation {
	a.Set("style", value)
	return a
}

func (a *Operation) TestIdBuilder(value interface{}) *Operation {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Operation) Testid(value interface{}) *Operation {
	a.Set("testid", value)
	return a
}

func (a *Operation) Type(value interface{}) *Operation {
	a.Set("type", value)
	return a
}

func (a *Operation) UseMobileUI(value interface{}) *Operation {
	a.Set("useMobileUI", value)
	return a
}

func (a *Operation) Visible(value interface{}) *Operation {
	a.Set("visible", value)
	return a
}

func (a *Operation) VisibleOn(value interface{}) *Operation {
	a.Set("visibleOn", value)
	return a
}

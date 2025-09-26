package renderers

type Password struct {
	*BaseRenderer
}

func NewPassword() *Password {
	a := &Password{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "password")
	return a
}

func (a *Password) Set(name string, value interface{}) *Password {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Password) ClassName(value interface{}) *Password {
	a.Set("className", value)
	return a
}

func (a *Password) Disabled(value interface{}) *Password {
	a.Set("disabled", value)
	return a
}

func (a *Password) DisabledOn(value interface{}) *Password {
	a.Set("disabledOn", value)
	return a
}

func (a *Password) EditorSetting(value interface{}) *Password {
	a.Set("editorSetting", value)
	return a
}

func (a *Password) Hidden(value interface{}) *Password {
	a.Set("hidden", value)
	return a
}

func (a *Password) HiddenOn(value interface{}) *Password {
	a.Set("hiddenOn", value)
	return a
}

func (a *Password) Id(value interface{}) *Password {
	a.Set("id", value)
	return a
}

func (a *Password) MosaicText(value interface{}) *Password {
	a.Set("mosaicText", value)
	return a
}

func (a *Password) OnEvent(value interface{}) *Password {
	a.Set("onEvent", value)
	return a
}

func (a *Password) Static(value interface{}) *Password {
	a.Set("static", value)
	return a
}

func (a *Password) StaticClassName(value interface{}) *Password {
	a.Set("staticClassName", value)
	return a
}

func (a *Password) StaticInputClassName(value interface{}) *Password {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Password) StaticLabelClassName(value interface{}) *Password {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Password) StaticOn(value interface{}) *Password {
	a.Set("staticOn", value)
	return a
}

func (a *Password) StaticPlaceholder(value interface{}) *Password {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Password) StaticSchema(value interface{}) *Password {
	a.Set("staticSchema", value)
	return a
}

func (a *Password) Style(value interface{}) *Password {
	a.Set("style", value)
	return a
}

func (a *Password) TestIdBuilder(value interface{}) *Password {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Password) Testid(value interface{}) *Password {
	a.Set("testid", value)
	return a
}

func (a *Password) Type(value interface{}) *Password {
	a.Set("type", value)
	return a
}

func (a *Password) UseMobileUI(value interface{}) *Password {
	a.Set("useMobileUI", value)
	return a
}

func (a *Password) Visible(value interface{}) *Password {
	a.Set("visible", value)
	return a
}

func (a *Password) VisibleOn(value interface{}) *Password {
	a.Set("visibleOn", value)
	return a
}

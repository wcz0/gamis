package renderers

type Plain struct {
	*BaseRenderer
}

func NewPlain() *Plain {
	a := &Plain{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "plain")
	return a
}

func (a *Plain) Set(name string, value interface{}) *Plain {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Plain) ClassName(value interface{}) *Plain {
	a.Set("className", value)
	return a
}

func (a *Plain) Disabled(value interface{}) *Plain {
	a.Set("disabled", value)
	return a
}

func (a *Plain) DisabledOn(value interface{}) *Plain {
	a.Set("disabledOn", value)
	return a
}

func (a *Plain) EditorSetting(value interface{}) *Plain {
	a.Set("editorSetting", value)
	return a
}

func (a *Plain) Hidden(value interface{}) *Plain {
	a.Set("hidden", value)
	return a
}

func (a *Plain) HiddenOn(value interface{}) *Plain {
	a.Set("hiddenOn", value)
	return a
}

func (a *Plain) Id(value interface{}) *Plain {
	a.Set("id", value)
	return a
}

func (a *Plain) Inline(value interface{}) *Plain {
	a.Set("inline", value)
	return a
}

func (a *Plain) OnEvent(value interface{}) *Plain {
	a.Set("onEvent", value)
	return a
}

func (a *Plain) Placeholder(value interface{}) *Plain {
	a.Set("placeholder", value)
	return a
}

func (a *Plain) Static(value interface{}) *Plain {
	a.Set("static", value)
	return a
}

func (a *Plain) StaticClassName(value interface{}) *Plain {
	a.Set("staticClassName", value)
	return a
}

func (a *Plain) StaticInputClassName(value interface{}) *Plain {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Plain) StaticLabelClassName(value interface{}) *Plain {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Plain) StaticOn(value interface{}) *Plain {
	a.Set("staticOn", value)
	return a
}

func (a *Plain) StaticPlaceholder(value interface{}) *Plain {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Plain) StaticSchema(value interface{}) *Plain {
	a.Set("staticSchema", value)
	return a
}

func (a *Plain) Style(value interface{}) *Plain {
	a.Set("style", value)
	return a
}

func (a *Plain) TestIdBuilder(value interface{}) *Plain {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Plain) Testid(value interface{}) *Plain {
	a.Set("testid", value)
	return a
}

func (a *Plain) Text(value interface{}) *Plain {
	a.Set("text", value)
	return a
}

func (a *Plain) Tpl(value interface{}) *Plain {
	a.Set("tpl", value)
	return a
}

func (a *Plain) Type(value interface{}) *Plain {
	a.Set("type", value)
	return a
}

func (a *Plain) UseMobileUI(value interface{}) *Plain {
	a.Set("useMobileUI", value)
	return a
}

func (a *Plain) Visible(value interface{}) *Plain {
	a.Set("visible", value)
	return a
}

func (a *Plain) VisibleOn(value interface{}) *Plain {
	a.Set("visibleOn", value)
	return a
}

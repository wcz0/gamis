package renderers

type Audio struct {
	*BaseRenderer
}

func NewAudio() *Audio {
	a := &Audio{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "audio")
	return a
}

func (a *Audio) Set(name string, value interface{}) *Audio {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Audio) AutoPlay(value interface{}) *Audio {
	a.Set("autoPlay", value)
	return a
}

func (a *Audio) ClassName(value interface{}) *Audio {
	a.Set("className", value)
	return a
}

func (a *Audio) Controls(value interface{}) *Audio {
	a.Set("controls", value)
	return a
}

func (a *Audio) Disabled(value interface{}) *Audio {
	a.Set("disabled", value)
	return a
}

func (a *Audio) DisabledOn(value interface{}) *Audio {
	a.Set("disabledOn", value)
	return a
}

func (a *Audio) EditorSetting(value interface{}) *Audio {
	a.Set("editorSetting", value)
	return a
}

func (a *Audio) Hidden(value interface{}) *Audio {
	a.Set("hidden", value)
	return a
}

func (a *Audio) HiddenOn(value interface{}) *Audio {
	a.Set("hiddenOn", value)
	return a
}

func (a *Audio) Id(value interface{}) *Audio {
	a.Set("id", value)
	return a
}

func (a *Audio) Inline(value interface{}) *Audio {
	a.Set("inline", value)
	return a
}

func (a *Audio) Loop(value interface{}) *Audio {
	a.Set("loop", value)
	return a
}

func (a *Audio) OnEvent(value interface{}) *Audio {
	a.Set("onEvent", value)
	return a
}

func (a *Audio) Rates(value interface{}) *Audio {
	a.Set("rates", value)
	return a
}

func (a *Audio) Src(value interface{}) *Audio {
	a.Set("src", value)
	return a
}

func (a *Audio) Static(value interface{}) *Audio {
	a.Set("static", value)
	return a
}

func (a *Audio) StaticClassName(value interface{}) *Audio {
	a.Set("staticClassName", value)
	return a
}

func (a *Audio) StaticInputClassName(value interface{}) *Audio {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Audio) StaticLabelClassName(value interface{}) *Audio {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Audio) StaticOn(value interface{}) *Audio {
	a.Set("staticOn", value)
	return a
}

func (a *Audio) StaticPlaceholder(value interface{}) *Audio {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Audio) StaticSchema(value interface{}) *Audio {
	a.Set("staticSchema", value)
	return a
}

func (a *Audio) Style(value interface{}) *Audio {
	a.Set("style", value)
	return a
}

func (a *Audio) TestIdBuilder(value interface{}) *Audio {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Audio) Testid(value interface{}) *Audio {
	a.Set("testid", value)
	return a
}

func (a *Audio) Type(value interface{}) *Audio {
	a.Set("type", value)
	return a
}

func (a *Audio) UseMobileUI(value interface{}) *Audio {
	a.Set("useMobileUI", value)
	return a
}

func (a *Audio) Visible(value interface{}) *Audio {
	a.Set("visible", value)
	return a
}

func (a *Audio) VisibleOn(value interface{}) *Audio {
	a.Set("visibleOn", value)
	return a
}

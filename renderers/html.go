package renderers

type Html struct {
	*BaseRenderer
}

func NewHtml() *Html {
	a := &Html{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "tpl")
	return a
}

func (a *Html) Set(name string, value interface{}) *Html {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Html) Badge(value interface{}) *Html {
	a.Set("badge", value)
	return a
}

func (a *Html) ClassName(value interface{}) *Html {
	a.Set("className", value)
	return a
}

func (a *Html) Disabled(value interface{}) *Html {
	a.Set("disabled", value)
	return a
}

func (a *Html) DisabledOn(value interface{}) *Html {
	a.Set("disabledOn", value)
	return a
}

func (a *Html) EditorSetting(value interface{}) *Html {
	a.Set("editorSetting", value)
	return a
}

func (a *Html) Hidden(value interface{}) *Html {
	a.Set("hidden", value)
	return a
}

func (a *Html) HiddenOn(value interface{}) *Html {
	a.Set("hiddenOn", value)
	return a
}

func (a *Html) Html(value interface{}) *Html {
	a.Set("html", value)
	return a
}

func (a *Html) Id(value interface{}) *Html {
	a.Set("id", value)
	return a
}

func (a *Html) Inline(value interface{}) *Html {
	a.Set("inline", value)
	return a
}

func (a *Html) OnEvent(value interface{}) *Html {
	a.Set("onEvent", value)
	return a
}

func (a *Html) Raw(value interface{}) *Html {
	a.Set("raw", value)
	return a
}

func (a *Html) Static(value interface{}) *Html {
	a.Set("static", value)
	return a
}

func (a *Html) StaticClassName(value interface{}) *Html {
	a.Set("staticClassName", value)
	return a
}

func (a *Html) StaticInputClassName(value interface{}) *Html {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Html) StaticLabelClassName(value interface{}) *Html {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Html) StaticOn(value interface{}) *Html {
	a.Set("staticOn", value)
	return a
}

func (a *Html) StaticPlaceholder(value interface{}) *Html {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Html) StaticSchema(value interface{}) *Html {
	a.Set("staticSchema", value)
	return a
}

func (a *Html) Style(value interface{}) *Html {
	a.Set("style", value)
	return a
}

func (a *Html) TestIdBuilder(value interface{}) *Html {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Html) Testid(value interface{}) *Html {
	a.Set("testid", value)
	return a
}

func (a *Html) Text(value interface{}) *Html {
	a.Set("text", value)
	return a
}

func (a *Html) Tpl(value interface{}) *Html {
	a.Set("tpl", value)
	return a
}

func (a *Html) Type(value interface{}) *Html {
	a.Set("type", value)
	return a
}

func (a *Html) UseMobileUI(value interface{}) *Html {
	a.Set("useMobileUI", value)
	return a
}

func (a *Html) Visible(value interface{}) *Html {
	a.Set("visible", value)
	return a
}

func (a *Html) VisibleOn(value interface{}) *Html {
	a.Set("visibleOn", value)
	return a
}

func (a *Html) WrapperComponent(value interface{}) *Html {
	a.Set("wrapperComponent", value)
	return a
}

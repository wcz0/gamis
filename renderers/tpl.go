package renderers

type Tpl struct {
	*BaseRenderer
}

func NewTpl() *Tpl {
	a := &Tpl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "tpl")
	return a
}

func (a *Tpl) Set(name string, value interface{}) *Tpl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Tpl) Badge(value interface{}) *Tpl {
	a.Set("badge", value)
	return a
}

func (a *Tpl) ClassName(value interface{}) *Tpl {
	a.Set("className", value)
	return a
}

func (a *Tpl) Disabled(value interface{}) *Tpl {
	a.Set("disabled", value)
	return a
}

func (a *Tpl) DisabledOn(value interface{}) *Tpl {
	a.Set("disabledOn", value)
	return a
}

func (a *Tpl) EditorSetting(value interface{}) *Tpl {
	a.Set("editorSetting", value)
	return a
}

func (a *Tpl) Hidden(value interface{}) *Tpl {
	a.Set("hidden", value)
	return a
}

func (a *Tpl) HiddenOn(value interface{}) *Tpl {
	a.Set("hiddenOn", value)
	return a
}

func (a *Tpl) Html(value interface{}) *Tpl {
	a.Set("html", value)
	return a
}

func (a *Tpl) Id(value interface{}) *Tpl {
	a.Set("id", value)
	return a
}

func (a *Tpl) Inline(value interface{}) *Tpl {
	a.Set("inline", value)
	return a
}

func (a *Tpl) OnEvent(value interface{}) *Tpl {
	a.Set("onEvent", value)
	return a
}

func (a *Tpl) Raw(value interface{}) *Tpl {
	a.Set("raw", value)
	return a
}

func (a *Tpl) Static(value interface{}) *Tpl {
	a.Set("static", value)
	return a
}

func (a *Tpl) StaticClassName(value interface{}) *Tpl {
	a.Set("staticClassName", value)
	return a
}

func (a *Tpl) StaticInputClassName(value interface{}) *Tpl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Tpl) StaticLabelClassName(value interface{}) *Tpl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Tpl) StaticOn(value interface{}) *Tpl {
	a.Set("staticOn", value)
	return a
}

func (a *Tpl) StaticPlaceholder(value interface{}) *Tpl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Tpl) StaticSchema(value interface{}) *Tpl {
	a.Set("staticSchema", value)
	return a
}

func (a *Tpl) Style(value interface{}) *Tpl {
	a.Set("style", value)
	return a
}

func (a *Tpl) TestIdBuilder(value interface{}) *Tpl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Tpl) Testid(value interface{}) *Tpl {
	a.Set("testid", value)
	return a
}

func (a *Tpl) Text(value interface{}) *Tpl {
	a.Set("text", value)
	return a
}

func (a *Tpl) Tpl(value interface{}) *Tpl {
	a.Set("tpl", value)
	return a
}

func (a *Tpl) Type(value interface{}) *Tpl {
	a.Set("type", value)
	return a
}

func (a *Tpl) UseMobileUI(value interface{}) *Tpl {
	a.Set("useMobileUI", value)
	return a
}

func (a *Tpl) Visible(value interface{}) *Tpl {
	a.Set("visible", value)
	return a
}

func (a *Tpl) VisibleOn(value interface{}) *Tpl {
	a.Set("visibleOn", value)
	return a
}

func (a *Tpl) WrapperComponent(value interface{}) *Tpl {
	a.Set("wrapperComponent", value)
	return a
}

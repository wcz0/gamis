package renderers

type Divider struct {
	*BaseRenderer
}

func NewDivider() *Divider {
	a := &Divider{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "divider")
	return a
}

func (a *Divider) Set(name string, value interface{}) *Divider {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Divider) ClassName(value interface{}) *Divider {
	a.Set("className", value)
	return a
}

func (a *Divider) Color(value interface{}) *Divider {
	a.Set("color", value)
	return a
}

func (a *Divider) Direction(value interface{}) *Divider {
	a.Set("direction", value)
	return a
}

func (a *Divider) Disabled(value interface{}) *Divider {
	a.Set("disabled", value)
	return a
}

func (a *Divider) DisabledOn(value interface{}) *Divider {
	a.Set("disabledOn", value)
	return a
}

func (a *Divider) EditorSetting(value interface{}) *Divider {
	a.Set("editorSetting", value)
	return a
}

func (a *Divider) Hidden(value interface{}) *Divider {
	a.Set("hidden", value)
	return a
}

func (a *Divider) HiddenOn(value interface{}) *Divider {
	a.Set("hiddenOn", value)
	return a
}

func (a *Divider) Id(value interface{}) *Divider {
	a.Set("id", value)
	return a
}

func (a *Divider) LineStyle(value interface{}) *Divider {
	a.Set("lineStyle", value)
	return a
}

func (a *Divider) OnEvent(value interface{}) *Divider {
	a.Set("onEvent", value)
	return a
}

func (a *Divider) Rotate(value interface{}) *Divider {
	a.Set("rotate", value)
	return a
}

func (a *Divider) Static(value interface{}) *Divider {
	a.Set("static", value)
	return a
}

func (a *Divider) StaticClassName(value interface{}) *Divider {
	a.Set("staticClassName", value)
	return a
}

func (a *Divider) StaticInputClassName(value interface{}) *Divider {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Divider) StaticLabelClassName(value interface{}) *Divider {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Divider) StaticOn(value interface{}) *Divider {
	a.Set("staticOn", value)
	return a
}

func (a *Divider) StaticPlaceholder(value interface{}) *Divider {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Divider) StaticSchema(value interface{}) *Divider {
	a.Set("staticSchema", value)
	return a
}

func (a *Divider) Style(value interface{}) *Divider {
	a.Set("style", value)
	return a
}

func (a *Divider) TestIdBuilder(value interface{}) *Divider {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Divider) Testid(value interface{}) *Divider {
	a.Set("testid", value)
	return a
}

func (a *Divider) Title(value interface{}) *Divider {
	a.Set("title", value)
	return a
}

func (a *Divider) TitleClassName(value interface{}) *Divider {
	a.Set("titleClassName", value)
	return a
}

func (a *Divider) TitlePosition(value interface{}) *Divider {
	a.Set("titlePosition", value)
	return a
}

func (a *Divider) Type(value interface{}) *Divider {
	a.Set("type", value)
	return a
}

func (a *Divider) UseMobileUI(value interface{}) *Divider {
	a.Set("useMobileUI", value)
	return a
}

func (a *Divider) Visible(value interface{}) *Divider {
	a.Set("visible", value)
	return a
}

func (a *Divider) VisibleOn(value interface{}) *Divider {
	a.Set("visibleOn", value)
	return a
}

package renderers

type Grid struct {
	*BaseRenderer
}

func NewGrid() *Grid {
	a := &Grid{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "grid")
	return a
}

func (a *Grid) Set(name string, value interface{}) *Grid {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Grid) Align(value interface{}) *Grid {
	a.Set("align", value)
	return a
}

func (a *Grid) ClassName(value interface{}) *Grid {
	a.Set("className", value)
	return a
}

func (a *Grid) Columns(value interface{}) *Grid {
	a.Set("columns", value)
	return a
}

func (a *Grid) Disabled(value interface{}) *Grid {
	a.Set("disabled", value)
	return a
}

func (a *Grid) DisabledOn(value interface{}) *Grid {
	a.Set("disabledOn", value)
	return a
}

func (a *Grid) EditorSetting(value interface{}) *Grid {
	a.Set("editorSetting", value)
	return a
}

func (a *Grid) Gap(value interface{}) *Grid {
	a.Set("gap", value)
	return a
}

func (a *Grid) Hidden(value interface{}) *Grid {
	a.Set("hidden", value)
	return a
}

func (a *Grid) HiddenOn(value interface{}) *Grid {
	a.Set("hiddenOn", value)
	return a
}

func (a *Grid) Id(value interface{}) *Grid {
	a.Set("id", value)
	return a
}

func (a *Grid) OnEvent(value interface{}) *Grid {
	a.Set("onEvent", value)
	return a
}

func (a *Grid) Static(value interface{}) *Grid {
	a.Set("static", value)
	return a
}

func (a *Grid) StaticClassName(value interface{}) *Grid {
	a.Set("staticClassName", value)
	return a
}

func (a *Grid) StaticInputClassName(value interface{}) *Grid {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Grid) StaticLabelClassName(value interface{}) *Grid {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Grid) StaticOn(value interface{}) *Grid {
	a.Set("staticOn", value)
	return a
}

func (a *Grid) StaticPlaceholder(value interface{}) *Grid {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Grid) StaticSchema(value interface{}) *Grid {
	a.Set("staticSchema", value)
	return a
}

func (a *Grid) Style(value interface{}) *Grid {
	a.Set("style", value)
	return a
}

func (a *Grid) TestIdBuilder(value interface{}) *Grid {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Grid) Testid(value interface{}) *Grid {
	a.Set("testid", value)
	return a
}

func (a *Grid) Type(value interface{}) *Grid {
	a.Set("type", value)
	return a
}

func (a *Grid) UseMobileUI(value interface{}) *Grid {
	a.Set("useMobileUI", value)
	return a
}

func (a *Grid) Valign(value interface{}) *Grid {
	a.Set("valign", value)
	return a
}

func (a *Grid) Visible(value interface{}) *Grid {
	a.Set("visible", value)
	return a
}

func (a *Grid) VisibleOn(value interface{}) *Grid {
	a.Set("visibleOn", value)
	return a
}

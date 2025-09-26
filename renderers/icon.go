package renderers

type Icon struct {
	*BaseRenderer
}

func NewIcon() *Icon {
	a := &Icon{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "icon")
	return a
}

func (a *Icon) Set(name string, value interface{}) *Icon {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Icon) Badge(value interface{}) *Icon {
	a.Set("badge", value)
	return a
}

func (a *Icon) ClassName(value interface{}) *Icon {
	a.Set("className", value)
	return a
}

func (a *Icon) Disabled(value interface{}) *Icon {
	a.Set("disabled", value)
	return a
}

func (a *Icon) DisabledOn(value interface{}) *Icon {
	a.Set("disabledOn", value)
	return a
}

func (a *Icon) EditorSetting(value interface{}) *Icon {
	a.Set("editorSetting", value)
	return a
}

func (a *Icon) Hidden(value interface{}) *Icon {
	a.Set("hidden", value)
	return a
}

func (a *Icon) HiddenOn(value interface{}) *Icon {
	a.Set("hiddenOn", value)
	return a
}

func (a *Icon) Icon(value interface{}) *Icon {
	a.Set("icon", value)
	return a
}

func (a *Icon) Id(value interface{}) *Icon {
	a.Set("id", value)
	return a
}

func (a *Icon) OnEvent(value interface{}) *Icon {
	a.Set("onEvent", value)
	return a
}

func (a *Icon) Static(value interface{}) *Icon {
	a.Set("static", value)
	return a
}

func (a *Icon) StaticClassName(value interface{}) *Icon {
	a.Set("staticClassName", value)
	return a
}

func (a *Icon) StaticInputClassName(value interface{}) *Icon {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Icon) StaticLabelClassName(value interface{}) *Icon {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Icon) StaticOn(value interface{}) *Icon {
	a.Set("staticOn", value)
	return a
}

func (a *Icon) StaticPlaceholder(value interface{}) *Icon {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Icon) StaticSchema(value interface{}) *Icon {
	a.Set("staticSchema", value)
	return a
}

func (a *Icon) Style(value interface{}) *Icon {
	a.Set("style", value)
	return a
}

func (a *Icon) TestIdBuilder(value interface{}) *Icon {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Icon) Testid(value interface{}) *Icon {
	a.Set("testid", value)
	return a
}

func (a *Icon) Type(value interface{}) *Icon {
	a.Set("type", value)
	return a
}

func (a *Icon) UseMobileUI(value interface{}) *Icon {
	a.Set("useMobileUI", value)
	return a
}

func (a *Icon) Vendor(value interface{}) *Icon {
	a.Set("vendor", value)
	return a
}

func (a *Icon) Visible(value interface{}) *Icon {
	a.Set("visible", value)
	return a
}

func (a *Icon) VisibleOn(value interface{}) *Icon {
	a.Set("visibleOn", value)
	return a
}

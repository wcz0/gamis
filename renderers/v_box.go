package renderers

type VBox struct {
	*BaseRenderer
}

func NewVBox() *VBox {
	a := &VBox{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "vbox")
	return a
}

func (a *VBox) Set(name string, value interface{}) *VBox {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *VBox) ClassName(value interface{}) *VBox {
	a.Set("className", value)
	return a
}

func (a *VBox) Disabled(value interface{}) *VBox {
	a.Set("disabled", value)
	return a
}

func (a *VBox) DisabledOn(value interface{}) *VBox {
	a.Set("disabledOn", value)
	return a
}

func (a *VBox) EditorSetting(value interface{}) *VBox {
	a.Set("editorSetting", value)
	return a
}

func (a *VBox) Hidden(value interface{}) *VBox {
	a.Set("hidden", value)
	return a
}

func (a *VBox) HiddenOn(value interface{}) *VBox {
	a.Set("hiddenOn", value)
	return a
}

func (a *VBox) Id(value interface{}) *VBox {
	a.Set("id", value)
	return a
}

func (a *VBox) OnEvent(value interface{}) *VBox {
	a.Set("onEvent", value)
	return a
}

func (a *VBox) Rows(value interface{}) *VBox {
	a.Set("rows", value)
	return a
}

func (a *VBox) Static(value interface{}) *VBox {
	a.Set("static", value)
	return a
}

func (a *VBox) StaticClassName(value interface{}) *VBox {
	a.Set("staticClassName", value)
	return a
}

func (a *VBox) StaticInputClassName(value interface{}) *VBox {
	a.Set("staticInputClassName", value)
	return a
}

func (a *VBox) StaticLabelClassName(value interface{}) *VBox {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *VBox) StaticOn(value interface{}) *VBox {
	a.Set("staticOn", value)
	return a
}

func (a *VBox) StaticPlaceholder(value interface{}) *VBox {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *VBox) StaticSchema(value interface{}) *VBox {
	a.Set("staticSchema", value)
	return a
}

func (a *VBox) Style(value interface{}) *VBox {
	a.Set("style", value)
	return a
}

func (a *VBox) TestIdBuilder(value interface{}) *VBox {
	a.Set("testIdBuilder", value)
	return a
}

func (a *VBox) Testid(value interface{}) *VBox {
	a.Set("testid", value)
	return a
}

func (a *VBox) Type(value interface{}) *VBox {
	a.Set("type", value)
	return a
}

func (a *VBox) UseMobileUI(value interface{}) *VBox {
	a.Set("useMobileUI", value)
	return a
}

func (a *VBox) Visible(value interface{}) *VBox {
	a.Set("visible", value)
	return a
}

func (a *VBox) VisibleOn(value interface{}) *VBox {
	a.Set("visibleOn", value)
	return a
}

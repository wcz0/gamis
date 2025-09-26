package renderers

type HBox struct {
	*BaseRenderer
}

func NewHBox() *HBox {
	a := &HBox{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "hbox")
	return a
}

func (a *HBox) Set(name string, value interface{}) *HBox {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *HBox) Align(value interface{}) *HBox {
	a.Set("align", value)
	return a
}

func (a *HBox) ClassName(value interface{}) *HBox {
	a.Set("className", value)
	return a
}

func (a *HBox) Columns(value interface{}) *HBox {
	a.Set("columns", value)
	return a
}

func (a *HBox) Disabled(value interface{}) *HBox {
	a.Set("disabled", value)
	return a
}

func (a *HBox) DisabledOn(value interface{}) *HBox {
	a.Set("disabledOn", value)
	return a
}

func (a *HBox) EditorSetting(value interface{}) *HBox {
	a.Set("editorSetting", value)
	return a
}

func (a *HBox) Gap(value interface{}) *HBox {
	a.Set("gap", value)
	return a
}

func (a *HBox) Hidden(value interface{}) *HBox {
	a.Set("hidden", value)
	return a
}

func (a *HBox) HiddenOn(value interface{}) *HBox {
	a.Set("hiddenOn", value)
	return a
}

func (a *HBox) Id(value interface{}) *HBox {
	a.Set("id", value)
	return a
}

func (a *HBox) OnEvent(value interface{}) *HBox {
	a.Set("onEvent", value)
	return a
}

func (a *HBox) Static(value interface{}) *HBox {
	a.Set("static", value)
	return a
}

func (a *HBox) StaticClassName(value interface{}) *HBox {
	a.Set("staticClassName", value)
	return a
}

func (a *HBox) StaticInputClassName(value interface{}) *HBox {
	a.Set("staticInputClassName", value)
	return a
}

func (a *HBox) StaticLabelClassName(value interface{}) *HBox {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *HBox) StaticOn(value interface{}) *HBox {
	a.Set("staticOn", value)
	return a
}

func (a *HBox) StaticPlaceholder(value interface{}) *HBox {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *HBox) StaticSchema(value interface{}) *HBox {
	a.Set("staticSchema", value)
	return a
}

func (a *HBox) Style(value interface{}) *HBox {
	a.Set("style", value)
	return a
}

func (a *HBox) SubFormHorizontal(value interface{}) *HBox {
	a.Set("subFormHorizontal", value)
	return a
}

func (a *HBox) SubFormMode(value interface{}) *HBox {
	a.Set("subFormMode", value)
	return a
}

func (a *HBox) TestIdBuilder(value interface{}) *HBox {
	a.Set("testIdBuilder", value)
	return a
}

func (a *HBox) Testid(value interface{}) *HBox {
	a.Set("testid", value)
	return a
}

func (a *HBox) Type(value interface{}) *HBox {
	a.Set("type", value)
	return a
}

func (a *HBox) UseMobileUI(value interface{}) *HBox {
	a.Set("useMobileUI", value)
	return a
}

func (a *HBox) Valign(value interface{}) *HBox {
	a.Set("valign", value)
	return a
}

func (a *HBox) Visible(value interface{}) *HBox {
	a.Set("visible", value)
	return a
}

func (a *HBox) VisibleOn(value interface{}) *HBox {
	a.Set("visibleOn", value)
	return a
}

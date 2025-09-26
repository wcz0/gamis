package renderers

type Remark struct {
	*BaseRenderer
}

func NewRemark() *Remark {
	a := &Remark{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "remark")
	return a
}

func (a *Remark) Set(name string, value interface{}) *Remark {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Remark) ClassName(value interface{}) *Remark {
	a.Set("className", value)
	return a
}

func (a *Remark) Content(value interface{}) *Remark {
	a.Set("content", value)
	return a
}

func (a *Remark) Disabled(value interface{}) *Remark {
	a.Set("disabled", value)
	return a
}

func (a *Remark) DisabledOn(value interface{}) *Remark {
	a.Set("disabledOn", value)
	return a
}

func (a *Remark) EditorSetting(value interface{}) *Remark {
	a.Set("editorSetting", value)
	return a
}

func (a *Remark) Hidden(value interface{}) *Remark {
	a.Set("hidden", value)
	return a
}

func (a *Remark) HiddenOn(value interface{}) *Remark {
	a.Set("hiddenOn", value)
	return a
}

func (a *Remark) Icon(value interface{}) *Remark {
	a.Set("icon", value)
	return a
}

func (a *Remark) Id(value interface{}) *Remark {
	a.Set("id", value)
	return a
}

func (a *Remark) Label(value interface{}) *Remark {
	a.Set("label", value)
	return a
}

func (a *Remark) OnEvent(value interface{}) *Remark {
	a.Set("onEvent", value)
	return a
}

func (a *Remark) Placement(value interface{}) *Remark {
	a.Set("placement", value)
	return a
}

func (a *Remark) RootClose(value interface{}) *Remark {
	a.Set("rootClose", value)
	return a
}

func (a *Remark) Shape(value interface{}) *Remark {
	a.Set("shape", value)
	return a
}

func (a *Remark) Static(value interface{}) *Remark {
	a.Set("static", value)
	return a
}

func (a *Remark) StaticClassName(value interface{}) *Remark {
	a.Set("staticClassName", value)
	return a
}

func (a *Remark) StaticInputClassName(value interface{}) *Remark {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Remark) StaticLabelClassName(value interface{}) *Remark {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Remark) StaticOn(value interface{}) *Remark {
	a.Set("staticOn", value)
	return a
}

func (a *Remark) StaticPlaceholder(value interface{}) *Remark {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Remark) StaticSchema(value interface{}) *Remark {
	a.Set("staticSchema", value)
	return a
}

func (a *Remark) Style(value interface{}) *Remark {
	a.Set("style", value)
	return a
}

func (a *Remark) TestIdBuilder(value interface{}) *Remark {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Remark) Testid(value interface{}) *Remark {
	a.Set("testid", value)
	return a
}

func (a *Remark) Title(value interface{}) *Remark {
	a.Set("title", value)
	return a
}

func (a *Remark) TooltipClassName(value interface{}) *Remark {
	a.Set("tooltipClassName", value)
	return a
}

func (a *Remark) Trigger(value interface{}) *Remark {
	a.Set("trigger", value)
	return a
}

func (a *Remark) Type(value interface{}) *Remark {
	a.Set("type", value)
	return a
}

func (a *Remark) UseMobileUI(value interface{}) *Remark {
	a.Set("useMobileUI", value)
	return a
}

func (a *Remark) Visible(value interface{}) *Remark {
	a.Set("visible", value)
	return a
}

func (a *Remark) VisibleOn(value interface{}) *Remark {
	a.Set("visibleOn", value)
	return a
}

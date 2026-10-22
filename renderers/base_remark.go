package renderers

type BaseRemark struct {
	*BaseRenderer
}

func NewBaseRemark() *BaseRemark {
	a := &BaseRemark{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *BaseRemark) Set(name string, value interface{}) *BaseRemark {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *BaseRemark) ClassName(value interface{}) *BaseRemark {
	a.Set("className", value)
	return a
}

func (a *BaseRemark) Content(value interface{}) *BaseRemark {
	a.Set("content", value)
	return a
}

func (a *BaseRemark) Disabled(value interface{}) *BaseRemark {
	a.Set("disabled", value)
	return a
}

func (a *BaseRemark) DisabledOn(value interface{}) *BaseRemark {
	a.Set("disabledOn", value)
	return a
}

func (a *BaseRemark) EditorSetting(value interface{}) *BaseRemark {
	a.Set("editorSetting", value)
	return a
}

func (a *BaseRemark) Hidden(value interface{}) *BaseRemark {
	a.Set("hidden", value)
	return a
}

func (a *BaseRemark) HiddenOn(value interface{}) *BaseRemark {
	a.Set("hiddenOn", value)
	return a
}

func (a *BaseRemark) Icon(value interface{}) *BaseRemark {
	a.Set("icon", value)
	return a
}

func (a *BaseRemark) Id(value interface{}) *BaseRemark {
	a.Set("id", value)
	return a
}

func (a *BaseRemark) Label(value interface{}) *BaseRemark {
	a.Set("label", value)
	return a
}

func (a *BaseRemark) OnEvent(value interface{}) *BaseRemark {
	a.Set("onEvent", value)
	return a
}

func (a *BaseRemark) Placement(value interface{}) *BaseRemark {
	a.Set("placement", value)
	return a
}

func (a *BaseRemark) RootClose(value interface{}) *BaseRemark {
	a.Set("rootClose", value)
	return a
}

func (a *BaseRemark) Shape(value interface{}) *BaseRemark {
	a.Set("shape", value)
	return a
}

func (a *BaseRemark) Static(value interface{}) *BaseRemark {
	a.Set("static", value)
	return a
}

func (a *BaseRemark) StaticClassName(value interface{}) *BaseRemark {
	a.Set("staticClassName", value)
	return a
}

func (a *BaseRemark) StaticInputClassName(value interface{}) *BaseRemark {
	a.Set("staticInputClassName", value)
	return a
}

func (a *BaseRemark) StaticLabelClassName(value interface{}) *BaseRemark {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *BaseRemark) StaticOn(value interface{}) *BaseRemark {
	a.Set("staticOn", value)
	return a
}

func (a *BaseRemark) StaticPlaceholder(value interface{}) *BaseRemark {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *BaseRemark) StaticSchema(value interface{}) *BaseRemark {
	a.Set("staticSchema", value)
	return a
}

func (a *BaseRemark) Style(value interface{}) *BaseRemark {
	a.Set("style", value)
	return a
}

func (a *BaseRemark) Testid(value interface{}) *BaseRemark {
	a.Set("testid", value)
	return a
}

func (a *BaseRemark) Title(value interface{}) *BaseRemark {
	a.Set("title", value)
	return a
}

func (a *BaseRemark) TooltipClassName(value interface{}) *BaseRemark {
	a.Set("tooltipClassName", value)
	return a
}

func (a *BaseRemark) Trigger(value interface{}) *BaseRemark {
	a.Set("trigger", value)
	return a
}

func (a *BaseRemark) Type(value interface{}) *BaseRemark {
	a.Set("type", value)
	return a
}

func (a *BaseRemark) UseMobileUI(value interface{}) *BaseRemark {
	a.Set("useMobileUI", value)
	return a
}

func (a *BaseRemark) Visible(value interface{}) *BaseRemark {
	a.Set("visible", value)
	return a
}

func (a *BaseRemark) VisibleOn(value interface{}) *BaseRemark {
	a.Set("visibleOn", value)
	return a
}

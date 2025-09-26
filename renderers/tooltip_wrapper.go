package renderers

type TooltipWrapper struct {
	*BaseRenderer
}

func NewTooltipWrapper() *TooltipWrapper {
	a := &TooltipWrapper{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "tooltip-wrapper")
	return a
}

func (a *TooltipWrapper) Set(name string, value interface{}) *TooltipWrapper {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *TooltipWrapper) Body(value interface{}) *TooltipWrapper {
	a.Set("body", value)
	return a
}

func (a *TooltipWrapper) ClassName(value interface{}) *TooltipWrapper {
	a.Set("className", value)
	return a
}

func (a *TooltipWrapper) Content(value interface{}) *TooltipWrapper {
	a.Set("content", value)
	return a
}

func (a *TooltipWrapper) Disabled(value interface{}) *TooltipWrapper {
	a.Set("disabled", value)
	return a
}

func (a *TooltipWrapper) DisabledOn(value interface{}) *TooltipWrapper {
	a.Set("disabledOn", value)
	return a
}

func (a *TooltipWrapper) EditorSetting(value interface{}) *TooltipWrapper {
	a.Set("editorSetting", value)
	return a
}

func (a *TooltipWrapper) Enterable(value interface{}) *TooltipWrapper {
	a.Set("enterable", value)
	return a
}

func (a *TooltipWrapper) Hidden(value interface{}) *TooltipWrapper {
	a.Set("hidden", value)
	return a
}

func (a *TooltipWrapper) HiddenOn(value interface{}) *TooltipWrapper {
	a.Set("hiddenOn", value)
	return a
}

func (a *TooltipWrapper) Id(value interface{}) *TooltipWrapper {
	a.Set("id", value)
	return a
}

func (a *TooltipWrapper) Inline(value interface{}) *TooltipWrapper {
	a.Set("inline", value)
	return a
}

func (a *TooltipWrapper) MouseEnterDelay(value interface{}) *TooltipWrapper {
	a.Set("mouseEnterDelay", value)
	return a
}

func (a *TooltipWrapper) MouseLeaveDelay(value interface{}) *TooltipWrapper {
	a.Set("mouseLeaveDelay", value)
	return a
}

func (a *TooltipWrapper) Offset(value interface{}) *TooltipWrapper {
	a.Set("offset", value)
	return a
}

func (a *TooltipWrapper) OnEvent(value interface{}) *TooltipWrapper {
	a.Set("onEvent", value)
	return a
}

func (a *TooltipWrapper) Placement(value interface{}) *TooltipWrapper {
	a.Set("placement", value)
	return a
}

func (a *TooltipWrapper) RootClose(value interface{}) *TooltipWrapper {
	a.Set("rootClose", value)
	return a
}

func (a *TooltipWrapper) ShowArrow(value interface{}) *TooltipWrapper {
	a.Set("showArrow", value)
	return a
}

func (a *TooltipWrapper) Static(value interface{}) *TooltipWrapper {
	a.Set("static", value)
	return a
}

func (a *TooltipWrapper) StaticClassName(value interface{}) *TooltipWrapper {
	a.Set("staticClassName", value)
	return a
}

func (a *TooltipWrapper) StaticInputClassName(value interface{}) *TooltipWrapper {
	a.Set("staticInputClassName", value)
	return a
}

func (a *TooltipWrapper) StaticLabelClassName(value interface{}) *TooltipWrapper {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *TooltipWrapper) StaticOn(value interface{}) *TooltipWrapper {
	a.Set("staticOn", value)
	return a
}

func (a *TooltipWrapper) StaticPlaceholder(value interface{}) *TooltipWrapper {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *TooltipWrapper) StaticSchema(value interface{}) *TooltipWrapper {
	a.Set("staticSchema", value)
	return a
}

func (a *TooltipWrapper) Style(value interface{}) *TooltipWrapper {
	a.Set("style", value)
	return a
}

func (a *TooltipWrapper) TestIdBuilder(value interface{}) *TooltipWrapper {
	a.Set("testIdBuilder", value)
	return a
}

func (a *TooltipWrapper) Testid(value interface{}) *TooltipWrapper {
	a.Set("testid", value)
	return a
}

func (a *TooltipWrapper) Title(value interface{}) *TooltipWrapper {
	a.Set("title", value)
	return a
}

func (a *TooltipWrapper) Tooltip(value interface{}) *TooltipWrapper {
	a.Set("tooltip", value)
	return a
}

func (a *TooltipWrapper) TooltipArrowClassName(value interface{}) *TooltipWrapper {
	a.Set("tooltipArrowClassName", value)
	return a
}

func (a *TooltipWrapper) TooltipClassName(value interface{}) *TooltipWrapper {
	a.Set("tooltipClassName", value)
	return a
}

func (a *TooltipWrapper) TooltipStyle(value interface{}) *TooltipWrapper {
	a.Set("tooltipStyle", value)
	return a
}

func (a *TooltipWrapper) TooltipTheme(value interface{}) *TooltipWrapper {
	a.Set("tooltipTheme", value)
	return a
}

func (a *TooltipWrapper) Trigger(value interface{}) *TooltipWrapper {
	a.Set("trigger", value)
	return a
}

func (a *TooltipWrapper) Type(value interface{}) *TooltipWrapper {
	a.Set("type", value)
	return a
}

func (a *TooltipWrapper) UseMobileUI(value interface{}) *TooltipWrapper {
	a.Set("useMobileUI", value)
	return a
}

func (a *TooltipWrapper) Visible(value interface{}) *TooltipWrapper {
	a.Set("visible", value)
	return a
}

func (a *TooltipWrapper) VisibleOn(value interface{}) *TooltipWrapper {
	a.Set("visibleOn", value)
	return a
}

func (a *TooltipWrapper) WrapperComponent(value interface{}) *TooltipWrapper {
	a.Set("wrapperComponent", value)
	return a
}

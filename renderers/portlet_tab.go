package renderers

type PortletTab struct {
	*BaseRenderer
}

func NewPortletTab() *PortletTab {
	a := &PortletTab{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *PortletTab) Set(name string, value interface{}) *PortletTab {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *PortletTab) Body(value interface{}) *PortletTab {
	a.Set("body", value)
	return a
}

func (a *PortletTab) ClassName(value interface{}) *PortletTab {
	a.Set("className", value)
	return a
}

func (a *PortletTab) Disabled(value interface{}) *PortletTab {
	a.Set("disabled", value)
	return a
}

func (a *PortletTab) DisabledOn(value interface{}) *PortletTab {
	a.Set("disabledOn", value)
	return a
}

func (a *PortletTab) EditorSetting(value interface{}) *PortletTab {
	a.Set("editorSetting", value)
	return a
}

func (a *PortletTab) Hidden(value interface{}) *PortletTab {
	a.Set("hidden", value)
	return a
}

func (a *PortletTab) HiddenOn(value interface{}) *PortletTab {
	a.Set("hiddenOn", value)
	return a
}

func (a *PortletTab) Icon(value interface{}) *PortletTab {
	a.Set("icon", value)
	return a
}

func (a *PortletTab) IconPosition(value interface{}) *PortletTab {
	a.Set("iconPosition", value)
	return a
}

func (a *PortletTab) Id(value interface{}) *PortletTab {
	a.Set("id", value)
	return a
}

func (a *PortletTab) MountOnEnter(value interface{}) *PortletTab {
	a.Set("mountOnEnter", value)
	return a
}

func (a *PortletTab) OnEvent(value interface{}) *PortletTab {
	a.Set("onEvent", value)
	return a
}

func (a *PortletTab) Reload(value interface{}) *PortletTab {
	a.Set("reload", value)
	return a
}

func (a *PortletTab) Static(value interface{}) *PortletTab {
	a.Set("static", value)
	return a
}

func (a *PortletTab) StaticClassName(value interface{}) *PortletTab {
	a.Set("staticClassName", value)
	return a
}

func (a *PortletTab) StaticInputClassName(value interface{}) *PortletTab {
	a.Set("staticInputClassName", value)
	return a
}

func (a *PortletTab) StaticLabelClassName(value interface{}) *PortletTab {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *PortletTab) StaticOn(value interface{}) *PortletTab {
	a.Set("staticOn", value)
	return a
}

func (a *PortletTab) StaticPlaceholder(value interface{}) *PortletTab {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *PortletTab) StaticSchema(value interface{}) *PortletTab {
	a.Set("staticSchema", value)
	return a
}

func (a *PortletTab) Style(value interface{}) *PortletTab {
	a.Set("style", value)
	return a
}

func (a *PortletTab) Tab(value interface{}) *PortletTab {
	a.Set("tab", value)
	return a
}

func (a *PortletTab) TestIdBuilder(value interface{}) *PortletTab {
	a.Set("testIdBuilder", value)
	return a
}

func (a *PortletTab) Testid(value interface{}) *PortletTab {
	a.Set("testid", value)
	return a
}

func (a *PortletTab) Title(value interface{}) *PortletTab {
	a.Set("title", value)
	return a
}

func (a *PortletTab) Toolbar(value interface{}) *PortletTab {
	a.Set("toolbar", value)
	return a
}

func (a *PortletTab) UnmountOnExit(value interface{}) *PortletTab {
	a.Set("unmountOnExit", value)
	return a
}

func (a *PortletTab) UseMobileUI(value interface{}) *PortletTab {
	a.Set("useMobileUI", value)
	return a
}

func (a *PortletTab) Visible(value interface{}) *PortletTab {
	a.Set("visible", value)
	return a
}

func (a *PortletTab) VisibleOn(value interface{}) *PortletTab {
	a.Set("visibleOn", value)
	return a
}

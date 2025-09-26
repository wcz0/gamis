package renderers

type Tabs struct {
	*BaseRenderer
}

func NewTabs() *Tabs {
	a := &Tabs{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "tabs")
	return a
}

func (a *Tabs) Set(name string, value interface{}) *Tabs {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Tabs) ActiveKey(value interface{}) *Tabs {
	a.Set("activeKey", value)
	return a
}

func (a *Tabs) AddBtnText(value interface{}) *Tabs {
	a.Set("addBtnText", value)
	return a
}

func (a *Tabs) Addable(value interface{}) *Tabs {
	a.Set("addable", value)
	return a
}

func (a *Tabs) ClassName(value interface{}) *Tabs {
	a.Set("className", value)
	return a
}

func (a *Tabs) Closable(value interface{}) *Tabs {
	a.Set("closable", value)
	return a
}

func (a *Tabs) CollapseBtnLabel(value interface{}) *Tabs {
	a.Set("collapseBtnLabel", value)
	return a
}

func (a *Tabs) CollapseOnExceed(value interface{}) *Tabs {
	a.Set("collapseOnExceed", value)
	return a
}

func (a *Tabs) ContentClassName(value interface{}) *Tabs {
	a.Set("contentClassName", value)
	return a
}

func (a *Tabs) DefaultKey(value interface{}) *Tabs {
	a.Set("defaultKey", value)
	return a
}

func (a *Tabs) Disabled(value interface{}) *Tabs {
	a.Set("disabled", value)
	return a
}

func (a *Tabs) DisabledOn(value interface{}) *Tabs {
	a.Set("disabledOn", value)
	return a
}

func (a *Tabs) Draggable(value interface{}) *Tabs {
	a.Set("draggable", value)
	return a
}

func (a *Tabs) Editable(value interface{}) *Tabs {
	a.Set("editable", value)
	return a
}

func (a *Tabs) EditorSetting(value interface{}) *Tabs {
	a.Set("editorSetting", value)
	return a
}

func (a *Tabs) Hidden(value interface{}) *Tabs {
	a.Set("hidden", value)
	return a
}

func (a *Tabs) HiddenOn(value interface{}) *Tabs {
	a.Set("hiddenOn", value)
	return a
}

func (a *Tabs) Id(value interface{}) *Tabs {
	a.Set("id", value)
	return a
}

func (a *Tabs) LinksClassName(value interface{}) *Tabs {
	a.Set("linksClassName", value)
	return a
}

func (a *Tabs) MountOnEnter(value interface{}) *Tabs {
	a.Set("mountOnEnter", value)
	return a
}

func (a *Tabs) OnEvent(value interface{}) *Tabs {
	a.Set("onEvent", value)
	return a
}

func (a *Tabs) Scrollable(value interface{}) *Tabs {
	a.Set("scrollable", value)
	return a
}

func (a *Tabs) ShowTip(value interface{}) *Tabs {
	a.Set("showTip", value)
	return a
}

func (a *Tabs) ShowTipClassName(value interface{}) *Tabs {
	a.Set("showTipClassName", value)
	return a
}

func (a *Tabs) SidePosition(value interface{}) *Tabs {
	a.Set("sidePosition", value)
	return a
}

func (a *Tabs) Source(value interface{}) *Tabs {
	a.Set("source", value)
	return a
}

func (a *Tabs) Static(value interface{}) *Tabs {
	a.Set("static", value)
	return a
}

func (a *Tabs) StaticClassName(value interface{}) *Tabs {
	a.Set("staticClassName", value)
	return a
}

func (a *Tabs) StaticInputClassName(value interface{}) *Tabs {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Tabs) StaticLabelClassName(value interface{}) *Tabs {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Tabs) StaticOn(value interface{}) *Tabs {
	a.Set("staticOn", value)
	return a
}

func (a *Tabs) StaticPlaceholder(value interface{}) *Tabs {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Tabs) StaticSchema(value interface{}) *Tabs {
	a.Set("staticSchema", value)
	return a
}

func (a *Tabs) Style(value interface{}) *Tabs {
	a.Set("style", value)
	return a
}

func (a *Tabs) SubFormHorizontal(value interface{}) *Tabs {
	a.Set("subFormHorizontal", value)
	return a
}

func (a *Tabs) SubFormMode(value interface{}) *Tabs {
	a.Set("subFormMode", value)
	return a
}

func (a *Tabs) Swipeable(value interface{}) *Tabs {
	a.Set("swipeable", value)
	return a
}

func (a *Tabs) Tabs(value interface{}) *Tabs {
	a.Set("tabs", value)
	return a
}

func (a *Tabs) TabsMode(value interface{}) *Tabs {
	a.Set("tabsMode", value)
	return a
}

func (a *Tabs) TestIdBuilder(value interface{}) *Tabs {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Tabs) Testid(value interface{}) *Tabs {
	a.Set("testid", value)
	return a
}

func (a *Tabs) Toolbar(value interface{}) *Tabs {
	a.Set("toolbar", value)
	return a
}

func (a *Tabs) Type(value interface{}) *Tabs {
	a.Set("type", value)
	return a
}

func (a *Tabs) UnmountOnExit(value interface{}) *Tabs {
	a.Set("unmountOnExit", value)
	return a
}

func (a *Tabs) UseMobileUI(value interface{}) *Tabs {
	a.Set("useMobileUI", value)
	return a
}

func (a *Tabs) Visible(value interface{}) *Tabs {
	a.Set("visible", value)
	return a
}

func (a *Tabs) VisibleOn(value interface{}) *Tabs {
	a.Set("visibleOn", value)
	return a
}

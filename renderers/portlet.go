package renderers

type Portlet struct {
	*BaseRenderer
}

func NewPortlet() *Portlet {
	a := &Portlet{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "portlet")
	return a
}

func (a *Portlet) Set(name string, value interface{}) *Portlet {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Portlet) ClassName(value interface{}) *Portlet {
	a.Set("className", value)
	return a
}

func (a *Portlet) ContentClassName(value interface{}) *Portlet {
	a.Set("contentClassName", value)
	return a
}

func (a *Portlet) Description(value interface{}) *Portlet {
	a.Set("description", value)
	return a
}

func (a *Portlet) Disabled(value interface{}) *Portlet {
	a.Set("disabled", value)
	return a
}

func (a *Portlet) DisabledOn(value interface{}) *Portlet {
	a.Set("disabledOn", value)
	return a
}

func (a *Portlet) Divider(value interface{}) *Portlet {
	a.Set("divider", value)
	return a
}

func (a *Portlet) EditorSetting(value interface{}) *Portlet {
	a.Set("editorSetting", value)
	return a
}

func (a *Portlet) Hidden(value interface{}) *Portlet {
	a.Set("hidden", value)
	return a
}

func (a *Portlet) HiddenOn(value interface{}) *Portlet {
	a.Set("hiddenOn", value)
	return a
}

func (a *Portlet) HideHeader(value interface{}) *Portlet {
	a.Set("hideHeader", value)
	return a
}

func (a *Portlet) Id(value interface{}) *Portlet {
	a.Set("id", value)
	return a
}

func (a *Portlet) LinksClassName(value interface{}) *Portlet {
	a.Set("linksClassName", value)
	return a
}

func (a *Portlet) MountOnEnter(value interface{}) *Portlet {
	a.Set("mountOnEnter", value)
	return a
}

func (a *Portlet) OnEvent(value interface{}) *Portlet {
	a.Set("onEvent", value)
	return a
}

func (a *Portlet) Scrollable(value interface{}) *Portlet {
	a.Set("scrollable", value)
	return a
}

func (a *Portlet) Source(value interface{}) *Portlet {
	a.Set("source", value)
	return a
}

func (a *Portlet) Static(value interface{}) *Portlet {
	a.Set("static", value)
	return a
}

func (a *Portlet) StaticClassName(value interface{}) *Portlet {
	a.Set("staticClassName", value)
	return a
}

func (a *Portlet) StaticInputClassName(value interface{}) *Portlet {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Portlet) StaticLabelClassName(value interface{}) *Portlet {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Portlet) StaticOn(value interface{}) *Portlet {
	a.Set("staticOn", value)
	return a
}

func (a *Portlet) StaticPlaceholder(value interface{}) *Portlet {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Portlet) StaticSchema(value interface{}) *Portlet {
	a.Set("staticSchema", value)
	return a
}

func (a *Portlet) Style(value interface{}) *Portlet {
	a.Set("style", value)
	return a
}

func (a *Portlet) Tabs(value interface{}) *Portlet {
	a.Set("tabs", value)
	return a
}

func (a *Portlet) TabsClassName(value interface{}) *Portlet {
	a.Set("tabsClassName", value)
	return a
}

func (a *Portlet) TabsMode(value interface{}) *Portlet {
	a.Set("tabsMode", value)
	return a
}

func (a *Portlet) TestIdBuilder(value interface{}) *Portlet {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Portlet) Testid(value interface{}) *Portlet {
	a.Set("testid", value)
	return a
}

func (a *Portlet) Toolbar(value interface{}) *Portlet {
	a.Set("toolbar", value)
	return a
}

func (a *Portlet) Type(value interface{}) *Portlet {
	a.Set("type", value)
	return a
}

func (a *Portlet) UnmountOnExit(value interface{}) *Portlet {
	a.Set("unmountOnExit", value)
	return a
}

func (a *Portlet) UseMobileUI(value interface{}) *Portlet {
	a.Set("useMobileUI", value)
	return a
}

func (a *Portlet) Visible(value interface{}) *Portlet {
	a.Set("visible", value)
	return a
}

func (a *Portlet) VisibleOn(value interface{}) *Portlet {
	a.Set("visibleOn", value)
	return a
}

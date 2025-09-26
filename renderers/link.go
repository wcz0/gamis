package renderers

type Link struct {
	*BaseRenderer
}

func NewLink() *Link {
	a := &Link{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "link")
	return a
}

func (a *Link) Set(name string, value interface{}) *Link {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Link) Badge(value interface{}) *Link {
	a.Set("badge", value)
	return a
}

func (a *Link) Blank(value interface{}) *Link {
	a.Set("blank", value)
	return a
}

func (a *Link) Body(value interface{}) *Link {
	a.Set("body", value)
	return a
}

func (a *Link) ClassName(value interface{}) *Link {
	a.Set("className", value)
	return a
}

func (a *Link) Disabled(value interface{}) *Link {
	a.Set("disabled", value)
	return a
}

func (a *Link) DisabledOn(value interface{}) *Link {
	a.Set("disabledOn", value)
	return a
}

func (a *Link) EditorSetting(value interface{}) *Link {
	a.Set("editorSetting", value)
	return a
}

func (a *Link) Hidden(value interface{}) *Link {
	a.Set("hidden", value)
	return a
}

func (a *Link) HiddenOn(value interface{}) *Link {
	a.Set("hiddenOn", value)
	return a
}

func (a *Link) Href(value interface{}) *Link {
	a.Set("href", value)
	return a
}

func (a *Link) HtmlTarget(value interface{}) *Link {
	a.Set("htmlTarget", value)
	return a
}

func (a *Link) Icon(value interface{}) *Link {
	a.Set("icon", value)
	return a
}

func (a *Link) Id(value interface{}) *Link {
	a.Set("id", value)
	return a
}

func (a *Link) OnEvent(value interface{}) *Link {
	a.Set("onEvent", value)
	return a
}

func (a *Link) RightIcon(value interface{}) *Link {
	a.Set("rightIcon", value)
	return a
}

func (a *Link) Static(value interface{}) *Link {
	a.Set("static", value)
	return a
}

func (a *Link) StaticClassName(value interface{}) *Link {
	a.Set("staticClassName", value)
	return a
}

func (a *Link) StaticInputClassName(value interface{}) *Link {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Link) StaticLabelClassName(value interface{}) *Link {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Link) StaticOn(value interface{}) *Link {
	a.Set("staticOn", value)
	return a
}

func (a *Link) StaticPlaceholder(value interface{}) *Link {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Link) StaticSchema(value interface{}) *Link {
	a.Set("staticSchema", value)
	return a
}

func (a *Link) Style(value interface{}) *Link {
	a.Set("style", value)
	return a
}

func (a *Link) TestIdBuilder(value interface{}) *Link {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Link) Testid(value interface{}) *Link {
	a.Set("testid", value)
	return a
}

func (a *Link) Type(value interface{}) *Link {
	a.Set("type", value)
	return a
}

func (a *Link) UseMobileUI(value interface{}) *Link {
	a.Set("useMobileUI", value)
	return a
}

func (a *Link) Visible(value interface{}) *Link {
	a.Set("visible", value)
	return a
}

func (a *Link) VisibleOn(value interface{}) *Link {
	a.Set("visibleOn", value)
	return a
}

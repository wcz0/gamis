package renderers

type AnchorNav struct {
	*BaseRenderer
}

func NewAnchorNav() *AnchorNav {
	a := &AnchorNav{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "anchor-nav")
	return a
}

func (a *AnchorNav) Set(name string, value interface{}) *AnchorNav {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *AnchorNav) Active(value interface{}) *AnchorNav {
	a.Set("active", value)
	return a
}

func (a *AnchorNav) ClassName(value interface{}) *AnchorNav {
	a.Set("className", value)
	return a
}

func (a *AnchorNav) Direction(value interface{}) *AnchorNav {
	a.Set("direction", value)
	return a
}

func (a *AnchorNav) Disabled(value interface{}) *AnchorNav {
	a.Set("disabled", value)
	return a
}

func (a *AnchorNav) DisabledOn(value interface{}) *AnchorNav {
	a.Set("disabledOn", value)
	return a
}

func (a *AnchorNav) EditorSetting(value interface{}) *AnchorNav {
	a.Set("editorSetting", value)
	return a
}

func (a *AnchorNav) Hidden(value interface{}) *AnchorNav {
	a.Set("hidden", value)
	return a
}

func (a *AnchorNav) HiddenOn(value interface{}) *AnchorNav {
	a.Set("hiddenOn", value)
	return a
}

func (a *AnchorNav) Id(value interface{}) *AnchorNav {
	a.Set("id", value)
	return a
}

func (a *AnchorNav) LinkClassName(value interface{}) *AnchorNav {
	a.Set("linkClassName", value)
	return a
}

func (a *AnchorNav) Links(value interface{}) *AnchorNav {
	a.Set("links", value)
	return a
}

func (a *AnchorNav) OnEvent(value interface{}) *AnchorNav {
	a.Set("onEvent", value)
	return a
}

func (a *AnchorNav) SectionClassName(value interface{}) *AnchorNav {
	a.Set("sectionClassName", value)
	return a
}

func (a *AnchorNav) Static(value interface{}) *AnchorNav {
	a.Set("static", value)
	return a
}

func (a *AnchorNav) StaticClassName(value interface{}) *AnchorNav {
	a.Set("staticClassName", value)
	return a
}

func (a *AnchorNav) StaticInputClassName(value interface{}) *AnchorNav {
	a.Set("staticInputClassName", value)
	return a
}

func (a *AnchorNav) StaticLabelClassName(value interface{}) *AnchorNav {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *AnchorNav) StaticOn(value interface{}) *AnchorNav {
	a.Set("staticOn", value)
	return a
}

func (a *AnchorNav) StaticPlaceholder(value interface{}) *AnchorNav {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *AnchorNav) StaticSchema(value interface{}) *AnchorNav {
	a.Set("staticSchema", value)
	return a
}

func (a *AnchorNav) Style(value interface{}) *AnchorNav {
	a.Set("style", value)
	return a
}

func (a *AnchorNav) TestIdBuilder(value interface{}) *AnchorNav {
	a.Set("testIdBuilder", value)
	return a
}

func (a *AnchorNav) Testid(value interface{}) *AnchorNav {
	a.Set("testid", value)
	return a
}

func (a *AnchorNav) Type(value interface{}) *AnchorNav {
	a.Set("type", value)
	return a
}

func (a *AnchorNav) UseMobileUI(value interface{}) *AnchorNav {
	a.Set("useMobileUI", value)
	return a
}

func (a *AnchorNav) Visible(value interface{}) *AnchorNav {
	a.Set("visible", value)
	return a
}

func (a *AnchorNav) VisibleOn(value interface{}) *AnchorNav {
	a.Set("visibleOn", value)
	return a
}

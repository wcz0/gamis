package renderers

type AnchorNavSection struct {
	*BaseRenderer
}

func NewAnchorNavSection() *AnchorNavSection {
	a := &AnchorNavSection{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *AnchorNavSection) Set(name string, value interface{}) *AnchorNavSection {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *AnchorNavSection) Body(value interface{}) *AnchorNavSection {
	a.Set("body", value)
	return a
}

func (a *AnchorNavSection) Children(value interface{}) *AnchorNavSection {
	a.Set("children", value)
	return a
}

func (a *AnchorNavSection) ClassName(value interface{}) *AnchorNavSection {
	a.Set("className", value)
	return a
}

func (a *AnchorNavSection) Disabled(value interface{}) *AnchorNavSection {
	a.Set("disabled", value)
	return a
}

func (a *AnchorNavSection) DisabledOn(value interface{}) *AnchorNavSection {
	a.Set("disabledOn", value)
	return a
}

func (a *AnchorNavSection) EditorSetting(value interface{}) *AnchorNavSection {
	a.Set("editorSetting", value)
	return a
}

func (a *AnchorNavSection) Hidden(value interface{}) *AnchorNavSection {
	a.Set("hidden", value)
	return a
}

func (a *AnchorNavSection) HiddenOn(value interface{}) *AnchorNavSection {
	a.Set("hiddenOn", value)
	return a
}

func (a *AnchorNavSection) Href(value interface{}) *AnchorNavSection {
	a.Set("href", value)
	return a
}

func (a *AnchorNavSection) Id(value interface{}) *AnchorNavSection {
	a.Set("id", value)
	return a
}

func (a *AnchorNavSection) OnEvent(value interface{}) *AnchorNavSection {
	a.Set("onEvent", value)
	return a
}

func (a *AnchorNavSection) Static(value interface{}) *AnchorNavSection {
	a.Set("static", value)
	return a
}

func (a *AnchorNavSection) StaticClassName(value interface{}) *AnchorNavSection {
	a.Set("staticClassName", value)
	return a
}

func (a *AnchorNavSection) StaticInputClassName(value interface{}) *AnchorNavSection {
	a.Set("staticInputClassName", value)
	return a
}

func (a *AnchorNavSection) StaticLabelClassName(value interface{}) *AnchorNavSection {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *AnchorNavSection) StaticOn(value interface{}) *AnchorNavSection {
	a.Set("staticOn", value)
	return a
}

func (a *AnchorNavSection) StaticPlaceholder(value interface{}) *AnchorNavSection {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *AnchorNavSection) StaticSchema(value interface{}) *AnchorNavSection {
	a.Set("staticSchema", value)
	return a
}

func (a *AnchorNavSection) Style(value interface{}) *AnchorNavSection {
	a.Set("style", value)
	return a
}

func (a *AnchorNavSection) TestIdBuilder(value interface{}) *AnchorNavSection {
	a.Set("testIdBuilder", value)
	return a
}

func (a *AnchorNavSection) Testid(value interface{}) *AnchorNavSection {
	a.Set("testid", value)
	return a
}

func (a *AnchorNavSection) Title(value interface{}) *AnchorNavSection {
	a.Set("title", value)
	return a
}

func (a *AnchorNavSection) UseMobileUI(value interface{}) *AnchorNavSection {
	a.Set("useMobileUI", value)
	return a
}

func (a *AnchorNavSection) Visible(value interface{}) *AnchorNavSection {
	a.Set("visible", value)
	return a
}

func (a *AnchorNavSection) VisibleOn(value interface{}) *AnchorNavSection {
	a.Set("visibleOn", value)
	return a
}

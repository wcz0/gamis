package renderers

type Timeline struct {
	*BaseRenderer
}

func NewTimeline() *Timeline {
	a := &Timeline{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "timeline")
	return a
}

func (a *Timeline) Set(name string, value interface{}) *Timeline {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Timeline) CardSchema(value interface{}) *Timeline {
	a.Set("cardSchema", value)
	return a
}

func (a *Timeline) ClassName(value interface{}) *Timeline {
	a.Set("className", value)
	return a
}

func (a *Timeline) DetailClassName(value interface{}) *Timeline {
	a.Set("detailClassName", value)
	return a
}

func (a *Timeline) Direction(value interface{}) *Timeline {
	a.Set("direction", value)
	return a
}

func (a *Timeline) Disabled(value interface{}) *Timeline {
	a.Set("disabled", value)
	return a
}

func (a *Timeline) DisabledOn(value interface{}) *Timeline {
	a.Set("disabledOn", value)
	return a
}

func (a *Timeline) EditorSetting(value interface{}) *Timeline {
	a.Set("editorSetting", value)
	return a
}

func (a *Timeline) Hidden(value interface{}) *Timeline {
	a.Set("hidden", value)
	return a
}

func (a *Timeline) HiddenOn(value interface{}) *Timeline {
	a.Set("hiddenOn", value)
	return a
}

func (a *Timeline) IconClassName(value interface{}) *Timeline {
	a.Set("iconClassName", value)
	return a
}

func (a *Timeline) Id(value interface{}) *Timeline {
	a.Set("id", value)
	return a
}

func (a *Timeline) ItemTitleSchema(value interface{}) *Timeline {
	a.Set("itemTitleSchema", value)
	return a
}

func (a *Timeline) Items(value interface{}) *Timeline {
	a.Set("items", value)
	return a
}

func (a *Timeline) Mode(value interface{}) *Timeline {
	a.Set("mode", value)
	return a
}

func (a *Timeline) OnEvent(value interface{}) *Timeline {
	a.Set("onEvent", value)
	return a
}

func (a *Timeline) Reverse(value interface{}) *Timeline {
	a.Set("reverse", value)
	return a
}

func (a *Timeline) Source(value interface{}) *Timeline {
	a.Set("source", value)
	return a
}

func (a *Timeline) Static(value interface{}) *Timeline {
	a.Set("static", value)
	return a
}

func (a *Timeline) StaticClassName(value interface{}) *Timeline {
	a.Set("staticClassName", value)
	return a
}

func (a *Timeline) StaticInputClassName(value interface{}) *Timeline {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Timeline) StaticLabelClassName(value interface{}) *Timeline {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Timeline) StaticOn(value interface{}) *Timeline {
	a.Set("staticOn", value)
	return a
}

func (a *Timeline) StaticPlaceholder(value interface{}) *Timeline {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Timeline) StaticSchema(value interface{}) *Timeline {
	a.Set("staticSchema", value)
	return a
}

func (a *Timeline) Style(value interface{}) *Timeline {
	a.Set("style", value)
	return a
}

func (a *Timeline) TestIdBuilder(value interface{}) *Timeline {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Timeline) Testid(value interface{}) *Timeline {
	a.Set("testid", value)
	return a
}

func (a *Timeline) TimeClassName(value interface{}) *Timeline {
	a.Set("timeClassName", value)
	return a
}

func (a *Timeline) TitleClassName(value interface{}) *Timeline {
	a.Set("titleClassName", value)
	return a
}

func (a *Timeline) Type(value interface{}) *Timeline {
	a.Set("type", value)
	return a
}

func (a *Timeline) UseMobileUI(value interface{}) *Timeline {
	a.Set("useMobileUI", value)
	return a
}

func (a *Timeline) Visible(value interface{}) *Timeline {
	a.Set("visible", value)
	return a
}

func (a *Timeline) VisibleOn(value interface{}) *Timeline {
	a.Set("visibleOn", value)
	return a
}

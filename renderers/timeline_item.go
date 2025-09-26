package renderers

type TimelineItem struct {
	*BaseRenderer
}

func NewTimelineItem() *TimelineItem {
	a := &TimelineItem{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *TimelineItem) Set(name string, value interface{}) *TimelineItem {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *TimelineItem) CardSchema(value interface{}) *TimelineItem {
	a.Set("cardSchema", value)
	return a
}

func (a *TimelineItem) ClassName(value interface{}) *TimelineItem {
	a.Set("className", value)
	return a
}

func (a *TimelineItem) Color(value interface{}) *TimelineItem {
	a.Set("color", value)
	return a
}

func (a *TimelineItem) Detail(value interface{}) *TimelineItem {
	a.Set("detail", value)
	return a
}

func (a *TimelineItem) DetailClassName(value interface{}) *TimelineItem {
	a.Set("detailClassName", value)
	return a
}

func (a *TimelineItem) DetailCollapsedText(value interface{}) *TimelineItem {
	a.Set("detailCollapsedText", value)
	return a
}

func (a *TimelineItem) DetailExpandedText(value interface{}) *TimelineItem {
	a.Set("detailExpandedText", value)
	return a
}

func (a *TimelineItem) Disabled(value interface{}) *TimelineItem {
	a.Set("disabled", value)
	return a
}

func (a *TimelineItem) DisabledOn(value interface{}) *TimelineItem {
	a.Set("disabledOn", value)
	return a
}

func (a *TimelineItem) DotSize(value interface{}) *TimelineItem {
	a.Set("dotSize", value)
	return a
}

func (a *TimelineItem) EditorSetting(value interface{}) *TimelineItem {
	a.Set("editorSetting", value)
	return a
}

func (a *TimelineItem) Hidden(value interface{}) *TimelineItem {
	a.Set("hidden", value)
	return a
}

func (a *TimelineItem) HiddenOn(value interface{}) *TimelineItem {
	a.Set("hiddenOn", value)
	return a
}

func (a *TimelineItem) HideDot(value interface{}) *TimelineItem {
	a.Set("hideDot", value)
	return a
}

func (a *TimelineItem) Icon(value interface{}) *TimelineItem {
	a.Set("icon", value)
	return a
}

func (a *TimelineItem) IconClassName(value interface{}) *TimelineItem {
	a.Set("iconClassName", value)
	return a
}

func (a *TimelineItem) Id(value interface{}) *TimelineItem {
	a.Set("id", value)
	return a
}

func (a *TimelineItem) LineColor(value interface{}) *TimelineItem {
	a.Set("lineColor", value)
	return a
}

func (a *TimelineItem) OnEvent(value interface{}) *TimelineItem {
	a.Set("onEvent", value)
	return a
}

func (a *TimelineItem) Static(value interface{}) *TimelineItem {
	a.Set("static", value)
	return a
}

func (a *TimelineItem) StaticClassName(value interface{}) *TimelineItem {
	a.Set("staticClassName", value)
	return a
}

func (a *TimelineItem) StaticInputClassName(value interface{}) *TimelineItem {
	a.Set("staticInputClassName", value)
	return a
}

func (a *TimelineItem) StaticLabelClassName(value interface{}) *TimelineItem {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *TimelineItem) StaticOn(value interface{}) *TimelineItem {
	a.Set("staticOn", value)
	return a
}

func (a *TimelineItem) StaticPlaceholder(value interface{}) *TimelineItem {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *TimelineItem) StaticSchema(value interface{}) *TimelineItem {
	a.Set("staticSchema", value)
	return a
}

func (a *TimelineItem) Style(value interface{}) *TimelineItem {
	a.Set("style", value)
	return a
}

func (a *TimelineItem) TestIdBuilder(value interface{}) *TimelineItem {
	a.Set("testIdBuilder", value)
	return a
}

func (a *TimelineItem) Testid(value interface{}) *TimelineItem {
	a.Set("testid", value)
	return a
}

func (a *TimelineItem) Time(value interface{}) *TimelineItem {
	a.Set("time", value)
	return a
}

func (a *TimelineItem) TimeClassName(value interface{}) *TimelineItem {
	a.Set("timeClassName", value)
	return a
}

func (a *TimelineItem) Title(value interface{}) *TimelineItem {
	a.Set("title", value)
	return a
}

func (a *TimelineItem) TitleClassName(value interface{}) *TimelineItem {
	a.Set("titleClassName", value)
	return a
}

func (a *TimelineItem) UseMobileUI(value interface{}) *TimelineItem {
	a.Set("useMobileUI", value)
	return a
}

func (a *TimelineItem) Visible(value interface{}) *TimelineItem {
	a.Set("visible", value)
	return a
}

func (a *TimelineItem) VisibleOn(value interface{}) *TimelineItem {
	a.Set("visibleOn", value)
	return a
}

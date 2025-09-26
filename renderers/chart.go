package renderers

type Chart struct {
	*BaseRenderer
}

func NewChart() *Chart {
	a := &Chart{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "chart")
	return a
}

func (a *Chart) Set(name string, value interface{}) *Chart {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Chart) Api(value interface{}) *Chart {
	a.Set("api", value)
	return a
}

func (a *Chart) ChartTheme(value interface{}) *Chart {
	a.Set("chartTheme", value)
	return a
}

func (a *Chart) ClassName(value interface{}) *Chart {
	a.Set("className", value)
	return a
}

func (a *Chart) ClickAction(value interface{}) *Chart {
	a.Set("clickAction", value)
	return a
}

func (a *Chart) Config(value interface{}) *Chart {
	a.Set("config", value)
	return a
}

func (a *Chart) DataFilter(value interface{}) *Chart {
	a.Set("dataFilter", value)
	return a
}

func (a *Chart) DisableDataMapping(value interface{}) *Chart {
	a.Set("disableDataMapping", value)
	return a
}

func (a *Chart) Disabled(value interface{}) *Chart {
	a.Set("disabled", value)
	return a
}

func (a *Chart) DisabledOn(value interface{}) *Chart {
	a.Set("disabledOn", value)
	return a
}

func (a *Chart) EditorSetting(value interface{}) *Chart {
	a.Set("editorSetting", value)
	return a
}

func (a *Chart) Height(value interface{}) *Chart {
	a.Set("height", value)
	return a
}

func (a *Chart) Hidden(value interface{}) *Chart {
	a.Set("hidden", value)
	return a
}

func (a *Chart) HiddenOn(value interface{}) *Chart {
	a.Set("hiddenOn", value)
	return a
}

func (a *Chart) Id(value interface{}) *Chart {
	a.Set("id", value)
	return a
}

func (a *Chart) InitFetch(value interface{}) *Chart {
	a.Set("initFetch", value)
	return a
}

func (a *Chart) InitFetchOn(value interface{}) *Chart {
	a.Set("initFetchOn", value)
	return a
}

func (a *Chart) Interval(value interface{}) *Chart {
	a.Set("interval", value)
	return a
}

func (a *Chart) LoadBaiduMap(value interface{}) *Chart {
	a.Set("loadBaiduMap", value)
	return a
}

func (a *Chart) MapName(value interface{}) *Chart {
	a.Set("mapName", value)
	return a
}

func (a *Chart) MapURL(value interface{}) *Chart {
	a.Set("mapURL", value)
	return a
}

func (a *Chart) Name(value interface{}) *Chart {
	a.Set("name", value)
	return a
}

func (a *Chart) OnEvent(value interface{}) *Chart {
	a.Set("onEvent", value)
	return a
}

func (a *Chart) ReplaceChartOption(value interface{}) *Chart {
	a.Set("replaceChartOption", value)
	return a
}

func (a *Chart) Source(value interface{}) *Chart {
	a.Set("source", value)
	return a
}

func (a *Chart) Static(value interface{}) *Chart {
	a.Set("static", value)
	return a
}

func (a *Chart) StaticClassName(value interface{}) *Chart {
	a.Set("staticClassName", value)
	return a
}

func (a *Chart) StaticInputClassName(value interface{}) *Chart {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Chart) StaticLabelClassName(value interface{}) *Chart {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Chart) StaticOn(value interface{}) *Chart {
	a.Set("staticOn", value)
	return a
}

func (a *Chart) StaticPlaceholder(value interface{}) *Chart {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Chart) StaticSchema(value interface{}) *Chart {
	a.Set("staticSchema", value)
	return a
}

func (a *Chart) Style(value interface{}) *Chart {
	a.Set("style", value)
	return a
}

func (a *Chart) TestIdBuilder(value interface{}) *Chart {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Chart) Testid(value interface{}) *Chart {
	a.Set("testid", value)
	return a
}

func (a *Chart) TrackExpression(value interface{}) *Chart {
	a.Set("trackExpression", value)
	return a
}

func (a *Chart) Type(value interface{}) *Chart {
	a.Set("type", value)
	return a
}

func (a *Chart) UnMountOnHidden(value interface{}) *Chart {
	a.Set("unMountOnHidden", value)
	return a
}

func (a *Chart) UseMobileUI(value interface{}) *Chart {
	a.Set("useMobileUI", value)
	return a
}

func (a *Chart) Visible(value interface{}) *Chart {
	a.Set("visible", value)
	return a
}

func (a *Chart) VisibleOn(value interface{}) *Chart {
	a.Set("visibleOn", value)
	return a
}

func (a *Chart) Width(value interface{}) *Chart {
	a.Set("width", value)
	return a
}

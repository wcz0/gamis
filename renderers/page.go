package renderers

type Page struct {
	*BaseRenderer
}

func NewPage() *Page {
	a := &Page{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "page")
	return a
}

func (a *Page) Set(name string, value interface{}) *Page {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Page) Aside(value interface{}) *Page {
	a.Set("aside", value)
	return a
}

func (a *Page) AsideClassName(value interface{}) *Page {
	a.Set("asideClassName", value)
	return a
}

func (a *Page) AsideMaxWidth(value interface{}) *Page {
	a.Set("asideMaxWidth", value)
	return a
}

func (a *Page) AsideMinWidth(value interface{}) *Page {
	a.Set("asideMinWidth", value)
	return a
}

func (a *Page) AsidePosition(value interface{}) *Page {
	a.Set("asidePosition", value)
	return a
}

func (a *Page) AsideResizor(value interface{}) *Page {
	a.Set("asideResizor", value)
	return a
}

func (a *Page) AsideSticky(value interface{}) *Page {
	a.Set("asideSticky", value)
	return a
}

func (a *Page) Body(value interface{}) *Page {
	a.Set("body", value)
	return a
}

func (a *Page) BodyClassName(value interface{}) *Page {
	a.Set("bodyClassName", value)
	return a
}

func (a *Page) ClassName(value interface{}) *Page {
	a.Set("className", value)
	return a
}

func (a *Page) Css(value interface{}) *Page {
	a.Set("css", value)
	return a
}

func (a *Page) CssVars(value interface{}) *Page {
	a.Set("cssVars", value)
	return a
}

func (a *Page) Data(value interface{}) *Page {
	a.Set("data", value)
	return a
}

func (a *Page) Definitions(value interface{}) *Page {
	a.Set("definitions", value)
	return a
}

func (a *Page) Disabled(value interface{}) *Page {
	a.Set("disabled", value)
	return a
}

func (a *Page) DisabledOn(value interface{}) *Page {
	a.Set("disabledOn", value)
	return a
}

func (a *Page) EditorSetting(value interface{}) *Page {
	a.Set("editorSetting", value)
	return a
}

func (a *Page) HeaderClassName(value interface{}) *Page {
	a.Set("headerClassName", value)
	return a
}

func (a *Page) Hidden(value interface{}) *Page {
	a.Set("hidden", value)
	return a
}

func (a *Page) HiddenOn(value interface{}) *Page {
	a.Set("hiddenOn", value)
	return a
}

func (a *Page) Id(value interface{}) *Page {
	a.Set("id", value)
	return a
}

func (a *Page) InitApi(value interface{}) *Page {
	a.Set("initApi", value)
	return a
}

func (a *Page) InitFetch(value interface{}) *Page {
	a.Set("initFetch", value)
	return a
}

func (a *Page) InitFetchOn(value interface{}) *Page {
	a.Set("initFetchOn", value)
	return a
}

func (a *Page) Interval(value interface{}) *Page {
	a.Set("interval", value)
	return a
}

func (a *Page) LoadingConfig(value interface{}) *Page {
	a.Set("loadingConfig", value)
	return a
}

func (a *Page) Messages(value interface{}) *Page {
	a.Set("messages", value)
	return a
}

func (a *Page) MobileCSS(value interface{}) *Page {
	a.Set("mobileCSS", value)
	return a
}

func (a *Page) Name(value interface{}) *Page {
	a.Set("name", value)
	return a
}

func (a *Page) OnEvent(value interface{}) *Page {
	a.Set("onEvent", value)
	return a
}

func (a *Page) PullRefresh(value interface{}) *Page {
	a.Set("pullRefresh", value)
	return a
}

func (a *Page) Regions(value interface{}) *Page {
	a.Set("regions", value)
	return a
}

func (a *Page) Remark(value interface{}) *Page {
	a.Set("remark", value)
	return a
}

func (a *Page) ShowErrorMsg(value interface{}) *Page {
	a.Set("showErrorMsg", value)
	return a
}

func (a *Page) SilentPolling(value interface{}) *Page {
	a.Set("silentPolling", value)
	return a
}

func (a *Page) Static(value interface{}) *Page {
	a.Set("static", value)
	return a
}

func (a *Page) StaticClassName(value interface{}) *Page {
	a.Set("staticClassName", value)
	return a
}

func (a *Page) StaticInputClassName(value interface{}) *Page {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Page) StaticLabelClassName(value interface{}) *Page {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Page) StaticOn(value interface{}) *Page {
	a.Set("staticOn", value)
	return a
}

func (a *Page) StaticPlaceholder(value interface{}) *Page {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Page) StaticSchema(value interface{}) *Page {
	a.Set("staticSchema", value)
	return a
}

func (a *Page) StopAutoRefreshWhen(value interface{}) *Page {
	a.Set("stopAutoRefreshWhen", value)
	return a
}

func (a *Page) Style(value interface{}) *Page {
	a.Set("style", value)
	return a
}

func (a *Page) SubTitle(value interface{}) *Page {
	a.Set("subTitle", value)
	return a
}

func (a *Page) TestIdBuilder(value interface{}) *Page {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Page) Testid(value interface{}) *Page {
	a.Set("testid", value)
	return a
}

func (a *Page) Title(value interface{}) *Page {
	a.Set("title", value)
	return a
}

func (a *Page) Toolbar(value interface{}) *Page {
	a.Set("toolbar", value)
	return a
}

func (a *Page) ToolbarClassName(value interface{}) *Page {
	a.Set("toolbarClassName", value)
	return a
}

func (a *Page) Type(value interface{}) *Page {
	a.Set("type", value)
	return a
}

func (a *Page) UseMobileUI(value interface{}) *Page {
	a.Set("useMobileUI", value)
	return a
}

func (a *Page) Visible(value interface{}) *Page {
	a.Set("visible", value)
	return a
}

func (a *Page) VisibleOn(value interface{}) *Page {
	a.Set("visibleOn", value)
	return a
}

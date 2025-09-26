package renderers

type Root struct {
	*BaseRenderer
}

func NewRoot() *Root {
	a := &Root{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "page")
	return a
}

func (a *Root) Set(name string, value interface{}) *Root {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Root) Aside(value interface{}) *Root {
	a.Set("aside", value)
	return a
}

func (a *Root) AsideClassName(value interface{}) *Root {
	a.Set("asideClassName", value)
	return a
}

func (a *Root) AsideMaxWidth(value interface{}) *Root {
	a.Set("asideMaxWidth", value)
	return a
}

func (a *Root) AsideMinWidth(value interface{}) *Root {
	a.Set("asideMinWidth", value)
	return a
}

func (a *Root) AsidePosition(value interface{}) *Root {
	a.Set("asidePosition", value)
	return a
}

func (a *Root) AsideResizor(value interface{}) *Root {
	a.Set("asideResizor", value)
	return a
}

func (a *Root) AsideSticky(value interface{}) *Root {
	a.Set("asideSticky", value)
	return a
}

func (a *Root) Body(value interface{}) *Root {
	a.Set("body", value)
	return a
}

func (a *Root) BodyClassName(value interface{}) *Root {
	a.Set("bodyClassName", value)
	return a
}

func (a *Root) ClassName(value interface{}) *Root {
	a.Set("className", value)
	return a
}

func (a *Root) Css(value interface{}) *Root {
	a.Set("css", value)
	return a
}

func (a *Root) CssVars(value interface{}) *Root {
	a.Set("cssVars", value)
	return a
}

func (a *Root) Data(value interface{}) *Root {
	a.Set("data", value)
	return a
}

func (a *Root) Definitions(value interface{}) *Root {
	a.Set("definitions", value)
	return a
}

func (a *Root) Disabled(value interface{}) *Root {
	a.Set("disabled", value)
	return a
}

func (a *Root) DisabledOn(value interface{}) *Root {
	a.Set("disabledOn", value)
	return a
}

func (a *Root) EditorSetting(value interface{}) *Root {
	a.Set("editorSetting", value)
	return a
}

func (a *Root) HeaderClassName(value interface{}) *Root {
	a.Set("headerClassName", value)
	return a
}

func (a *Root) Hidden(value interface{}) *Root {
	a.Set("hidden", value)
	return a
}

func (a *Root) HiddenOn(value interface{}) *Root {
	a.Set("hiddenOn", value)
	return a
}

func (a *Root) Id(value interface{}) *Root {
	a.Set("id", value)
	return a
}

func (a *Root) InitApi(value interface{}) *Root {
	a.Set("initApi", value)
	return a
}

func (a *Root) InitFetch(value interface{}) *Root {
	a.Set("initFetch", value)
	return a
}

func (a *Root) InitFetchOn(value interface{}) *Root {
	a.Set("initFetchOn", value)
	return a
}

func (a *Root) Interval(value interface{}) *Root {
	a.Set("interval", value)
	return a
}

func (a *Root) LoadingConfig(value interface{}) *Root {
	a.Set("loadingConfig", value)
	return a
}

func (a *Root) Messages(value interface{}) *Root {
	a.Set("messages", value)
	return a
}

func (a *Root) MobileCSS(value interface{}) *Root {
	a.Set("mobileCSS", value)
	return a
}

func (a *Root) Name(value interface{}) *Root {
	a.Set("name", value)
	return a
}

func (a *Root) OnEvent(value interface{}) *Root {
	a.Set("onEvent", value)
	return a
}

func (a *Root) PullRefresh(value interface{}) *Root {
	a.Set("pullRefresh", value)
	return a
}

func (a *Root) Regions(value interface{}) *Root {
	a.Set("regions", value)
	return a
}

func (a *Root) Remark(value interface{}) *Root {
	a.Set("remark", value)
	return a
}

func (a *Root) ShowErrorMsg(value interface{}) *Root {
	a.Set("showErrorMsg", value)
	return a
}

func (a *Root) SilentPolling(value interface{}) *Root {
	a.Set("silentPolling", value)
	return a
}

func (a *Root) Static(value interface{}) *Root {
	a.Set("static", value)
	return a
}

func (a *Root) StaticClassName(value interface{}) *Root {
	a.Set("staticClassName", value)
	return a
}

func (a *Root) StaticInputClassName(value interface{}) *Root {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Root) StaticLabelClassName(value interface{}) *Root {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Root) StaticOn(value interface{}) *Root {
	a.Set("staticOn", value)
	return a
}

func (a *Root) StaticPlaceholder(value interface{}) *Root {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Root) StaticSchema(value interface{}) *Root {
	a.Set("staticSchema", value)
	return a
}

func (a *Root) StopAutoRefreshWhen(value interface{}) *Root {
	a.Set("stopAutoRefreshWhen", value)
	return a
}

func (a *Root) Style(value interface{}) *Root {
	a.Set("style", value)
	return a
}

func (a *Root) SubTitle(value interface{}) *Root {
	a.Set("subTitle", value)
	return a
}

func (a *Root) TestIdBuilder(value interface{}) *Root {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Root) Testid(value interface{}) *Root {
	a.Set("testid", value)
	return a
}

func (a *Root) Title(value interface{}) *Root {
	a.Set("title", value)
	return a
}

func (a *Root) Toolbar(value interface{}) *Root {
	a.Set("toolbar", value)
	return a
}

func (a *Root) ToolbarClassName(value interface{}) *Root {
	a.Set("toolbarClassName", value)
	return a
}

func (a *Root) Type(value interface{}) *Root {
	a.Set("type", value)
	return a
}

func (a *Root) UseMobileUI(value interface{}) *Root {
	a.Set("useMobileUI", value)
	return a
}

func (a *Root) Visible(value interface{}) *Root {
	a.Set("visible", value)
	return a
}

func (a *Root) VisibleOn(value interface{}) *Root {
	a.Set("visibleOn", value)
	return a
}

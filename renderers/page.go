package renderers


/**
 * amis Page 渲染器。详情请见：https://aisuda.bce.baidu.com/amis/zh-CN/components/page 一个页面只允许有一个 Page 渲染器。

 * @author wcz0
 * @version 6.2.2
 */
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
/**
 * asideResizor
 */
func (a *Page) AsideResizor(value interface{}) *Page {
    a.Set("asideResizor", value)
    return a
}

/**
 * asideSticky
 */
func (a *Page) AsideSticky(value interface{}) *Page {
    a.Set("asideSticky", value)
    return a
}

/**
 * cssVars
 */
func (a *Page) CssVars(value interface{}) *Page {
    a.Set("cssVars", value)
    return a
}

/**
 * id
 */
func (a *Page) Id(value interface{}) *Page {
    a.Set("id", value)
    return a
}

/**
 * static
 */
func (a *Page) Static(value interface{}) *Page {
    a.Set("static", value)
    return a
}

/**
 * title
 */
func (a *Page) Title(value interface{}) *Page {
    a.Set("title", value)
    return a
}

/**
 * asidePosition
 */
func (a *Page) AsidePosition(value interface{}) *Page {
    a.Set("asidePosition", value)
    return a
}

/**
 * css
 */
func (a *Page) Css(value interface{}) *Page {
    a.Set("css", value)
    return a
}

/**
 * pullRefresh
 */
func (a *Page) PullRefresh(value interface{}) *Page {
    a.Set("pullRefresh", value)
    return a
}

/**
 * loadingConfig
 */
func (a *Page) LoadingConfig(value interface{}) *Page {
    a.Set("loadingConfig", value)
    return a
}

/**
 * staticSchema
 */
func (a *Page) StaticSchema(value interface{}) *Page {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *Page) Type(value interface{}) *Page {
    a.Set("type", value)
    return a
}

/**
 * mobileCSS
 */
func (a *Page) MobileCSS(value interface{}) *Page {
    a.Set("mobileCSS", value)
    return a
}

/**
 * headerClassName
 */
func (a *Page) HeaderClassName(value interface{}) *Page {
    a.Set("headerClassName", value)
    return a
}

/**
 * initFetchOn
 */
func (a *Page) InitFetchOn(value interface{}) *Page {
    a.Set("initFetchOn", value)
    return a
}

/**
 * toolbarClassName
 */
func (a *Page) ToolbarClassName(value interface{}) *Page {
    a.Set("toolbarClassName", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Page) StaticPlaceholder(value interface{}) *Page {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Page) StaticLabelClassName(value interface{}) *Page {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Page) StaticInputClassName(value interface{}) *Page {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * testid
 */
func (a *Page) Testid(value interface{}) *Page {
    a.Set("testid", value)
    return a
}

/**
 * initApi
 */
func (a *Page) InitApi(value interface{}) *Page {
    a.Set("initApi", value)
    return a
}

/**
 * stopAutoRefreshWhen
 */
func (a *Page) StopAutoRefreshWhen(value interface{}) *Page {
    a.Set("stopAutoRefreshWhen", value)
    return a
}

/**
 * disabled
 */
func (a *Page) Disabled(value interface{}) *Page {
    a.Set("disabled", value)
    return a
}

/**
 * disabledOn
 */
func (a *Page) DisabledOn(value interface{}) *Page {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticOn
 */
func (a *Page) StaticOn(value interface{}) *Page {
    a.Set("staticOn", value)
    return a
}

/**
 * editorSetting
 */
func (a *Page) EditorSetting(value interface{}) *Page {
    a.Set("editorSetting", value)
    return a
}

/**
 * remark
 */
func (a *Page) Remark(value interface{}) *Page {
    a.Set("remark", value)
    return a
}

/**
 * name
 */
func (a *Page) Name(value interface{}) *Page {
    a.Set("name", value)
    return a
}

/**
 * toolbar
 */
func (a *Page) Toolbar(value interface{}) *Page {
    a.Set("toolbar", value)
    return a
}

/**
 * definitions
 */
func (a *Page) Definitions(value interface{}) *Page {
    a.Set("definitions", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Page) HiddenOn(value interface{}) *Page {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visible
 */
func (a *Page) Visible(value interface{}) *Page {
    a.Set("visible", value)
    return a
}

/**
 * visibleOn
 */
func (a *Page) VisibleOn(value interface{}) *Page {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *Page) StaticClassName(value interface{}) *Page {
    a.Set("staticClassName", value)
    return a
}

/**
 * style
 */
func (a *Page) Style(value interface{}) *Page {
    a.Set("style", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Page) UseMobileUI(value interface{}) *Page {
    a.Set("useMobileUI", value)
    return a
}

/**
 * subTitle
 */
func (a *Page) SubTitle(value interface{}) *Page {
    a.Set("subTitle", value)
    return a
}

/**
 * onEvent
 */
func (a *Page) OnEvent(value interface{}) *Page {
    a.Set("onEvent", value)
    return a
}

/**
 * bodyClassName
 */
func (a *Page) BodyClassName(value interface{}) *Page {
    a.Set("bodyClassName", value)
    return a
}

/**
 * asideMaxWidth
 */
func (a *Page) AsideMaxWidth(value interface{}) *Page {
    a.Set("asideMaxWidth", value)
    return a
}

/**
 * asideClassName
 */
func (a *Page) AsideClassName(value interface{}) *Page {
    a.Set("asideClassName", value)
    return a
}

/**
 * initFetch
 */
func (a *Page) InitFetch(value interface{}) *Page {
    a.Set("initFetch", value)
    return a
}

/**
 * interval
 */
func (a *Page) Interval(value interface{}) *Page {
    a.Set("interval", value)
    return a
}

/**
 * silentPolling
 */
func (a *Page) SilentPolling(value interface{}) *Page {
    a.Set("silentPolling", value)
    return a
}

/**
 * regions
 */
func (a *Page) Regions(value interface{}) *Page {
    a.Set("regions", value)
    return a
}

/**
 * className
 */
func (a *Page) ClassName(value interface{}) *Page {
    a.Set("className", value)
    return a
}

/**
 * hidden
 */
func (a *Page) Hidden(value interface{}) *Page {
    a.Set("hidden", value)
    return a
}

/**
 * body
 */
func (a *Page) Body(value interface{}) *Page {
    a.Set("body", value)
    return a
}

/**
 * asideMinWidth
 */
func (a *Page) AsideMinWidth(value interface{}) *Page {
    a.Set("asideMinWidth", value)
    return a
}

/**
 * data
 */
func (a *Page) Data(value interface{}) *Page {
    a.Set("data", value)
    return a
}

/**
 * messages
 */
func (a *Page) Messages(value interface{}) *Page {
    a.Set("messages", value)
    return a
}

/**
 * showErrorMsg
 */
func (a *Page) ShowErrorMsg(value interface{}) *Page {
    a.Set("showErrorMsg", value)
    return a
}

/**
 * aside
 */
func (a *Page) Aside(value interface{}) *Page {
    a.Set("aside", value)
    return a
}

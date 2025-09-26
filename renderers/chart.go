package renderers


/**
 * Chart 图表渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/chart

 * @author wcz0
 * @version 6.2.2
 */
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
/**
 * initFetch
 */
func (a *Chart) InitFetch(value interface{}) *Chart {
    a.Set("initFetch", value)
    return a
}

/**
 * source
 */
func (a *Chart) Source(value interface{}) *Chart {
    a.Set("source", value)
    return a
}

/**
 * replaceChartOption
 */
func (a *Chart) ReplaceChartOption(value interface{}) *Chart {
    a.Set("replaceChartOption", value)
    return a
}

/**
 * disabledOn
 */
func (a *Chart) DisabledOn(value interface{}) *Chart {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Chart) StaticInputClassName(value interface{}) *Chart {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * testid
 */
func (a *Chart) Testid(value interface{}) *Chart {
    a.Set("testid", value)
    return a
}

/**
 * chartTheme
 */
func (a *Chart) ChartTheme(value interface{}) *Chart {
    a.Set("chartTheme", value)
    return a
}

/**
 * api
 */
func (a *Chart) Api(value interface{}) *Chart {
    a.Set("api", value)
    return a
}

/**
 * config
 */
func (a *Chart) Config(value interface{}) *Chart {
    a.Set("config", value)
    return a
}

/**
 * interval
 */
func (a *Chart) Interval(value interface{}) *Chart {
    a.Set("interval", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Chart) HiddenOn(value interface{}) *Chart {
    a.Set("hiddenOn", value)
    return a
}

/**
 * onEvent
 */
func (a *Chart) OnEvent(value interface{}) *Chart {
    a.Set("onEvent", value)
    return a
}

/**
 * staticClassName
 */
func (a *Chart) StaticClassName(value interface{}) *Chart {
    a.Set("staticClassName", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Chart) UseMobileUI(value interface{}) *Chart {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Chart) Type(value interface{}) *Chart {
    a.Set("type", value)
    return a
}

/**
 * height
 */
func (a *Chart) Height(value interface{}) *Chart {
    a.Set("height", value)
    return a
}

/**
 * name
 */
func (a *Chart) Name(value interface{}) *Chart {
    a.Set("name", value)
    return a
}

/**
 * dataFilter
 */
func (a *Chart) DataFilter(value interface{}) *Chart {
    a.Set("dataFilter", value)
    return a
}

/**
 * disableDataMapping
 */
func (a *Chart) DisableDataMapping(value interface{}) *Chart {
    a.Set("disableDataMapping", value)
    return a
}

/**
 * mapURL
 */
func (a *Chart) MapURL(value interface{}) *Chart {
    a.Set("mapURL", value)
    return a
}

/**
 * visible
 */
func (a *Chart) Visible(value interface{}) *Chart {
    a.Set("visible", value)
    return a
}

/**
 * visibleOn
 */
func (a *Chart) VisibleOn(value interface{}) *Chart {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Chart) StaticPlaceholder(value interface{}) *Chart {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Chart) StaticLabelClassName(value interface{}) *Chart {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *Chart) EditorSetting(value interface{}) *Chart {
    a.Set("editorSetting", value)
    return a
}

/**
 * mapName
 */
func (a *Chart) MapName(value interface{}) *Chart {
    a.Set("mapName", value)
    return a
}

/**
 * static
 */
func (a *Chart) Static(value interface{}) *Chart {
    a.Set("static", value)
    return a
}

/**
 * initFetchOn
 */
func (a *Chart) InitFetchOn(value interface{}) *Chart {
    a.Set("initFetchOn", value)
    return a
}

/**
 * trackExpression
 */
func (a *Chart) TrackExpression(value interface{}) *Chart {
    a.Set("trackExpression", value)
    return a
}

/**
 * unMountOnHidden
 */
func (a *Chart) UnMountOnHidden(value interface{}) *Chart {
    a.Set("unMountOnHidden", value)
    return a
}

/**
 * className
 */
func (a *Chart) ClassName(value interface{}) *Chart {
    a.Set("className", value)
    return a
}

/**
 * id
 */
func (a *Chart) Id(value interface{}) *Chart {
    a.Set("id", value)
    return a
}

/**
 * hidden
 */
func (a *Chart) Hidden(value interface{}) *Chart {
    a.Set("hidden", value)
    return a
}

/**
 * staticOn
 */
func (a *Chart) StaticOn(value interface{}) *Chart {
    a.Set("staticOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *Chart) StaticSchema(value interface{}) *Chart {
    a.Set("staticSchema", value)
    return a
}

/**
 * width
 */
func (a *Chart) Width(value interface{}) *Chart {
    a.Set("width", value)
    return a
}

/**
 * clickAction
 */
func (a *Chart) ClickAction(value interface{}) *Chart {
    a.Set("clickAction", value)
    return a
}

/**
 * loadBaiduMap
 */
func (a *Chart) LoadBaiduMap(value interface{}) *Chart {
    a.Set("loadBaiduMap", value)
    return a
}

/**
 * disabled
 */
func (a *Chart) Disabled(value interface{}) *Chart {
    a.Set("disabled", value)
    return a
}

/**
 * style
 */
func (a *Chart) Style(value interface{}) *Chart {
    a.Set("style", value)
    return a
}

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
 * 是否禁用表达式
 */
func (a *Chart) DisabledOn(value interface{}) *Chart {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Chart) Hidden(value interface{}) *Chart {
    a.Set("hidden", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Chart) OnEvent(value interface{}) *Chart {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Chart) StaticOn(value interface{}) *Chart {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Chart) StaticClassName(value interface{}) *Chart {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Chart) StaticInputClassName(value interface{}) *Chart {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *Chart) StaticSchema(value interface{}) *Chart {
    a.Set("staticSchema", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Chart) ClassName(value interface{}) *Chart {
    a.Set("className", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Chart) StaticPlaceholder(value interface{}) *Chart {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Chart) EditorSetting(value interface{}) *Chart {
    a.Set("editorSetting", value)
    return a
}

/**
 * Chart 主题配置
 */
func (a *Chart) ChartTheme(value interface{}) *Chart {
    a.Set("chartTheme", value)
    return a
}

/**
 * 获取 geo json 文件的地址
 */
func (a *Chart) MapURL(value interface{}) *Chart {
    a.Set("mapURL", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Chart) HiddenOn(value interface{}) *Chart {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *Chart) Visible(value interface{}) *Chart {
    a.Set("visible", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Chart) UseMobileUI(value interface{}) *Chart {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否初始加载用表达式来配置
 */
func (a *Chart) InitFetchOn(value interface{}) *Chart {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 宽度设置
 */
func (a *Chart) Width(value interface{}) *Chart {
    a.Set("width", value)
    return a
}

/**
 */
func (a *Chart) Source(value interface{}) *Chart {
    a.Set("source", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Chart) Static(value interface{}) *Chart {
    a.Set("static", value)
    return a
}

/**
 * style样式
 */
func (a *Chart) Style(value interface{}) *Chart {
    a.Set("style", value)
    return a
}

/**
 */
func (a *Chart) Testid(value interface{}) *Chart {
    a.Set("testid", value)
    return a
}

/**
 * 是否初始加载。
 */
func (a *Chart) InitFetch(value interface{}) *Chart {
    a.Set("initFetch", value)
    return a
}

/**
 * 默认配置时追加的，如果更新配置想完全替换配置请配置为 true.
 */
func (a *Chart) ReplaceChartOption(value interface{}) *Chart {
    a.Set("replaceChartOption", value)
    return a
}

/**
 * 不可见的时候隐藏
 */
func (a *Chart) UnMountOnHidden(value interface{}) *Chart {
    a.Set("unMountOnHidden", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Chart) Id(value interface{}) *Chart {
    a.Set("id", value)
    return a
}

/**
 */
func (a *Chart) TestIdBuilder(value interface{}) *Chart {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 默认开启 Config 中的数据映射，如果想关闭，请开启此功能。
 */
func (a *Chart) DisableDataMapping(value interface{}) *Chart {
    a.Set("disableDataMapping", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Chart) Disabled(value interface{}) *Chart {
    a.Set("disabled", value)
    return a
}

/**
 * 指定为 chart 类型
 */
func (a *Chart) Type(value interface{}) *Chart {
    a.Set("type", value)
    return a
}

/**
 * 图表配置接口
 */
func (a *Chart) Api(value interface{}) *Chart {
    a.Set("api", value)
    return a
}

/**
 * 跟踪表达式，如果这个表达式的运行结果发生变化了，则会更新 Echart，当 config 中用了数据映射时有用。
 */
func (a *Chart) TrackExpression(value interface{}) *Chart {
    a.Set("trackExpression", value)
    return a
}

/**
 * 高度设置
 */
func (a *Chart) Height(value interface{}) *Chart {
    a.Set("height", value)
    return a
}

/**
 * 加载百度地图
 */
func (a *Chart) LoadBaiduMap(value interface{}) *Chart {
    a.Set("loadBaiduMap", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Chart) StaticLabelClassName(value interface{}) *Chart {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 配置echart的config，支持数据映射。如果用了数据映射，为了同步更新，请设置 trackExpression
 */
func (a *Chart) Config(value interface{}) *Chart {
    a.Set("config", value)
    return a
}

/**
 */
func (a *Chart) Name(value interface{}) *Chart {
    a.Set("name", value)
    return a
}

/**
 * 地图名称
 */
func (a *Chart) MapName(value interface{}) *Chart {
    a.Set("mapName", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Chart) VisibleOn(value interface{}) *Chart {
    a.Set("visibleOn", value)
    return a
}

/**
 * 刷新时间
 */
func (a *Chart) Interval(value interface{}) *Chart {
    a.Set("interval", value)
    return a
}

/**
 */
func (a *Chart) DataFilter(value interface{}) *Chart {
    a.Set("dataFilter", value)
    return a
}

/**
 * 点击行为配置，可以用来满足下钻操作等。
 */
func (a *Chart) ClickAction(value interface{}) *Chart {
    a.Set("clickAction", value)
    return a
}

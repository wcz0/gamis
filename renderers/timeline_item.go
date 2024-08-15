package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
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
/**
 * 是否隐藏
 */
func (a *TimelineItem) Hidden(value interface{}) *TimelineItem {
    a.Set("hidden", value)
    return a
}

/**
 * 组件样式
 */
func (a *TimelineItem) Style(value interface{}) *TimelineItem {
    a.Set("style", value)
    return a
}

/**
 * 时间点圆圈颜色
 */
func (a *TimelineItem) Color(value interface{}) *TimelineItem {
    a.Set("color", value)
    return a
}

/**
 * 时间点
 */
func (a *TimelineItem) Time(value interface{}) *TimelineItem {
    a.Set("time", value)
    return a
}

/**
 * 时间节点标题
 */
func (a *TimelineItem) Title(value interface{}) *TimelineItem {
    a.Set("title", value)
    return a
}

/**
 * 节点时间的CSS类名（优先级高于统一配置的timeClassName）
 */
func (a *TimelineItem) Timeclassname(value interface{}) *TimelineItem {
    a.Set("timeClassName", value)
    return a
}

/**
 */
func (a *TimelineItem) Testid(value interface{}) *TimelineItem {
    a.Set("testid", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *TimelineItem) Visibleon(value interface{}) *TimelineItem {
    a.Set("visibleOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *TimelineItem) Id(value interface{}) *TimelineItem {
    a.Set("id", value)
    return a
}

/**
 */
func (a *TimelineItem) Testidbuilder(value interface{}) *TimelineItem {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *TimelineItem) Classname(value interface{}) *TimelineItem {
    a.Set("className", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *TimelineItem) Editorsetting(value interface{}) *TimelineItem {
    a.Set("editorSetting", value)
    return a
}

/**
 * 节点标题的CSS类名（优先级高于统一配置的titleClassName）
 */
func (a *TimelineItem) Titleclassname(value interface{}) *TimelineItem {
    a.Set("titleClassName", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *TimelineItem) Staticclassname(value interface{}) *TimelineItem {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *TimelineItem) Staticlabelclassname(value interface{}) *TimelineItem {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *TimelineItem) Staticinputclassname(value interface{}) *TimelineItem {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * detail展开时文案
 */
func (a *TimelineItem) Detailexpandedtext(value interface{}) *TimelineItem {
    a.Set("detailExpandedText", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *TimelineItem) Disabledon(value interface{}) *TimelineItem {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *TimelineItem) Hiddenon(value interface{}) *TimelineItem {
    a.Set("hiddenOn", value)
    return a
}

/**
 */
func (a *TimelineItem) Staticschema(value interface{}) *TimelineItem {
    a.Set("staticSchema", value)
    return a
}

/**
 * detail折叠时文案
 */
func (a *TimelineItem) Detailcollapsedtext(value interface{}) *TimelineItem {
    a.Set("detailCollapsedText", value)
    return a
}

/**
 * 节点详情的CSS类名（优先级高于统一配置的detailClassName）
 */
func (a *TimelineItem) Detailclassname(value interface{}) *TimelineItem {
    a.Set("detailClassName", value)
    return a
}

/**
 * 是否显示
 */
func (a *TimelineItem) Visible(value interface{}) *TimelineItem {
    a.Set("visible", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *TimelineItem) Staticon(value interface{}) *TimelineItem {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *TimelineItem) Staticplaceholder(value interface{}) *TimelineItem {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 详细内容
 */
func (a *TimelineItem) Detail(value interface{}) *TimelineItem {
    a.Set("detail", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *TimelineItem) Usemobileui(value interface{}) *TimelineItem {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 图标
 */
func (a *TimelineItem) Icon(value interface{}) *TimelineItem {
    a.Set("icon", value)
    return a
}

/**
 * 图标的CSS类名
 */
func (a *TimelineItem) Iconclassname(value interface{}) *TimelineItem {
    a.Set("iconClassName", value)
    return a
}

/**
 * 是否禁用
 */
func (a *TimelineItem) Disabled(value interface{}) *TimelineItem {
    a.Set("disabled", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *TimelineItem) Onevent(value interface{}) *TimelineItem {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *TimelineItem) Static(value interface{}) *TimelineItem {
    a.Set("static", value)
    return a
}

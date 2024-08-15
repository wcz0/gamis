package renderers


/**
 * 状态展示控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/status

 * @author wcz0
 * @version 6.2.2
 */
type Status struct {
	*BaseRenderer
}

func NewStatus() *Status {
    a := &Status{
        BaseRenderer: NewBaseRenderer(),
    }

func (a *Status) Set(name string, value interface{}) *Status {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}

    a.Set("type", "status")
    return a
}

/**
 * 是否显示
 */
func (a *Status) Visible(value interface{}) *Status {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Status) Staticclassname(value interface{}) *Status {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *Status) Testidbuilder(value interface{}) *Status {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 指定为状态展示控件
 */
func (a *Status) Type(value interface{}) *Status {
    a.Set("type", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Status) Static(value interface{}) *Status {
    a.Set("static", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Status) Hiddenon(value interface{}) *Status {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Status) Staticlabelclassname(value interface{}) *Status {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Status) Disabled(value interface{}) *Status {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Status) Visibleon(value interface{}) *Status {
    a.Set("visibleOn", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Status) Classname(value interface{}) *Status {
    a.Set("className", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Status) Id(value interface{}) *Status {
    a.Set("id", value)
    return a
}

/**
 * 占位符
 */
func (a *Status) Placeholder(value interface{}) *Status {
    a.Set("placeholder", value)
    return a
}

/**
 * 状态图标映射关系
 */
func (a *Status) Map(value interface{}) *Status {
    a.Set("map", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Status) Disabledon(value interface{}) *Status {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Status) Hidden(value interface{}) *Status {
    a.Set("hidden", value)
    return a
}

/**
 * 组件样式
 */
func (a *Status) Style(value interface{}) *Status {
    a.Set("style", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Status) Usemobileui(value interface{}) *Status {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Status) Testid(value interface{}) *Status {
    a.Set("testid", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Status) Editorsetting(value interface{}) *Status {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Status) Staticon(value interface{}) *Status {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Status) Staticplaceholder(value interface{}) *Status {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Status) Staticinputclassname(value interface{}) *Status {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *Status) Staticschema(value interface{}) *Status {
    a.Set("staticSchema", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Status) Onevent(value interface{}) *Status {
    a.Set("onEvent", value)
    return a
}

/**
 * 文字映射关系
 */
func (a *Status) Labelmap(value interface{}) *Status {
    a.Set("labelMap", value)
    return a
}

/**
 * 新版配置映射源的字段 可以兼容新版icon并且配置颜色 2.8.0 新增
 */
func (a *Status) Source(value interface{}) *Status {
    a.Set("source", value)
    return a
}

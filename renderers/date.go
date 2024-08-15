package renderers


/**
 * Date 展示渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/date

 * @author wcz0
 * @version 6.2.2
 */
type Date struct {
	*BaseRenderer
}

func NewDate() *Date {
    a := &Date{
        BaseRenderer: NewBaseRenderer(),
    }

func (a *Date) Set(name string, value interface{}) *Date {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}

    a.Set("type", "date")
    return a
}

/**
 * 组件样式
 */
func (a *Date) Style(value interface{}) *Date {
    a.Set("style", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Date) Disabled(value interface{}) *Date {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Date) Visibleon(value interface{}) *Date {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Date) Staticclassname(value interface{}) *Date {
    a.Set("staticClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Date) Usemobileui(value interface{}) *Date {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Date) Classname(value interface{}) *Date {
    a.Set("className", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Date) Staticplaceholder(value interface{}) *Date {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Date) Static(value interface{}) *Date {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Date) Staticon(value interface{}) *Date {
    a.Set("staticOn", value)
    return a
}

/**
 * 时区
 */
func (a *Date) Displaytimezone(value interface{}) *Date {
    a.Set("displayTimeZone", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Date) Id(value interface{}) *Date {
    a.Set("id", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Date) Onevent(value interface{}) *Date {
    a.Set("onEvent", value)
    return a
}

/**
 * 展示的时间格式，参考 moment 中的格式说明。
 */
func (a *Date) Format(value interface{}) *Date {
    a.Set("format", value)
    return a
}

/**
 * 更新频率， 默认为1分钟
 */
func (a *Date) Updatefrequency(value interface{}) *Date {
    a.Set("updateFrequency", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Date) Disabledon(value interface{}) *Date {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *Date) Visible(value interface{}) *Date {
    a.Set("visible", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Date) Hidden(value interface{}) *Date {
    a.Set("hidden", value)
    return a
}

/**
 * 指定为日期展示类型
 * 可选值: date | datetime | time | static-date | static-datetime | static-time
 */
func (a *Date) Type(value interface{}) *Date {
    a.Set("type", value)
    return a
}

/**
 * 显示成相对时间，比如1分钟前
 */
func (a *Date) Fromnow(value interface{}) *Date {
    a.Set("fromNow", value)
    return a
}

/**
 * 占位符
 */
func (a *Date) Placeholder(value interface{}) *Date {
    a.Set("placeholder", value)
    return a
}

/**
 * 值的时间格式，参考 moment 中的格式说明。
 */
func (a *Date) Valueformat(value interface{}) *Date {
    a.Set("valueFormat", value)
    return a
}

/**
 */
func (a *Date) Staticschema(value interface{}) *Date {
    a.Set("staticSchema", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Date) Editorsetting(value interface{}) *Date {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *Date) Testidbuilder(value interface{}) *Date {
    a.Set("testIdBuilder", value)
    return a
}

/**
 */
func (a *Date) Testid(value interface{}) *Date {
    a.Set("testid", value)
    return a
}

/**
 * 展示的时间格式，参考 moment 中的格式说明。（新：同format）
 */
func (a *Date) Displayformat(value interface{}) *Date {
    a.Set("displayFormat", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Date) Hiddenon(value interface{}) *Date {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Date) Staticlabelclassname(value interface{}) *Date {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Date) Staticinputclassname(value interface{}) *Date {
    a.Set("staticInputClassName", value)
    return a
}

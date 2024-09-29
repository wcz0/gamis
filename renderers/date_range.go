package renderers


/**
 * DateRange 展示渲染器。

 * @author wcz0
 * @version 6.2.2
 */
type DateRange struct {
	*BaseRenderer
}

func NewDateRange() *DateRange {
    a := &DateRange{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "date-range")
    return a
}


func (a *DateRange) Set(name string, value interface{}) *DateRange {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *DateRange) Id(value interface{}) *DateRange {
    a.Set("id", value)
    return a
}

/**
 */
func (a *DateRange) StaticSchema(value interface{}) *DateRange {
    a.Set("staticSchema", value)
    return a
}

/**
 * 展示的时间格式，参考 moment 中的格式说明。
 */
func (a *DateRange) Format(value interface{}) *DateRange {
    a.Set("format", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *DateRange) StaticOn(value interface{}) *DateRange {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *DateRange) StaticClassName(value interface{}) *DateRange {
    a.Set("staticClassName", value)
    return a
}

/**
 * 值的时间格式，参考 moment 中的格式说明。
 */
func (a *DateRange) ValueFormat(value interface{}) *DateRange {
    a.Set("valueFormat", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *DateRange) ClassName(value interface{}) *DateRange {
    a.Set("className", value)
    return a
}

/**
 * 组件样式
 */
func (a *DateRange) Style(value interface{}) *DateRange {
    a.Set("style", value)
    return a
}

/**
 * 连接符
 */
func (a *DateRange) Connector(value interface{}) *DateRange {
    a.Set("connector", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *DateRange) DisabledOn(value interface{}) *DateRange {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *DateRange) Hidden(value interface{}) *DateRange {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示
 */
func (a *DateRange) Visible(value interface{}) *DateRange {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *DateRange) StaticLabelClassName(value interface{}) *DateRange {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 是否禁用
 */
func (a *DateRange) Disabled(value interface{}) *DateRange {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *DateRange) HiddenOn(value interface{}) *DateRange {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *DateRange) VisibleOn(value interface{}) *DateRange {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *DateRange) StaticInputClassName(value interface{}) *DateRange {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *DateRange) TestIdBuilder(value interface{}) *DateRange {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *DateRange) EditorSetting(value interface{}) *DateRange {
    a.Set("editorSetting", value)
    return a
}

/**
 * 指定为日期展示类型
 */
func (a *DateRange) Type(value interface{}) *DateRange {
    a.Set("type", value)
    return a
}

/**
 * 展示的时间格式，参考 moment 中的格式说明。（新：同format）
 */
func (a *DateRange) DisplayFormat(value interface{}) *DateRange {
    a.Set("displayFormat", value)
    return a
}

/**
 * 分割符
 */
func (a *DateRange) Delimiter(value interface{}) *DateRange {
    a.Set("delimiter", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *DateRange) OnEvent(value interface{}) *DateRange {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *DateRange) Static(value interface{}) *DateRange {
    a.Set("static", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *DateRange) StaticPlaceholder(value interface{}) *DateRange {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *DateRange) UseMobileUI(value interface{}) *DateRange {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *DateRange) Testid(value interface{}) *DateRange {
    a.Set("testid", value)
    return a
}

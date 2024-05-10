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
    a.Set("type", "status")
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
 * 是否禁用表达式
 */
func (a *Status) DisabledOn(value interface{}) *Status {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Status) HiddenOn(value interface{}) *Status {
    a.Set("hiddenOn", value)
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
 * 组件唯一 id，主要用于日志采集
 */
func (a *Status) Id(value interface{}) *Status {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Status) StaticOn(value interface{}) *Status {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Status) StaticLabelClassName(value interface{}) *Status {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *Status) StaticSchema(value interface{}) *Status {
    a.Set("staticSchema", value)
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
 * 占位符
 */
func (a *Status) Placeholder(value interface{}) *Status {
    a.Set("placeholder", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Status) ClassName(value interface{}) *Status {
    a.Set("className", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Status) VisibleOn(value interface{}) *Status {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Status) StaticClassName(value interface{}) *Status {
    a.Set("staticClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Status) EditorSetting(value interface{}) *Status {
    a.Set("editorSetting", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Status) UseMobileUI(value interface{}) *Status {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Status) OnEvent(value interface{}) *Status {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Status) StaticPlaceholder(value interface{}) *Status {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 文字映射关系
 */
func (a *Status) LabelMap(value interface{}) *Status {
    a.Set("labelMap", value)
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
 * 是否静态展示
 */
func (a *Status) Static(value interface{}) *Status {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Status) StaticInputClassName(value interface{}) *Status {
    a.Set("staticInputClassName", value)
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
 * 状态图标映射关系
 */
func (a *Status) Map(value interface{}) *Status {
    a.Set("map", value)
    return a
}

/**
 * 新版配置映射源的字段 可以兼容新版icon并且配置颜色 2.8.0 新增
 */
func (a *Status) Source(value interface{}) *Status {
    a.Set("source", value)
    return a
}

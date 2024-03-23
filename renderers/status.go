package renderers


/**
 * 状态展示控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/status
 *

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
 * 是否隐藏
 */
func (a *Status) Hidden(value string) *Status {
    a.Set("hidden", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Status) Id(value string) *Status {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Status) StaticClassName(value string) *Status {
    a.Set("staticClassName", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Status) DisabledOn(value string) *Status {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Status) Static(value string) *Status {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Status) StaticOn(value string) *Status {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Status) StaticLabelClassName(value string) *Status {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 新版配置映射源的字段 可以兼容新版icon并且配置颜色 2.8.0 新增
 */
func (a *Status) Source(value string) *Status {
    a.Set("source", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Status) UseMobileUI(value string) *Status {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Status) Disabled(value string) *Status {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Status) HiddenOn(value string) *Status {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *Status) Visible(value string) *Status {
    a.Set("visible", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Status) VisibleOn(value string) *Status {
    a.Set("visibleOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Status) OnEvent(value string) *Status {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Status) StaticInputClassName(value string) *Status {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *Status) Style(value string) *Status {
    a.Set("style", value)
    return a
}

/**
 * 文字映射关系
 */
func (a *Status) LabelMap(value string) *Status {
    a.Set("labelMap", value)
    return a
}

/**
 * 状态图标映射关系
 */
func (a *Status) Map(value string) *Status {
    a.Set("map", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Status) ClassName(value string) *Status {
    a.Set("className", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Status) StaticPlaceholder(value string) *Status {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *Status) StaticSchema(value string) *Status {
    a.Set("staticSchema", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Status) EditorSetting(value string) *Status {
    a.Set("editorSetting", value)
    return a
}

/**
 * 指定为状态展示控件
 */
func (a *Status) Type(value string) *Status {
    a.Set("type", value)
    return a
}

/**
 * 占位符
 */
func (a *Status) Placeholder(value string) *Status {
    a.Set("placeholder", value)
    return a
}

package renderers


/**
 * 提示渲染器，默认会显示个小图标，鼠标放上来的时候显示配置的内容。
 *

*/
type Remark struct {
	*BaseRenderer
}

func NewRemark() *Remark {
    a := &Remark{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "remark")
    return a
}
/**
 * 提示内容
 */
func (a *Remark) Content(value string) *Remark {
    a.Set("content", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Remark) VisibleOn(value string) *Remark {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Remark) StaticOn(value string) *Remark {
    a.Set("staticOn", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Remark) UseMobileUI(value string) *Remark {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Remark) Icon(value string) *Remark {
    a.Set("icon", value)
    return a
}

/**
 * 触发规则
 */
func (a *Remark) Trigger(value string) *Remark {
    a.Set("trigger", value)
    return a
}

/**
 * 显示位置
 * 可选值: top | right | bottom | left
 */
func (a *Remark) Placement(value string) *Remark {
    a.Set("placement", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Remark) HiddenOn(value string) *Remark {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Remark) Id(value string) *Remark {
    a.Set("id", value)
    return a
}

/**
 */
func (a *Remark) StaticSchema(value string) *Remark {
    a.Set("staticSchema", value)
    return a
}

/**
 * 指定为提示类型
 */
func (a *Remark) Type(value string) *Remark {
    a.Set("type", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Remark) ClassName(value string) *Remark {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Remark) Disabled(value string) *Remark {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Remark) Hidden(value string) *Remark {
    a.Set("hidden", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Remark) StaticPlaceholder(value string) *Remark {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *Remark) TooltipClassName(value string) *Remark {
    a.Set("tooltipClassName", value)
    return a
}

/**
 * icon的形状
 * 可选值: circle | square
 */
func (a *Remark) Shape(value string) *Remark {
    a.Set("shape", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Remark) StaticLabelClassName(value string) *Remark {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Remark) StaticInputClassName(value string) *Remark {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *Remark) Style(value string) *Remark {
    a.Set("style", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Remark) EditorSetting(value string) *Remark {
    a.Set("editorSetting", value)
    return a
}

/**
 * 提示标题
 */
func (a *Remark) Title(value string) *Remark {
    a.Set("title", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Remark) StaticClassName(value string) *Remark {
    a.Set("staticClassName", value)
    return a
}

/**
 * 点击其他内容时是否关闭弹框信息
 */
func (a *Remark) RootClose(value string) *Remark {
    a.Set("rootClose", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Remark) DisabledOn(value string) *Remark {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *Remark) Visible(value string) *Remark {
    a.Set("visible", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Remark) OnEvent(value string) *Remark {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Remark) Static(value string) *Remark {
    a.Set("static", value)
    return a
}

/**
 */
func (a *Remark) Label(value string) *Remark {
    a.Set("label", value)
    return a
}

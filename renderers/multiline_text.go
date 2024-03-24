package renderers


/**
 * MultilineText
 *

*/
type MultilineText struct {
	*BaseRenderer
}

func NewMultilineText() *MultilineText {
    a := &MultilineText{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "multiline-text")
    return a
}
/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *MultilineText) Id(value string) *MultilineText {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *MultilineText) StaticOn(value string) *MultilineText {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *MultilineText) StaticClassName(value string) *MultilineText {
    a.Set("staticClassName", value)
    return a
}

/**
 * 文本
 */
func (a *MultilineText) Text(value string) *MultilineText {
    a.Set("text", value)
    return a
}

/**
 * 最大行数
 */
func (a *MultilineText) MaxRows(value string) *MultilineText {
    a.Set("maxRows", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *MultilineText) ClassName(value string) *MultilineText {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *MultilineText) HiddenOn(value string) *MultilineText {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *MultilineText) OnEvent(value string) *MultilineText {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *MultilineText) StaticPlaceholder(value string) *MultilineText {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *MultilineText) StaticSchema(value string) *MultilineText {
    a.Set("staticSchema", value)
    return a
}

/**
 * 组件样式
 */
func (a *MultilineText) Style(value string) *MultilineText {
    a.Set("style", value)
    return a
}

/**
 * 收起按钮文本
 */
func (a *MultilineText) CollapseButtonText(value string) *MultilineText {
    a.Set("collapseButtonText", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *MultilineText) DisabledOn(value string) *MultilineText {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *MultilineText) Hidden(value string) *MultilineText {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示
 */
func (a *MultilineText) Visible(value string) *MultilineText {
    a.Set("visible", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *MultilineText) Static(value string) *MultilineText {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *MultilineText) StaticInputClassName(value string) *MultilineText {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否禁用
 */
func (a *MultilineText) Disabled(value string) *MultilineText {
    a.Set("disabled", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *MultilineText) EditorSetting(value string) *MultilineText {
    a.Set("editorSetting", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *MultilineText) UseMobileUI(value string) *MultilineText {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *MultilineText) Type(value string) *MultilineText {
    a.Set("type", value)
    return a
}

/**
 * 展开按钮文本
 */
func (a *MultilineText) ExpendButtonText(value string) *MultilineText {
    a.Set("expendButtonText", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *MultilineText) VisibleOn(value string) *MultilineText {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *MultilineText) StaticLabelClassName(value string) *MultilineText {
    a.Set("staticLabelClassName", value)
    return a
}

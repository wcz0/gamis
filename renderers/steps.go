package renderers


/**

*/
type Steps struct {
	*BaseRenderer
}

func NewSteps() *Steps {
    a := &Steps{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "steps")
    return a
}
/**
 * 是否禁用表达式
 */
func (a *Steps) DisabledOn(value string) *Steps {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Steps) VisibleOn(value string) *Steps {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Steps) StaticPlaceholder(value string) *Steps {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *Steps) StaticSchema(value string) *Steps {
    a.Set("staticSchema", value)
    return a
}

/**
 * 展示模式
 * 可选值: horizontal | vertical
 */
func (a *Steps) Mode(value string) *Steps {
    a.Set("mode", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Steps) Static(value string) *Steps {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Steps) StaticOn(value string) *Steps {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Steps) StaticLabelClassName(value string) *Steps {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Steps) EditorSetting(value string) *Steps {
    a.Set("editorSetting", value)
    return a
}

/**
 * 指定当前步骤
 */
func (a *Steps) Value(value string) *Steps {
    a.Set("value", value)
    return a
}

/**
 * 变量映射
 */
func (a *Steps) Name(value string) *Steps {
    a.Set("name", value)
    return a
}

/**
 * 标签放置位置
 * 可选值: horizontal | vertical
 */
func (a *Steps) LabelPlacement(value string) *Steps {
    a.Set("labelPlacement", value)
    return a
}

/**
 * 是否显示
 */
func (a *Steps) Visible(value string) *Steps {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Steps) StaticInputClassName(value string) *Steps {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 步骤
 */
func (a *Steps) Steps(value string) *Steps {
    a.Set("steps", value)
    return a
}

/**
 * API 或 数据映射
 */
func (a *Steps) Source(value string) *Steps {
    a.Set("source", value)
    return a
}

/**
 */
func (a *Steps) Status(value string) *Steps {
    a.Set("status", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Steps) HiddenOn(value string) *Steps {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Steps) OnEvent(value string) *Steps {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Steps) Disabled(value string) *Steps {
    a.Set("disabled", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Steps) Id(value string) *Steps {
    a.Set("id", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Steps) ClassName(value string) *Steps {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Steps) Hidden(value string) *Steps {
    a.Set("hidden", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Steps) UseMobileUI(value string) *Steps {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 指定为 Steps 步骤条渲染器
 */
func (a *Steps) Type(value string) *Steps {
    a.Set("type", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Steps) StaticClassName(value string) *Steps {
    a.Set("staticClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *Steps) Style(value string) *Steps {
    a.Set("style", value)
    return a
}

/**
 * 点状步骤条
 */
func (a *Steps) ProgressDot(value string) *Steps {
    a.Set("progressDot", value)
    return a
}

package renderers


/**

*/
type Step struct {
	*BaseRenderer
}

func NewStep() *Step {
    a := &Step{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}
/**
 * 子标题
 */
func (a *Step) SubTitle(value string) *Step {
    a.Set("subTitle", value)
    return a
}

/**
 * 图标
 */
func (a *Step) Icon(value string) *Step {
    a.Set("icon", value)
    return a
}

/**
 * 描述
 */
func (a *Step) Description(value string) *Step {
    a.Set("description", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Step) Disabled(value string) *Step {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Step) Hidden(value string) *Step {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示
 */
func (a *Step) Visible(value string) *Step {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Step) StaticPlaceholder(value string) *Step {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Step) EditorSetting(value string) *Step {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Step) DisabledOn(value string) *Step {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Step) Static(value string) *Step {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Step) StaticOn(value string) *Step {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *Step) Value(value string) *Step {
    a.Set("value", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Step) StaticLabelClassName(value string) *Step {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Step) StaticInputClassName(value string) *Step {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Step) UseMobileUI(value string) *Step {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Step) ClassName(value string) *Step {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Step) HiddenOn(value string) *Step {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Step) VisibleOn(value string) *Step {
    a.Set("visibleOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Step) Id(value string) *Step {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Step) StaticClassName(value string) *Step {
    a.Set("staticClassName", value)
    return a
}

/**
 * 标题
 */
func (a *Step) Title(value string) *Step {
    a.Set("title", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Step) OnEvent(value string) *Step {
    a.Set("onEvent", value)
    return a
}

/**
 */
func (a *Step) StaticSchema(value string) *Step {
    a.Set("staticSchema", value)
    return a
}

/**
 * 组件样式
 */
func (a *Step) Style(value string) *Step {
    a.Set("style", value)
    return a
}

package renderers


/**

*/
type State struct {
	*BaseRenderer
}

func NewState() *State {
    a := &State{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}
/**
 * 容器 css 类名
 */
func (a *State) ClassName(value string) *State {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用
 */
func (a *State) Disabled(value string) *State {
    a.Set("disabled", value)
    return a
}

/**
 * 显示条件
 */
func (a *State) VisibleOn(value string) *State {
    a.Set("visibleOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *State) Id(value string) *State {
    a.Set("id", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *State) OnEvent(value string) *State {
    a.Set("onEvent", value)
    return a
}

/**
 * 状态标题
 */
func (a *State) Title(value string) *State {
    a.Set("title", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *State) DisabledOn(value string) *State {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *State) Visible(value string) *State {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *State) StaticClassName(value string) *State {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *State) StaticInputClassName(value string) *State {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *State) UseMobileUI(value string) *State {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 内容
 */
func (a *State) Body(value string) *State {
    a.Set("body", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *State) HiddenOn(value string) *State {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *State) Static(value string) *State {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *State) StaticLabelClassName(value string) *State {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *State) StaticSchema(value string) *State {
    a.Set("staticSchema", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *State) EditorSetting(value string) *State {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *State) Hidden(value string) *State {
    a.Set("hidden", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *State) StaticOn(value string) *State {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *State) StaticPlaceholder(value string) *State {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 组件样式
 */
func (a *State) Style(value string) *State {
    a.Set("style", value)
    return a
}

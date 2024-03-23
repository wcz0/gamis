package renderers


/**

*/
type NavItem struct {
	*BaseRenderer
}

func NewNavItem() *NavItem {
    a := &NavItem{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}
/**
 */
func (a *NavItem) DeferApi(value string) *NavItem {
    a.Set("deferApi", value)
    return a
}

/**
 */
func (a *NavItem) Children(value string) *NavItem {
    a.Set("children", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *NavItem) Hidden(value string) *NavItem {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *NavItem) VisibleOn(value string) *NavItem {
    a.Set("visibleOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *NavItem) OnEvent(value string) *NavItem {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *NavItem) StaticLabelClassName(value string) *NavItem {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *NavItem) Disabled(value string) *NavItem {
    a.Set("disabled", value)
    return a
}

/**
 * 组件样式
 */
func (a *NavItem) Style(value string) *NavItem {
    a.Set("style", value)
    return a
}

/**
 */
func (a *NavItem) Active(value string) *NavItem {
    a.Set("active", value)
    return a
}

/**
 */
func (a *NavItem) DisabledTip(value string) *NavItem {
    a.Set("disabledTip", value)
    return a
}

/**
 */
func (a *NavItem) To(value string) *NavItem {
    a.Set("to", value)
    return a
}

/**
 */
func (a *NavItem) Target(value string) *NavItem {
    a.Set("target", value)
    return a
}

/**
 */
func (a *NavItem) Defer(value string) *NavItem {
    a.Set("defer", value)
    return a
}

/**
 */
func (a *NavItem) ClassName(value string) *NavItem {
    a.Set("className", value)
    return a
}

/**
 * 是否显示
 */
func (a *NavItem) Visible(value string) *NavItem {
    a.Set("visible", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *NavItem) Id(value string) *NavItem {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *NavItem) StaticInputClassName(value string) *NavItem {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 文字说明
 */
func (a *NavItem) Label(value string) *NavItem {
    a.Set("label", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *NavItem) EditorSetting(value string) *NavItem {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *NavItem) Key(value string) *NavItem {
    a.Set("key", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *NavItem) DisabledOn(value string) *NavItem {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *NavItem) HiddenOn(value string) *NavItem {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *NavItem) Static(value string) *NavItem {
    a.Set("static", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *NavItem) StaticPlaceholder(value string) *NavItem {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *NavItem) StaticSchema(value string) *NavItem {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *NavItem) Unfolded(value string) *NavItem {
    a.Set("unfolded", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *NavItem) StaticOn(value string) *NavItem {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *NavItem) StaticClassName(value string) *NavItem {
    a.Set("staticClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *NavItem) UseMobileUI(value string) *NavItem {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 图标类名，参考 fontawesome 4。
 */
func (a *NavItem) Icon(value string) *NavItem {
    a.Set("icon", value)
    return a
}

/**
 */
func (a *NavItem) Mode(value string) *NavItem {
    a.Set("mode", value)
    return a
}

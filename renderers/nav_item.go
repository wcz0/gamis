package renderers


/**

 * @author wcz0
 * @version 6.2.2
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
func (a *NavItem) Key(value interface{}) *NavItem {
    a.Set("key", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *NavItem) StaticInputClassName(value interface{}) *NavItem {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *NavItem) UseMobileUI(value interface{}) *NavItem {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 图标类名，参考 fontawesome 4。
 */
func (a *NavItem) Icon(value interface{}) *NavItem {
    a.Set("icon", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *NavItem) EditorSetting(value interface{}) *NavItem {
    a.Set("editorSetting", value)
    return a
}

/**
 * 文字说明
 */
func (a *NavItem) Label(value interface{}) *NavItem {
    a.Set("label", value)
    return a
}

/**
 */
func (a *NavItem) Target(value interface{}) *NavItem {
    a.Set("target", value)
    return a
}

/**
 */
func (a *NavItem) Unfolded(value interface{}) *NavItem {
    a.Set("unfolded", value)
    return a
}

/**
 */
func (a *NavItem) DeferApi(value interface{}) *NavItem {
    a.Set("deferApi", value)
    return a
}

/**
 */
func (a *NavItem) ClassName(value interface{}) *NavItem {
    a.Set("className", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *NavItem) OnEvent(value interface{}) *NavItem {
    a.Set("onEvent", value)
    return a
}

/**
 */
func (a *NavItem) StaticSchema(value interface{}) *NavItem {
    a.Set("staticSchema", value)
    return a
}

/**
 * 是否显示
 */
func (a *NavItem) Visible(value interface{}) *NavItem {
    a.Set("visible", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *NavItem) StaticOn(value interface{}) *NavItem {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *NavItem) StaticPlaceholder(value interface{}) *NavItem {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *NavItem) Active(value interface{}) *NavItem {
    a.Set("active", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *NavItem) DisabledOn(value interface{}) *NavItem {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *NavItem) Hidden(value interface{}) *NavItem {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *NavItem) VisibleOn(value interface{}) *NavItem {
    a.Set("visibleOn", value)
    return a
}

/**
 */
func (a *NavItem) Mode(value interface{}) *NavItem {
    a.Set("mode", value)
    return a
}

/**
 */
func (a *NavItem) DisabledTip(value interface{}) *NavItem {
    a.Set("disabledTip", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *NavItem) Id(value interface{}) *NavItem {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *NavItem) Static(value interface{}) *NavItem {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *NavItem) StaticLabelClassName(value interface{}) *NavItem {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *NavItem) Style(value interface{}) *NavItem {
    a.Set("style", value)
    return a
}

/**
 */
func (a *NavItem) Disabled(value interface{}) *NavItem {
    a.Set("disabled", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *NavItem) StaticClassName(value interface{}) *NavItem {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *NavItem) Children(value interface{}) *NavItem {
    a.Set("children", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *NavItem) HiddenOn(value interface{}) *NavItem {
    a.Set("hiddenOn", value)
    return a
}

/**
 */
func (a *NavItem) To(value interface{}) *NavItem {
    a.Set("to", value)
    return a
}

/**
 */
func (a *NavItem) Defer(value interface{}) *NavItem {
    a.Set("defer", value)
    return a
}

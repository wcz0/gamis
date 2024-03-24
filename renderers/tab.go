package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type Tab struct {
	*BaseRenderer
}

func NewTab() *Tab {
    a := &Tab{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}
/**
 * 是否禁用表达式
 */
func (a *Tab) DisabledOn(value interface{}) *Tab {
    a.Set("disabledOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Tab) EditorSetting(value interface{}) *Tab {
    a.Set("editorSetting", value)
    return a
}

/**
 * 配置子表单项默认的展示方式。
 * 可选值: normal | inline | horizontal
 */
func (a *Tab) Mode(value interface{}) *Tab {
    a.Set("mode", value)
    return a
}

/**
 * 可选值: left | right
 */
func (a *Tab) IconPosition(value interface{}) *Tab {
    a.Set("iconPosition", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Tab) Disabled(value interface{}) *Tab {
    a.Set("disabled", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Tab) StaticClassName(value interface{}) *Tab {
    a.Set("staticClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Tab) UseMobileUI(value interface{}) *Tab {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 内容
 */
func (a *Tab) Tab(value interface{}) *Tab {
    a.Set("tab", value)
    return a
}

/**
 * 内容
 */
func (a *Tab) Body(value interface{}) *Tab {
    a.Set("body", value)
    return a
}

/**
 * 设置以后将跟url的hash对应
 */
func (a *Tab) Hash(value interface{}) *Tab {
    a.Set("hash", value)
    return a
}

/**
 * 按钮图标
 */
func (a *Tab) Icon(value interface{}) *Tab {
    a.Set("icon", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Tab) Hidden(value interface{}) *Tab {
    a.Set("hidden", value)
    return a
}

/**
 * 卡片隐藏就销毁卡片节点。
 */
func (a *Tab) UnmountOnExit(value interface{}) *Tab {
    a.Set("unmountOnExit", value)
    return a
}

/**
 * 是否可关闭，优先级高于 tabs 的 closable
 */
func (a *Tab) Closable(value interface{}) *Tab {
    a.Set("closable", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Tab) HiddenOn(value interface{}) *Tab {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Tab) Id(value interface{}) *Tab {
    a.Set("id", value)
    return a
}

/**
 * 徽标
 */
func (a *Tab) Badge(value interface{}) *Tab {
    a.Set("badge", value)
    return a
}

/**
 * 点开时才加载卡片内容
 */
func (a *Tab) MountOnEnter(value interface{}) *Tab {
    a.Set("mountOnEnter", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Tab) StaticLabelClassName(value interface{}) *Tab {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Tab) VisibleOn(value interface{}) *Tab {
    a.Set("visibleOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Tab) OnEvent(value interface{}) *Tab {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Tab) StaticInputClassName(value interface{}) *Tab {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 设置以后内容每次都会重新渲染
 */
func (a *Tab) Reload(value interface{}) *Tab {
    a.Set("reload", value)
    return a
}

/**
 * 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
 */
func (a *Tab) Horizontal(value interface{}) *Tab {
    a.Set("horizontal", value)
    return a
}

/**
 * 是否显示
 */
func (a *Tab) Visible(value interface{}) *Tab {
    a.Set("visible", value)
    return a
}

/**
 */
func (a *Tab) StaticSchema(value interface{}) *Tab {
    a.Set("staticSchema", value)
    return a
}

/**
 * 组件样式
 */
func (a *Tab) Style(value interface{}) *Tab {
    a.Set("style", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Tab) ClassName(value interface{}) *Tab {
    a.Set("className", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Tab) Static(value interface{}) *Tab {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Tab) StaticOn(value interface{}) *Tab {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Tab) StaticPlaceholder(value interface{}) *Tab {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * Tab 标题
 */
func (a *Tab) Title(value interface{}) *Tab {
    a.Set("title", value)
    return a
}

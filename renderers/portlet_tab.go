package renderers


/**
 * 栏目容器渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/portlet

 * @author wcz0
 * @version 6.2.2
 */
type PortletTab struct {
	*BaseRenderer
}

func NewPortletTab() *PortletTab {
    a := &PortletTab{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}
/**
 * 是否静态展示
 */
func (a *PortletTab) Static(value interface{}) *PortletTab {
    a.Set("static", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *PortletTab) EditorSetting(value interface{}) *PortletTab {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *PortletTab) Hidden(value interface{}) *PortletTab {
    a.Set("hidden", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *PortletTab) Id(value interface{}) *PortletTab {
    a.Set("id", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *PortletTab) StaticPlaceholder(value interface{}) *PortletTab {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *PortletTab) StaticInputClassName(value interface{}) *PortletTab {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * Tab 标题
 */
func (a *PortletTab) Title(value interface{}) *PortletTab {
    a.Set("title", value)
    return a
}

/**
 * 内容
 */
func (a *PortletTab) Tab(value interface{}) *PortletTab {
    a.Set("tab", value)
    return a
}

/**
 * 内容
 */
func (a *PortletTab) Body(value interface{}) *PortletTab {
    a.Set("body", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *PortletTab) HiddenOn(value interface{}) *PortletTab {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *PortletTab) OnEvent(value interface{}) *PortletTab {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *PortletTab) StaticOn(value interface{}) *PortletTab {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *PortletTab) StaticSchema(value interface{}) *PortletTab {
    a.Set("staticSchema", value)
    return a
}

/**
 * 卡片隐藏就销毁卡片节点。
 */
func (a *PortletTab) UnmountOnExit(value interface{}) *PortletTab {
    a.Set("unmountOnExit", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *PortletTab) DisabledOn(value interface{}) *PortletTab {
    a.Set("disabledOn", value)
    return a
}

/**
 * 设置以后内容每次都会重新渲染
 */
func (a *PortletTab) Reload(value interface{}) *PortletTab {
    a.Set("reload", value)
    return a
}

/**
 * 点开时才加载卡片内容
 */
func (a *PortletTab) MountOnEnter(value interface{}) *PortletTab {
    a.Set("mountOnEnter", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *PortletTab) ClassName(value interface{}) *PortletTab {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用
 */
func (a *PortletTab) Disabled(value interface{}) *PortletTab {
    a.Set("disabled", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *PortletTab) StaticLabelClassName(value interface{}) *PortletTab {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *PortletTab) UseMobileUI(value interface{}) *PortletTab {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 可以在右侧配置点其他功能按钮，随着tab切换而切换
 */
func (a *PortletTab) Toolbar(value interface{}) *PortletTab {
    a.Set("toolbar", value)
    return a
}

/**
 * 是否显示
 */
func (a *PortletTab) Visible(value interface{}) *PortletTab {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *PortletTab) StaticClassName(value interface{}) *PortletTab {
    a.Set("staticClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *PortletTab) Style(value interface{}) *PortletTab {
    a.Set("style", value)
    return a
}

/**
 * 按钮图标
 */
func (a *PortletTab) Icon(value interface{}) *PortletTab {
    a.Set("icon", value)
    return a
}

/**
 * 可选值: left | right
 */
func (a *PortletTab) IconPosition(value interface{}) *PortletTab {
    a.Set("iconPosition", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *PortletTab) VisibleOn(value interface{}) *PortletTab {
    a.Set("visibleOn", value)
    return a
}

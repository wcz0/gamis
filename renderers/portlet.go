package renderers


/**

*/
type Portlet struct {
	*BaseRenderer
}

func NewPortlet() *Portlet {
    a := &Portlet{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "portlet")
    return a
}
/**
 * 隐藏头部
 */
func (a *Portlet) HideHeader(value string) *Portlet {
    a.Set("hideHeader", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Portlet) Hidden(value string) *Portlet {
    a.Set("hidden", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Portlet) StaticLabelClassName(value string) *Portlet {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 可以在右侧配置点其他功能按钮。不会随着tab切换
 */
func (a *Portlet) Toolbar(value string) *Portlet {
    a.Set("toolbar", value)
    return a
}

/**
 * 标题右侧的描述
 */
func (a *Portlet) Description(value string) *Portlet {
    a.Set("description", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Portlet) StaticInputClassName(value string) *Portlet {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 指定为 portlet 类型
 */
func (a *Portlet) Type(value string) *Portlet {
    a.Set("type", value)
    return a
}

/**
 * 卡片是否只有在点开的时候加载？
 */
func (a *Portlet) MountOnEnter(value string) *Portlet {
    a.Set("mountOnEnter", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Portlet) StaticOn(value string) *Portlet {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *Portlet) StaticSchema(value string) *Portlet {
    a.Set("staticSchema", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Portlet) EditorSetting(value string) *Portlet {
    a.Set("editorSetting", value)
    return a
}

/**
 * header和内容是否展示分割线
 */
func (a *Portlet) Divider(value string) *Portlet {
    a.Set("divider", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Portlet) DisabledOn(value string) *Portlet {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *Portlet) Visible(value string) *Portlet {
    a.Set("visible", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Portlet) Static(value string) *Portlet {
    a.Set("static", value)
    return a
}

/**
 * 是否支持溢出滚动
 */
func (a *Portlet) Scrollable(value string) *Portlet {
    a.Set("scrollable", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Portlet) ClassName(value string) *Portlet {
    a.Set("className", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Portlet) UseMobileUI(value string) *Portlet {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 类名
 */
func (a *Portlet) TabsClassName(value string) *Portlet {
    a.Set("tabsClassName", value)
    return a
}

/**
 * 关联已有数据，选项卡直接根据目标数据重复。
 */
func (a *Portlet) Source(value string) *Portlet {
    a.Set("source", value)
    return a
}

/**
 * 展示形式
 * 可选值:  | line | card | radio | vertical | tiled
 */
func (a *Portlet) TabsMode(value string) *Portlet {
    a.Set("tabsMode", value)
    return a
}

/**
 * 内容类名
 */
func (a *Portlet) ContentClassName(value string) *Portlet {
    a.Set("contentClassName", value)
    return a
}

/**
 * 链接外层类名
 */
func (a *Portlet) LinksClassName(value string) *Portlet {
    a.Set("linksClassName", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Portlet) HiddenOn(value string) *Portlet {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Portlet) Id(value string) *Portlet {
    a.Set("id", value)
    return a
}

/**
 */
func (a *Portlet) Tabs(value string) *Portlet {
    a.Set("tabs", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Portlet) StaticClassName(value string) *Portlet {
    a.Set("staticClassName", value)
    return a
}

/**
 * 卡片隐藏的时候是否销毁卡片内容
 */
func (a *Portlet) UnmountOnExit(value string) *Portlet {
    a.Set("unmountOnExit", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Portlet) VisibleOn(value string) *Portlet {
    a.Set("visibleOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Portlet) OnEvent(value string) *Portlet {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Portlet) StaticPlaceholder(value string) *Portlet {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Portlet) Disabled(value string) *Portlet {
    a.Set("disabled", value)
    return a
}

/**
 * 自定义样式
 */
func (a *Portlet) Style(value string) *Portlet {
    a.Set("style", value)
    return a
}

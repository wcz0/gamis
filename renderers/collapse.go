package renderers


/**
 * Collapse 折叠渲染器，格式说明。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/collapse

* @author wcz0
* @version 6.2.2
*/
type Collapse struct {
	*BaseRenderer
}

func NewCollapse() *Collapse {
    a := &Collapse{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "collapse")
    return a
}
/**
 * 是否静态展示
 */
func (a *Collapse) Static(value string) *Collapse {
    a.Set("static", value)
    return a
}

/**
 * 标题内容分割线
 */
func (a *Collapse) DivideLine(value string) *Collapse {
    a.Set("divideLine", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Collapse) DisabledOn(value string) *Collapse {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Collapse) Hidden(value string) *Collapse {
    a.Set("hidden", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Collapse) StaticPlaceholder(value string) *Collapse {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Collapse) StaticClassName(value string) *Collapse {
    a.Set("staticClassName", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Collapse) Id(value string) *Collapse {
    a.Set("id", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Collapse) OnEvent(value string) *Collapse {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Collapse) StaticOn(value string) *Collapse {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *Collapse) StaticSchema(value string) *Collapse {
    a.Set("staticSchema", value)
    return a
}

/**
 * 控件大小
 * 可选值: xs | sm | md | lg | base
 */
func (a *Collapse) Size(value string) *Collapse {
    a.Set("size", value)
    return a
}

/**
 * 卡片隐藏就销毁内容。
 */
func (a *Collapse) UnmountOnExit(value string) *Collapse {
    a.Set("unmountOnExit", value)
    return a
}

/**
 * 是否显示
 */
func (a *Collapse) Visible(value string) *Collapse {
    a.Set("visible", value)
    return a
}

/**
 * 组件样式
 */
func (a *Collapse) Style(value string) *Collapse {
    a.Set("style", value)
    return a
}

/**
 * 内容区域
 */
func (a *Collapse) Body(value string) *Collapse {
    a.Set("body", value)
    return a
}

/**
 * 是否可折叠
 */
func (a *Collapse) Collapsable(value string) *Collapse {
    a.Set("collapsable", value)
    return a
}

/**
 * 默认是否折叠
 */
func (a *Collapse) Collapsed(value string) *Collapse {
    a.Set("collapsed", value)
    return a
}

/**
 * 图标是否展示
 */
func (a *Collapse) ShowArrow(value string) *Collapse {
    a.Set("showArrow", value)
    return a
}

/**
 * 自定义切换图标
 */
func (a *Collapse) ExpandIcon(value string) *Collapse {
    a.Set("expandIcon", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Collapse) ClassName(value string) *Collapse {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Collapse) HiddenOn(value string) *Collapse {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Collapse) StaticInputClassName(value string) *Collapse {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Collapse) UseMobileUI(value string) *Collapse {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 标题展示位置
 * 可选值: top | bottom
 */
func (a *Collapse) HeaderPosition(value string) *Collapse {
    a.Set("headerPosition", value)
    return a
}

/**
 * 配置 Body 容器 className
 */
func (a *Collapse) BodyClassName(value string) *Collapse {
    a.Set("bodyClassName", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Collapse) Disabled(value string) *Collapse {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Collapse) VisibleOn(value string) *Collapse {
    a.Set("visibleOn", value)
    return a
}

/**
 * 指定为折叠器类型
 */
func (a *Collapse) Type(value string) *Collapse {
    a.Set("type", value)
    return a
}

/**
 * 标题
 */
func (a *Collapse) Header(value string) *Collapse {
    a.Set("header", value)
    return a
}

/**
 * 标题 CSS 类名
 */
func (a *Collapse) HeadingClassName(value string) *Collapse {
    a.Set("headingClassName", value)
    return a
}

/**
 * 点开时才加载内容
 */
func (a *Collapse) MountOnEnter(value string) *Collapse {
    a.Set("mountOnEnter", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Collapse) StaticLabelClassName(value string) *Collapse {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Collapse) EditorSetting(value string) *Collapse {
    a.Set("editorSetting", value)
    return a
}

/**
 * 标识
 */
func (a *Collapse) Key(value string) *Collapse {
    a.Set("key", value)
    return a
}

/**
 * 收起的标题
 */
func (a *Collapse) CollapseHeader(value string) *Collapse {
    a.Set("collapseHeader", value)
    return a
}

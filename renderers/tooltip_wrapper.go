package renderers


/**

*/
type TooltipWrapper struct {
	*BaseRenderer
}

func NewTooltipWrapper() *TooltipWrapper {
    a := &TooltipWrapper{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "tooltip-wrapper")
    return a
}
/**
 * 是否显示
 */
func (a *TooltipWrapper) Visible(value string) *TooltipWrapper {
    a.Set("visible", value)
    return a
}

/**
 * 文字提示标题
 */
func (a *TooltipWrapper) Title(value string) *TooltipWrapper {
    a.Set("title", value)
    return a
}

/**
 * 浮层位置相对偏移量
 */
func (a *TooltipWrapper) Offset(value string) *TooltipWrapper {
    a.Set("offset", value)
    return a
}

/**
 * 自定义提示浮层样式
 */
func (a *TooltipWrapper) TooltipStyle(value string) *TooltipWrapper {
    a.Set("tooltipStyle", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *TooltipWrapper) DisabledOn(value string) *TooltipWrapper {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *TooltipWrapper) Hidden(value string) *TooltipWrapper {
    a.Set("hidden", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *TooltipWrapper) StaticOn(value string) *TooltipWrapper {
    a.Set("staticOn", value)
    return a
}

/**
 * 文字提示容器
 */
func (a *TooltipWrapper) Type(value string) *TooltipWrapper {
    a.Set("type", value)
    return a
}

/**
 * 文字提示内容，兼容 tooltip，但建议通过 content 来实现提示内容
 */
func (a *TooltipWrapper) Content(value string) *TooltipWrapper {
    a.Set("content", value)
    return a
}

/**
 * 是否点击非内容区域关闭提示，默认为true
 */
func (a *TooltipWrapper) RootClose(value string) *TooltipWrapper {
    a.Set("rootClose", value)
    return a
}

/**
 * 内容区包裹标签
 */
func (a *TooltipWrapper) WrapperComponent(value string) *TooltipWrapper {
    a.Set("wrapperComponent", value)
    return a
}

/**
 * 主题样式， 默认为light
 * 可选值: light | dark
 */
func (a *TooltipWrapper) TooltipTheme(value string) *TooltipWrapper {
    a.Set("tooltipTheme", value)
    return a
}

/**
 * 内容区CSS类名
 */
func (a *TooltipWrapper) ClassName(value string) *TooltipWrapper {
    a.Set("className", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *TooltipWrapper) VisibleOn(value string) *TooltipWrapper {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *TooltipWrapper) StaticInputClassName(value string) *TooltipWrapper {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *TooltipWrapper) EditorSetting(value string) *TooltipWrapper {
    a.Set("editorSetting", value)
    return a
}

/**
 * 浮层延迟隐藏时间, 单位 ms
 */
func (a *TooltipWrapper) MouseLeaveDelay(value string) *TooltipWrapper {
    a.Set("mouseLeaveDelay", value)
    return a
}

/**
 * 是否禁用提示
 */
func (a *TooltipWrapper) Disabled(value string) *TooltipWrapper {
    a.Set("disabled", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *TooltipWrapper) OnEvent(value string) *TooltipWrapper {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *TooltipWrapper) StaticClassName(value string) *TooltipWrapper {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *TooltipWrapper) Tooltip(value string) *TooltipWrapper {
    a.Set("tooltip", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *TooltipWrapper) Static(value string) *TooltipWrapper {
    a.Set("static", value)
    return a
}

/**
 */
func (a *TooltipWrapper) StaticSchema(value string) *TooltipWrapper {
    a.Set("staticSchema", value)
    return a
}

/**
 * 文字提示浮层出现位置，默认为top
 * 可选值: top | right | bottom | left
 */
func (a *TooltipWrapper) Placement(value string) *TooltipWrapper {
    a.Set("placement", value)
    return a
}

/**
 * 内容区是否内联显示，默认为false
 */
func (a *TooltipWrapper) Inline(value string) *TooltipWrapper {
    a.Set("inline", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *TooltipWrapper) StaticPlaceholder(value string) *TooltipWrapper {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *TooltipWrapper) UseMobileUI(value string) *TooltipWrapper {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 浮层触发方式，默认为hover
 */
func (a *TooltipWrapper) Trigger(value string) *TooltipWrapper {
    a.Set("trigger", value)
    return a
}

/**
 * 内容区域
 */
func (a *TooltipWrapper) Body(value string) *TooltipWrapper {
    a.Set("body", value)
    return a
}

/**
 * 内容区自定义样式
 */
func (a *TooltipWrapper) Style(value string) *TooltipWrapper {
    a.Set("style", value)
    return a
}

/**
 * 是否展示浮层指向箭头
 */
func (a *TooltipWrapper) ShowArrow(value string) *TooltipWrapper {
    a.Set("showArrow", value)
    return a
}

/**
 * 浮层延迟显示时间, 单位 ms
 */
func (a *TooltipWrapper) MouseEnterDelay(value string) *TooltipWrapper {
    a.Set("mouseEnterDelay", value)
    return a
}

/**
 * 文字提示浮层CSS类名
 */
func (a *TooltipWrapper) TooltipClassName(value string) *TooltipWrapper {
    a.Set("tooltipClassName", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *TooltipWrapper) HiddenOn(value string) *TooltipWrapper {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *TooltipWrapper) Id(value string) *TooltipWrapper {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *TooltipWrapper) StaticLabelClassName(value string) *TooltipWrapper {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 是否可以移入浮层中, 默认true
 */
func (a *TooltipWrapper) Enterable(value string) *TooltipWrapper {
    a.Set("enterable", value)
    return a
}

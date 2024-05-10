package renderers


/**

 * @author wcz0
 * @version 6.2.2
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
 * 是否禁用表达式
 */
func (a *TooltipWrapper) DisabledOn(value interface{}) *TooltipWrapper {
    a.Set("disabledOn", value)
    return a
}

/**
 * 文字提示标题
 */
func (a *TooltipWrapper) Title(value interface{}) *TooltipWrapper {
    a.Set("title", value)
    return a
}

/**
 * 浮层延迟隐藏时间, 单位 ms
 */
func (a *TooltipWrapper) MouseLeaveDelay(value interface{}) *TooltipWrapper {
    a.Set("mouseLeaveDelay", value)
    return a
}

/**
 * 自定义提示浮层样式
 */
func (a *TooltipWrapper) TooltipStyle(value interface{}) *TooltipWrapper {
    a.Set("tooltipStyle", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *TooltipWrapper) Hidden(value interface{}) *TooltipWrapper {
    a.Set("hidden", value)
    return a
}

/**
 */
func (a *TooltipWrapper) StaticSchema(value interface{}) *TooltipWrapper {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *TooltipWrapper) Tooltip(value interface{}) *TooltipWrapper {
    a.Set("tooltip", value)
    return a
}

/**
 * 是否可以移入浮层中, 默认true
 */
func (a *TooltipWrapper) Enterable(value interface{}) *TooltipWrapper {
    a.Set("enterable", value)
    return a
}

/**
 * 内容区CSS类名
 */
func (a *TooltipWrapper) ClassName(value interface{}) *TooltipWrapper {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *TooltipWrapper) HiddenOn(value interface{}) *TooltipWrapper {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *TooltipWrapper) StaticPlaceholder(value interface{}) *TooltipWrapper {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 文字提示浮层CSS类名
 */
func (a *TooltipWrapper) TooltipClassName(value interface{}) *TooltipWrapper {
    a.Set("tooltipClassName", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *TooltipWrapper) Id(value interface{}) *TooltipWrapper {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *TooltipWrapper) Static(value interface{}) *TooltipWrapper {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *TooltipWrapper) StaticInputClassName(value interface{}) *TooltipWrapper {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *TooltipWrapper) EditorSetting(value interface{}) *TooltipWrapper {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否展示浮层指向箭头
 */
func (a *TooltipWrapper) ShowArrow(value interface{}) *TooltipWrapper {
    a.Set("showArrow", value)
    return a
}

/**
 * 浮层触发方式，默认为hover
 */
func (a *TooltipWrapper) Trigger(value interface{}) *TooltipWrapper {
    a.Set("trigger", value)
    return a
}

/**
 * 是否显示
 */
func (a *TooltipWrapper) Visible(value interface{}) *TooltipWrapper {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *TooltipWrapper) StaticClassName(value interface{}) *TooltipWrapper {
    a.Set("staticClassName", value)
    return a
}

/**
 * 浮层延迟显示时间, 单位 ms
 */
func (a *TooltipWrapper) MouseEnterDelay(value interface{}) *TooltipWrapper {
    a.Set("mouseEnterDelay", value)
    return a
}

/**
 * 内容区是否内联显示，默认为false
 */
func (a *TooltipWrapper) Inline(value interface{}) *TooltipWrapper {
    a.Set("inline", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *TooltipWrapper) VisibleOn(value interface{}) *TooltipWrapper {
    a.Set("visibleOn", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *TooltipWrapper) UseMobileUI(value interface{}) *TooltipWrapper {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 文字提示浮层出现位置，默认为top
 * 可选值: top | right | bottom | left
 */
func (a *TooltipWrapper) Placement(value interface{}) *TooltipWrapper {
    a.Set("placement", value)
    return a
}

/**
 * 内容区域
 */
func (a *TooltipWrapper) Body(value interface{}) *TooltipWrapper {
    a.Set("body", value)
    return a
}

/**
 * 内容区包裹标签
 */
func (a *TooltipWrapper) WrapperComponent(value interface{}) *TooltipWrapper {
    a.Set("wrapperComponent", value)
    return a
}

/**
 * 是否禁用提示
 */
func (a *TooltipWrapper) Disabled(value interface{}) *TooltipWrapper {
    a.Set("disabled", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *TooltipWrapper) StaticLabelClassName(value interface{}) *TooltipWrapper {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 内容区自定义样式
 */
func (a *TooltipWrapper) Style(value interface{}) *TooltipWrapper {
    a.Set("style", value)
    return a
}

/**
 * 文字提示容器
 */
func (a *TooltipWrapper) Type(value interface{}) *TooltipWrapper {
    a.Set("type", value)
    return a
}

/**
 * 文字提示内容，兼容 tooltip，但建议通过 content 来实现提示内容
 */
func (a *TooltipWrapper) Content(value interface{}) *TooltipWrapper {
    a.Set("content", value)
    return a
}

/**
 * 是否点击非内容区域关闭提示，默认为true
 */
func (a *TooltipWrapper) RootClose(value interface{}) *TooltipWrapper {
    a.Set("rootClose", value)
    return a
}

/**
 * 主题样式， 默认为light
 * 可选值: light | dark
 */
func (a *TooltipWrapper) TooltipTheme(value interface{}) *TooltipWrapper {
    a.Set("tooltipTheme", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *TooltipWrapper) OnEvent(value interface{}) *TooltipWrapper {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *TooltipWrapper) StaticOn(value interface{}) *TooltipWrapper {
    a.Set("staticOn", value)
    return a
}

/**
 * 浮层位置相对偏移量
 */
func (a *TooltipWrapper) Offset(value interface{}) *TooltipWrapper {
    a.Set("offset", value)
    return a
}

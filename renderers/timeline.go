package renderers


/**

*/
type Timeline struct {
	*BaseRenderer
}

func NewTimeline() *Timeline {
    a := &Timeline{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "timeline")
    return a
}
/**
 * 图标的CSS类名
 */
func (a *Timeline) IconClassName(value string) *Timeline {
    a.Set("iconClassName", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Timeline) StaticClassName(value string) *Timeline {
    a.Set("staticClassName", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Timeline) Hidden(value string) *Timeline {
    a.Set("hidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Timeline) HiddenOn(value string) *Timeline {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *Timeline) Visible(value string) *Timeline {
    a.Set("visible", value)
    return a
}

/**
 * 展示方向
 * 可选值: horizontal | vertical
 */
func (a *Timeline) Direction(value string) *Timeline {
    a.Set("direction", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Timeline) Disabled(value string) *Timeline {
    a.Set("disabled", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Timeline) OnEvent(value string) *Timeline {
    a.Set("onEvent", value)
    return a
}

/**
 * 文字相对于时间轴展示方向
 * 可选值: left | right | alternate
 */
func (a *Timeline) Mode(value string) *Timeline {
    a.Set("mode", value)
    return a
}

/**
 * 节点时间的CSS类名
 */
func (a *Timeline) TimeClassName(value string) *Timeline {
    a.Set("timeClassName", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Timeline) Id(value string) *Timeline {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Timeline) StaticInputClassName(value string) *Timeline {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *Timeline) Style(value string) *Timeline {
    a.Set("style", value)
    return a
}

/**
 * 节点title自定一展示模板
 */
func (a *Timeline) ItemTitleSchema(value string) *Timeline {
    a.Set("itemTitleSchema", value)
    return a
}

/**
 * 节点详情的CSS类名
 */
func (a *Timeline) DetailClassName(value string) *Timeline {
    a.Set("detailClassName", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Timeline) Static(value string) *Timeline {
    a.Set("static", value)
    return a
}

/**
 * API 或 数据映射
 */
func (a *Timeline) Source(value string) *Timeline {
    a.Set("source", value)
    return a
}

/**
 * 节点倒序
 */
func (a *Timeline) Reverse(value string) *Timeline {
    a.Set("reverse", value)
    return a
}

/**
 * 节点标题的CSS类名
 */
func (a *Timeline) TitleClassName(value string) *Timeline {
    a.Set("titleClassName", value)
    return a
}

/**
 * 节点数据
 */
func (a *Timeline) Items(value string) *Timeline {
    a.Set("items", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Timeline) StaticLabelClassName(value string) *Timeline {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Timeline) EditorSetting(value string) *Timeline {
    a.Set("editorSetting", value)
    return a
}

/**
 * 指定为 Timeline 时间轴渲染器
 */
func (a *Timeline) Type(value string) *Timeline {
    a.Set("type", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Timeline) DisabledOn(value string) *Timeline {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Timeline) VisibleOn(value string) *Timeline {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Timeline) StaticOn(value string) *Timeline {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Timeline) StaticPlaceholder(value string) *Timeline {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Timeline) ClassName(value string) *Timeline {
    a.Set("className", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Timeline) UseMobileUI(value string) *Timeline {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Timeline) StaticSchema(value string) *Timeline {
    a.Set("staticSchema", value)
    return a
}

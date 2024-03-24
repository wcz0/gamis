package renderers


/**
 * AnchorNavSection 锚点区域渲染器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/anchor-nav

* @author wcz0
* @version 6.2.2
*/
type AnchorNavSection struct {
	*BaseRenderer
}

func NewAnchorNavSection() *AnchorNavSection {
    a := &AnchorNavSection{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}
/**
 * 编辑器配置，运行时可以忽略
 */
func (a *AnchorNavSection) EditorSetting(value string) *AnchorNavSection {
    a.Set("editorSetting", value)
    return a
}

/**
 * 锚点链接
 */
func (a *AnchorNavSection) Href(value string) *AnchorNavSection {
    a.Set("href", value)
    return a
}

/**
 * 是否禁用
 */
func (a *AnchorNavSection) Disabled(value string) *AnchorNavSection {
    a.Set("disabled", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *AnchorNavSection) Id(value string) *AnchorNavSection {
    a.Set("id", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *AnchorNavSection) StaticPlaceholder(value string) *AnchorNavSection {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *AnchorNavSection) StaticSchema(value string) *AnchorNavSection {
    a.Set("staticSchema", value)
    return a
}

/**
 * 是否显示
 */
func (a *AnchorNavSection) Visible(value string) *AnchorNavSection {
    a.Set("visible", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *AnchorNavSection) Static(value string) *AnchorNavSection {
    a.Set("static", value)
    return a
}

/**
 * 组件样式
 */
func (a *AnchorNavSection) Style(value string) *AnchorNavSection {
    a.Set("style", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *AnchorNavSection) ClassName(value string) *AnchorNavSection {
    a.Set("className", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *AnchorNavSection) VisibleOn(value string) *AnchorNavSection {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *AnchorNavSection) StaticLabelClassName(value string) *AnchorNavSection {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 子节点
 */
func (a *AnchorNavSection) Children(value string) *AnchorNavSection {
    a.Set("children", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *AnchorNavSection) OnEvent(value string) *AnchorNavSection {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *AnchorNavSection) StaticOn(value string) *AnchorNavSection {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *AnchorNavSection) StaticClassName(value string) *AnchorNavSection {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *AnchorNavSection) StaticInputClassName(value string) *AnchorNavSection {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *AnchorNavSection) DisabledOn(value string) *AnchorNavSection {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *AnchorNavSection) Hidden(value string) *AnchorNavSection {
    a.Set("hidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *AnchorNavSection) HiddenOn(value string) *AnchorNavSection {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *AnchorNavSection) UseMobileUI(value string) *AnchorNavSection {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 导航文字说明
 */
func (a *AnchorNavSection) Title(value string) *AnchorNavSection {
    a.Set("title", value)
    return a
}

/**
 * 内容
 */
func (a *AnchorNavSection) Body(value string) *AnchorNavSection {
    a.Set("body", value)
    return a
}

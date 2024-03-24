package renderers


/**
 * CollapseGroup 折叠渲染器，格式说明。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/collapse

* @author wcz0
* @version 6.2.2
*/
type CollapseGroup struct {
	*BaseRenderer
}

func NewCollapseGroup() *CollapseGroup {
    a := &CollapseGroup{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "collapse-group")
    return a
}
/**
 * 静态展示空值占位
 */
func (a *CollapseGroup) StaticPlaceholder(value string) *CollapseGroup {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *CollapseGroup) StaticSchema(value string) *CollapseGroup {
    a.Set("staticSchema", value)
    return a
}

/**
 * 内容区域
 */
func (a *CollapseGroup) Body(value string) *CollapseGroup {
    a.Set("body", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *CollapseGroup) ClassName(value string) *CollapseGroup {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *CollapseGroup) Hidden(value string) *CollapseGroup {
    a.Set("hidden", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *CollapseGroup) UseMobileUI(value string) *CollapseGroup {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *CollapseGroup) StaticOn(value string) *CollapseGroup {
    a.Set("staticOn", value)
    return a
}

/**
 * 设置图标位置
 * 可选值: left | right
 */
func (a *CollapseGroup) ExpandIconPosition(value string) *CollapseGroup {
    a.Set("expandIconPosition", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *CollapseGroup) HiddenOn(value string) *CollapseGroup {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *CollapseGroup) Id(value string) *CollapseGroup {
    a.Set("id", value)
    return a
}

/**
 * 手风琴模式
 */
func (a *CollapseGroup) Accordion(value string) *CollapseGroup {
    a.Set("accordion", value)
    return a
}

/**
 * 当Collapse作为Form组件的子元素时，开启该属性后组件样式设置为FieldSet组件的样式，默认开启
 */
func (a *CollapseGroup) EnableFieldSetStyle(value string) *CollapseGroup {
    a.Set("enableFieldSetStyle", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *CollapseGroup) DisabledOn(value string) *CollapseGroup {
    a.Set("disabledOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *CollapseGroup) OnEvent(value string) *CollapseGroup {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *CollapseGroup) StaticClassName(value string) *CollapseGroup {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *CollapseGroup) StaticLabelClassName(value string) *CollapseGroup {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 指定为折叠器类型
 */
func (a *CollapseGroup) Type(value string) *CollapseGroup {
    a.Set("type", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *CollapseGroup) Static(value string) *CollapseGroup {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *CollapseGroup) StaticInputClassName(value string) *CollapseGroup {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *CollapseGroup) Style(value string) *CollapseGroup {
    a.Set("style", value)
    return a
}

/**
 * 自定义切换图标
 */
func (a *CollapseGroup) ExpandIcon(value string) *CollapseGroup {
    a.Set("expandIcon", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *CollapseGroup) VisibleOn(value string) *CollapseGroup {
    a.Set("visibleOn", value)
    return a
}

/**
 * 激活面板
 */
func (a *CollapseGroup) ActiveKey(value string) *CollapseGroup {
    a.Set("activeKey", value)
    return a
}

/**
 * 是否禁用
 */
func (a *CollapseGroup) Disabled(value string) *CollapseGroup {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示
 */
func (a *CollapseGroup) Visible(value string) *CollapseGroup {
    a.Set("visible", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *CollapseGroup) EditorSetting(value string) *CollapseGroup {
    a.Set("editorSetting", value)
    return a
}

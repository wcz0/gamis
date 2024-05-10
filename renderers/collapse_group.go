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
 * 可以组件级别用来关闭移动端样式
 */
func (a *CollapseGroup) UseMobileUI(value interface{}) *CollapseGroup {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 内容区域
 */
func (a *CollapseGroup) Body(value interface{}) *CollapseGroup {
    a.Set("body", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *CollapseGroup) ClassName(value interface{}) *CollapseGroup {
    a.Set("className", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *CollapseGroup) Id(value interface{}) *CollapseGroup {
    a.Set("id", value)
    return a
}

/**
 * 组件样式
 */
func (a *CollapseGroup) Style(value interface{}) *CollapseGroup {
    a.Set("style", value)
    return a
}

/**
 * 激活面板
 */
func (a *CollapseGroup) ActiveKey(value interface{}) *CollapseGroup {
    a.Set("activeKey", value)
    return a
}

/**
 * 是否禁用
 */
func (a *CollapseGroup) Disabled(value interface{}) *CollapseGroup {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *CollapseGroup) Hidden(value interface{}) *CollapseGroup {
    a.Set("hidden", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *CollapseGroup) StaticOn(value interface{}) *CollapseGroup {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *CollapseGroup) StaticLabelClassName(value interface{}) *CollapseGroup {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 是否显示
 */
func (a *CollapseGroup) Visible(value interface{}) *CollapseGroup {
    a.Set("visible", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *CollapseGroup) VisibleOn(value interface{}) *CollapseGroup {
    a.Set("visibleOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *CollapseGroup) EditorSetting(value interface{}) *CollapseGroup {
    a.Set("editorSetting", value)
    return a
}

/**
 * 指定为折叠器类型
 */
func (a *CollapseGroup) Type(value interface{}) *CollapseGroup {
    a.Set("type", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *CollapseGroup) HiddenOn(value interface{}) *CollapseGroup {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *CollapseGroup) StaticPlaceholder(value interface{}) *CollapseGroup {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *CollapseGroup) StaticInputClassName(value interface{}) *CollapseGroup {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 当Collapse作为Form组件的子元素时，开启该属性后组件样式设置为FieldSet组件的样式，默认开启
 */
func (a *CollapseGroup) EnableFieldSetStyle(value interface{}) *CollapseGroup {
    a.Set("enableFieldSetStyle", value)
    return a
}

/**
 * 设置图标位置
 * 可选值: left | right
 */
func (a *CollapseGroup) ExpandIconPosition(value interface{}) *CollapseGroup {
    a.Set("expandIconPosition", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *CollapseGroup) OnEvent(value interface{}) *CollapseGroup {
    a.Set("onEvent", value)
    return a
}

/**
 */
func (a *CollapseGroup) StaticSchema(value interface{}) *CollapseGroup {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *CollapseGroup) Testid(value interface{}) *CollapseGroup {
    a.Set("testid", value)
    return a
}

/**
 * 自定义切换图标
 */
func (a *CollapseGroup) ExpandIcon(value interface{}) *CollapseGroup {
    a.Set("expandIcon", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *CollapseGroup) Static(value interface{}) *CollapseGroup {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *CollapseGroup) StaticClassName(value interface{}) *CollapseGroup {
    a.Set("staticClassName", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *CollapseGroup) DisabledOn(value interface{}) *CollapseGroup {
    a.Set("disabledOn", value)
    return a
}

/**
 */
func (a *CollapseGroup) TestIdBuilder(value interface{}) *CollapseGroup {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 手风琴模式
 */
func (a *CollapseGroup) Accordion(value interface{}) *CollapseGroup {
    a.Set("accordion", value)
    return a
}

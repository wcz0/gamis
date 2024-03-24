package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type ListItem struct {
	*BaseRenderer
}

func NewListItem() *ListItem {
    a := &ListItem{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}
/**
 * 静态展示表单项Value类名
 */
func (a *ListItem) StaticInputClassName(value interface{}) *ListItem {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *ListItem) Actions(value interface{}) *ListItem {
    a.Set("actions", value)
    return a
}

/**
 * 图片地址
 */
func (a *ListItem) Avatar(value interface{}) *ListItem {
    a.Set("avatar", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *ListItem) OnEvent(value interface{}) *ListItem {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *ListItem) Static(value interface{}) *ListItem {
    a.Set("static", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *ListItem) StaticPlaceholder(value interface{}) *ListItem {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *ListItem) StaticSchema(value interface{}) *ListItem {
    a.Set("staticSchema", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *ListItem) Hidden(value interface{}) *ListItem {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示
 */
func (a *ListItem) Visible(value interface{}) *ListItem {
    a.Set("visible", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *ListItem) DisabledOn(value interface{}) *ListItem {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *ListItem) StaticOn(value interface{}) *ListItem {
    a.Set("staticOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *ListItem) VisibleOn(value interface{}) *ListItem {
    a.Set("visibleOn", value)
    return a
}

/**
 * 组件样式
 */
func (a *ListItem) Style(value interface{}) *ListItem {
    a.Set("style", value)
    return a
}

/**
 * 副标题
 */
func (a *ListItem) SubTitle(value interface{}) *ListItem {
    a.Set("subTitle", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *ListItem) ClassName(value interface{}) *ListItem {
    a.Set("className", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *ListItem) StaticLabelClassName(value interface{}) *ListItem {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 内容区域
 */
func (a *ListItem) Body(value interface{}) *ListItem {
    a.Set("body", value)
    return a
}

/**
 * 描述
 */
func (a *ListItem) Desc(value interface{}) *ListItem {
    a.Set("desc", value)
    return a
}

/**
 * tooltip 说明
 */
func (a *ListItem) Remark(value interface{}) *ListItem {
    a.Set("remark", value)
    return a
}

/**
 * 标题
 */
func (a *ListItem) Title(value interface{}) *ListItem {
    a.Set("title", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *ListItem) HiddenOn(value interface{}) *ListItem {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *ListItem) EditorSetting(value interface{}) *ListItem {
    a.Set("editorSetting", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *ListItem) UseMobileUI(value interface{}) *ListItem {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 操作位置，默认在右侧，可以设置成左侧。
 * 可选值: left | right
 */
func (a *ListItem) ActionsPosition(value interface{}) *ListItem {
    a.Set("actionsPosition", value)
    return a
}

/**
 * 是否禁用
 */
func (a *ListItem) Disabled(value interface{}) *ListItem {
    a.Set("disabled", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *ListItem) Id(value interface{}) *ListItem {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *ListItem) StaticClassName(value interface{}) *ListItem {
    a.Set("staticClassName", value)
    return a
}

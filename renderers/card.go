package renderers


/**
 * Card 卡片渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/card
 *

*/
type Card struct {
	*BaseRenderer
}

func NewCard() *Card {
    a := &Card{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "card")
    return a
}
/**
 * 头部配置
 */
func (a *Card) Header(value string) *Card {
    a.Set("header", value)
    return a
}

/**
 * 卡片内容区的表单项label是否使用Card内部的样式，默认为true
 */
func (a *Card) UseCardLabel(value string) *Card {
    a.Set("useCardLabel", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Card) DisabledOn(value string) *Card {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *Card) Visible(value string) *Card {
    a.Set("visible", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Card) Static(value string) *Card {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Card) StaticLabelClassName(value string) *Card {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *Card) Style(value string) *Card {
    a.Set("style", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Card) EditorSetting(value string) *Card {
    a.Set("editorSetting", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Card) UseMobileUI(value string) *Card {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 指定为 card 类型
 */
func (a *Card) Type(value string) *Card {
    a.Set("type", value)
    return a
}

/**
 * 内容区域
 */
func (a *Card) Body(value string) *Card {
    a.Set("body", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Card) ClassName(value string) *Card {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Card) Disabled(value string) *Card {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Card) HiddenOn(value string) *Card {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Card) StaticOn(value string) *Card {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Card) StaticInputClassName(value string) *Card {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Card) VisibleOn(value string) *Card {
    a.Set("visibleOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Card) OnEvent(value string) *Card {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Card) StaticClassName(value string) *Card {
    a.Set("staticClassName", value)
    return a
}

/**
 * 多媒体区域
 */
func (a *Card) Media(value string) *Card {
    a.Set("media", value)
    return a
}

/**
 * 工具栏按钮
 */
func (a *Card) Toolbar(value string) *Card {
    a.Set("toolbar", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Card) Hidden(value string) *Card {
    a.Set("hidden", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Card) Id(value string) *Card {
    a.Set("id", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Card) StaticPlaceholder(value string) *Card {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 底部按钮集合。
 */
func (a *Card) Actions(value string) *Card {
    a.Set("actions", value)
    return a
}

/**
 * 次要说明
 */
func (a *Card) Secondary(value string) *Card {
    a.Set("secondary", value)
    return a
}

/**
 */
func (a *Card) StaticSchema(value string) *Card {
    a.Set("staticSchema", value)
    return a
}
package renderers


/**
 * Card 卡片渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/card

 * @author wcz0
 * @version 6.2.2
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
 * 是否隐藏
 */
func (a *Card) Hidden(value interface{}) *Card {
    a.Set("hidden", value)
    return a
}

/**
 */
func (a *Card) StaticSchema(value interface{}) *Card {
    a.Set("staticSchema", value)
    return a
}

/**
 * 卡片内容区的表单项label是否使用Card内部的样式，默认为true
 */
func (a *Card) UseCardLabel(value interface{}) *Card {
    a.Set("useCardLabel", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Card) StaticOn(value interface{}) *Card {
    a.Set("staticOn", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Card) UseMobileUI(value interface{}) *Card {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Card) HiddenOn(value interface{}) *Card {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Card) VisibleOn(value interface{}) *Card {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Card) Static(value interface{}) *Card {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Card) StaticClassName(value interface{}) *Card {
    a.Set("staticClassName", value)
    return a
}

/**
 * 指定为 card 类型
 */
func (a *Card) Type(value interface{}) *Card {
    a.Set("type", value)
    return a
}

/**
 * 头部配置
 */
func (a *Card) Header(value interface{}) *Card {
    a.Set("header", value)
    return a
}

/**
 * 内容区域
 */
func (a *Card) Body(value interface{}) *Card {
    a.Set("body", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Card) Disabled(value interface{}) *Card {
    a.Set("disabled", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Card) DisabledOn(value interface{}) *Card {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *Card) Visible(value interface{}) *Card {
    a.Set("visible", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Card) Id(value interface{}) *Card {
    a.Set("id", value)
    return a
}

/**
 * 多媒体区域
 */
func (a *Card) Media(value interface{}) *Card {
    a.Set("media", value)
    return a
}

/**
 * 次要说明
 */
func (a *Card) Secondary(value interface{}) *Card {
    a.Set("secondary", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Card) OnEvent(value interface{}) *Card {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Card) StaticPlaceholder(value interface{}) *Card {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Card) EditorSetting(value interface{}) *Card {
    a.Set("editorSetting", value)
    return a
}

/**
 * 组件样式
 */
func (a *Card) Style(value interface{}) *Card {
    a.Set("style", value)
    return a
}

/**
 */
func (a *Card) TestIdBuilder(value interface{}) *Card {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 工具栏按钮
 */
func (a *Card) Toolbar(value interface{}) *Card {
    a.Set("toolbar", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Card) ClassName(value interface{}) *Card {
    a.Set("className", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Card) StaticLabelClassName(value interface{}) *Card {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Card) StaticInputClassName(value interface{}) *Card {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *Card) Testid(value interface{}) *Card {
    a.Set("testid", value)
    return a
}

/**
 * 底部按钮集合。
 */
func (a *Card) Actions(value interface{}) *Card {
    a.Set("actions", value)
    return a
}

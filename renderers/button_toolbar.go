package renderers


/**
 * Button Toolar 渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/button-toolbar
 *

*/
type ButtonToolbar struct {
	*BaseRenderer
}

func NewButtonToolbar() *ButtonToolbar {
    a := &ButtonToolbar{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "button-toolbar")
    return a
}
/**
 * 容器 css 类名
 */
func (a *ButtonToolbar) ClassName(value string) *ButtonToolbar {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *ButtonToolbar) DisabledOn(value string) *ButtonToolbar {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *ButtonToolbar) Static(value string) *ButtonToolbar {
    a.Set("static", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *ButtonToolbar) UseMobileUI(value string) *ButtonToolbar {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 指定为按钮工具集合类型
 */
func (a *ButtonToolbar) Type(value string) *ButtonToolbar {
    a.Set("type", value)
    return a
}

/**
 */
func (a *ButtonToolbar) StaticSchema(value string) *ButtonToolbar {
    a.Set("staticSchema", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *ButtonToolbar) Hidden(value string) *ButtonToolbar {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *ButtonToolbar) VisibleOn(value string) *ButtonToolbar {
    a.Set("visibleOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *ButtonToolbar) OnEvent(value string) *ButtonToolbar {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *ButtonToolbar) StaticOn(value string) *ButtonToolbar {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *ButtonToolbar) StaticPlaceholder(value string) *ButtonToolbar {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *ButtonToolbar) StaticClassName(value string) *ButtonToolbar {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *ButtonToolbar) StaticInputClassName(value string) *ButtonToolbar {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *ButtonToolbar) EditorSetting(value string) *ButtonToolbar {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *ButtonToolbar) Buttons(value string) *ButtonToolbar {
    a.Set("buttons", value)
    return a
}

/**
 * 是否禁用
 */
func (a *ButtonToolbar) Disabled(value string) *ButtonToolbar {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *ButtonToolbar) HiddenOn(value string) *ButtonToolbar {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *ButtonToolbar) Id(value string) *ButtonToolbar {
    a.Set("id", value)
    return a
}

/**
 * 组件样式
 */
func (a *ButtonToolbar) Style(value string) *ButtonToolbar {
    a.Set("style", value)
    return a
}

/**
 * 是否显示
 */
func (a *ButtonToolbar) Visible(value string) *ButtonToolbar {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *ButtonToolbar) StaticLabelClassName(value string) *ButtonToolbar {
    a.Set("staticLabelClassName", value)
    return a
}

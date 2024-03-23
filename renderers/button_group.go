package renderers


/**
 * Button Group 渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/button-group
 *

*/
type ButtonGroup struct {
	*BaseRenderer
}

func NewButtonGroup() *ButtonGroup {
    a := &ButtonGroup{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "button-group")
    return a
}
/**
 * 是否隐藏表达式
 */
func (a *ButtonGroup) HiddenOn(value string) *ButtonGroup {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *ButtonGroup) StaticInputClassName(value string) *ButtonGroup {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 按钮样式级别
 */
func (a *ButtonGroup) BtnLevel(value string) *ButtonGroup {
    a.Set("btnLevel", value)
    return a
}

/**
 */
func (a *ButtonGroup) BtnActiveClassName(value string) *ButtonGroup {
    a.Set("btnActiveClassName", value)
    return a
}

/**
 * 平铺展示？
 */
func (a *ButtonGroup) Tiled(value string) *ButtonGroup {
    a.Set("tiled", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *ButtonGroup) ClassName(value string) *ButtonGroup {
    a.Set("className", value)
    return a
}

/**
 * 是否显示
 */
func (a *ButtonGroup) Visible(value string) *ButtonGroup {
    a.Set("visible", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *ButtonGroup) Static(value string) *ButtonGroup {
    a.Set("static", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *ButtonGroup) UseMobileUI(value string) *ButtonGroup {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *ButtonGroup) EditorSetting(value string) *ButtonGroup {
    a.Set("editorSetting", value)
    return a
}

/**
 * 通过 JS 表达式来配置当前表单项是否显示
 */
func (a *ButtonGroup) VisibleOn(value string) *ButtonGroup {
    a.Set("visibleOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *ButtonGroup) Id(value string) *ButtonGroup {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *ButtonGroup) StaticClassName(value string) *ButtonGroup {
    a.Set("staticClassName", value)
    return a
}

/**
 * 指定为提交按钮类型
 */
func (a *ButtonGroup) Type(value string) *ButtonGroup {
    a.Set("type", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *ButtonGroup) OnEvent(value string) *ButtonGroup {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *ButtonGroup) StaticPlaceholder(value string) *ButtonGroup {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *ButtonGroup) StaticSchema(value string) *ButtonGroup {
    a.Set("staticSchema", value)
    return a
}

/**
 * 按钮大小
 * 可选值: xs | sm | md | lg
 */
func (a *ButtonGroup) Size(value string) *ButtonGroup {
    a.Set("size", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *ButtonGroup) StaticLabelClassName(value string) *ButtonGroup {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *ButtonGroup) Style(value string) *ButtonGroup {
    a.Set("style", value)
    return a
}

/**
 */
func (a *ButtonGroup) BtnClassName(value string) *ButtonGroup {
    a.Set("btnClassName", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *ButtonGroup) Hidden(value string) *ButtonGroup {
    a.Set("hidden", value)
    return a
}

/**
 * 按钮集合
 */
func (a *ButtonGroup) Buttons(value string) *ButtonGroup {
    a.Set("buttons", value)
    return a
}

/**
 * 垂直展示？
 */
func (a *ButtonGroup) Vertical(value string) *ButtonGroup {
    a.Set("vertical", value)
    return a
}

/**
 * 是否为禁用状态。
 */
func (a *ButtonGroup) Disabled(value string) *ButtonGroup {
    a.Set("disabled", value)
    return a
}

/**
 * 通过 JS 表达式来配置当前表单项的禁用状态。
 */
func (a *ButtonGroup) DisabledOn(value string) *ButtonGroup {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *ButtonGroup) StaticOn(value string) *ButtonGroup {
    a.Set("staticOn", value)
    return a
}

/**
 * 按钮选中的样式级别
 */
func (a *ButtonGroup) BtnActiveLevel(value string) *ButtonGroup {
    a.Set("btnActiveLevel", value)
    return a
}

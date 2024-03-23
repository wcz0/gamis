package renderers


/**
 * 下拉按钮渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/dropdown-button
 *

*/
type DropdownButton struct {
	*BaseRenderer
}

func NewDropdownButton() *DropdownButton {
    a := &DropdownButton{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "dropdown-button")
    return a
}
/**
 * 对齐方式
 * 可选值: left | right
 */
func (a *DropdownButton) Align(value string) *DropdownButton {
    a.Set("align", value)
    return a
}

/**
 * 是否显示下拉按钮
 */
func (a *DropdownButton) HideCaret(value string) *DropdownButton {
    a.Set("hideCaret", value)
    return a
}

/**
 * 是否禁用
 */
func (a *DropdownButton) Disabled(value string) *DropdownButton {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *DropdownButton) VisibleOn(value string) *DropdownButton {
    a.Set("visibleOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *DropdownButton) OnEvent(value string) *DropdownButton {
    a.Set("onEvent", value)
    return a
}

/**
 * 按钮级别，样式
 * 可选值: info | success | danger | warning | primary | link
 */
func (a *DropdownButton) Level(value string) *DropdownButton {
    a.Set("level", value)
    return a
}

/**
 * 按钮大小
 * 可选值: xs | sm | md | lg
 */
func (a *DropdownButton) Size(value string) *DropdownButton {
    a.Set("size", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *DropdownButton) StaticOn(value string) *DropdownButton {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *DropdownButton) StaticInputClassName(value string) *DropdownButton {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 按钮集合，支持分组
 */
func (a *DropdownButton) Buttons(value string) *DropdownButton {
    a.Set("buttons", value)
    return a
}

/**
 * 点击外部是否关闭
 */
func (a *DropdownButton) CloseOnOutside(value string) *DropdownButton {
    a.Set("closeOnOutside", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *DropdownButton) ClassName(value string) *DropdownButton {
    a.Set("className", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *DropdownButton) UseMobileUI(value string) *DropdownButton {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *DropdownButton) StaticClassName(value string) *DropdownButton {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *DropdownButton) StaticLabelClassName(value string) *DropdownButton {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 给 Button 配置 className。
 */
func (a *DropdownButton) BtnClassName(value string) *DropdownButton {
    a.Set("btnClassName", value)
    return a
}

/**
 * 右侧图标
 */
func (a *DropdownButton) RightIcon(value string) *DropdownButton {
    a.Set("rightIcon", value)
    return a
}

/**
 * 菜单 CSS 样式
 */
func (a *DropdownButton) MenuClassName(value string) *DropdownButton {
    a.Set("menuClassName", value)
    return a
}

/**
 */
func (a *DropdownButton) OverlayPlacement(value string) *DropdownButton {
    a.Set("overlayPlacement", value)
    return a
}

/**
 * 是否显示
 */
func (a *DropdownButton) Visible(value string) *DropdownButton {
    a.Set("visible", value)
    return a
}

/**
 */
func (a *DropdownButton) StaticSchema(value string) *DropdownButton {
    a.Set("staticSchema", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *DropdownButton) EditorSetting(value string) *DropdownButton {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否独占一行 `display: block`
 */
func (a *DropdownButton) Block(value string) *DropdownButton {
    a.Set("block", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *DropdownButton) HiddenOn(value string) *DropdownButton {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否只显示图标。
 */
func (a *DropdownButton) IconOnly(value string) *DropdownButton {
    a.Set("iconOnly", value)
    return a
}

/**
 * 触发条件，默认是 click
 * 可选值: click | hover
 */
func (a *DropdownButton) Trigger(value string) *DropdownButton {
    a.Set("trigger", value)
    return a
}

/**
 * 指定为 DropDown Button 类型
 */
func (a *DropdownButton) Type(value string) *DropdownButton {
    a.Set("type", value)
    return a
}

/**
 * 按钮文字
 */
func (a *DropdownButton) Label(value string) *DropdownButton {
    a.Set("label", value)
    return a
}

/**
 * 点击内容是否关闭
 */
func (a *DropdownButton) CloseOnClick(value string) *DropdownButton {
    a.Set("closeOnClick", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *DropdownButton) StaticPlaceholder(value string) *DropdownButton {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 组件样式
 */
func (a *DropdownButton) Style(value string) *DropdownButton {
    a.Set("style", value)
    return a
}

/**
 */
func (a *DropdownButton) Testid(value string) *DropdownButton {
    a.Set("testid", value)
    return a
}

/**
 * 内容区域
 */
func (a *DropdownButton) Body(value string) *DropdownButton {
    a.Set("body", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *DropdownButton) DisabledOn(value string) *DropdownButton {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *DropdownButton) Hidden(value string) *DropdownButton {
    a.Set("hidden", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *DropdownButton) Id(value string) *DropdownButton {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *DropdownButton) Static(value string) *DropdownButton {
    a.Set("static", value)
    return a
}

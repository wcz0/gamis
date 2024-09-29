package renderers


/**
 * 下拉按钮渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/dropdown-button

 * @author wcz0
 * @version 6.2.2
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


func (a *DropdownButton) Set(name string, value interface{}) *DropdownButton {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *DropdownButton) UseMobileUI(value interface{}) *DropdownButton {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *DropdownButton) TestIdBuilder(value interface{}) *DropdownButton {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 指定为 DropDown Button 类型
 */
func (a *DropdownButton) Type(value interface{}) *DropdownButton {
    a.Set("type", value)
    return a
}

/**
 * 按钮级别，样式
 * 可选值: info | success | danger | warning | primary | link
 */
func (a *DropdownButton) Level(value interface{}) *DropdownButton {
    a.Set("level", value)
    return a
}

/**
 * 右侧图标
 */
func (a *DropdownButton) RightIcon(value interface{}) *DropdownButton {
    a.Set("rightIcon", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *DropdownButton) VisibleOn(value interface{}) *DropdownButton {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *DropdownButton) StaticPlaceholder(value interface{}) *DropdownButton {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *DropdownButton) StaticSchema(value interface{}) *DropdownButton {
    a.Set("staticSchema", value)
    return a
}

/**
 * 点击内容是否关闭
 */
func (a *DropdownButton) CloseOnClick(value interface{}) *DropdownButton {
    a.Set("closeOnClick", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *DropdownButton) Id(value interface{}) *DropdownButton {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *DropdownButton) StaticInputClassName(value interface{}) *DropdownButton {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否独占一行 `display: block`
 */
func (a *DropdownButton) Block(value interface{}) *DropdownButton {
    a.Set("block", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *DropdownButton) ClassName(value interface{}) *DropdownButton {
    a.Set("className", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *DropdownButton) StaticClassName(value interface{}) *DropdownButton {
    a.Set("staticClassName", value)
    return a
}

/**
 * 对齐方式
 * 可选值: left | right
 */
func (a *DropdownButton) Align(value interface{}) *DropdownButton {
    a.Set("align", value)
    return a
}

/**
 * 按钮集合，支持分组
 */
func (a *DropdownButton) Buttons(value interface{}) *DropdownButton {
    a.Set("buttons", value)
    return a
}

/**
 */
func (a *DropdownButton) OverlayPlacement(value interface{}) *DropdownButton {
    a.Set("overlayPlacement", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *DropdownButton) HiddenOn(value interface{}) *DropdownButton {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *DropdownButton) Visible(value interface{}) *DropdownButton {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *DropdownButton) StaticLabelClassName(value interface{}) *DropdownButton {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 内容区域
 */
func (a *DropdownButton) Body(value interface{}) *DropdownButton {
    a.Set("body", value)
    return a
}

/**
 * 触发条件，默认是 click
 * 可选值: click | hover
 */
func (a *DropdownButton) Trigger(value interface{}) *DropdownButton {
    a.Set("trigger", value)
    return a
}

/**
 * 是否禁用
 */
func (a *DropdownButton) Disabled(value interface{}) *DropdownButton {
    a.Set("disabled", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *DropdownButton) StaticOn(value interface{}) *DropdownButton {
    a.Set("staticOn", value)
    return a
}

/**
 * 给 Button 配置 className。
 */
func (a *DropdownButton) BtnClassName(value interface{}) *DropdownButton {
    a.Set("btnClassName", value)
    return a
}

/**
 * 按钮文字
 */
func (a *DropdownButton) Label(value interface{}) *DropdownButton {
    a.Set("label", value)
    return a
}

/**
 * 是否显示下拉按钮
 */
func (a *DropdownButton) HideCaret(value interface{}) *DropdownButton {
    a.Set("hideCaret", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *DropdownButton) Hidden(value interface{}) *DropdownButton {
    a.Set("hidden", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *DropdownButton) OnEvent(value interface{}) *DropdownButton {
    a.Set("onEvent", value)
    return a
}

/**
 */
func (a *DropdownButton) Testid(value interface{}) *DropdownButton {
    a.Set("testid", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *DropdownButton) Static(value interface{}) *DropdownButton {
    a.Set("static", value)
    return a
}

/**
 * 按钮大小
 * 可选值: xs | sm | md | lg
 */
func (a *DropdownButton) Size(value interface{}) *DropdownButton {
    a.Set("size", value)
    return a
}

/**
 * 菜单 CSS 样式
 */
func (a *DropdownButton) MenuClassName(value interface{}) *DropdownButton {
    a.Set("menuClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *DropdownButton) Style(value interface{}) *DropdownButton {
    a.Set("style", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *DropdownButton) EditorSetting(value interface{}) *DropdownButton {
    a.Set("editorSetting", value)
    return a
}

/**
 * 点击外部是否关闭
 */
func (a *DropdownButton) CloseOnOutside(value interface{}) *DropdownButton {
    a.Set("closeOnOutside", value)
    return a
}

/**
 * 是否只显示图标。
 */
func (a *DropdownButton) IconOnly(value interface{}) *DropdownButton {
    a.Set("iconOnly", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *DropdownButton) DisabledOn(value interface{}) *DropdownButton {
    a.Set("disabledOn", value)
    return a
}

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
 * popOverContainerSelector
 */
func (a *DropdownButton) PopOverContainerSelector(value interface{}) *DropdownButton {
    a.Set("popOverContainerSelector", value)
    return a
}

/**
 * menuClassName
 */
func (a *DropdownButton) MenuClassName(value interface{}) *DropdownButton {
    a.Set("menuClassName", value)
    return a
}

/**
 * staticOn
 */
func (a *DropdownButton) StaticOn(value interface{}) *DropdownButton {
    a.Set("staticOn", value)
    return a
}

/**
 * visible
 */
func (a *DropdownButton) Visible(value interface{}) *DropdownButton {
    a.Set("visible", value)
    return a
}

/**
 * closeOnOutside
 */
func (a *DropdownButton) CloseOnOutside(value interface{}) *DropdownButton {
    a.Set("closeOnOutside", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *DropdownButton) StaticPlaceholder(value interface{}) *DropdownButton {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * style
 */
func (a *DropdownButton) Style(value interface{}) *DropdownButton {
    a.Set("style", value)
    return a
}

/**
 * iconOnly
 */
func (a *DropdownButton) IconOnly(value interface{}) *DropdownButton {
    a.Set("iconOnly", value)
    return a
}

/**
 * disabledOn
 */
func (a *DropdownButton) DisabledOn(value interface{}) *DropdownButton {
    a.Set("disabledOn", value)
    return a
}

/**
 * static
 */
func (a *DropdownButton) Static(value interface{}) *DropdownButton {
    a.Set("static", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *DropdownButton) StaticInputClassName(value interface{}) *DropdownButton {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *DropdownButton) EditorSetting(value interface{}) *DropdownButton {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *DropdownButton) Type(value interface{}) *DropdownButton {
    a.Set("type", value)
    return a
}

/**
 * level
 */
func (a *DropdownButton) Level(value interface{}) *DropdownButton {
    a.Set("level", value)
    return a
}

/**
 * onEvent
 */
func (a *DropdownButton) OnEvent(value interface{}) *DropdownButton {
    a.Set("onEvent", value)
    return a
}

/**
 * useMobileUI
 */
func (a *DropdownButton) UseMobileUI(value interface{}) *DropdownButton {
    a.Set("useMobileUI", value)
    return a
}

/**
 * label
 */
func (a *DropdownButton) Label(value interface{}) *DropdownButton {
    a.Set("label", value)
    return a
}

/**
 * hideCaret
 */
func (a *DropdownButton) HideCaret(value interface{}) *DropdownButton {
    a.Set("hideCaret", value)
    return a
}

/**
 * className
 */
func (a *DropdownButton) ClassName(value interface{}) *DropdownButton {
    a.Set("className", value)
    return a
}

/**
 * hiddenOn
 */
func (a *DropdownButton) HiddenOn(value interface{}) *DropdownButton {
    a.Set("hiddenOn", value)
    return a
}

/**
 * id
 */
func (a *DropdownButton) Id(value interface{}) *DropdownButton {
    a.Set("id", value)
    return a
}

/**
 * staticSchema
 */
func (a *DropdownButton) StaticSchema(value interface{}) *DropdownButton {
    a.Set("staticSchema", value)
    return a
}

/**
 * block
 */
func (a *DropdownButton) Block(value interface{}) *DropdownButton {
    a.Set("block", value)
    return a
}

/**
 * body
 */
func (a *DropdownButton) Body(value interface{}) *DropdownButton {
    a.Set("body", value)
    return a
}

/**
 * disabled
 */
func (a *DropdownButton) Disabled(value interface{}) *DropdownButton {
    a.Set("disabled", value)
    return a
}

/**
 * hidden
 */
func (a *DropdownButton) Hidden(value interface{}) *DropdownButton {
    a.Set("hidden", value)
    return a
}

/**
 * testid
 */
func (a *DropdownButton) Testid(value interface{}) *DropdownButton {
    a.Set("testid", value)
    return a
}

/**
 * btnClassName
 */
func (a *DropdownButton) BtnClassName(value interface{}) *DropdownButton {
    a.Set("btnClassName", value)
    return a
}

/**
 * size
 */
func (a *DropdownButton) Size(value interface{}) *DropdownButton {
    a.Set("size", value)
    return a
}

/**
 * align
 */
func (a *DropdownButton) Align(value interface{}) *DropdownButton {
    a.Set("align", value)
    return a
}

/**
 * trigger
 */
func (a *DropdownButton) Trigger(value interface{}) *DropdownButton {
    a.Set("trigger", value)
    return a
}

/**
 * overlayPlacement
 */
func (a *DropdownButton) OverlayPlacement(value interface{}) *DropdownButton {
    a.Set("overlayPlacement", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *DropdownButton) StaticLabelClassName(value interface{}) *DropdownButton {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * closeOnClick
 */
func (a *DropdownButton) CloseOnClick(value interface{}) *DropdownButton {
    a.Set("closeOnClick", value)
    return a
}

/**
 * visibleOn
 */
func (a *DropdownButton) VisibleOn(value interface{}) *DropdownButton {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *DropdownButton) StaticClassName(value interface{}) *DropdownButton {
    a.Set("staticClassName", value)
    return a
}

/**
 * buttons
 */
func (a *DropdownButton) Buttons(value interface{}) *DropdownButton {
    a.Set("buttons", value)
    return a
}

/**
 * rightIcon
 */
func (a *DropdownButton) RightIcon(value interface{}) *DropdownButton {
    a.Set("rightIcon", value)
    return a
}

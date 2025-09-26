package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type TooltipWrapper struct {
	*BaseRenderer
}

func NewTooltipWrapper() *TooltipWrapper {
    a := &TooltipWrapper{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "tooltip-wrapper")
    return a
}


func (a *TooltipWrapper) Set(name string, value interface{}) *TooltipWrapper {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * tooltipArrowClassName
 */
func (a *TooltipWrapper) TooltipArrowClassName(value interface{}) *TooltipWrapper {
    a.Set("tooltipArrowClassName", value)
    return a
}

/**
 * static
 */
func (a *TooltipWrapper) Static(value interface{}) *TooltipWrapper {
    a.Set("static", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *TooltipWrapper) StaticLabelClassName(value interface{}) *TooltipWrapper {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * testid
 */
func (a *TooltipWrapper) Testid(value interface{}) *TooltipWrapper {
    a.Set("testid", value)
    return a
}

/**
 * trigger
 */
func (a *TooltipWrapper) Trigger(value interface{}) *TooltipWrapper {
    a.Set("trigger", value)
    return a
}

/**
 * useMobileUI
 */
func (a *TooltipWrapper) UseMobileUI(value interface{}) *TooltipWrapper {
    a.Set("useMobileUI", value)
    return a
}

/**
 * content
 */
func (a *TooltipWrapper) Content(value interface{}) *TooltipWrapper {
    a.Set("content", value)
    return a
}

/**
 * mouseLeaveDelay
 */
func (a *TooltipWrapper) MouseLeaveDelay(value interface{}) *TooltipWrapper {
    a.Set("mouseLeaveDelay", value)
    return a
}

/**
 * staticClassName
 */
func (a *TooltipWrapper) StaticClassName(value interface{}) *TooltipWrapper {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *TooltipWrapper) StaticSchema(value interface{}) *TooltipWrapper {
    a.Set("staticSchema", value)
    return a
}

/**
 * editorSetting
 */
func (a *TooltipWrapper) EditorSetting(value interface{}) *TooltipWrapper {
    a.Set("editorSetting", value)
    return a
}

/**
 * showArrow
 */
func (a *TooltipWrapper) ShowArrow(value interface{}) *TooltipWrapper {
    a.Set("showArrow", value)
    return a
}

/**
 * wrapperComponent
 */
func (a *TooltipWrapper) WrapperComponent(value interface{}) *TooltipWrapper {
    a.Set("wrapperComponent", value)
    return a
}

/**
 * inline
 */
func (a *TooltipWrapper) Inline(value interface{}) *TooltipWrapper {
    a.Set("inline", value)
    return a
}

/**
 * tooltipStyle
 */
func (a *TooltipWrapper) TooltipStyle(value interface{}) *TooltipWrapper {
    a.Set("tooltipStyle", value)
    return a
}

/**
 * disabledOn
 */
func (a *TooltipWrapper) DisabledOn(value interface{}) *TooltipWrapper {
    a.Set("disabledOn", value)
    return a
}

/**
 * visible
 */
func (a *TooltipWrapper) Visible(value interface{}) *TooltipWrapper {
    a.Set("visible", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *TooltipWrapper) StaticPlaceholder(value interface{}) *TooltipWrapper {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * tooltip
 */
func (a *TooltipWrapper) Tooltip(value interface{}) *TooltipWrapper {
    a.Set("tooltip", value)
    return a
}

/**
 * mouseEnterDelay
 */
func (a *TooltipWrapper) MouseEnterDelay(value interface{}) *TooltipWrapper {
    a.Set("mouseEnterDelay", value)
    return a
}

/**
 * rootClose
 */
func (a *TooltipWrapper) RootClose(value interface{}) *TooltipWrapper {
    a.Set("rootClose", value)
    return a
}

/**
 * tooltipTheme
 */
func (a *TooltipWrapper) TooltipTheme(value interface{}) *TooltipWrapper {
    a.Set("tooltipTheme", value)
    return a
}

/**
 * visibleOn
 */
func (a *TooltipWrapper) VisibleOn(value interface{}) *TooltipWrapper {
    a.Set("visibleOn", value)
    return a
}

/**
 * id
 */
func (a *TooltipWrapper) Id(value interface{}) *TooltipWrapper {
    a.Set("id", value)
    return a
}

/**
 * title
 */
func (a *TooltipWrapper) Title(value interface{}) *TooltipWrapper {
    a.Set("title", value)
    return a
}

/**
 * placement
 */
func (a *TooltipWrapper) Placement(value interface{}) *TooltipWrapper {
    a.Set("placement", value)
    return a
}

/**
 * enterable
 */
func (a *TooltipWrapper) Enterable(value interface{}) *TooltipWrapper {
    a.Set("enterable", value)
    return a
}

/**
 * tooltipClassName
 */
func (a *TooltipWrapper) TooltipClassName(value interface{}) *TooltipWrapper {
    a.Set("tooltipClassName", value)
    return a
}

/**
 * disabled
 */
func (a *TooltipWrapper) Disabled(value interface{}) *TooltipWrapper {
    a.Set("disabled", value)
    return a
}

/**
 * hidden
 */
func (a *TooltipWrapper) Hidden(value interface{}) *TooltipWrapper {
    a.Set("hidden", value)
    return a
}

/**
 * hiddenOn
 */
func (a *TooltipWrapper) HiddenOn(value interface{}) *TooltipWrapper {
    a.Set("hiddenOn", value)
    return a
}

/**
 * style
 */
func (a *TooltipWrapper) Style(value interface{}) *TooltipWrapper {
    a.Set("style", value)
    return a
}

/**
 * staticOn
 */
func (a *TooltipWrapper) StaticOn(value interface{}) *TooltipWrapper {
    a.Set("staticOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *TooltipWrapper) StaticInputClassName(value interface{}) *TooltipWrapper {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *TooltipWrapper) Type(value interface{}) *TooltipWrapper {
    a.Set("type", value)
    return a
}

/**
 * className
 */
func (a *TooltipWrapper) ClassName(value interface{}) *TooltipWrapper {
    a.Set("className", value)
    return a
}

/**
 * onEvent
 */
func (a *TooltipWrapper) OnEvent(value interface{}) *TooltipWrapper {
    a.Set("onEvent", value)
    return a
}

/**
 * offset
 */
func (a *TooltipWrapper) Offset(value interface{}) *TooltipWrapper {
    a.Set("offset", value)
    return a
}

/**
 * body
 */
func (a *TooltipWrapper) Body(value interface{}) *TooltipWrapper {
    a.Set("body", value)
    return a
}

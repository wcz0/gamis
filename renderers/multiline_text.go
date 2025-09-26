package renderers


/**
 * MultilineText

 * @author wcz0
 * @version 6.2.2
 */
type MultilineText struct {
	*BaseRenderer
}

func NewMultilineText() *MultilineText {
    a := &MultilineText{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "multiline-text")
    return a
}


func (a *MultilineText) Set(name string, value interface{}) *MultilineText {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * visibleOn
 */
func (a *MultilineText) VisibleOn(value interface{}) *MultilineText {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *MultilineText) StaticSchema(value interface{}) *MultilineText {
    a.Set("staticSchema", value)
    return a
}

/**
 * maxRows
 */
func (a *MultilineText) MaxRows(value interface{}) *MultilineText {
    a.Set("maxRows", value)
    return a
}

/**
 * expendButtonText
 */
func (a *MultilineText) ExpendButtonText(value interface{}) *MultilineText {
    a.Set("expendButtonText", value)
    return a
}

/**
 * collapseButtonText
 */
func (a *MultilineText) CollapseButtonText(value interface{}) *MultilineText {
    a.Set("collapseButtonText", value)
    return a
}

/**
 * disabled
 */
func (a *MultilineText) Disabled(value interface{}) *MultilineText {
    a.Set("disabled", value)
    return a
}

/**
 * disabledOn
 */
func (a *MultilineText) DisabledOn(value interface{}) *MultilineText {
    a.Set("disabledOn", value)
    return a
}

/**
 * hidden
 */
func (a *MultilineText) Hidden(value interface{}) *MultilineText {
    a.Set("hidden", value)
    return a
}

/**
 * hiddenOn
 */
func (a *MultilineText) HiddenOn(value interface{}) *MultilineText {
    a.Set("hiddenOn", value)
    return a
}

/**
 * id
 */
func (a *MultilineText) Id(value interface{}) *MultilineText {
    a.Set("id", value)
    return a
}

/**
 * onEvent
 */
func (a *MultilineText) OnEvent(value interface{}) *MultilineText {
    a.Set("onEvent", value)
    return a
}

/**
 * style
 */
func (a *MultilineText) Style(value interface{}) *MultilineText {
    a.Set("style", value)
    return a
}

/**
 * testid
 */
func (a *MultilineText) Testid(value interface{}) *MultilineText {
    a.Set("testid", value)
    return a
}

/**
 * visible
 */
func (a *MultilineText) Visible(value interface{}) *MultilineText {
    a.Set("visible", value)
    return a
}

/**
 * static
 */
func (a *MultilineText) Static(value interface{}) *MultilineText {
    a.Set("static", value)
    return a
}

/**
 * staticOn
 */
func (a *MultilineText) StaticOn(value interface{}) *MultilineText {
    a.Set("staticOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *MultilineText) StaticLabelClassName(value interface{}) *MultilineText {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * useMobileUI
 */
func (a *MultilineText) UseMobileUI(value interface{}) *MultilineText {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *MultilineText) Type(value interface{}) *MultilineText {
    a.Set("type", value)
    return a
}

/**
 * text
 */
func (a *MultilineText) Text(value interface{}) *MultilineText {
    a.Set("text", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *MultilineText) StaticPlaceholder(value interface{}) *MultilineText {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * className
 */
func (a *MultilineText) ClassName(value interface{}) *MultilineText {
    a.Set("className", value)
    return a
}

/**
 * staticClassName
 */
func (a *MultilineText) StaticClassName(value interface{}) *MultilineText {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *MultilineText) StaticInputClassName(value interface{}) *MultilineText {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *MultilineText) EditorSetting(value interface{}) *MultilineText {
    a.Set("editorSetting", value)
    return a
}

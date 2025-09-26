package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type Steps struct {
	*BaseRenderer
}

func NewSteps() *Steps {
    a := &Steps{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "steps")
    return a
}


func (a *Steps) Set(name string, value interface{}) *Steps {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * className
 */
func (a *Steps) ClassName(value interface{}) *Steps {
    a.Set("className", value)
    return a
}

/**
 * disabled
 */
func (a *Steps) Disabled(value interface{}) *Steps {
    a.Set("disabled", value)
    return a
}

/**
 * disabledOn
 */
func (a *Steps) DisabledOn(value interface{}) *Steps {
    a.Set("disabledOn", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Steps) HiddenOn(value interface{}) *Steps {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Steps) StaticInputClassName(value interface{}) *Steps {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * style
 */
func (a *Steps) Style(value interface{}) *Steps {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *Steps) EditorSetting(value interface{}) *Steps {
    a.Set("editorSetting", value)
    return a
}

/**
 * source
 */
func (a *Steps) Source(value interface{}) *Steps {
    a.Set("source", value)
    return a
}

/**
 * visible
 */
func (a *Steps) Visible(value interface{}) *Steps {
    a.Set("visible", value)
    return a
}

/**
 * staticOn
 */
func (a *Steps) StaticOn(value interface{}) *Steps {
    a.Set("staticOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Steps) StaticLabelClassName(value interface{}) *Steps {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * progressDot
 */
func (a *Steps) ProgressDot(value interface{}) *Steps {
    a.Set("progressDot", value)
    return a
}

/**
 * visibleOn
 */
func (a *Steps) VisibleOn(value interface{}) *Steps {
    a.Set("visibleOn", value)
    return a
}

/**
 * static
 */
func (a *Steps) Static(value interface{}) *Steps {
    a.Set("static", value)
    return a
}

/**
 */
func (a *Steps) Type(value interface{}) *Steps {
    a.Set("type", value)
    return a
}

/**
 * hidden
 */
func (a *Steps) Hidden(value interface{}) *Steps {
    a.Set("hidden", value)
    return a
}

/**
 * onEvent
 */
func (a *Steps) OnEvent(value interface{}) *Steps {
    a.Set("onEvent", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Steps) UseMobileUI(value interface{}) *Steps {
    a.Set("useMobileUI", value)
    return a
}

/**
 * mode
 */
func (a *Steps) Mode(value interface{}) *Steps {
    a.Set("mode", value)
    return a
}

/**
 * labelPlacement
 */
func (a *Steps) LabelPlacement(value interface{}) *Steps {
    a.Set("labelPlacement", value)
    return a
}

/**
 * iconPosition
 */
func (a *Steps) IconPosition(value interface{}) *Steps {
    a.Set("iconPosition", value)
    return a
}

/**
 * staticSchema
 */
func (a *Steps) StaticSchema(value interface{}) *Steps {
    a.Set("staticSchema", value)
    return a
}

/**
 * testid
 */
func (a *Steps) Testid(value interface{}) *Steps {
    a.Set("testid", value)
    return a
}

/**
 * steps
 */
func (a *Steps) Steps(value interface{}) *Steps {
    a.Set("steps", value)
    return a
}

/**
 * value
 */
func (a *Steps) Value(value interface{}) *Steps {
    a.Set("value", value)
    return a
}

/**
 * id
 */
func (a *Steps) Id(value interface{}) *Steps {
    a.Set("id", value)
    return a
}

/**
 * name
 */
func (a *Steps) Name(value interface{}) *Steps {
    a.Set("name", value)
    return a
}

/**
 * status
 */
func (a *Steps) Status(value interface{}) *Steps {
    a.Set("status", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Steps) StaticPlaceholder(value interface{}) *Steps {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticClassName
 */
func (a *Steps) StaticClassName(value interface{}) *Steps {
    a.Set("staticClassName", value)
    return a
}

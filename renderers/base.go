package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type Base struct {
	*BaseRenderer
}

func NewBase() *Base {
    a := &Base{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *Base) Set(name string, value interface{}) *Base {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * testid
 */
func (a *Base) Testid(value interface{}) *Base {
    a.Set("testid", value)
    return a
}

/**
 * disabledOn
 */
func (a *Base) DisabledOn(value interface{}) *Base {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticOn
 */
func (a *Base) StaticOn(value interface{}) *Base {
    a.Set("staticOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *Base) StaticClassName(value interface{}) *Base {
    a.Set("staticClassName", value)
    return a
}

/**
 * type
 */
func (a *Base) Type(value interface{}) *Base {
    a.Set("type", value)
    return a
}

/**
 * visibleOn
 */
func (a *Base) VisibleOn(value interface{}) *Base {
    a.Set("visibleOn", value)
    return a
}

/**
 * id
 */
func (a *Base) Id(value interface{}) *Base {
    a.Set("id", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Base) StaticPlaceholder(value interface{}) *Base {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Base) StaticLabelClassName(value interface{}) *Base {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * style
 */
func (a *Base) Style(value interface{}) *Base {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *Base) EditorSetting(value interface{}) *Base {
    a.Set("editorSetting", value)
    return a
}

/**
 * disabled
 */
func (a *Base) Disabled(value interface{}) *Base {
    a.Set("disabled", value)
    return a
}

/**
 * hidden
 */
func (a *Base) Hidden(value interface{}) *Base {
    a.Set("hidden", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Base) HiddenOn(value interface{}) *Base {
    a.Set("hiddenOn", value)
    return a
}

/**
 * static
 */
func (a *Base) Static(value interface{}) *Base {
    a.Set("static", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Base) UseMobileUI(value interface{}) *Base {
    a.Set("useMobileUI", value)
    return a
}

/**
 * className
 */
func (a *Base) ClassName(value interface{}) *Base {
    a.Set("className", value)
    return a
}

/**
 * visible
 */
func (a *Base) Visible(value interface{}) *Base {
    a.Set("visible", value)
    return a
}

/**
 * onEvent
 */
func (a *Base) OnEvent(value interface{}) *Base {
    a.Set("onEvent", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Base) StaticInputClassName(value interface{}) *Base {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *Base) StaticSchema(value interface{}) *Base {
    a.Set("staticSchema", value)
    return a
}

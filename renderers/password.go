package renderers


/**
 * Password

 * @author wcz0
 * @version 6.2.2
 */
type Password struct {
	*BaseRenderer
}

func NewPassword() *Password {
    a := &Password{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "password")
    return a
}


func (a *Password) Set(name string, value interface{}) *Password {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * staticPlaceholder
 */
func (a *Password) StaticPlaceholder(value interface{}) *Password {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticClassName
 */
func (a *Password) StaticClassName(value interface{}) *Password {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Password) StaticInputClassName(value interface{}) *Password {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * className
 */
func (a *Password) ClassName(value interface{}) *Password {
    a.Set("className", value)
    return a
}

/**
 * hidden
 */
func (a *Password) Hidden(value interface{}) *Password {
    a.Set("hidden", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Password) HiddenOn(value interface{}) *Password {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *Password) StaticSchema(value interface{}) *Password {
    a.Set("staticSchema", value)
    return a
}

/**
 * editorSetting
 */
func (a *Password) EditorSetting(value interface{}) *Password {
    a.Set("editorSetting", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Password) UseMobileUI(value interface{}) *Password {
    a.Set("useMobileUI", value)
    return a
}

/**
 * disabled
 */
func (a *Password) Disabled(value interface{}) *Password {
    a.Set("disabled", value)
    return a
}

/**
 * visible
 */
func (a *Password) Visible(value interface{}) *Password {
    a.Set("visible", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Password) StaticLabelClassName(value interface{}) *Password {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *Password) Type(value interface{}) *Password {
    a.Set("type", value)
    return a
}

/**
 * static
 */
func (a *Password) Static(value interface{}) *Password {
    a.Set("static", value)
    return a
}

/**
 * staticOn
 */
func (a *Password) StaticOn(value interface{}) *Password {
    a.Set("staticOn", value)
    return a
}

/**
 * style
 */
func (a *Password) Style(value interface{}) *Password {
    a.Set("style", value)
    return a
}

/**
 * testid
 */
func (a *Password) Testid(value interface{}) *Password {
    a.Set("testid", value)
    return a
}

/**
 * mosaicText
 */
func (a *Password) MosaicText(value interface{}) *Password {
    a.Set("mosaicText", value)
    return a
}

/**
 * disabledOn
 */
func (a *Password) DisabledOn(value interface{}) *Password {
    a.Set("disabledOn", value)
    return a
}

/**
 * visibleOn
 */
func (a *Password) VisibleOn(value interface{}) *Password {
    a.Set("visibleOn", value)
    return a
}

/**
 * id
 */
func (a *Password) Id(value interface{}) *Password {
    a.Set("id", value)
    return a
}

/**
 * onEvent
 */
func (a *Password) OnEvent(value interface{}) *Password {
    a.Set("onEvent", value)
    return a
}

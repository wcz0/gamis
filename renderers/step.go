package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type Step struct {
	*BaseRenderer
}

func NewStep() *Step {
    a := &Step{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *Step) Set(name string, value interface{}) *Step {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * value
 */
func (a *Step) Value(value interface{}) *Step {
    a.Set("value", value)
    return a
}

/**
 * hidden
 */
func (a *Step) Hidden(value interface{}) *Step {
    a.Set("hidden", value)
    return a
}

/**
 * visible
 */
func (a *Step) Visible(value interface{}) *Step {
    a.Set("visible", value)
    return a
}

/**
 * visibleOn
 */
func (a *Step) VisibleOn(value interface{}) *Step {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *Step) StaticSchema(value interface{}) *Step {
    a.Set("staticSchema", value)
    return a
}

/**
 * style
 */
func (a *Step) Style(value interface{}) *Step {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *Step) EditorSetting(value interface{}) *Step {
    a.Set("editorSetting", value)
    return a
}

/**
 * title
 */
func (a *Step) Title(value interface{}) *Step {
    a.Set("title", value)
    return a
}

/**
 * icon
 */
func (a *Step) Icon(value interface{}) *Step {
    a.Set("icon", value)
    return a
}

/**
 * className
 */
func (a *Step) ClassName(value interface{}) *Step {
    a.Set("className", value)
    return a
}

/**
 * id
 */
func (a *Step) Id(value interface{}) *Step {
    a.Set("id", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Step) UseMobileUI(value interface{}) *Step {
    a.Set("useMobileUI", value)
    return a
}

/**
 * subTitle
 */
func (a *Step) SubTitle(value interface{}) *Step {
    a.Set("subTitle", value)
    return a
}

/**
 * disabled
 */
func (a *Step) Disabled(value interface{}) *Step {
    a.Set("disabled", value)
    return a
}

/**
 * staticOn
 */
func (a *Step) StaticOn(value interface{}) *Step {
    a.Set("staticOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Step) StaticPlaceholder(value interface{}) *Step {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticClassName
 */
func (a *Step) StaticClassName(value interface{}) *Step {
    a.Set("staticClassName", value)
    return a
}

/**
 * description
 */
func (a *Step) Description(value interface{}) *Step {
    a.Set("description", value)
    return a
}

/**
 * disabledOn
 */
func (a *Step) DisabledOn(value interface{}) *Step {
    a.Set("disabledOn", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Step) HiddenOn(value interface{}) *Step {
    a.Set("hiddenOn", value)
    return a
}

/**
 * onEvent
 */
func (a *Step) OnEvent(value interface{}) *Step {
    a.Set("onEvent", value)
    return a
}

/**
 * static
 */
func (a *Step) Static(value interface{}) *Step {
    a.Set("static", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Step) StaticLabelClassName(value interface{}) *Step {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Step) StaticInputClassName(value interface{}) *Step {
    a.Set("staticInputClassName", value)
    return a
}

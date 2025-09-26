package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type State struct {
	*BaseRenderer
}

func NewState() *State {
    a := &State{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *State) Set(name string, value interface{}) *State {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * hiddenOn
 */
func (a *State) HiddenOn(value interface{}) *State {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visible
 */
func (a *State) Visible(value interface{}) *State {
    a.Set("visible", value)
    return a
}

/**
 * onEvent
 */
func (a *State) OnEvent(value interface{}) *State {
    a.Set("onEvent", value)
    return a
}

/**
 * staticSchema
 */
func (a *State) StaticSchema(value interface{}) *State {
    a.Set("staticSchema", value)
    return a
}

/**
 * className
 */
func (a *State) ClassName(value interface{}) *State {
    a.Set("className", value)
    return a
}

/**
 * staticClassName
 */
func (a *State) StaticClassName(value interface{}) *State {
    a.Set("staticClassName", value)
    return a
}

/**
 * title
 */
func (a *State) Title(value interface{}) *State {
    a.Set("title", value)
    return a
}

/**
 * body
 */
func (a *State) Body(value interface{}) *State {
    a.Set("body", value)
    return a
}

/**
 * visibleOn
 */
func (a *State) VisibleOn(value interface{}) *State {
    a.Set("visibleOn", value)
    return a
}

/**
 * static
 */
func (a *State) Static(value interface{}) *State {
    a.Set("static", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *State) StaticPlaceholder(value interface{}) *State {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *State) StaticLabelClassName(value interface{}) *State {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *State) EditorSetting(value interface{}) *State {
    a.Set("editorSetting", value)
    return a
}

/**
 * useMobileUI
 */
func (a *State) UseMobileUI(value interface{}) *State {
    a.Set("useMobileUI", value)
    return a
}

/**
 * id
 */
func (a *State) Id(value interface{}) *State {
    a.Set("id", value)
    return a
}

/**
 * staticOn
 */
func (a *State) StaticOn(value interface{}) *State {
    a.Set("staticOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *State) StaticInputClassName(value interface{}) *State {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * style
 */
func (a *State) Style(value interface{}) *State {
    a.Set("style", value)
    return a
}

/**
 * disabled
 */
func (a *State) Disabled(value interface{}) *State {
    a.Set("disabled", value)
    return a
}

/**
 * disabledOn
 */
func (a *State) DisabledOn(value interface{}) *State {
    a.Set("disabledOn", value)
    return a
}

/**
 * hidden
 */
func (a *State) Hidden(value interface{}) *State {
    a.Set("hidden", value)
    return a
}

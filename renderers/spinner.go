package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type Spinner struct {
	*BaseRenderer
}

func NewSpinner() *Spinner {
    a := &Spinner{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "spinner")
    return a
}


func (a *Spinner) Set(name string, value interface{}) *Spinner {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * onEvent
 */
func (a *Spinner) OnEvent(value interface{}) *Spinner {
    a.Set("onEvent", value)
    return a
}

/**
 * loadingConfig
 */
func (a *Spinner) LoadingConfig(value interface{}) *Spinner {
    a.Set("loadingConfig", value)
    return a
}

/**
 * className
 */
func (a *Spinner) ClassName(value interface{}) *Spinner {
    a.Set("className", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Spinner) UseMobileUI(value interface{}) *Spinner {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Spinner) Type(value interface{}) *Spinner {
    a.Set("type", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Spinner) HiddenOn(value interface{}) *Spinner {
    a.Set("hiddenOn", value)
    return a
}

/**
 * spinnerClassName
 */
func (a *Spinner) SpinnerClassName(value interface{}) *Spinner {
    a.Set("spinnerClassName", value)
    return a
}

/**
 * mode
 */
func (a *Spinner) Mode(value interface{}) *Spinner {
    a.Set("mode", value)
    return a
}

/**
 * size
 */
func (a *Spinner) Size(value interface{}) *Spinner {
    a.Set("size", value)
    return a
}

/**
 * delay
 */
func (a *Spinner) Delay(value interface{}) *Spinner {
    a.Set("delay", value)
    return a
}

/**
 * hidden
 */
func (a *Spinner) Hidden(value interface{}) *Spinner {
    a.Set("hidden", value)
    return a
}

/**
 * visible
 */
func (a *Spinner) Visible(value interface{}) *Spinner {
    a.Set("visible", value)
    return a
}

/**
 * editorSetting
 */
func (a *Spinner) EditorSetting(value interface{}) *Spinner {
    a.Set("editorSetting", value)
    return a
}

/**
 * show
 */
func (a *Spinner) Show(value interface{}) *Spinner {
    a.Set("show", value)
    return a
}

/**
 * spinnerWrapClassName
 */
func (a *Spinner) SpinnerWrapClassName(value interface{}) *Spinner {
    a.Set("spinnerWrapClassName", value)
    return a
}

/**
 * tipPlacement
 */
func (a *Spinner) TipPlacement(value interface{}) *Spinner {
    a.Set("tipPlacement", value)
    return a
}

/**
 * staticOn
 */
func (a *Spinner) StaticOn(value interface{}) *Spinner {
    a.Set("staticOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *Spinner) StaticClassName(value interface{}) *Spinner {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Spinner) StaticLabelClassName(value interface{}) *Spinner {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Spinner) StaticInputClassName(value interface{}) *Spinner {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * body
 */
func (a *Spinner) Body(value interface{}) *Spinner {
    a.Set("body", value)
    return a
}

/**
 * disabled
 */
func (a *Spinner) Disabled(value interface{}) *Spinner {
    a.Set("disabled", value)
    return a
}

/**
 * visibleOn
 */
func (a *Spinner) VisibleOn(value interface{}) *Spinner {
    a.Set("visibleOn", value)
    return a
}

/**
 * id
 */
func (a *Spinner) Id(value interface{}) *Spinner {
    a.Set("id", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Spinner) StaticPlaceholder(value interface{}) *Spinner {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * testid
 */
func (a *Spinner) Testid(value interface{}) *Spinner {
    a.Set("testid", value)
    return a
}

/**
 * overlay
 */
func (a *Spinner) Overlay(value interface{}) *Spinner {
    a.Set("overlay", value)
    return a
}

/**
 * disabledOn
 */
func (a *Spinner) DisabledOn(value interface{}) *Spinner {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *Spinner) StaticSchema(value interface{}) *Spinner {
    a.Set("staticSchema", value)
    return a
}

/**
 * icon
 */
func (a *Spinner) Icon(value interface{}) *Spinner {
    a.Set("icon", value)
    return a
}

/**
 * tip
 */
func (a *Spinner) Tip(value interface{}) *Spinner {
    a.Set("tip", value)
    return a
}

/**
 * static
 */
func (a *Spinner) Static(value interface{}) *Spinner {
    a.Set("static", value)
    return a
}

/**
 * style
 */
func (a *Spinner) Style(value interface{}) *Spinner {
    a.Set("style", value)
    return a
}

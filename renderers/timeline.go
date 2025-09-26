package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type Timeline struct {
	*BaseRenderer
}

func NewTimeline() *Timeline {
    a := &Timeline{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "timeline")
    return a
}


func (a *Timeline) Set(name string, value interface{}) *Timeline {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * reverse
 */
func (a *Timeline) Reverse(value interface{}) *Timeline {
    a.Set("reverse", value)
    return a
}

/**
 * items
 */
func (a *Timeline) Items(value interface{}) *Timeline {
    a.Set("items", value)
    return a
}

/**
 * itemTitleSchema
 */
func (a *Timeline) ItemTitleSchema(value interface{}) *Timeline {
    a.Set("itemTitleSchema", value)
    return a
}

/**
 * timeClassName
 */
func (a *Timeline) TimeClassName(value interface{}) *Timeline {
    a.Set("timeClassName", value)
    return a
}

/**
 * disabledOn
 */
func (a *Timeline) DisabledOn(value interface{}) *Timeline {
    a.Set("disabledOn", value)
    return a
}

/**
 */
func (a *Timeline) Type(value interface{}) *Timeline {
    a.Set("type", value)
    return a
}

/**
 * source
 */
func (a *Timeline) Source(value interface{}) *Timeline {
    a.Set("source", value)
    return a
}

/**
 * titleClassName
 */
func (a *Timeline) TitleClassName(value interface{}) *Timeline {
    a.Set("titleClassName", value)
    return a
}

/**
 * cardSchema
 */
func (a *Timeline) CardSchema(value interface{}) *Timeline {
    a.Set("cardSchema", value)
    return a
}

/**
 * hidden
 */
func (a *Timeline) Hidden(value interface{}) *Timeline {
    a.Set("hidden", value)
    return a
}

/**
 * static
 */
func (a *Timeline) Static(value interface{}) *Timeline {
    a.Set("static", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Timeline) StaticPlaceholder(value interface{}) *Timeline {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticClassName
 */
func (a *Timeline) StaticClassName(value interface{}) *Timeline {
    a.Set("staticClassName", value)
    return a
}

/**
 * className
 */
func (a *Timeline) ClassName(value interface{}) *Timeline {
    a.Set("className", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Timeline) StaticInputClassName(value interface{}) *Timeline {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * mode
 */
func (a *Timeline) Mode(value interface{}) *Timeline {
    a.Set("mode", value)
    return a
}

/**
 * direction
 */
func (a *Timeline) Direction(value interface{}) *Timeline {
    a.Set("direction", value)
    return a
}

/**
 * id
 */
func (a *Timeline) Id(value interface{}) *Timeline {
    a.Set("id", value)
    return a
}

/**
 * style
 */
func (a *Timeline) Style(value interface{}) *Timeline {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *Timeline) EditorSetting(value interface{}) *Timeline {
    a.Set("editorSetting", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Timeline) UseMobileUI(value interface{}) *Timeline {
    a.Set("useMobileUI", value)
    return a
}

/**
 * iconClassName
 */
func (a *Timeline) IconClassName(value interface{}) *Timeline {
    a.Set("iconClassName", value)
    return a
}

/**
 * visibleOn
 */
func (a *Timeline) VisibleOn(value interface{}) *Timeline {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Timeline) StaticLabelClassName(value interface{}) *Timeline {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * testid
 */
func (a *Timeline) Testid(value interface{}) *Timeline {
    a.Set("testid", value)
    return a
}

/**
 * detailClassName
 */
func (a *Timeline) DetailClassName(value interface{}) *Timeline {
    a.Set("detailClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *Timeline) StaticSchema(value interface{}) *Timeline {
    a.Set("staticSchema", value)
    return a
}

/**
 * disabled
 */
func (a *Timeline) Disabled(value interface{}) *Timeline {
    a.Set("disabled", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Timeline) HiddenOn(value interface{}) *Timeline {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visible
 */
func (a *Timeline) Visible(value interface{}) *Timeline {
    a.Set("visible", value)
    return a
}

/**
 * onEvent
 */
func (a *Timeline) OnEvent(value interface{}) *Timeline {
    a.Set("onEvent", value)
    return a
}

/**
 * staticOn
 */
func (a *Timeline) StaticOn(value interface{}) *Timeline {
    a.Set("staticOn", value)
    return a
}

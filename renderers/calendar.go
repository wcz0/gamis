package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type Calendar struct {
	*BaseRenderer
}

func NewCalendar() *Calendar {
    a := &Calendar{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "calendar")
    return a
}


func (a *Calendar) Set(name string, value interface{}) *Calendar {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * staticClassName
 */
func (a *Calendar) StaticClassName(value interface{}) *Calendar {
    a.Set("staticClassName", value)
    return a
}

/**
 * style
 */
func (a *Calendar) Style(value interface{}) *Calendar {
    a.Set("style", value)
    return a
}

/**
 * scheduleClassNames
 */
func (a *Calendar) ScheduleClassNames(value interface{}) *Calendar {
    a.Set("scheduleClassNames", value)
    return a
}

/**
 * largeMode
 */
func (a *Calendar) LargeMode(value interface{}) *Calendar {
    a.Set("largeMode", value)
    return a
}

/**
 * disabledOn
 */
func (a *Calendar) DisabledOn(value interface{}) *Calendar {
    a.Set("disabledOn", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Calendar) HiddenOn(value interface{}) *Calendar {
    a.Set("hiddenOn", value)
    return a
}

/**
 * onEvent
 */
func (a *Calendar) OnEvent(value interface{}) *Calendar {
    a.Set("onEvent", value)
    return a
}

/**
 * testid
 */
func (a *Calendar) Testid(value interface{}) *Calendar {
    a.Set("testid", value)
    return a
}

/**
 * scheduleAction
 */
func (a *Calendar) ScheduleAction(value interface{}) *Calendar {
    a.Set("scheduleAction", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Calendar) StaticPlaceholder(value interface{}) *Calendar {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticSchema
 */
func (a *Calendar) StaticSchema(value interface{}) *Calendar {
    a.Set("staticSchema", value)
    return a
}

/**
 * editorSetting
 */
func (a *Calendar) EditorSetting(value interface{}) *Calendar {
    a.Set("editorSetting", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Calendar) UseMobileUI(value interface{}) *Calendar {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Calendar) Type(value interface{}) *Calendar {
    a.Set("type", value)
    return a
}

/**
 * todayActiveStyle
 */
func (a *Calendar) TodayActiveStyle(value interface{}) *Calendar {
    a.Set("todayActiveStyle", value)
    return a
}

/**
 * disabled
 */
func (a *Calendar) Disabled(value interface{}) *Calendar {
    a.Set("disabled", value)
    return a
}

/**
 * hidden
 */
func (a *Calendar) Hidden(value interface{}) *Calendar {
    a.Set("hidden", value)
    return a
}

/**
 * visible
 */
func (a *Calendar) Visible(value interface{}) *Calendar {
    a.Set("visible", value)
    return a
}

/**
 * staticOn
 */
func (a *Calendar) StaticOn(value interface{}) *Calendar {
    a.Set("staticOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Calendar) StaticLabelClassName(value interface{}) *Calendar {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Calendar) StaticInputClassName(value interface{}) *Calendar {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * schedules
 */
func (a *Calendar) Schedules(value interface{}) *Calendar {
    a.Set("schedules", value)
    return a
}

/**
 * className
 */
func (a *Calendar) ClassName(value interface{}) *Calendar {
    a.Set("className", value)
    return a
}

/**
 * visibleOn
 */
func (a *Calendar) VisibleOn(value interface{}) *Calendar {
    a.Set("visibleOn", value)
    return a
}

/**
 * id
 */
func (a *Calendar) Id(value interface{}) *Calendar {
    a.Set("id", value)
    return a
}

/**
 * static
 */
func (a *Calendar) Static(value interface{}) *Calendar {
    a.Set("static", value)
    return a
}

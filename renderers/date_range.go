package renderers


/**
 * DateRange 展示渲染器。

 * @author wcz0
 * @version 6.2.2
 */
type DateRange struct {
	*BaseRenderer
}

func NewDateRange() *DateRange {
    a := &DateRange{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "date-range")
    return a
}


func (a *DateRange) Set(name string, value interface{}) *DateRange {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * hidden
 */
func (a *DateRange) Hidden(value interface{}) *DateRange {
    a.Set("hidden", value)
    return a
}

/**
 * visible
 */
func (a *DateRange) Visible(value interface{}) *DateRange {
    a.Set("visible", value)
    return a
}

/**
 * staticOn
 */
func (a *DateRange) StaticOn(value interface{}) *DateRange {
    a.Set("staticOn", value)
    return a
}

/**
 * style
 */
func (a *DateRange) Style(value interface{}) *DateRange {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *DateRange) EditorSetting(value interface{}) *DateRange {
    a.Set("editorSetting", value)
    return a
}

/**
 * displayFormat
 */
func (a *DateRange) DisplayFormat(value interface{}) *DateRange {
    a.Set("displayFormat", value)
    return a
}

/**
 * disabled
 */
func (a *DateRange) Disabled(value interface{}) *DateRange {
    a.Set("disabled", value)
    return a
}

/**
 * static
 */
func (a *DateRange) Static(value interface{}) *DateRange {
    a.Set("static", value)
    return a
}

/**
 * staticSchema
 */
func (a *DateRange) StaticSchema(value interface{}) *DateRange {
    a.Set("staticSchema", value)
    return a
}

/**
 * delimiter
 */
func (a *DateRange) Delimiter(value interface{}) *DateRange {
    a.Set("delimiter", value)
    return a
}

/**
 * hiddenOn
 */
func (a *DateRange) HiddenOn(value interface{}) *DateRange {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *DateRange) StaticPlaceholder(value interface{}) *DateRange {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *DateRange) StaticLabelClassName(value interface{}) *DateRange {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * useMobileUI
 */
func (a *DateRange) UseMobileUI(value interface{}) *DateRange {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *DateRange) Type(value interface{}) *DateRange {
    a.Set("type", value)
    return a
}

/**
 * valueFormat
 */
func (a *DateRange) ValueFormat(value interface{}) *DateRange {
    a.Set("valueFormat", value)
    return a
}

/**
 * connector
 */
func (a *DateRange) Connector(value interface{}) *DateRange {
    a.Set("connector", value)
    return a
}

/**
 * disabledOn
 */
func (a *DateRange) DisabledOn(value interface{}) *DateRange {
    a.Set("disabledOn", value)
    return a
}

/**
 * visibleOn
 */
func (a *DateRange) VisibleOn(value interface{}) *DateRange {
    a.Set("visibleOn", value)
    return a
}

/**
 * id
 */
func (a *DateRange) Id(value interface{}) *DateRange {
    a.Set("id", value)
    return a
}

/**
 * onEvent
 */
func (a *DateRange) OnEvent(value interface{}) *DateRange {
    a.Set("onEvent", value)
    return a
}

/**
 * staticClassName
 */
func (a *DateRange) StaticClassName(value interface{}) *DateRange {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *DateRange) StaticInputClassName(value interface{}) *DateRange {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * testid
 */
func (a *DateRange) Testid(value interface{}) *DateRange {
    a.Set("testid", value)
    return a
}

/**
 * format
 */
func (a *DateRange) Format(value interface{}) *DateRange {
    a.Set("format", value)
    return a
}

/**
 * className
 */
func (a *DateRange) ClassName(value interface{}) *DateRange {
    a.Set("className", value)
    return a
}

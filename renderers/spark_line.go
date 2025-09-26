package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type SparkLine struct {
	*BaseRenderer
}

func NewSparkLine() *SparkLine {
    a := &SparkLine{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "sparkline")
    return a
}


func (a *SparkLine) Set(name string, value interface{}) *SparkLine {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * disabled
 */
func (a *SparkLine) Disabled(value interface{}) *SparkLine {
    a.Set("disabled", value)
    return a
}

/**
 * hiddenOn
 */
func (a *SparkLine) HiddenOn(value interface{}) *SparkLine {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visible
 */
func (a *SparkLine) Visible(value interface{}) *SparkLine {
    a.Set("visible", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *SparkLine) StaticLabelClassName(value interface{}) *SparkLine {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * onEvent
 */
func (a *SparkLine) OnEvent(value interface{}) *SparkLine {
    a.Set("onEvent", value)
    return a
}

/**
 * staticClassName
 */
func (a *SparkLine) StaticClassName(value interface{}) *SparkLine {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *SparkLine) StaticInputClassName(value interface{}) *SparkLine {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * staticOn
 */
func (a *SparkLine) StaticOn(value interface{}) *SparkLine {
    a.Set("staticOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *SparkLine) StaticPlaceholder(value interface{}) *SparkLine {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * style
 */
func (a *SparkLine) Style(value interface{}) *SparkLine {
    a.Set("style", value)
    return a
}

/**
 * useMobileUI
 */
func (a *SparkLine) UseMobileUI(value interface{}) *SparkLine {
    a.Set("useMobileUI", value)
    return a
}

/**
 * testid
 */
func (a *SparkLine) Testid(value interface{}) *SparkLine {
    a.Set("testid", value)
    return a
}

/**
 * name
 */
func (a *SparkLine) Name(value interface{}) *SparkLine {
    a.Set("name", value)
    return a
}

/**
 * width
 */
func (a *SparkLine) Width(value interface{}) *SparkLine {
    a.Set("width", value)
    return a
}

/**
 * value
 */
func (a *SparkLine) Value(value interface{}) *SparkLine {
    a.Set("value", value)
    return a
}

/**
 * disabledOn
 */
func (a *SparkLine) DisabledOn(value interface{}) *SparkLine {
    a.Set("disabledOn", value)
    return a
}

/**
 * visibleOn
 */
func (a *SparkLine) VisibleOn(value interface{}) *SparkLine {
    a.Set("visibleOn", value)
    return a
}

/**
 * static
 */
func (a *SparkLine) Static(value interface{}) *SparkLine {
    a.Set("static", value)
    return a
}

/**
 */
func (a *SparkLine) Type(value interface{}) *SparkLine {
    a.Set("type", value)
    return a
}

/**
 * height
 */
func (a *SparkLine) Height(value interface{}) *SparkLine {
    a.Set("height", value)
    return a
}

/**
 * clickAction
 */
func (a *SparkLine) ClickAction(value interface{}) *SparkLine {
    a.Set("clickAction", value)
    return a
}

/**
 * placeholder
 */
func (a *SparkLine) Placeholder(value interface{}) *SparkLine {
    a.Set("placeholder", value)
    return a
}

/**
 * editorSetting
 */
func (a *SparkLine) EditorSetting(value interface{}) *SparkLine {
    a.Set("editorSetting", value)
    return a
}

/**
 * className
 */
func (a *SparkLine) ClassName(value interface{}) *SparkLine {
    a.Set("className", value)
    return a
}

/**
 * hidden
 */
func (a *SparkLine) Hidden(value interface{}) *SparkLine {
    a.Set("hidden", value)
    return a
}

/**
 * id
 */
func (a *SparkLine) Id(value interface{}) *SparkLine {
    a.Set("id", value)
    return a
}

/**
 * staticSchema
 */
func (a *SparkLine) StaticSchema(value interface{}) *SparkLine {
    a.Set("staticSchema", value)
    return a
}

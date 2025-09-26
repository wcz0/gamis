package renderers


/**
 * Date 展示渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/date

 * @author wcz0
 * @version 6.2.2
 */
type Date struct {
	*BaseRenderer
}

func NewDate() *Date {
    a := &Date{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "date")
    return a
}


func (a *Date) Set(name string, value interface{}) *Date {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * disabledOn
 */
func (a *Date) DisabledOn(value interface{}) *Date {
    a.Set("disabledOn", value)
    return a
}

/**
 * hidden
 */
func (a *Date) Hidden(value interface{}) *Date {
    a.Set("hidden", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Date) HiddenOn(value interface{}) *Date {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticOn
 */
func (a *Date) StaticOn(value interface{}) *Date {
    a.Set("staticOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Date) StaticInputClassName(value interface{}) *Date {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * format
 */
func (a *Date) Format(value interface{}) *Date {
    a.Set("format", value)
    return a
}

/**
 * className
 */
func (a *Date) ClassName(value interface{}) *Date {
    a.Set("className", value)
    return a
}

/**
 * id
 */
func (a *Date) Id(value interface{}) *Date {
    a.Set("id", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Date) StaticPlaceholder(value interface{}) *Date {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticSchema
 */
func (a *Date) StaticSchema(value interface{}) *Date {
    a.Set("staticSchema", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Date) UseMobileUI(value interface{}) *Date {
    a.Set("useMobileUI", value)
    return a
}

/**
 * updateFrequency
 */
func (a *Date) UpdateFrequency(value interface{}) *Date {
    a.Set("updateFrequency", value)
    return a
}

/**
 * visibleOn
 */
func (a *Date) VisibleOn(value interface{}) *Date {
    a.Set("visibleOn", value)
    return a
}

/**
 * 可选值: date | datetime | time | static-date | static-datetime | static-time
 */
func (a *Date) Type(value interface{}) *Date {
    a.Set("type", value)
    return a
}

/**
 * valueFormat
 */
func (a *Date) ValueFormat(value interface{}) *Date {
    a.Set("valueFormat", value)
    return a
}

/**
 * displayTimeZone
 */
func (a *Date) DisplayTimeZone(value interface{}) *Date {
    a.Set("displayTimeZone", value)
    return a
}

/**
 * visible
 */
func (a *Date) Visible(value interface{}) *Date {
    a.Set("visible", value)
    return a
}

/**
 * editorSetting
 */
func (a *Date) EditorSetting(value interface{}) *Date {
    a.Set("editorSetting", value)
    return a
}

/**
 * fromNow
 */
func (a *Date) FromNow(value interface{}) *Date {
    a.Set("fromNow", value)
    return a
}

/**
 * disabled
 */
func (a *Date) Disabled(value interface{}) *Date {
    a.Set("disabled", value)
    return a
}

/**
 * staticClassName
 */
func (a *Date) StaticClassName(value interface{}) *Date {
    a.Set("staticClassName", value)
    return a
}

/**
 * testid
 */
func (a *Date) Testid(value interface{}) *Date {
    a.Set("testid", value)
    return a
}

/**
 * onEvent
 */
func (a *Date) OnEvent(value interface{}) *Date {
    a.Set("onEvent", value)
    return a
}

/**
 * displayFormat
 */
func (a *Date) DisplayFormat(value interface{}) *Date {
    a.Set("displayFormat", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Date) StaticLabelClassName(value interface{}) *Date {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * style
 */
func (a *Date) Style(value interface{}) *Date {
    a.Set("style", value)
    return a
}

/**
 * static
 */
func (a *Date) Static(value interface{}) *Date {
    a.Set("static", value)
    return a
}

/**
 * placeholder
 */
func (a *Date) Placeholder(value interface{}) *Date {
    a.Set("placeholder", value)
    return a
}

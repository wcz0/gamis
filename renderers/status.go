package renderers


/**
 * 状态展示控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/status

 * @author wcz0
 * @version 6.2.2
 */
type Status struct {
	*BaseRenderer
}

func NewStatus() *Status {
    a := &Status{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "status")
    return a
}


func (a *Status) Set(name string, value interface{}) *Status {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * id
 */
func (a *Status) Id(value interface{}) *Status {
    a.Set("id", value)
    return a
}

/**
 * staticOn
 */
func (a *Status) StaticOn(value interface{}) *Status {
    a.Set("staticOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *Status) StaticClassName(value interface{}) *Status {
    a.Set("staticClassName", value)
    return a
}

/**
 * testid
 */
func (a *Status) Testid(value interface{}) *Status {
    a.Set("testid", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Status) HiddenOn(value interface{}) *Status {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Status) StaticPlaceholder(value interface{}) *Status {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Status) StaticLabelClassName(value interface{}) *Status {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Status) StaticInputClassName(value interface{}) *Status {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * map
 */
func (a *Status) Map(value interface{}) *Status {
    a.Set("map", value)
    return a
}

/**
 * source
 */
func (a *Status) Source(value interface{}) *Status {
    a.Set("source", value)
    return a
}

/**
 * className
 */
func (a *Status) ClassName(value interface{}) *Status {
    a.Set("className", value)
    return a
}

/**
 * disabledOn
 */
func (a *Status) DisabledOn(value interface{}) *Status {
    a.Set("disabledOn", value)
    return a
}

/**
 * onEvent
 */
func (a *Status) OnEvent(value interface{}) *Status {
    a.Set("onEvent", value)
    return a
}

/**
 * static
 */
func (a *Status) Static(value interface{}) *Status {
    a.Set("static", value)
    return a
}

/**
 * staticSchema
 */
func (a *Status) StaticSchema(value interface{}) *Status {
    a.Set("staticSchema", value)
    return a
}

/**
 * style
 */
func (a *Status) Style(value interface{}) *Status {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *Status) EditorSetting(value interface{}) *Status {
    a.Set("editorSetting", value)
    return a
}

/**
 * disabled
 */
func (a *Status) Disabled(value interface{}) *Status {
    a.Set("disabled", value)
    return a
}

/**
 * hidden
 */
func (a *Status) Hidden(value interface{}) *Status {
    a.Set("hidden", value)
    return a
}

/**
 * visible
 */
func (a *Status) Visible(value interface{}) *Status {
    a.Set("visible", value)
    return a
}

/**
 * visibleOn
 */
func (a *Status) VisibleOn(value interface{}) *Status {
    a.Set("visibleOn", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Status) UseMobileUI(value interface{}) *Status {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Status) Type(value interface{}) *Status {
    a.Set("type", value)
    return a
}

/**
 * placeholder
 */
func (a *Status) Placeholder(value interface{}) *Status {
    a.Set("placeholder", value)
    return a
}

/**
 * labelMap
 */
func (a *Status) LabelMap(value interface{}) *Status {
    a.Set("labelMap", value)
    return a
}

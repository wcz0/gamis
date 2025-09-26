package renderers


/**
 * Color 显示渲染器，格式说明。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/color

 * @author wcz0
 * @version 6.2.2
 */
type Color struct {
	*BaseRenderer
}

func NewColor() *Color {
    a := &Color{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "color")
    return a
}


func (a *Color) Set(name string, value interface{}) *Color {
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
func (a *Color) StaticClassName(value interface{}) *Color {
    a.Set("staticClassName", value)
    return a
}

/**
 * onEvent
 */
func (a *Color) OnEvent(value interface{}) *Color {
    a.Set("onEvent", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Color) StaticInputClassName(value interface{}) *Color {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * testid
 */
func (a *Color) Testid(value interface{}) *Color {
    a.Set("testid", value)
    return a
}

/**
 * defaultColor
 */
func (a *Color) DefaultColor(value interface{}) *Color {
    a.Set("defaultColor", value)
    return a
}

/**
 * showValue
 */
func (a *Color) ShowValue(value interface{}) *Color {
    a.Set("showValue", value)
    return a
}

/**
 * staticOn
 */
func (a *Color) StaticOn(value interface{}) *Color {
    a.Set("staticOn", value)
    return a
}

/**
 * disabled
 */
func (a *Color) Disabled(value interface{}) *Color {
    a.Set("disabled", value)
    return a
}

/**
 * disabledOn
 */
func (a *Color) DisabledOn(value interface{}) *Color {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Color) StaticLabelClassName(value interface{}) *Color {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *Color) EditorSetting(value interface{}) *Color {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *Color) Type(value interface{}) *Color {
    a.Set("type", value)
    return a
}

/**
 * visibleOn
 */
func (a *Color) VisibleOn(value interface{}) *Color {
    a.Set("visibleOn", value)
    return a
}

/**
 * static
 */
func (a *Color) Static(value interface{}) *Color {
    a.Set("static", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Color) StaticPlaceholder(value interface{}) *Color {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticSchema
 */
func (a *Color) StaticSchema(value interface{}) *Color {
    a.Set("staticSchema", value)
    return a
}

/**
 * style
 */
func (a *Color) Style(value interface{}) *Color {
    a.Set("style", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Color) UseMobileUI(value interface{}) *Color {
    a.Set("useMobileUI", value)
    return a
}

/**
 * id
 */
func (a *Color) Id(value interface{}) *Color {
    a.Set("id", value)
    return a
}

/**
 * className
 */
func (a *Color) ClassName(value interface{}) *Color {
    a.Set("className", value)
    return a
}

/**
 * hidden
 */
func (a *Color) Hidden(value interface{}) *Color {
    a.Set("hidden", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Color) HiddenOn(value interface{}) *Color {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visible
 */
func (a *Color) Visible(value interface{}) *Color {
    a.Set("visible", value)
    return a
}
